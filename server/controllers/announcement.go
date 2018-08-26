package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	Mongo "tkBot/database/mongo"
	Utils "tkBot/utils"
)

func CreateAnnouncement(c *gin.Context){
	var(
		ann *Mongo.Announcement
		err error
	)
	err=c.BindJSON(&ann)
	//log.Println("receive ann:",ann)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else {

		err = Mongo.GetAnnouncementDB().Insert(ann)

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

func GetAnnouncements(c *gin.Context){
	var(
		ann_list []*Mongo.Announcement
		err error
	)
	exchange_id := Utils.ToInt(c.DefaultQuery("exchange_id","-1"))
	limit:=Utils.ToInt(c.DefaultQuery("limit","20"))
	last_id:=Utils.ToInt(c.DefaultQuery("last_id","999999999999"))

	ann_list,err=Mongo.GetAnnouncementDB().FindAnns(exchange_id,last_id,limit)


	if ann_list==nil{
		ann_list=make([]*Mongo.Announcement,0)
	}
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),"data":ann_list,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"success":true,"message":"成功","data":ann_list,
		})
	}

}

