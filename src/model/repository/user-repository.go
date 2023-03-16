package repository

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *Error.RestError)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepositoryInterface {
	return &userRepository{db: db}
}
