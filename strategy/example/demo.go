package example

import (

	Utils "tkBot/utils"

	"time"
	"log"
)

const websocketServer = "ws://localhost:8888/robotws"



type Strategy struct {
	robotDetect *Utils.RobotDetect

}

func NewStrategy() *Strategy{
	return &Strategy{
		Utils.NewRobotDetect(websocketServer),
		}
}

func (s Strategy) Run(){
	for{
		log.Println("buy 100")
		time.Sleep(5*time.Second)
	}
}

func (s Strategy)Start(){
	go s.robotDetect.Start()
	s.Run()
}
