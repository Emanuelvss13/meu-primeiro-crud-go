package mongodb

import (
	"context"
	"os"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGODBURL = "MONGODB_URL"
var MONGODBDATABASE = "MONGODB_DATABASE"

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	mongo_uri := os.Getenv(MONGODBURL)
	mongo_database := os.Getenv(MONGODBDATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connection with mongo successfuly!")

	return client.Database(mongo_database), nil
}
