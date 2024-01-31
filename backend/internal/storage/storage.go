package storage

import (
	"backend/utils"
	"context"
	"firebase.google.com/go/storage"
	"github.com/robfig/cron/v3"
	"google.golang.org/api/iterator"
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

const (
	smtpHost       = "your-smtp-host"
	smtpPort       = 587
	firebaseBucket = "cloudsharex-b8353.appspot.com"
	emailFrom      = "cloud"
	smtpUsername   = "smtp-Username"
	smtpPassword   = "smtp-password"
)

func StartStorageCheck() {
	app, err := utils.InitializeFirebase()
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
		expiryDateString, ok := objAttrs.Metadata["expiry_date"]
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

func sendEmail(deletedFileName string, emailTo string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "File Deleted Notification")
	m.SetBody("text/plain", "The following file has been deleted: "+deletedFileName)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Println("Email notification sent successfully")
	return nil
}
