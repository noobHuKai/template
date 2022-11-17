package initialize

import (
	"github.com/noobHuKai/app/g"
	"github.com/spf13/viper"
	"log"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %w", err)
	}

	err = viper.Unmarshal(&g.Cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	g.JWTSecret = []byte(g.Cfg.Server.Secret)
}

func initWebRoutes()  {
	viper.SetConfigName("routes")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error router file: %w", err)
	}
	err = viper.Unmarshal(&g.WebRouter)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
