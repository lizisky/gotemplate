package commonHandlers

import (
	"net/http"

	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/qiniuUtil"
)

//
// 对与存储在七牛云上的资源，需要通过七牛云的SDK来上传和下载。
// client需要一个完整的上传token，才能够上传资源。
// 本接口用于获取一个完整的上传token。
//

// define http request param
type requestParam_getQiniuUploadToken struct {
	// SessionToken string `json:"sessionToken,omitempty"` // session token，如果用户已经login，则输入用户当前的 session token。如果没有login，设置为空即可
}

// validate request parameters
func (param *requestParam_getQiniuUploadToken) IsValid() httpdata.HttpError {
	// sinfo := sessionmgr.GetSession(param.SessionToken)
	// if sinfo == nil {
	// 	return httpdata.RcSessionExpired
	// }
	return nil
}

type requestResponse_getQiniuUploadToken struct {
	UploadToken string `json:"uploadToken,omitempty"` // qiniu upload token
	Expires     uint64 `json:"expires,omitempty"`     // qiniu upload token expires time in milliseconds
}

// requestHandler_getQiniuUploadToken implements the HttpHandler interface
type requestHandler_getQiniuUploadToken struct {
}

func (handler *requestHandler_getQiniuUploadToken) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_getQiniuUploadToken) URL() string {
	return httpurl.Url_Common_get_qiniu_upload_token
}

func (handler *requestHandler_getQiniuUploadToken) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_getQiniuUploadToken
	return &param
}

func (handler *requestHandler_getQiniuUploadToken) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	token, expires_milliseconds := qiniuUtil.BuildUploadToken()
	response := httpdata.NewResponseSucc(&requestResponse_getQiniuUploadToken{
		UploadToken: token,
		Expires:     expires_milliseconds,
	})

	return response, nil
}
