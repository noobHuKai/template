package db_model

import (
	"gorm.io/gorm"
	"time"
)

// Model ID --> UUID
type Model struct {
	UUID      uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
