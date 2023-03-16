package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/repository/entity"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *Error.RestError) {
	logger.Info("Init findUseByID repository")

	collectionName := os.Getenv(MONGODB_COLLECTION)

	collection := ur.db.Collection(collectionName)

	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectID}}

	userEntity := &entity.UserEntity{}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			message := fmt.Sprintf("User not found with this id %v", id)
			return nil, Error.NewNotFoundError(message)
		}

		message := "Error trying to find user"
		return nil, Error.NewInternalServerError(message)

	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *Error.RestError) {
	logger.Info("Init findUseByEmail repository")

	collectionName := os.Getenv(MONGODB_COLLECTION)

	collection := ur.db.Collection(collectionName)

	filter := bson.D{{Key: "email", Value: email}}

	userEntity := &entity.UserEntity{}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			message := fmt.Sprintf("User not found with this email %v", email)
			return nil, Error.NewNotFoundError(message)
		}

		message := "Error trying to find user"
		return nil, Error.NewInternalServerError(message)

	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}
