package accountHandler

import "lizisky.com/lizisky/src/http/httpServer"

// RegisterHandlers register http request handlers
func RegisterHandlers() {
	httpServer.RegisterHandler(&requestHandler_updateUserInfo{})
	httpServer.RegisterHandler(&requestHandler_searchUser{})
}
