package admin_service

import (
	"errors"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/models/db_model"
	"gorm.io/gorm"
)

// QueryUserService 查询用户是否存在
func QueryUserService(username, password string) (bool, error) {
	var user db_model.User
	result := g.DB.Where(&db_model.User{Username: username, Password: password}).First(&user)
	if result.Error != nil {
		// not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

// InsertUserService 插入用户
func InsertUserService() {

}

// DeleteUserService 删除用户
func DeleteUserService() {

}
