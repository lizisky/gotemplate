package httphandlers

import (
	"fmt"
	"net/http"
	"time"

	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/http/httpurl"
)

type requestParam_echoData struct {
	Info string `json:"info,omitempty" validate:"min=2"`
}

func (ed *requestParam_echoData) IsValid() httpdata.HttpError {
	// if err := utils.IsValidStruct(ed); err != nil {
	// 	return httpdata.RcSyetemErr.WithArgs(httpdata.RC_InvalidHttpRequestParams)
	// }
	return nil
}

// requestHandler_echoData implements the "Echo message" interface
type requestHandler_echoData struct {
}

func (handler *requestHandler_echoData) Method() string {
	return http.MethodPost
}

func (handler *requestHandler_echoData) URL() string {
	return httpurl.EchoMessage
}

func (handler *requestHandler_echoData) PrepareRequestData() httpdata.HttpDataWithValidator {
	var echo requestParam_echoData
	return &echo
}

func (handler *requestHandler_echoData) Handle(rdata *httpdata.RequestData) (interface{}, httpdata.HttpError) {

	info, ok := rdata.Data.(*requestParam_echoData)
	if !ok {
		return nil, httpdata.RcConvertRequestParamDataErr
	}

	fmt.Printf("verify echo message, length:%d, data is:%s\n", len(info.Info), info.Info)
	response := httpdata.NewResponseSucc()
	response.Msg = fmt.Sprintf("LiziCS Server: [%s], Your Time Is: [%s]", info.Info, time.UnixMilli(rdata.Base.Time))

	return response, nil
}
