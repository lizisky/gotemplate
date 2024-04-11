package httpdata

import (
	"fmt"

	"lizisky.com/lizisky/src/sessionmgr/sessiondata"
)

/**
 * common request & response data structure
 **/

type (
	// http execution error interface
	HttpError interface {
		Code() int
		Error() string
	}

	// HttpErrorData define the http error data structure
	HttpErrorData struct {
		ErrCode int    // Error Code
		Msg     string // Error Message
	}

	// each http API data with this validator
	HttpDataWithValidator interface {
		IsValid() HttpError
	}

	// HttpHeaderParams contains all data in the http request
	HttpHeaderParams struct {
		Sign string // header: Signature
	}

	// HTTP Request format definitions of base partion
	RequestBase struct {
		Time         int64  `json:"time,omitempty"`  // Timestamp, UNIX时间戳,毫秒(13位整数),Beijing 时间, UTC+8
		Aid          uint64 `json:"aid,omitempty"`   // account ID
		SessionToken string `json:"token,omitempty"` // session token，
	}

	RequestMagic struct {
		SessionInfo *sessiondata.SessionInfo
	}

	// RequestData suggest each request handle define this one individually
	RequestData struct {
		Base  RequestBase           `json:"base,omitempty"`
		Data  HttpDataWithValidator `json:"data,omitempty"`
		Magic RequestMagic          `json:"-"`
	}

	// ResponseData defines the common http response format
	ResponseData struct {
		RC   int         `json:"rc"`             // return code
		Msg  string      `json:"msg,omitempty"`  // return information
		Data interface{} `json:"data,omitempty"` // specific data for each interface
		// DebugInfo string      `json:"debugInfo,omitempty"` // debug information for development phase ONLY
	}
)

// ----------------------------------------------------------------------------
// implement HttpError interface for HttpErrorData structure

func (err *HttpErrorData) Code() int {
	return err.ErrCode
}

func (err *HttpErrorData) Error() string {
	return err.Msg
}

func (err *HttpErrorData) WithArgs(arg ...any) HttpError {
	return &HttpErrorData{
		ErrCode: err.ErrCode,
		Msg:     fmt.Sprintf(err.Msg, arg...),
	}
}

// ----------------------------------------------------------------------------
// func (header *HttpHeaderParams) IsValid() HttpError {
// 	// if constants.LengthSignature == len(header.Sign) {
// 	return nil
// 	// }

// 	// return errors.New("invalid signature")
// }

// ----------------------------------------------------------------------------
// implement HttpDataWithValidator interface for RequestBase structure

func (base *RequestBase) IsValid() HttpError {
	return nil
}

// ----------------------------------------------------------------------------
// implement HttpDataWithValidator interface for RequestData structure

func (data *RequestData) IsValid() HttpError {
	if err := data.Base.IsValid(); err != nil {
		return err
	}

	if err := data.Data.IsValid(); err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------------------------------------
// helper functions to create http request response data easily

func NewResponseSucc(arg ...interface{}) *ResponseData {
	response := &ResponseData{RC: RC_OkCode}
	if len(arg) > 0 {
		response.Data = arg[0]
	}
	return response
}

func NewResponseFromError(err HttpError, rspData interface{}, arg ...any) *ResponseData {
	response := &ResponseData{RC: err.Code(), Data: rspData}
	if len(arg) > 0 {
		response.Msg = fmt.Sprintf(err.Error(), arg...)
	} else {
		response.Msg = err.Error()
	}
	return response
}

// ----------------------------------------------------------------------------
// helper functions to create http error data easily for data validation

func NewError(code int, err error) HttpError {
	return &HttpErrorData{
		ErrCode: code,
		Msg:     err.Error(),
	}
}

// ----------------------------------------------------------------------------
