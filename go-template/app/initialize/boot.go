package initialize

func InitServer() {
	initConfig()
	initLogger()
	initDatabase()
	initRedis()
	initCron()
}
