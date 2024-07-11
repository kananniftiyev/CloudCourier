package test

import (
	storage "backend/storage-service/internal"
	"testing"

	"github.com/kananniftiyev/cloudcourier-lib/shared"
)

func TestInitializeFirebase(t *testing.T) {
	// Set the environment variable for the Firebase JSON path
	shared.LoadEnv()

	app, err := storage.InitializeFirebase()
	if err != nil {
		t.Fatalf("Failed to initialize Firebase: %v", err)
	}

	if app == nil {
		t.Fatalf("Expected a Firebase app instance, got nil")
	}

}
