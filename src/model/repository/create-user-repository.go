package repository

import (
	"context"
	"os"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
)

var (
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *Error.RestError) {
	logger.Info("Init createUser repository")

	collectionName := os.Getenv(MONGODB_COLLECTION)

	collection := ur.db.Collection(collectionName)

	value, err := userDomain.GetJSONValue()

	if err != nil {
		return nil, Error.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, Error.NewInternalServerError(err.Error())
	}

	userDomain.ChangeID(result.InsertedID.(string))

	return userDomain, nil
}
