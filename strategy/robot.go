package strategy

import (



	Utils "tkBot/utils"

	WebSocket "tkBot/server/websocket"

)



const websocketServer = "ws://localhost:8888/api/v1/ws/robot"
const Heartbeat_Interval=2

type RobotInstance interface {
	 Run()
}



type Robot struct {
	RobotDetect *Utils.RobotDetect


}

func NewRobot() *Robot{
	return &Robot{
		Utils.NewRobotDetect(websocketServer),

	}
}

func (r *Robot)LiveDetect(msg *WebSocket.RobotMsg){
	r.RobotDetect.Start(Heartbeat_Interval,msg)
}
