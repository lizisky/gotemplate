package httphandlers

import (
	"lizisky.com/lizisky/src/http/httpServer"
	"lizisky.com/lizisky/src/http/httphandlers/accountHandler"
	"lizisky.com/lizisky/src/http/httphandlers/commonHandlers"
	"lizisky.com/lizisky/src/http/httphandlers/organization"
	"lizisky.com/lizisky/src/http/httphandlers/wechat"
)

// RegisterHandlers register http request handlers
func RegisterHandlers() {

	organization.RegisterHandlers()
	commonHandlers.RegisterHandlers()
	accountHandler.RegisterHandlers()
	wechat.RegisterHandlers()

	httpServer.RegisterHandler_without_session(&requestHandler_echoData{})
}
