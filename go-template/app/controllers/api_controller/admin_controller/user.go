package admin_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/models/request_model"
	response "github.com/noobHuKai/app/models/response_model"
	"github.com/noobHuKai/app/services/api_service/admin_service"
	"github.com/noobHuKai/app/utils"
	"time"
)

func Login(c *gin.Context) {
	var req request_model.AdminUserLoginReq
	if err := c.BindJSON(&req); err != nil {
		response.FailWithMsg(c, "request format error")
		return
	}

	user, err := admin_service.QueryUserService(req.Username, req.Password)
	if err != nil {
		g.Logger.Error(err.Error())
		response.FailWithMsg(c, "database query error")
		return
	}
	if user == nil {
		response.FailWithMsg(c, "not found user")
		return
	}

	res := &response.AdminUserLoginRes{
		AccessToken:  utils.CreateToken(user.ID, time.Hour*3),
		RefreshToken: utils.CreateToken(user.ID, time.Hour*7),
	}

	response.OkWithData(c, res)
}
