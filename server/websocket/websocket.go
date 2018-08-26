package websocket

import (
	"net/http"
	"log"
	"time"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"

	GlobalVar "tkBot/server/Global"
	Mongo "tkBot/database/mongo"
)

var HubWsConn=make(map[string]*WsConnection)

func processHeartBeat(msg *RobotHubMsg){
	if GlobalVar.Status[msg.RobotHubName]==nil{
		GlobalVar.Status[msg.RobotHubName]=&GlobalVar.RobotHubStatus{
			msg.RobotHubName,
			0,
			make(map[string]*GlobalVar.RobotStatus)}
	}
	var robots []string
	json.Unmarshal([]byte(msg.Data),&robots)
	GlobalVar.Status[msg.RobotHubName].LastLogin=time.Now().Unix()
	for _,robot :=range robots{
		if GlobalVar.Status[msg.RobotHubName].Robot[robot]==nil{
			log.Println("robot not in globalvar")
			return
		}
		GlobalVar.Status[msg.RobotHubName].Robot[robot].LastLogin=time.Now().Unix()
	}

}

func processRegister(msg *RobotHubMsg){
	strategy_name:=msg.RobotHubName
	data:=msg.Data
	var datamap map[string]interface{}
	strategy:=&Mongo.Strategy{}
	strategy.Id=bson.NewObjectId()
	strategy.StrategyName=strategy_name
	strategy.Para=make(map[string]interface{})
	json.Unmarshal([]byte(data),&datamap)
	for k,v:=range datamap{
		strategy.Para[k]=v
	}

	if err:=Mongo.GetStrategyDB().Insert(strategy);err!=nil{
		log.Println(err)
	}
	GlobalVar.AddRobotHub(strategy) //todo process error

}

func processPushData(msg *RobotHubMsg){
	strategy_name:=msg.RobotHubName
	switch strategy_name {
	case "公告爬虫":
		data:=msg.Data
		var ann Mongo.Announcement
		if err:=json.Unmarshal([]byte(data),&ann);err!=nil{
			log.Println(err)
			return
		}

		if err:=Mongo.GetAnnouncementDB().Insert(&ann);err!=nil{
			log.Println(err)
		}


	}

}

func processLog(msg *RobotHubMsg){

}

func ProcessRobotMsg(msg *RobotHubMsg){
	//log.Printf("server receive %v \n",msg)
	switch msg.Cmd {
	case CmdType_HeartBeat:
		processHeartBeat(msg)
	case CmdType_Register:
		processRegister(msg)
	case CmdType_Log:
		processLog(msg)
	case CmdType_PushData:
		processPushData(msg)
	}



}

func WsHandlerServer(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	client_name:=req.URL.Query().Get("client_name")
	log.Println(client_name)
	wsConn:=NewWsConnection(client_name,wsSocket)
	HubWsConn[wsConn.clientName]=wsConn

	// 处理器
	go wsConn.ProcLoop(ProcessRobotMsg)
	// 读协程
	go wsConn.WsReadLoop()
	// 写协程
	go wsConn.WsWriteLoop()
}

