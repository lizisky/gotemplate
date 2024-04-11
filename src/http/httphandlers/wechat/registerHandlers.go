package wechat

import "lizisky.com/lizisky/src/http/httpServer"

// RegisterHandlers register http request handlers
func RegisterHandlers() {
	httpServer.RegisterHandler_without_session(&requestHandler_wxLoginAuthCode{})

	httpServer.RegisterHandler(&requestHandler_wxGetPhoneNumber{})
}
