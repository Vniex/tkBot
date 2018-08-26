package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Mongo "tkBot/database/mongo"
	GlobalVar "tkBot/server/Global"

	WebSocket "tkBot/server/websocket"
	"encoding/json"
	"log"
)

func CreateRobot(c *gin.Context){
	var(
		strategy *Mongo.Strategy
		err error
	)
	err=c.BindJSON(&strategy)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else {
		if WebSocket.HubWsConn[strategy.StrategyName]!=nil{
			p,_:=json.Marshal(strategy.Para)
			msg:=WebSocket.NewRobotMsg("server",WebSocket.CmdType_Start,string(p))
			GlobalVar.Status[strategy.StrategyName].Robot[strategy.Para["robot_name"].(string)]=GlobalVar.NewRobotStatus(
				strategy.Para["robot_name"].(string),strategy.Para,
			)

			if err:=WebSocket.HubWsConn[strategy.StrategyName].WsWrite(msg);err!=nil{
				log.Println(err)
				c.JSON(http.StatusOK,gin.H{
					"success":false,"message":err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK,gin.H{
				"success":true,"message":"成功",
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"success":false,"message":"没有对应策略模板",
			})
		}
	}
}

func GetRobots(c *gin.Context){
	data:=make([]*GlobalVar.RobotHubStatus,0)
	for _,v :=range GlobalVar.Status{
		data=append(data, v)
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,"message":"成功","data":data,
	})
}

func DeleteRobot( c *gin.Context){
	var(
		strategy *Mongo.Strategy
		err error
	)
	err=c.BindJSON(&strategy)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else{
		delete(GlobalVar.Status[strategy.StrategyName].Robot, strategy.Para["robot_name"].(string))
		c.JSON(http.StatusOK,gin.H{
			"success":true,"message":"成功",
		})
	}
}


func StopRobot(c *gin.Context){
	//data:=make([]*GlobalVar.RobotHubStatus,0)
	var(
		strategy *Mongo.Strategy
		err error
	)
	err=c.BindJSON(&strategy)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else {
		if WebSocket.HubWsConn[strategy.StrategyName]!=nil{
			p,_:=json.Marshal([]string{strategy.Para["robot_name"].(string)})

			msg:=WebSocket.NewRobotMsg("server",WebSocket.CmdType_Kill,string(p))
			//GlobalVar.Status[strategy.StrategyName].Robot[strategy.Para["robot_name"].(string)]=GlobalVar.NewRobotStatus(
			//	strategy.Para["robot_name"].(string),strategy.Para,
			//)

			if err:=WebSocket.HubWsConn[strategy.StrategyName].WsWrite(msg);err!=nil{
				log.Println(err)
				c.JSON(http.StatusOK,gin.H{
					"success":false,"message":err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK,gin.H{
				"success":true,"message":"成功",
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"success":false,"message":"没有对应策略模板",
			})
		}
	}
}