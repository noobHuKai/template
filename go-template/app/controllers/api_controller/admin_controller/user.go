package admin_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/models/request_model"
	response "github.com/noobHuKai/app/models/response_model"
)

// Login godoc
// @Summary Login
// @Description Administrator Login
// @Tags admin
// @Accept json
// @Produce json
//@Param    user  body      request_model.AdminUserLoginReq  true  "login"
// @Success 200 {string} aaa
// @Router /api/admin/login [post]
func Login(c *gin.Context) {
	var req request_model.AdminUserLoginReq
	if c.BindJSON(&req) == nil {
		fmt.Println(req.Username)
		fmt.Println(req.Password)
	}
	response.Ok(c)
}
