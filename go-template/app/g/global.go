package g

import (
	"github.com/go-redis/redis/v9"
	"github.com/noobHuKai/app/model/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Cfg       *config.Config
	Logger    *zap.Logger
	JWTSecret []byte
	DB        *gorm.DB
	RDB       *redis.Client
)
