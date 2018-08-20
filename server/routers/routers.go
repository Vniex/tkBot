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
	v1 := router.Group("/api/v1")
	{
		announcement:=v1.Group("/announcement")
		{
			announcement.POST("/", Controllers.CreateAnnouncement)
			announcement.GET("/", Controllers.GetAnnouncements)
			announcement.GET("/:id", nil)
			announcement.PUT("/:id", nil)
			announcement.DELETE("/:id", nil)
		}

		robot:=v1.Group("/robot")
		{

			robot.POST("/", Controllers.CreateRobot)
			robot.GET("/", Controllers.GetRobots)
			robot.GET("/:id", nil)
			robot.PUT("/:id", nil)
			robot.DELETE("/:id", nil)

		}

		strategy:=v1.Group("/strategy")
		{

			strategy.POST("/", Controllers.CreateStrategy)
			strategy.GET("/", Controllers.GetStrategies)
			strategy.GET("/:id", nil)
			strategy.PUT("/:id", nil)
			strategy.DELETE("/:id", nil)

		}

		ws:=v1.Group("/ws")
		{
			ws.GET("/robot", func(c *gin.Context) {
				WebSocket.WsHandlerServer(c.Writer, c.Request)
			})

		}


	}



	return router

}

