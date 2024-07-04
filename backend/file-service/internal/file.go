package file_upload

import (
	"backend/file-service/internal/database"
	"context"
	"encoding/base64"

	"github.com/google/uuid"
)


func GetFileURL(bucketName, objectName string) (string, error) {
	app, err := database.InitializeFirebase()
	if err != nil {
		return "", err
	}

	// Create a Firebase Storage client
	storageClient, err := app.Storage(context.Background())
	if err != nil {
		return "", err
	}

	// Get a reference to the Firebase Storage bucket
	bucket, err := storageClient.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	fileRef := bucket.Object(objectName)

	// Get the download URL for the file
	fileURL, err := fileRef.Attrs(context.Background())
	if err != nil {
		return "", err
	}

	// Return the file's download URL
	return fileURL.MediaLink, nil
}

func DecodeUUID(base64UUID string) (uuid.UUID, error) {
	// Decode the Base64 string into bytes
	uuidBytes, err := base64.StdEncoding.DecodeString(base64UUID)
	if err != nil {
		return uuid.Nil, err
	}

	// Create a UUID from the decoded bytes
	u, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return uuid.Nil, err
	}

	return u, nil
}
