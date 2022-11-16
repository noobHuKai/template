package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/router/v1/system"
)

type RouterGroup struct {
	System system.RouterGroup
}


func (r *RouterGroup) InitV1Router(router *gin.RouterGroup) {
	systemRouterGroup := router.Group("system")
	{
		r.System.InitSystemRouter(systemRouterGroup)
	}
}