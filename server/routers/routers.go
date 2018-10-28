package routers

import (
	"github.com/gin-gonic/gin"
	Controllers "tkBot/server/controllers"
	WebSocket "tkBot/server/websocket"
	Config "tkBot/config"
	Middleware "tkBot/server/middleware"

)







func InitRouter() *gin.Engine {
	if Config.ProductionEnv{
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(Middleware.CorsMiddleware())
	v1 := router.Group("/api/v1")
	{
		user:=v1.Group("/user")
		{
			user.GET("/",nil)
			user.POST("/register",Controllers.Register)
			user.POST("/login",Controllers.Login)

		}

		// out for public
		v1.GET("/announcement/", Controllers.GetAnnouncements)
		announcement:=v1.Group("/announcement")
		announcement.Use(Middleware.AUTH())
		{
			announcement.POST("/", Controllers.CreateAnnouncement)
			announcement.GET("/:id", nil)
			announcement.PUT("/:id", nil)
			announcement.DELETE("/:id", nil)
		}

		robot:=v1.Group("/robot")
		robot.Use(Middleware.AUTH())
		{

			robot.POST("/", Controllers.CreateRobot)
			robot.GET("/", Controllers.GetRobots)
			robot.GET("/:id", nil)
			robot.PUT("/", Controllers.StopRobot)
			robot.DELETE("/", Controllers.DeleteRobot)

		}

		strategy:=v1.Group("/strategy")
		strategy.Use(Middleware.AUTH())
		{

			strategy.POST("/", nil)
			strategy.GET("/", Controllers.GetStrategies)
			strategy.GET("/:id", nil)
			strategy.PUT("/:id", nil)
			strategy.DELETE("/:id", nil)

		}

		asset:=v1.Group("/asset")
		asset.Use(Middleware.AUTH())
		{
			asset.GET("/",Controllers.GetAsset)
			asset.POST("/",Controllers.CreatAsset)
		}

		order:=v1.Group("/order")
		order.Use(Middleware.AUTH())
		{
			order.GET("/",nil)
			order.POST("/",Controllers.CreatOrder)
		}

		loginfo:=v1.Group("/loginfo")
		loginfo.Use(Middleware.AUTH())
		{
			loginfo.GET("/",Controllers.GetLogInfo)
			loginfo.POST("/",Controllers.CreatLogInfo)
		}


		test:=v1.Group("/test")
		test.Use(Middleware.AUTH())
		{

			test.POST("/", Controllers.TestApi)


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

