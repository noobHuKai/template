package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/api/v1"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(router *gin.RouterGroup) {
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi

	router.GET("basic_info", baseApi.GetSysBasicInfo)
	router.GET("monitor_info", baseApi.GetSysMonitorInfo)
}
