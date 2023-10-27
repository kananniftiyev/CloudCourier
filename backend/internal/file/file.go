package file_upload

import (
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

const SECRET_KEY = "secret"

type CustomClaims struct {
	jwt.StandardClaims
	UserID   uint
	Username string

	// Add other custom claims as needed
}

func InitializeFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("C:\\Users\\kenan\\Documents\\GitHub\\CloudShareX\\backend\\firebase-x.json") // Replace with your Firebase Admin SDK credentials file

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
		return nil, err
	}

	return app, nil
}

func GetFileURL(bucketName, objectName string) (string, error) {
	app, err := InitializeFirebase()
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
	token, err := jwt.ParseWithClaims(cookie.Value, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		log.Print(err)
		return 0, "", err
	}

	if !token.Valid {
		log.Println(err)
		return 0, "", err
	}
	claims, ok := token.Claims.(*CustomClaims)
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
