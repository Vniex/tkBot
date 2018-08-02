package main

import (
	Server "tkBot/server"
)

func main() {
	httpServer:=new(Server.HttpServer)
	httpServer.Start()
}
