package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/api/v1"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		// 获取用户信息
		router.GET("info", userApi.GetUserInfo)
		// 获取用户路由
		router.GET("routes", userApi.GetUserRoutes)
	}
}
