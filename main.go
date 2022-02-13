package main

import (
	"kry-go/config"
	"kry-go/server"
	"kry-go/server/routes"
	"kry-go/utils"
	"log"
)

func main() {
	cfg := config.Load()
	app := server.Init(cfg)

	routes.InitRouter(app)
	utils.MigrateDatabase(app.DB)

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
