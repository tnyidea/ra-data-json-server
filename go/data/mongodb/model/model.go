package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

func NewMongoDbSession(url string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	tokens := strings.Split(url, "/")
	dbName := tokens[len(tokens)-1]

	return client.Database(dbName), nil
}
