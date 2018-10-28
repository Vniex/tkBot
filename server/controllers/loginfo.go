package controllers

import (
	"github.com/gin-gonic/gin"
	Mongo "tkBot/database/mongo"
	"net/http"
	log "github.com/sirupsen/logrus"
	Global "tkBot/global"
	"tkBot/utils"
	"strconv"
)

func CreatLogInfo(c *gin.Context){

	var (
		logInfo *Mongo.LogInfo
		err error
		code int
		exist,
		tmpExist bool
		robot_name ,
		level,
		msg,
		time_stamp string
	)
	robot_name,exist=c.GetPostForm("robot_name")
	level,tmpExist=c.GetPostForm("level")
	exist=exist&&tmpExist
	time_stamp,tmpExist=c.GetPostForm("time_stamp")
	exist=exist&&tmpExist
	msg,tmpExist=c.GetPostForm("msg")
	exist=exist&&tmpExist
	if !exist{
		code=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
	}


	logInfo=Mongo.NewLogInfo(robot_name,utils.ToInt(time_stamp),utils.ToInt(level),msg)


	db:=Mongo.GetLogInfoDB()
	if err=db.Insert(logInfo);err!=nil{
		log.Error(err)
		code:=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
		return
	}
	code=Global.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : Global.GetMsg(code),
		"data" : nil,
	})

}

func GetLogInfo(c *gin.Context){
	var (
		err error
	)



	db:=Mongo.GetLogInfoDB()
	robot_name := c.DefaultQuery("robot_name","")
	limit,err:= strconv.Atoi( c.DefaultQuery("limit","20"))
	if err!=nil{
		log.Error(err)
		code:=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
		return
	}
	res,err:=db.FindlastestLogInfoByRobot(robot_name,utils.ToInt(limit))
	if err!=nil{
		log.Error(err)
		code:=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
		return
	}

	code:=Global.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : Global.GetMsg(code),
		"data" : res,
	})
}
