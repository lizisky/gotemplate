package httpServer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/lizisky/liziutils/utils"
	"lizisky.com/lizisky/src/http/httpdata"
	"lizisky.com/lizisky/src/sessionmgr"
)

func handleHttpRequest(handler httpdata.HttpHandler, ctx *gin.Context, needValidateSession bool) {
	// 1: read raw parameters body from http request
	// 2: parse parameters from raw data
	var requestParam *httpdata.RequestData = parseRequestParams(handler, ctx)
	if requestParam == nil {
		return
	}

	// 3: validate parameters
	if !validateParams(ctx, requestParam) {
		return
	}

	if needValidateSession {
		if !validateSession(ctx, requestParam) {
			return
		}
	}

	// 4: execute request
	executeRequest(handler, ctx, requestParam)
}

// 1: read raw parameters body from http request
// func readRequestBodyData(ctx *gin.Context) []byte {
// 	if ctx.Request.ContentLength < 1 {
// 		glog.Infoln("received http request: %s, but the content length is 0\n", ctx.Request.URL)
// 		return nil
// 	}

// 	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
// 	if err != nil {
// 		glog.Infoln("received http request: %s, but read body data failed: %s\n", ctx.Request.URL, err.Error())
// 		return nil
// 	}
// 	if len(bodyBytes) < 1 {
// 		glog.Infoln("received http request: %s, but you didn't provide body data\n", ctx.Request.URL)
// 		return nil
// 	}

// 	glog.Infoln("received http request data: ", string(bodyBytes))

// 	return bodyBytes
// }

// 2: parse parameters from raw data
// func parseRequestParams(handler httpdata.HttpHandler, ctx *gin.Context, rawParametersBody []byte) *httpdata.RequestData {
// 	requestParamLocal := httpdata.RequestData{Data: handler.PrepareRequestData()}
// 	err := json.Unmarshal(rawParametersBody, &requestParamLocal)
// 	if err != nil {
// 		glog.Infof("Parse http body failed: %s\n", err.Error())
// 		httpFeedback := httpdata.ResponseData{
// 			RC:  httpdata.RC_ParseHttpRequestParamsFailed,
// 			Msg: err.Error(),
// 		}
// 		flushJSONData2Client(httpFeedback, ctx.Writer)
// 		return nil
// 	}

// 	return &requestParamLocal
// }

// merge step 1 & step 2
// 1: read raw parameters body from http request
// 2: parse parameters from raw data
func parseRequestParams(handler httpdata.HttpHandler, ctx *gin.Context) *httpdata.RequestData {

	// 1: read raw parameters body from http request
	var rawParametersBody []byte
	{
		if ctx.Request.ContentLength < 1 {
			glog.Infof("received http request: %s, but the content length is 0\n", ctx.Request.URL)
			return nil
		}

		bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			glog.Infof("received http request: %s, but read body data failed: %s\n", ctx.Request.URL, err.Error())
			return nil
		}
		if len(bodyBytes) < 1 {
			glog.Infof("received http request: %s, but you didn't provide body data\n", ctx.Request.URL)
			return nil
		}

		glog.Infoln("received http request data: ", string(bodyBytes))

		rawParametersBody = bodyBytes
	}

	// 2: parse parameters from raw data
	{
		requestParamLocal := httpdata.RequestData{Data: handler.PrepareRequestData()}
		err := json.Unmarshal(rawParametersBody, &requestParamLocal)
		if err != nil {
			glog.Infof("Parse http body failed: %s\n", err.Error())
			httpFeedback := httpdata.ResponseData{
				RC:  httpdata.RC_ParseHttpRequestParamsFailed,
				Msg: err.Error(),
			}
			flushJSONData2Client(httpFeedback, ctx.Writer)
			return nil
		}

		return &requestParamLocal
	}
}

