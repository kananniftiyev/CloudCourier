package database

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitializeFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("C:\\Users\\kenan\\Documents\\GitHub\\CloudShareX\\backend\\firebase-x.json") // Replace with your Firebase Admin SDK credentials file

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
		return nil, err
	}

	return app, nil
}
