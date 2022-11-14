package boot

import (
	"github.com/noobHuKai/app/constants"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/models/db_model"
	"github.com/noobHuKai/app/utils"
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
	initDatabaseData()
}

// 初始化数据库数据
func initDatabaseData() {
	root := db_model.User{
		Username: "root",
		Password: utils.MD5Encrypt("admin"),
		Nickname: "root",
		Age:      18,
		Status:   0,
		Role:     "super",
	}
	admin := db_model.User{
		Username: "admin",
		Password: utils.MD5Encrypt("admin"),
		Nickname: "admin",
		Age:      18,
		Status:   0,
		Role:     "admin",
	}
	noob := db_model.User{
		Username: "noob",
		Password: utils.MD5Encrypt("admin"),
		Nickname: "admin",
		Age:      18,
		Status:   0,
		Role:     "user",
	}
	g.DB.FirstOrCreate(&root)
	g.DB.FirstOrCreate(&admin)
	g.DB.FirstOrCreate(&noob)
}