// 3: validate parameters
func validateParams(ctx *gin.Context, requestParam *httpdata.RequestData) bool {
	// MUST validate all received data from client to avoid attack
	if err1 := requestParam.IsValid(); err1 != nil {
		httpFeedback := &httpdata.ResponseData{RC: err1.Code(), Msg: err1.Error()}
		flushJSONData2Client(httpFeedback, ctx.Writer)
		return false
	}
	if err2 := utils.IsValidStruct(requestParam.Data); err2 != nil {
		httpFeedback := &httpdata.ResponseData{RC: httpdata.RC_InvalidHttpRequestParams, Msg: err2.Error()}
		flushJSONData2Client(httpFeedback, ctx.Writer)
		return false
	}
	// if err != nil {
	// 	// 这里的执行次数非常高，需要采用最高效的方式
	// 	// httpFeedback := &httpdata.ResponseData{RC: err.Code(), Msg: err.Error()}
	// 	httpFeedback := &httpdata.ResponseData{RC: httpdata.RC_InvalidHttpRequestParams, Msg: err.Error()}
	// 	flushJSONData2Client(httpFeedback, ctx.Writer)
	// 	return
	// }
	// if err := utils.IsValidStruct(requestParam.Data); err != nil {
	// 	httpFeedback := &httpdata.ResponseData{RC: httpdata.RC_InvalidHttpRequestParams, Msg: err.Error()}
	// 	flushJSONData2Client(httpFeedback, ctx.Writer)
	// 	return
	// }

	return true
}

// validate session token
func validateSession(ctx *gin.Context, requestParam *httpdata.RequestData) bool {
	sinfo := sessionmgr.GetSession(requestParam.Base.SessionToken)
	// session 必须存在，而且必须是当前用户的 session
	if (sinfo == nil) || (sinfo.Aid != requestParam.Base.Aid) {
		httpFeedback := httpdata.NewResponseFromError(httpdata.RcSessionExpired, nil)
		flushJSONData2Client(httpFeedback, ctx.Writer)
		return false
	}

	{
		// update session last access time
		rightNow := uint64(time.Now().UnixMilli())
		sinfo.LastAccessTime = rightNow
	}

	requestParam.Magic.SessionInfo = sinfo

	return true
}

// 4: execute request
func executeRequest(handler httpdata.HttpHandler, ctx *gin.Context, requestParam *httpdata.RequestData) {
	httpFeedback, err := handler.Handle(requestParam)
	if err != nil {
		// 这里的执行次数非常高，需要采用最高效的方式
		httpFeedback = &httpdata.ResponseData{RC: err.Code(), Msg: err.Error()}
	}

	flushJSONData2Client(httpFeedback, ctx.Writer)

}

// flushJSONData2Client flush HTTP handler result to client
func flushJSONData2Client(data interface{}, writer http.ResponseWriter) {
	if (data == nil) || (writer == nil) {
		// logger.Info("FlushJSONData2Client: internel error, data or writer is nil pointer")
		return
	}

	writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	toClient, err := json.Marshal(data)
	if err == nil {
		writer.Write(toClient)
		glog.Infof("FlushJsonData2Clinet success: [%s] \n", string(toClient))
	} else {
		writer.Write([]byte(httpdata.RcSystemErrConveryResultFailed_json_str))
		glog.Infof("FlushJsonData2Clinet failed: [%s] \n", err.Error())
	}
}

func checkSecurity(rawBody []byte, requestData *httpdata.RequestData, requestHeader *httpdata.HttpHeaderParams) error {

	curMilliseconds := time.Now().UnixMilli()

	// 检查时间，这一次请求是否在当前时间前后%timeSlot_HttpRequestCache%的时间窗口内
	{
		if (requestData.Base.Time < curMilliseconds-timeSlot_HttpRequestCache) ||
			(requestData.Base.Time > curMilliseconds+timeSlot_HttpRequestCache) {
			// return httpdata.RcInvalidHttpRequestArgs
		}
	}

	// 看看是不是已经在 cache 里面了
	{
		rawHash := utils.Sha256(rawBody)
		if lCache.IsExisting(rawHash) {
			// this is replay attack
			// return httpdata.RcInvalidHttpRequestArgs
		}
		lCache.Add(curMilliseconds, rawHash)
	}

	// check signature
	{

	}

	return nil
}
