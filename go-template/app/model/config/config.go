package config

import "fmt"

type Config struct {
	Mode     string   `mapstructure:"mode"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
}

type Server struct {
	Port   int    `mapstructure:"port"`
	Secret string `mapstructure:"secret"`
}

func (s Server) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
