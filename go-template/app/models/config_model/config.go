package config_model

import "fmt"

type Config struct {
	Mode     string   `mapstructure:"mode"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Server struct {
	Port   int    `mapstructure:"port"`
	Secret string `mapstructure:"secret"`
}

func (s Server) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}
