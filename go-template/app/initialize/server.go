package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
)

func RunServer() {
	r := initRouter()

	if g.Cfg.Mode == "release" {
		gin.SetMode(gin.DebugMode)
	} else if g.Cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Run(g.Cfg.Server.Addr())
}
