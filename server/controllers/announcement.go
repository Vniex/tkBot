package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	DB "tkBot/database/mongo"
	Utils "tkBot/utils"
)

func CreateAnnouncement(c *gin.Context){
	var(
		ann *DB.Announcement
		err error
	)
	err=c.BindJSON(&ann)
	//log.Println("receive ann:",ann)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),
		})
	}else {
		annDB := DB.NewAnnouncementDB()
		err = annDB.Insert(ann)

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
		ann_list []*DB.Announcement
		err error
	)
	exchange_id := Utils.ToInt(c.DefaultQuery("exchange_id","-1"))
	limit:=Utils.ToInt(c.DefaultQuery("limit","300"))
	annDB:=DB.NewAnnouncementDB()
	if exchange_id==-1{ // for all exchange
		ann_list,err=annDB.FindAnns(limit)

	}else{
		ann_list,err=annDB.FindAnnsByExchangeId(exchange_id)
	}
	if ann_list==nil{
		ann_list=make([]*DB.Announcement,0)
	}
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"success":false,"message":err.Error(),"data":ann_list,
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,"message":"成功","data":ann_list,
	})
}

