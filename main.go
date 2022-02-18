package main

import (
	"kry-go/config"
	"kry-go/server"
	"kry-go/server/routes"
	"kry-go/utils"
	"log"
)

// @title kry-go
// @version 0.1
// @description kry(Kim-gi-dong, Ryo, Yun) is simple problem-solving contest server written in Go.
// @termsOfService http://swagger.io/terms/

// @contact.name Ryo
// @contact.url https://github.com/gwanryo/kry-go
// @contact.email gwanryo@gmail.com

// @license.name MIT
// @license.url

// @host
// @BasePath
// swag init --parseDependency --parseInternal
func main() {
	cfg := config.Load()
	app := server.Init(cfg)

	routes.InitRouter(app)
	utils.MigrateDatabase(app.DB)

	if err := app.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
