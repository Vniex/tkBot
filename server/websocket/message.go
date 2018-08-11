package websocket

import (
	"encoding/json"
	"log"
)

const (
	MSG_ROBOT  = 1+iota
	MSG_MARKET
	
)

type Message struct {
	MsgType int `json:"msg_type"`
	Data  interface{} `json:"data"`
}

func ParseMsg(message []byte) *Message {
	var data Message
	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Println("Fail to parse message:%v", err)
		return nil
	}
	return &data
}

func PackageMsg(msg_type int,data interface{}) []byte {
	var req = Message{
		MsgType:   msg_type,
		Data:  data,
	}

	msg, err := json.Marshal(req)
	if err != nil {
		log.Println("Fail to packageResponseMsg:%v", err)
		return nil
	}
	return msg
}

