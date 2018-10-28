package controllers

import (
	"github.com/gin-gonic/gin"
	Mongo "tkBot/database/mongo"
	"net/http"
	Global "tkBot/global"
	log "github.com/sirupsen/logrus"
	"tkBot/utils"
)

func CreatOrder(c *gin.Context){

	var (
		order *Mongo.Order
		err error
		code int
		exist,
		tmpExist bool
		robotName ,
		price,
		amount,
		avgPrice,
		fee,
		orderId,
		orderTime,
		pair,
		side string

	)
	robotName,exist=c.GetPostForm("robot_name")
	price,tmpExist=c.GetPostForm("price")
	exist=exist&&tmpExist
	amount,tmpExist=c.GetPostForm("amount")
	exist=exist&&tmpExist
	avgPrice,tmpExist=c.GetPostForm("avg_price")
	exist=exist&&tmpExist
	fee,tmpExist=c.GetPostForm("fee")
	exist=exist&&tmpExist
	orderId,tmpExist=c.GetPostForm("order_id")
	exist=exist&&tmpExist
	orderTime,tmpExist=c.GetPostForm("order_time")
	exist=exist&&tmpExist
	pair,tmpExist=c.GetPostForm("pair")
	exist=exist&&tmpExist
	side,tmpExist=c.GetPostForm("side")
	exist=exist&&tmpExist
	if !exist{
		code=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
	}


	order=&Mongo.Order{
		RobotName:robotName,
		Price:utils.ToFloat64(price),
		Amount:utils.ToFloat64(amount),
		AvgPrice:utils.ToFloat64(avgPrice),
		Fee:utils.ToFloat64(fee),
		OrderID:orderId,
		OrderTime:utils.ToInt(orderTime),
		Pair:pair,
		Side:side,
	}


	db:=Mongo.GetOrderDB()
	if err=db.Insert(order);err!=nil{
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

