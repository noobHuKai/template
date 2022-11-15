package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/middleware"
	"github.com/noobHuKai/app/router"
)

func initRouter() *gin.Engine {
	// init
	r := gin.New()

	// middleware
	// log
	r.Use(middleware.LoggerConfigMiddleware())
	// cors
	r.Use(middleware.CORSMiddleware())

	// router
	systemRouter := router.RouterGroupApp.System

	publicGroup := r.Group("")
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	privateGroup := r.Group("")
	{
		// user
		systemRouter.InitUserRouter(privateGroup)
	}

	return r
}
