package main

import (
	Server "tkBot/server"

	Global "tkBot/global"
	log "github.com/sirupsen/logrus"
)


func main() {
	Global.InitLog()
	log.Info("system start....")
	httpServer:=new(Server.HttpServer)
	httpServer.Start()

}
