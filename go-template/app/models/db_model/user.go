package db_model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string
	Password string

	Nickname string
	Age      int
	Status   int
	Role     string
}
