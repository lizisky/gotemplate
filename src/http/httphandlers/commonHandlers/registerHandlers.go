package commonHandlers

import "lizisky.com/lizisky/src/http/httpServer"

// RegisterHandlers register http request handlers
func RegisterHandlers() {
	httpServer.RegisterHandler(&requestHandler_pullBootupInfo{})
	httpServer.RegisterHandler(&requestHandler_getQiniuUploadToken{})
	httpServer.RegisterHandler(&requestHandler_readHomepageData{})
}
