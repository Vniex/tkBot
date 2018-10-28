package global

import (
	Mongo "tkBot/database/mongo"

	log "github.com/sirupsen/logrus"
	"os"
	Config "tkBot/config"
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
		log.Error(err)
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

func InitLog(){
	//MyLogger=logrus.New()
	if Config.ProductionEnv == true {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
	file, err := os.OpenFile(Config.LogFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetLevel(log.InfoLevel)
}