package main

import (
	"context"
	"log"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start application")

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	database, err := mongodb.NewMongoDBConnection(ctx)

	if err != nil {
		log.Fatalf(
			"Error trying to connect to MongoDB, error: %s", err.Error(),
		)
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
