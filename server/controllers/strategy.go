package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	Mongo "tkBot/database/mongo"
	log "github.com/sirupsen/logrus"
)

func CreateStrategy(c *gin.Context){
	var(
		strategy *Mongo.Strategy
		err error
	)
	err=c.BindJSON(&strategy)
	//log.Println("receive ann:",ann)
	if err!=nil{
		log.Error(err)
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else {
		err = Mongo.GetStrategyDB().Insert(strategy)

		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"success": false, "message": err.Error(),
			})

		}else{
			c.JSON(http.StatusOK, gin.H{
				"success": true, "message": "成功",
			})
		}
	}
}


func GetStrategies(c *gin.Context){
	var(
		strategies []*Mongo.Strategy
		err error
	)

		strategies,err =Mongo.GetStrategyDB().FindStrategies()
		if strategies ==nil{
			strategies=make([]*Mongo.Strategy,0)
		}
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"success": false, "message": err.Error(),"data":strategies,
			})

		}else{
			c.JSON(http.StatusOK, gin.H{
				"success": true, "message": "成功","data":strategies,
			})
		}

}