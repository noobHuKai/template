package boot

func InitServer() {
	initConfig()
	initLogger()
	initDatabase()
	initCron()
}
