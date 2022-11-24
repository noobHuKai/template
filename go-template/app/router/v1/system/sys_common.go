package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/api/v1"
)

type CommonRouter struct {
}

func (c *CommonRouter) InitCommonRouter(router *gin.RouterGroup) {
	commonApi := v1.ApiGroupApp.SystemApiGroup.CommonApi

	router.POST("login", commonApi.Login)

}
