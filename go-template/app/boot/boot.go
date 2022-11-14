package boot

func InitServer() {
	initConfig()
	initLogger()
	initDatabase()
	initRedis()
	initCron()
}
