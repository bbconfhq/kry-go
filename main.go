package main

import (
	"kry-go/config"
	"kry-go/config/db"
	"kry-go/routes"
	"kry-go/server"
	"log"
)

// swag init --parseDependency --parseInternal

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
	cfg := config.Load()
	app := server.Init(cfg)

	routes.InitRouter(app)
	db.MigrateDatabase(app.DB)

	if err := app.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}
