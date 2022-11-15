package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
	systemReq "github.com/noobHuKai/app/model/system/request"
	systemRes "github.com/noobHuKai/app/model/system/response"

	uuid "github.com/satori/go.uuid"
	"time"
)

type UserApi struct{}

func (u *UserApi) AdminLogin(c *gin.Context) {
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
	// generate token
	token := uuid.NewV4().String()
	// set redis
	g.RDB.Set(ctx, token, userInter.ID, time.Hour*7)
	// response
	res := systemRes.LoginResponse{Token: token}
	response.OkWithData(c, res)
}
