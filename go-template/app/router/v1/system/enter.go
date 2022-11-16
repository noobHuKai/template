package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/middleware"
)

type RouterGroup struct {
	BaseRouter
	UserRouter
}

func (r RouterGroup) InitSystemRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("base")
	{
		r.InitBaseRouter(baseRouter)
	}

	userRouter := router.Group("users")
	userRouter.Use(middleware.TokenAuthorizeMiddleware())
	{
		r.InitUserRouter(userRouter)
	}
}
