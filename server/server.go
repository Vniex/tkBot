package server

import (
	"sync"
	"net/http"
	"time"

	Config "tkBot/config"
	Routers "tkBot/server/routers"
	"tkBot/global"
)

type HttpServer struct {
	server *http.Server

	Tasks *sync.Map
}





func (h *HttpServer)Start(){
	global.InitGlobalVar()
	router:=Routers.InitRouter()
	h.server = &http.Server{
		Addr:           ":"+Config.ServerPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	h.server.ListenAndServe()
}


