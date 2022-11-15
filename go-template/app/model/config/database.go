package config

import (
	"fmt"
)

type Database struct {
	Type     string
	Host     string
	User     string
	Password string
	DbName   string
	Port     int
	SslMode  string
	TimeZone string
}

func (d Database) Dsn() string {
	var dsn string
	if d.Type == "postgres" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ",
			d.Host, d.User, d.Password, d.DbName, d.Port)
		if d.SslMode != "" {
			dsn += " sslmode=" + d.SslMode
		}
		if d.TimeZone != "" {
			dsn += " TimeZone=" + d.TimeZone
		}
	}

	return dsn
}
