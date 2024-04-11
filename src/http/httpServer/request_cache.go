package httpServer

import (
	"sync"
	"time"
)

const (
	// http 请求缓存的基本计数单位
	baseTimeUnit = 100 // 100 Milliseconds

	// 每一次http请求是否在当前时间前后 %timeSlot_HttpRequestCache% 的时间窗口内
	// 目前定义为5秒
	timeSlot_HttpRequestCache = 5 * 10 * baseTimeUnit // 5秒（单位：毫秒）
)

var (
	lCache *requestCache
)

type requestCache struct {
	// Key: hash of request body
	hashHistory map[string]struct{}

	// key:  time(second / 10), value: (hash of request body) in this second
	// 可以考虑把这个时间切分的更细，这样才过期之后做remove的时候，频率可以更高一些，会使得每一次执行的会更快。这有点儿类似于vm中的垃圾收集
	hashList map[int64][]string

	mutex sync.Mutex

	// 是否需要停止 cleaner ？
	needStopCleaner bool

	// 最后一次从 hashList 中 remove 的 key
	lastRemovedTimeSlotKey int64
}

func init() {
	lCache = &requestCache{
		hashHistory:            make(map[string]struct{}),
		hashList:               make(map[int64][]string),
		needStopCleaner:        true,
		lastRemovedTimeSlotKey: 0,
	}
}

// add a new request hash with timestamp in cache
// timestamp: 以毫秒为单位，在这个cache中，以100毫秒为一个 time slot 存储 hash list
func (cache *requestCache) Add(timestamp int64, hash string) {

	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// http request 的请求时间，除以基本时间单位，就得到 cache map 的 key
	keyOfRequestCache := timestamp / baseTimeUnit
	if len(cache.hashList) < 1 {
		// 如果之前 hashList 为空，则记录lastRemovedTimeSlotKey为当前key - 1
		cache.lastRemovedTimeSlotKey = keyOfRequestCache - 1
		go lCache.startCacheCleaner()
	}

	var resultList []string
	list, ok := cache.hashList[keyOfRequestCache]
	if ok {
		resultList = append(list, hash)
	} else {
		resultList = []string{hash}
	}
	cache.hashList[keyOfRequestCache] = resultList
}

// check a specific request hash existing in cache or not
func (cache *requestCache) IsExisting(hash string) bool {
	_, ok := cache.hashHistory[hash]
	return ok
}

// start to run cache cleaner
func (cache *requestCache) startCacheCleaner() {
	// run once per 100 millisecond
	cache.needStopCleaner = false
	ticker := time.NewTicker(time.Millisecond * baseTimeUnit)
	for {
		<-ticker.C
		cache.clean()
		if cache.needStopCleaner {
			break
		}
	}
	ticker.Stop()
}

//
func (cache *requestCache) clean() {

	key := (time.Now().UnixMilli() - timeSlot_HttpRequestCache) / baseTimeUnit
	if cache.lastRemovedTimeSlotKey+1 > key {
		// 说明在这一轮 clean 动作中，没有需要清理的工作
		return
	}

	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// lastRemovedTimeSlotKey 是上一轮 clean 动作时，已经 remove 的 key
	// 所以，这一次要从 +1 开始
	for cache.lastRemovedTimeSlotKey < key {
		cache.lastRemovedTimeSlotKey++
		list := cache.hashList[cache.lastRemovedTimeSlotKey]
		if len(list) < 1 {
			continue
		}

		delete(cache.hashList, cache.lastRemovedTimeSlotKey)
		for idx := len(list) - 1; idx > -1; idx-- {
			delete(cache.hashHistory, list[idx])
		}
	}

	if len(cache.hashList) < 1 {
		cache.needStopCleaner = true
	}
}
