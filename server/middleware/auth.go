package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	log "github.com/sirupsen/logrus"
	Utils "tkBot/utils"
	Global "tkBot/global"
	Mongo "tkBot/database/mongo"
	"strconv"
)




func AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = Global.SUCCESS

		token,exist := c.GetQuery("token")
		if exist { // for jwt auth
			if token == "" {
				code = Global.ERROR_PARAMETER
			} else {
				claims, err := Utils.ParseToken(token)
				if err != nil {
					code = Global.ERROR_AUTH_CHECK_TOKEN_FAIL
				} else if time.Now().Unix() > claims.ExpiresAt {
					code = Global.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}

			if code != Global.SUCCESS {
				c.JSON(http.StatusOK, gin.H{
					"code":    code,
					"message": Global.GetMsg(code),
					"data":    data,
				})

				c.Abort()
				return
			}

			c.Next()
		}else{ // for apikey secretkey auth
			var(
				err error
				exist bool
				tmp_exist bool
				para_time string
				para_time_int64 int64
				para_sign string
				para_apikey string
				user *Mongo.User
			)
			para_time,exist=c.GetPostForm("timestamp")

			para_sign,tmp_exist=c.GetPostForm("sign")
			exist=exist && tmp_exist

			para_apikey,tmp_exist=c.GetPostForm("apikey")
			exist=exist && tmp_exist

			if para_time_int64,err=strconv.ParseInt(para_time, 10, 64); !exist || err!=nil||
				((time.Now().Unix()-para_time_int64)>10){
					if !exist || err!=nil{
						c.JSON(http.StatusOK, gin.H{
							"code":    Global.ERROR_PARAMETER,
							"message": Global.GetMsg(Global.ERROR_PARAMETER),
							"data":    data,
						})
					}else{
						c.JSON(http.StatusOK, gin.H{
							"code":    Global.ERROR_TIMESTAMP,
							"message": Global.GetMsg(Global.ERROR_TIMESTAMP),
							"data":    data,
						})
					}

				c.Abort()
				return

			}


			user,err=Mongo.GetUserDB().FindUserById(para_apikey)
			if err!=nil{
				log.Error(err)
				c.JSON(http.StatusOK, gin.H{
					"code":    Global.ERROR_PARAMETER,
					"message": Global.GetMsg(Global.ERROR_PARAMETER),
					"data":    data,
				})
				c.Abort()
				return

			}
			if !(Utils.GetSHA256(para_apikey+user.SecretKey+para_time)==para_sign){
				c.JSON(http.StatusOK, gin.H{
					"code":    Global.ERROR_PARAMETER,
					"message": Global.GetMsg(Global.ERROR_PARAMETER),
					"data":    data,
				})

				c.Abort()
				return
			}
			c.Next()

		}
	}
}




