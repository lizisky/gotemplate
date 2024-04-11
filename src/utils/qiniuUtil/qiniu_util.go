package qiniuUtil

import (
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"lizisky.com/lizisky/src/basictypes/basictype"
	"lizisky.com/lizisky/src/config"
)

// build qiniu upload token
// @return token:  upload token
// @retuen expires: upload token expires time in milliseconds
func BuildUploadToken() (token string, expires uint64) {
	qiniuCfg := config.GetConfig().Qiniu
	mac := qbox.NewMac(qiniuCfg.AccessKey, qiniuCfg.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:   qiniuCfg.Bucket,
		Expires: qiniuCfg.UploadExpires * 60,
		// Expires: 600, // 10 minutes 有效期
	}

	upToken := putPolicy.UploadToken(mac)
	// fmt.Println("---", mac)
	// fmt.Println("qiniu file upload token:", upToken)

	// expire 时间用毫秒返回，并且不需要 client 再做计算
	exipres := qiniuCfg.UploadExpires*60*1000 + uint64(time.Now().UnixMilli())
	return upToken, exipres
}

// build qiniu download url
// @param key: qiniu file key
// @return downloadURL:  upload token
// @retuen expires: download token expires time in milliseconds
func BuildDownloadURL(key string) (downloadURL string) {
	qiniuCfg := config.GetConfig().Qiniu
	mac := qbox.NewMac(qiniuCfg.AccessKey, qiniuCfg.SecretKey)
	// expires := time.Duration(qiniuCfg.DownloadExpires * 60) // config 文件中的时间单位是分钟
	// deadline := time.Now().Add(expires).Unix()              //1小时有效期
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, qiniuCfg.Domain, key, deadline)
	// fmt.Println("privateAccessURL: ", privateAccessURL)
	return privateAccessURL
}

// func buildDownloadURL(key string) (downloadURL string, expires uint64) {
// 	qiniuCfg := config.GetConfig().Qiniu
// 	mac := qbox.NewMac(qiniuCfg.AccessKey, qiniuCfg.SecretKey)
// 	deadline := qiniuCfg.DownloadExpires * 60
// 	privateAccessURL := storage.MakePrivateURL(mac, qiniuCfg.Domain, key, int64(deadline))
// 	fmt.Println("privateAccessURL: ", privateAccessURL)
// 	return privateAccessURL, deadline * 1000
// }

func BuildDownloadURL_forImgSlice(imgSlice *basictype.ImageSlice) {
	if imgSlice == nil {
		return
	}
	for idx := len(*imgSlice) - 1; idx >= 0; idx-- {
		if len((*imgSlice)[idx]) > 0 {
			(*imgSlice)[idx] = BuildDownloadURL((*imgSlice)[idx])
		}
	}
}
