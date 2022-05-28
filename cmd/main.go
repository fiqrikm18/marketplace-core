package main

import (
	"github.com/fiqrikm18/marketplace/core_services/pkg/route"
	"log"

	"github.com/fiqrikm18/marketplace/core_services/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//* loading env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	appConfig := config.AppConfig{}
	appConfig.SetupConfiguration()

	//* configure database
	dbConf := config.DatabaseConfig{}
	err = dbConf.NewConnection()
	if err != nil {
		panic(err)
	}
	dbConf.Migrate()

	//* configure server
	srv := gin.Default()
	route.RegisterRouter(srv)

	err = srv.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
