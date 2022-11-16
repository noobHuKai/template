package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
	systemReq "github.com/noobHuKai/app/model/system/request"
	systemRes "github.com/noobHuKai/app/model/system/response"
	"github.com/noobHuKai/app/utils"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var req systemReq.LoginReq
	if err := c.BindJSON(&req); err != nil {
		errMsg := "Error Request Content"
		g.Logger.Error(errMsg)
		response.CreateFailUnprocessableEntityWithMsg(c, errMsg)
		return
	}
	userInter, err := userService.Login(req.Username, req.Password)
	if err != nil {
		g.Logger.Error(err.Error())
		response.CreateFailBadRequestWithMsg(c, err.Error())
		return
	}
	// generate token
	token := utils.SHA256Encrypt(uuid.NewV4().Bytes())
	// set redis
	g.RDB.Set(ctx, token, userInter.ID, g.TimeExpireToken)
	// response
	res := systemRes.LoginResponse{Token: token, Uid: strconv.Itoa(int(userInter.ID)), ExpiresAt: g.TimeExpireToken.String()}
	response.CreatedOkWithData(c, res)
}
