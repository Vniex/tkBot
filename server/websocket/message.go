package websocket

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)
const(
	CmdType_HeartBeat = iota
	CmdType_Kill
	CmdType_Start
	CmdType_Register
	CmdType_PushData
	CmdType_Log

)


type RobotHubMsg struct {
	RobotHubName string `json:"robot_hub_name"`
	Cmd   int      `json:"cmd"`
	Data  string `json:"data"`
}


func NewRobotMsg(robot_name string,cmd int,data string) *RobotHubMsg{
	return &RobotHubMsg{robot_name,cmd,data}
}

func (r *RobotHubMsg)ToBytes() ([]byte,error){
	msg, err := json.Marshal(r)
	if err != nil {
		log.Error("Fail to package robotMsg :%v", err)
		return nil,err
	}
	return msg,nil
}



func ParseRobotMsg(message []byte) *RobotHubMsg {
	var data RobotHubMsg
	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Error("Fail to parse message:%v", err)
		return nil
	}
	return &data
}

func PackageRobotMsg(robot_name string,cmd int,   data string) []byte {
	var req = RobotHubMsg{
		RobotHubName:robot_name,
		Cmd:   cmd,
		Data:  data,
	}

	msg, err := json.Marshal(req)
	if err != nil {
		log.Error("Fail to packageResponseMsg:%v", err)
		return nil
	}
	return msg
}
