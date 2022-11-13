package g

import (
	"github.com/noobHuKai/internal/models/config_model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Cfg       *config_model.Config
	Logger    *zap.Logger
	JWTSecret []byte
	DB        *gorm.DB
)
