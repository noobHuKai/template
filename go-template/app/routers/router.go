package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/middlewares"
)

func InitRouter() *gin.Engine {
	// init
	r := gin.New()

	// middleware
	r.Use(middlewares.LoggerConfigMiddleware())

	apiGroup := r.Group("/api")
	{
		// admin
		adminGroup := apiGroup.Group("/admin")
		initAdminRouter(adminGroup)
	}

	return r
}
