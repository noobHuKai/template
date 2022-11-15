package initialize

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/noobHuKai/app/g"
)

func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     g.Cfg.Redis.Addr(),
		Password: g.Cfg.Redis.Password,
		DB:       g.Cfg.Redis.DB,
	})
	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		g.Logger.Fatal(err.Error())
	}
	if result != "PONG" {
		g.Logger.Fatal("redis ping error")
	}
	g.RDB = rdb
}
