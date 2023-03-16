package repository

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MONGODB_COLLECTION = "MONGODB_USER_COLLECTION"
)

type UserRepositoryInterface interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *Error.RestError)
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *Error.RestError)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *Error.RestError)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepositoryInterface {
	return &userRepository{db: db}
}
