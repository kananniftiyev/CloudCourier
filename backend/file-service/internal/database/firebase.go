package database

import (
	"context"
	"log"
	"os"
	"sync"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firebaseAppInstance *firebase.App
var firebaseAppLock sync.Mutex

// TODO: Move firebase to Shared-lib
func InitializeFirebase() (*firebase.App, error) {
	if firebaseAppInstance == nil {
		firebaseAppLock.Lock()
		defer firebaseAppLock.Unlock()
		opt := option.WithCredentialsFile(os.Getenv("FIREBASE_JSON"))

		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("Error initializing Firebase: %v", err)
			return nil, err
		}
		firebaseAppInstance = app
	}

	return firebaseAppInstance, nil
}
