package httpServer

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"lizisky.com/lizisky/src/http/httpdata"
)

var (
	gRouter *gin.Engine
)

// Start http server to listen
func Start(addr string) {
	gServer := &http.Server{
		Addr:           addr,
		Handler:        gRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	gServer.ListenAndServe()
}

// RegisterHandler register all HTTP handlers into here
func RegisterHandler(handler httpdata.HttpHandler) {
	registerHandler_with_validate_session_flag(handler, true)
}

// RegisterHandler register all HTTP handlers into here
func RegisterHandler_without_session(handler httpdata.HttpHandler) {
	registerHandler_with_validate_session_flag(handler, false)
}

// RegisterHandler register all HTTP handlers into here
func registerHandler_with_validate_session_flag(handler httpdata.HttpHandler, needValidateSession bool) {
	glog.Infoln("Register http handler:", handler.URL())
	switch handler.Method() {
	case http.MethodGet:
		gRouter.GET(handler.URL(), func(ctx *gin.Context) {
			handleHttpRequest(handler, ctx, needValidateSession)
		})
	case http.MethodPost:
		gRouter.POST(handler.URL(), func(ctx *gin.Context) {
			handleHttpRequest(handler, ctx, needValidateSession)
		})
	default:
		break
	}
}

func globalRecover(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if rec := recover(); rec != nil {
			glog.Error("server panic: ", rec, string(debug.Stack()))

			response := httpdata.NewResponseFromError(httpdata.RcSyetemErr, nil, httpdata.RC_SystemCrashAndRecovered)
			ctx.JSON(http.StatusOK, response)
		}
	}(ctx)
	ctx.Next()
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	gRouter = gin.Default()
	gRouter.Use(globalRecover)
	gRouter.Use(logUrl)
}

func logUrl(c *gin.Context) {
	fmt.Printf("\n\n")
	// glog.Infoln("req url ---> ", c.Request.URL)
	c.Next()
}
