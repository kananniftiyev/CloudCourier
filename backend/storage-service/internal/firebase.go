package storage

import (
	"context"
	"log"
	"os"
	"sync"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	firebaseInstance *firebase.App
	firebaseLock     sync.Mutex
)

func InitializeFirebase() (*firebase.App, error) {
	if firebaseInstance == nil {
		firebaseLock.Lock()
		defer firebaseLock.Unlock()
		opt := option.WithCredentialsFile(os.Getenv("FIREBASE_JSON")) // Replace with your Firebase Admin SDK credentials file

		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("Error initializing Firebase: %v", err)
			return nil, err
		}
		firebaseInstance = app
	}

	return firebaseInstance, nil
}
