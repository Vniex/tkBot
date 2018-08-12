package websocket

import (
	"net/http"
	"log"
)

func WsHandlerServer(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	wsConn:=NewWsConnection(wsSocket)

	// 处理器
	go wsConn.ProcLoop(func (msg *RobotMsg){
		log.Printf("server receive %v \n",msg)
	})
	// 读协程
	go wsConn.WsReadLoop()
	// 写协程
	go wsConn.WsWriteLoop()
}

