package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	DB "tkBot/database/mongo"
	AnnSpider "tkBot/strategy/annSpider"
	Utils "tkBot/utils"

)

func CreateRobot(c *gin.Context){
	var(
		strategy *DB.Strategy
		err error
	)
	err=c.BindJSON(&strategy)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else {
		if strategy.StrategyName=="annSpider"{
			taskPara:=&AnnSpider.Parameters{}
			Utils.Map2Struct(strategy.Para,&taskPara,"json")
			task:=AnnSpider.NewStrategy(taskPara)
			go task.Run()
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

}