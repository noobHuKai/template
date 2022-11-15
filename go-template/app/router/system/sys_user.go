package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/api/v1"
	"github.com/noobHuKai/app/middleware"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	userRouter := router.Group("user")
	//  public router
	{
		userRouter.POST("/admin/login", userApi.AdminLogin)
	}
	// private router
	{
		userRouter.Use(middleware.TokenAuthorizeMiddleware())
	}
}
