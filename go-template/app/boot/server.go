package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/routers"

	docs "github.com/noobHuKai/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() {
	r := routers.InitRouter()

	if g.Cfg.Mode == "release" {
		gin.SetMode(gin.DebugMode)
	} else if g.Cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// swag
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(g.Cfg.Server.Addr())
}
