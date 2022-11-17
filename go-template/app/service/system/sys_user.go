package system

import (
	"errors"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/system"
	"gorm.io/gorm"
)

type UserService struct {
}

func (userService *UserService) Login(username, password string) (userInter *system.SysUser, err error) {
	var user system.SysUser

	err = g.DB.Where(&system.SysUser{Username: username}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found user")
		}
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("password error")
	}
	return &user, err
}

func (userService *UserService) GetUserInfo(uid uint64) (userInter *system.SysUser, err error) {
	var user system.SysUser

	err = g.DB.First(&user, uid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found user")
		}
		return nil, err
	}
	return &user, err
}
