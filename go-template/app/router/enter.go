package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/router/v1"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	V1 v1.RouterGroup
}

func (r *RouterGroup) InitApiRouter(router *gin.RouterGroup) {
	v1RouterGroup := router.Group("v1")
	{
		r.V1.InitV1Router(v1RouterGroup)
	}
}
