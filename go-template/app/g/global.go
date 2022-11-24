package g

import (
	"github.com/go-redis/redis/v9"
	jsoniter "github.com/json-iterator/go"
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
	WebRouter map[string]interface{}
)

var (
	JsonIter = jsoniter.ConfigFastest
)
