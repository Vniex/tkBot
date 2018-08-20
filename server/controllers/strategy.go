package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	DB "tkBot/database/mongo"
)

func CreateStrategy(c *gin.Context){
	var(
		strategy *DB.Strategy
		err error
	)
	err=c.BindJSON(&strategy)
	//log.Println("receive ann:",ann)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else {
		strategyDB := DB.NewStrategyDB()
		err = strategyDB.Insert(strategy)

		if err != nil {
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
		strategies []*DB.Strategy
		err error
	)



		strategyDB := DB.NewStrategyDB()
		strategies,err = strategyDB.FindStrategies()
		if strategies ==nil{
			strategies=make([]*DB.Strategy,0)
		}
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false, "message": err.Error(),"data":strategies,
			})

		}else{
			c.JSON(http.StatusOK, gin.H{
				"success": true, "message": "成功","data":strategies,
			})
		}

}