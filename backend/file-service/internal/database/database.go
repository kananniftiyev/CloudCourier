package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDBInstance *mongo.Database
var mongoDBLock sync.Mutex

func ConnectToMongoDB() *mongo.Database {
	if mongoDBInstance == nil {
		mongoDBLock.Lock()
		defer mongoDBLock.Unlock()
		port := os.Getenv("MONGO_DB_PORT")
		username := os.Getenv("MONGO_DB_USERNAME")
		password := os.Getenv("MONGO_DB_PASWORD")
		host := os.Getenv("MONGO_DB_HOST")
		if username == "" || password == "" || host == "" || port == "" {
			log.Fatal("One or more Mongo DB environment variables are empty or missing.")
		}

		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?ssl=true&retrywrites=false", username, password, host, port)

		client, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		mongoDBInstance = client.Database(os.Getenv("MONGO_DB_DATABASE_NAME"))
	}

	return mongoDBInstance
}
