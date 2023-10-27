package rest

import (
	file_upload "backend/internal/file"
	"backend/internal/file/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
)

// TODO: Fix this shit
// TODO: Add Password
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	userId, username, err := file_upload.GetUserFromJWT(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	app, err := file_upload.InitializeFirebase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a Firebase Storage client
	storageClient, err := app.Storage(context.Background())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to create Firebase Storage client", http.StatusInternalServerError)
		return
	}

	// Get a reference to the Firebase Storage bucket
	bucket, err := storageClient.Bucket("cloudsharex-b8353.appspot.com")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to get Firebase Storage bucket", http.StatusInternalServerError)
		return
	}

	// Parse the form data to extract the file
	err = r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get a reference to the file in Firebase Storage

	userRef := username + "/"
	fileRef := bucket.Object(userRef + handler.Filename)

	// Create a writer for the Firebase Storage object
	writer := fileRef.NewWriter(context.Background())

	// Copy the uploaded file's content to Firebase Storage
	_, err = io.Copy(writer, file)
	if err != nil {
		http.Error(w, "Unable to upload the file to Firebase Storage", http.StatusInternalServerError)
		return
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		fmt.Println(err)

		http.Error(w, "Unable to close the Firebase Storage writer", http.StatusInternalServerError)
		return
	}

	fileUUID := uuid.New()

	fileURL, err := file_upload.GetFileURL("cloudsharex-b8353.appspot.com", userRef+handler.Filename)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to get file URL", http.StatusInternalServerError)
		return
	}

	// Create New Mongo Record
	fileRepo := database.NewFileRepository(database.ConnectToMongoDB())
	newFileRecord := database.File{
		ID:         primitive.ObjectID{},
		UserID:     userId,
		Username:   username,
		FileName:   handler.Filename,
		FilePath:   fileURL,
		SpecialURL: fileUUID,
		Password:   "",
	}

	err = fileRepo.Create(context.Background(), &newFileRecord)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func FileRetrieveHandler(w http.ResponseWriter, r *http.Request) {

}
