package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/controllers/api_controller/admin_controller"
)

func initAdminRouter(r *gin.RouterGroup) {
	r.POST("/login", admin_controller.Login)
}
