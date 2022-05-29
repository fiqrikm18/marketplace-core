package main

import (
	"github.com/fiqrikm18/marketplace/core_services/pkg/route"
	"log"

	"github.com/fiqrikm18/marketplace/core_services/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	//* configure database
	dbConf, err := config.NewConnection()
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
