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

	systemRouterGroup := r.Group("system")

	{
		// user
		systemRouter.InitUserRouter(systemRouterGroup)
	}
	// private router
	{
		//
	}

	return r
}
