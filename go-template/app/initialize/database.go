package initialize

import (
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/system"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := g.Cfg.Database.Dsn()
	if g.Cfg.Database.Type == "postgres" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			g.Logger.Fatal(err.Error())
		}
		g.DB = db
	}

	migrate()
}

// Migrate the schema
func migrate() {
	err := g.DB.AutoMigrate(
		&system.SysUser{},
	)

	if err != nil {
		g.Logger.Fatal(err.Error())
	}
}
