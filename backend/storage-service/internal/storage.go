package storage

import (
	"context"
	"log"
	"os"
	"time"

	"firebase.google.com/go/storage"
	"github.com/kananniftiyev/cloudcourier-lib/shared"
	"github.com/robfig/cron/v3"
	"google.golang.org/api/iterator"
)

func StartStorageCheck() {
	shared.LoadEnv()
	app, err := InitializeFirebase()
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Run the task immediately when the function starts
	log.Println("Checking and deleting expired files...")
	if err := deleteExpiredFiles(client); err != nil {
		log.Printf("Error deleting expired files: %v\n", err)
	}

	c := cron.New()

	// Schedule a job to check and delete expired files every day at midnight
	_, err = c.AddFunc("@hourly", func() {
		log.Println("Checking and deleting expired files...")
		if err := deleteExpiredFiles(client); err != nil {
			log.Printf("Error deleting expired files: %v\n", err)
		}
	})
	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}

	c.Start()

	select {}
}

func deleteExpiredFiles(client *storage.Client) error {
	// Create a context for Firebase Storage operations
	ctx := context.Background()

	// Get a reference to the Firebase Storage bucket
	firebaseBucket := os.Getenv("FIREBASE_BUCKET")
	if firebaseBucket == "" {
		log.Fatal("FIREBASE_BUCKET environment variable is not set")
	}
	bucket, err := client.Bucket(firebaseBucket)
	if err != nil {
		return err
	}

	currentTime := time.Now()

	// Create an ObjectIterator to list objects
	objIter := bucket.Objects(ctx, nil)
	for {
		objAttrs, err := objIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		// Fetch object metadata
		const expiryDateKey = "expiry_date"
		expiryDateString, ok := objAttrs.Metadata[expiryDateKey]
		if !ok {
			continue // Skip objects without an expiry date
		}

		expiryDate, err := time.Parse(time.RFC3339, expiryDateString)
		if err != nil {
			return err
		}

		if currentTime.After(expiryDate) {
			// The file has expired, delete it
			if err := bucket.Object(objAttrs.Name).Delete(ctx); err != nil {
				return err
			}
			log.Printf("Deleted expired file: %s\n", objAttrs.Name)
		}
	}

	return nil
}
