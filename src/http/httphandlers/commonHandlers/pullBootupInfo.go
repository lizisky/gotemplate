package commonHandlers

import (
	"net/http"

	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
)

//
// Client 每次在App启动的时候，异步调用这个 API 一次即可，主要目的有两点
// 1，获取 client 的信息，当前 client 是否被支持，如果不被支持，则提示用户升级 client 版本
// 2，对于已经 login 的 client，获取与此 user 相关的 notification

// define http request param
type requestParam_pullBootupInfo struct {
	ClientType    int `json:"clientType,omitempty" validate:"gt=0"` // client Type, refer to: <roo/path>/src/basictypes/constants/common.go
	ClientVersion int `json:"clientVer,omitempty" validate:"gt=0"`  // client version, 用于后台校验是否支持当前的 client
	// SessionToken  string      `json:"sessionToken,omitempty"`               // session token，如果用户已经login，则输入用户当前的 session token。如果没有login，设置为空即可
	SystemInfo interface{} `json:"systemInfo,omitempty"` // client system information, 目前在微信小程序端，上传调用wx.getSystemInfo获取的信息即可
}

// validate request parameters
func (param *requestParam_pullBootupInfo) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(param); err != nil {
	// 	return httpdata.NewError(httpdata.RC_InvalidHttpRequestParams, err)
	// }
	return nil
}

// define http request response
type bootupInformation struct {
	InfoType int    `json:"infoType,omitempty"` // information type
	Info     string `json:"info,omitempty"`     // info
	URL      string `json:"url,omitempty"`      // info
}

type requestResponse_pullBootupInfo struct {
	InfoList []*bootupInformation `json:"infoList,omitempty"` // membership list of this org
}

// requestHandler_pullBootupInfo implements the HttpHandler interface
type requestHandler_pullBootupInfo struct {
}

func (handler *requestHandler_pullBootupInfo) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_pullBootupInfo) URL() string {
	return httpurl.Url_Common_Pull_Bootup_Info
}

func (handler *requestHandler_pullBootupInfo) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_pullBootupInfo
	return &param
}

func (handler *requestHandler_pullBootupInfo) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {

	_, ok := rdata.Data.(*requestParam_pullBootupInfo)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	// fmt.Println("receive:", utils.ToJSONIndent(rdata.Data))

	list := []*bootupInformation{
		{
			InfoType: 12345,
			Info:     "Description_11111",
			URL:      "http://lizitime.com/log.png",
		},
		{
			InfoType: 9878,
			Info:     "Description_23949812",
			URL:      "http://lizitime.com/userAggrement.html",
		},
	}
	response := httpdata.NewResponseSucc(requestResponse_pullBootupInfo{
		InfoList: list,
	})

	return response, nil
}
