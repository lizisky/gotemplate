package sessiondata

// SessionInfo define the session information
type SessionInfo struct {
	SessionKey     string // lizics Session Key, between LiziCS Server and LiziCS Client
	WxSessionKey   string // WeChat session key, between LiziCS Server and WX-Server
	Aid            uint64 // account id
	Nickname       string // account nick name
	AddTime        uint64 // 这个session的添加时间
	LastAccessTime uint64 // 最后一次读、写这个session数据的时间
}
