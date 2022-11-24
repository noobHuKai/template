package system

import (
	"context"
	"github.com/noobHuKai/app/service"
)

type ApiGroup struct {
	BaseApi
	UserApi
	CommonApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	baseService = service.ServiceGroupApp.SystemServiceGroup.BasicService
)

var (
	ctx = context.Background()
)
