package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/routers"
)

func RunServer() {
	r := routers.InitRouter()

	if g.Cfg.Mode == "release" {
		gin.SetMode(gin.DebugMode)
	} else if g.Cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Run(g.Cfg.Server.Addr())
}
