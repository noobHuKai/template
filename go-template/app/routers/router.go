package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/middlewares"
)

func InitRouter() *gin.Engine {
	// init
	r := gin.New()

	// middleware
	// log
	r.Use(middlewares.LoggerConfigMiddleware())
	// cors
	r.Use(middlewares.CORSMiddleware())

	apiGroup := r.Group("/api")
	{
		// admin
		adminGroup := apiGroup.Group("/admin")
		initAdminRouter(adminGroup)
	}

	return r
}
