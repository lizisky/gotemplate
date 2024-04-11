package wechat

import (
	"net/http"

	"github.com/golang/glog"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
	"lizisky.com/lizisky/src/utils/wxutil"
)

//
// 获取微信用户手机号码，
// https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
//

// define http request param
type requestParam_wxGetPhoneNumber struct {
	// 微信小程序获取 Phone Number 的 Code
	Code string `json:"code,omitempty" validate:"omitempty,min=10"`
}

func (acp *requestParam_wxGetPhoneNumber) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(acp); err != nil {
	// 	return httpdata.RcSyetemErr.WithArgs(httpdata.RC_InvalidHttpRequestParams)
	// }
	return nil
}

// define http request response
type requestResponse_wxGetPhoneNumber struct {
	Mobile string `json:"mobile,omitempty"` // phone number
}

// requestHandler_wxGetPhoneNumber implements the HttpHandler interface
type requestHandler_wxGetPhoneNumber struct {
}

func (handler *requestHandler_wxGetPhoneNumber) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_wxGetPhoneNumber) URL() string {
	return httpurl.Url_Wechat_GetPhoneNumber
}

func (handler *requestHandler_wxGetPhoneNumber) PrepareRequestData() httpdata.HttpDataWithValidator {
	var param requestParam_wxGetPhoneNumber
	return &param
}

// Handle http request
func (handler *requestHandler_wxGetPhoneNumber) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {
	params, ok := rdata.Data.(*requestParam_wxGetPhoneNumber)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	phoneNumber, err := wxutil.GetPhoneNumber(params.Code)
	// fmt.Println("----------- GetPhoneNumber 2", phoneNumber, err)
	if err != nil {
		glog.Info("get phone number failed:", err.Error())
		return nil, httpdata.RcGetPhoneNumberFailed
	}

	response := httpdata.NewResponseSucc(&requestResponse_wxGetPhoneNumber{
		Mobile: phoneNumber,
	})

	return response, nil
}
