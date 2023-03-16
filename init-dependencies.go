package main

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/controller"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/repository"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repository := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repository)
	return controller.NewUserController(service)

}
