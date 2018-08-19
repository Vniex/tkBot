package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	Controllers "tkBot/server/controllers"
	WebSocket "tkBot/server/websocket"

)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		isAccess:=true
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}



func InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(CorsMiddleware())
	v1 := router.Group("/api/v1/announcement")
	{
		v1.POST("/", Controllers.CreateAnnouncement)
		v1.GET("/", Controllers.GetAnnouncements)
		v1.GET("/:id", nil)
		v1.PUT("/:id", nil)
		v1.DELETE("/:id", nil)
	}


	router.GET("/robotws", func(c *gin.Context) {
		WebSocket.WsHandlerServer(c.Writer, c.Request)
	})

	return router

}

