package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Utils "tkBot/utils"
	Global "tkBot/global"
	"fmt"
	log "github.com/sirupsen/logrus"
	Mongo "tkBot/database/mongo"
)


type LoginInfo struct{
	Username string `json:"username" binding:"required"`
	Password string ` json:"password" binding:"required"`
}

func Register(c *gin.Context){
	var (
		loginInfo LoginInfo
		err error
	)

	if err=c.ShouldBindJSON(&loginInfo);err!=nil{
		log.Error(err)
		code:=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
		return
	}

	u:=Mongo.NewUser(loginInfo.Username,loginInfo.Password)
	db:=Mongo.GetUserDB()
	if err=db.Insert(u);err!=nil{
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
		"data" : nil,
	})
}



func Login(c *gin.Context)  {
	var (
		loginInfo LoginInfo
		err error
	)
	err=c.ShouldBindJSON(&loginInfo)
	if err!=nil{
		log.Println(err)
		code:=Global.ERROR_PARAMETER
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : Global.GetMsg(code),
			"data" : nil,
		})
		return

	}
	data := make(map[string]interface{})
	code := Global.ERROR_PARAMETER
	fmt.Println(loginInfo)
	isExist := CheckAuth(loginInfo.Username, loginInfo.Password)
	if isExist {
		token, err := Utils.GenerateToken(loginInfo.Username, loginInfo.Password)
		if err != nil {
			code = Global.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = Global.SUCCESS
		}

	} else {
		code = Global.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	}


	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : Global.GetMsg(code),
		"data" : data,
	})



}


func CheckAuth(username,password string)bool{
	return true
}