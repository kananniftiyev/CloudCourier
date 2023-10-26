package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectToMongoDB() *mongo.Database {
	uri := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("CloudShareX")
}
