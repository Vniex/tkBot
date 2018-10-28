package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	log "github.com/sirupsen/logrus"
	Mongo "tkBot/database/mongo"
	Global "tkBot/global"
	"strconv"
	"tkBot/utils"

)

type AssetResp struct {
	Asset float64 `json:"asset"`
	Timestamp int `json:"timestamp"`
}

func CreatAsset(c *gin.Context){

	var (
		asset *Mongo.Asset
		err error
		code int
		exist,
		tmpExist bool
		robot_name ,
		net_asset,
		time_stamp string
	)
	robot_name,exist=c.GetPostForm("robot_name")
	net_asset,tmpExist=c.GetPostForm("net_asset")
	exist=exist&&tmpExist
	time_stamp,tmpExist=c.GetPostForm("time_stamp")
	exist=exist&&tmpExist
	if !exist{
		code=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
	}


	asset=Mongo.NewAsset(robot_name,utils.ToFloat64(net_asset),utils.ToInt(time_stamp))


	db:=Mongo.GetAssetDB()
	if err=db.Insert(asset);err!=nil{
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


func GetAsset(c *gin.Context){
	var (
		err error
	)



	db:=Mongo.GetAssetDB()
	robot_name := c.DefaultQuery("robot_name","")
	time_stamp,err:= strconv.Atoi( c.DefaultQuery("time_stamp","0"))
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
	res,err:=db.FindAssetsInTimeByRobot(robot_name,time_stamp)
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
