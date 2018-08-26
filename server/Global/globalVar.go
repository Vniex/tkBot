package Global

import (
	Mongo "tkBot/database/mongo"
	"log"

)


var Status=make(map[string]*RobotHubStatus)


type RobotStatus struct {
	RobotName string `json:"robot_name"`
	LastLogin int64 `json:"last_login"`
	Para map[string]interface{} `json:"para"`
	
}

func NewRobotStatus(name string,para map[string]interface{}) *RobotStatus{
	return &RobotStatus{
		name,
		0,
		para,
	}
}

type RobotHubStatus struct {
	HubName string `json:"hub_name"`
	LastLogin int64 `json:"last_login"`
	Robot map[string]*RobotStatus `json:"robot"`
}

func NewRobotHubStatus(name string) *RobotHubStatus{
	return &RobotHubStatus{
		name,
		0,
		make(map[string]*RobotStatus),
	}
}





func InitGlobalVar(){
	strategies,err:=Mongo.GetStrategyDB().FindStrategies()
	if err!=nil{
		log.Println(err)
	}
	for _,strategy:=range strategies{
		Status[strategy.StrategyName]=NewRobotHubStatus(strategy.StrategyName)
	}
}

func AddRobotHub(strategy *Mongo.Strategy){
	if Status[strategy.StrategyName]==nil{
		Status[strategy.StrategyName]=NewRobotHubStatus(strategy.StrategyName)
	}

}

func DeleteRobotHub(strategy *Mongo.Strategy){
	panic("not implement")
}