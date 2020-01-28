package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redmeros/htrade/web/controllers"
	"github.com/redmeros/htrade/web/controllers/data"
	"github.com/redmeros/htrade/web/middlewares"
)

func setRouting(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.POST("/signup", controllers.SignUp)
		v1.POST("/login", controllers.Login)

	}
	authorized := router.Group("api/v1")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/welcome", controllers.Welcome)
		authorized.POST("/logout", controllers.Logout)
		authorized.POST("/refresh", controllers.Refresh)

		dataCollector := authorized.Group("/data_collector")
		{
			dataCollector.GET("/", controllers.CollectorStatus)
			dataCollector.POST("/", controllers.CollectorStart)
			dataCollector.DELETE("/", controllers.CollectorStop)
		}

		pairs := authorized.Group("/pairs")
		{
			pairs.GET("/:name", data.GetPairByName)
			pairs.DELETE("/:id", data.DeletePair)
			pairs.GET("/", data.GetPairs)
			pairs.POST("/", data.NewPair)
		}
	}
}
