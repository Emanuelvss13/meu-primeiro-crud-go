package main

import (
	"context"
	"log"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/controller"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/controller/routes"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start application")

	ctx := context.Background()

	mongodb.NewMongoDBConnection(ctx)

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserController(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
