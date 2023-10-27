package file_upload

import (
	"backend/utils"
	"context"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func GetFileURL(bucketName, objectName string) (string, error) {
	app, err := utils.InitializeFirebase()
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

func GetUserFromJWT(r *http.Request) (int, string, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		log.Print(err)
		return 0, "", err
	}
	customClaims := &utils.CustomClaims{}

	token, err := jwt.ParseWithClaims(cookie.Value, customClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SECRET_KEY), nil
	})

	if err != nil {
		log.Print(err)
		return 0, "", err
	}

	if !token.Valid {
		log.Println(err)
		return 0, "", err
	}
	claims, ok := token.Claims.(*utils.CustomClaims)
	if !ok {
		log.Printf("Failed to get custom claims from JWT token", err)
		return 0, "", err
	}
	return int(claims.UserID), claims.Username, nil
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
