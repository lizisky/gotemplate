package warRoom

import (
	"lizisky.com/lizisky/src/config"
	"lizisky.com/lizisky/src/dbpool"
	"lizisky.com/lizisky/src/http/httpServer"
	"lizisky.com/lizisky/src/http/httphandlers"
)

func Run() bool {
	if !config.LoadConfig() {
		return false
	}

	dbpool.InitDatabase()
	httphandlers.RegisterHandlers()
	go httpServer.Start(config.GetConfig().ServerAddr)
	return true
}
