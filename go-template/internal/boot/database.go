package boot

import (
	"github.com/noobHuKai/internal/constants"
	"github.com/noobHuKai/internal/g"
	"github.com/noobHuKai/internal/models/db_model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := g.Cfg.Database.Dsn()
	if g.Cfg.Database.Type == constants.DBTypePostgres {
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
		&db_model.User{},
	)

	if err != nil {
		g.Logger.Fatal(err.Error())
	}
}
