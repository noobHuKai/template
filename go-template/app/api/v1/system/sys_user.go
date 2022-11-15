package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
	systemReq "github.com/noobHuKai/app/model/system/request"
	uuid "github.com/satori/go.uuid"
	"time"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	var req systemReq.LoginReq
	if err := c.BindJSON(&req); err != nil {
		g.Logger.Error(err.Error())
		response.FailWithMessage(c, err.Error())
		return
	}
	userInter, err := userService.Login(req.Username, req.Password)
	if err != nil {
		g.Logger.Error(err.Error())
		response.FailWithMessage(c, err.Error())
		return
	}
	token := uuid.NewV4().String()
	g.RDB.Set(ctx, token, userInter.UUID, time.Hour*7)
	response.Ok(c)
}
