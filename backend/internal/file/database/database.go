package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"path/filepath"
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envFile := filepath.Join(dir, ".env")
	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectToMongoDB() *mongo.Database {
	port := os.Getenv("COSMOS_DB_PORT")
	username := os.Getenv("COSMOS_DB_USERNAME")
	password := os.Getenv("COSMOS_DB_PP")
	host := os.Getenv("COSMOS_DB_HOST")
	if username == "" || password == "" || host == "" || port == "" {
		log.Fatal("One or more Cosmos DB environment variables are empty or missing.")
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

	// Specify the database and collection you want to work with.
	databaseName := "your-database-name"
	return client.Database(databaseName)
}
