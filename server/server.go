package server

import (
	"sync"
	"net/http"
	"time"

	Global "tkBot/config"
	Routers "tkBot/server/routers"
)

type HttpServer struct {
	server *http.Server

	Tasks *sync.Map
}



func (h *HttpServer)Start(){
	router:=Routers.InitRouter()
	h.server = &http.Server{
		Addr:           ":"+Global.ServerPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	h.server.ListenAndServe()
}


