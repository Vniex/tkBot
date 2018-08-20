package utils

import (
	"net/url"
	"net/http"
	"log"

	Message "tkBot/server/websocket"
	WebSocket "github.com/gorilla/websocket"

)

func SendToWechat(SERVER_SCKEY,text,desp string) {

	wechatUrl:="https://sc.ftqq.com/"+SERVER_SCKEY+".send"
	params := url.Values{}
	params.Set("text",text)
	params.Set("desp",desp)
	http.PostForm(wechatUrl,params)
}


type RobotDetect struct {
	websocketServer string
	wsConn *Message.WsConnection
}


func NewRobotDetect(websocketServer string) *RobotDetect{
	conn, _, err := WebSocket.DefaultDialer.Dial(websocketServer, nil)
	if err != nil {
		log.Printf("Fail to dial: %v", err)
		return nil
	}

	return &RobotDetect{websocketServer,Message.NewWsConnection(conn)}
}


func (r *RobotDetect) Start(interval int,msg *Message.RobotMsg) {
	go r.wsConn.Heartbeat(interval,msg)
	go r.wsConn.ProcLoop(func(msg *Message.RobotMsg) {
		log.Printf("client receive %v \n",msg)
	})
	go r.wsConn.WsReadLoop()
	go r.wsConn.WsWriteLoop()
}
