package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
	"lizisky.com/lizisky/src/warRoom"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Infoln("starting Lizi Golang template Server ......")
	if !warRoom.Run() {
		return
	}
	wait()
	clean()
}

func wait() {
	glog.Infoln("Lizi Golang template Server has started successfully, enjoy it ......")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	glog.Infoln("Received signal, shutting down...")
}

func clean() {

}
