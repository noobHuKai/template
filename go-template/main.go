package main

import (
	"github.com/noobHuKai/app/boot"
)

// @title           Golang Template API
// @version         1.0
// @description     This is a go template server.

// @contact.name   noobHuKai
// @contact.url    https://github.com/noobHuKai
// @contact.email  hukai2943@aliyun.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8000
// @BasePath  /api
func main() {
	boot.InitServer()
	boot.RunServer()
}
