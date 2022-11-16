package system

import (
	"context"
	"github.com/noobHuKai/app/service"
)

type ApiGroup struct {
	BaseApi
	UserApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)

var (
	ctx = context.Background()
)
