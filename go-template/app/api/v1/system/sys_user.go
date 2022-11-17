package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
	"strconv"
)

type UserApi struct{}

func (u *UserApi) GetUserInfo(c *gin.Context) {
	uidStr, exists := c.Get("uid")
	if !exists {
		response.FailWithMsg(c, "获取用户失败")
		return
	}
	uid, err := strconv.ParseUint(uidStr.(string), 10, 64)
	if err != nil {
		g.Logger.Info(err.Error())
		response.FailWithMsg(c, err.Error())
		return
	}
	userInter, err := userService.GetUserInfo(uid)
	if err != nil {
		g.Logger.Error(err.Error())
		response.FailWithMsg(c, err.Error())
		return
	}
	response.OkWithData(c, userInter)
}
