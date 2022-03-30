package main

import (
	"kry-go/base"
	"kry-go/routes"
	"log"
)

// @title           kry-go
// @version         0.1
// @description     kry(Kim-gi-dong, Ryo, Yun) is simple problem-solving contest server written in Go.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Ryo
// @contact.url    https://github.com/gwanryo/kry-go
// @contact.email  gwanryo@gmail.com

// @license.name  MIT
// @license.url   https://github.com/gwanryo/kry-go/blob/main/LICENSE

// @host
// @BasePath
func main() {
	cfg := base.Config{}
	cfg.Load()

	app := base.Server{}
	app.Init(&cfg)

	routes.InitRouter(&app)

	if err := app.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}
