package rest

import (
	"backend/internal/file"
	"backend/internal/file/database"
	"backend/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

// TODO: Refactor Code
// Todo: Write code to check if there are file with same name if yes then do not let them do it.
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	userId, username, err := file_upload.GetUserFromJWT(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	password := r.FormValue("password")
	app, err := utils.InitializeFirebase()
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
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Get a reference to the file in Firebase Storage

	userRef := username + "/"
	fileRef := bucket.Object(userRef + handler.Filename)

	// Create a writer for the Firebase Storage object
	writer := fileRef.NewWriter(context.Background())

	expDate := time.Now().Add(12 * time.Hour)
	expirationDateString := expDate.Format(time.RFC3339)
	metadata := map[string]string{
		"expiry_date": expirationDateString,
	}

	writer.ObjectAttrs.Metadata = metadata
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	// Create New Mongo Record
	fileRepo := database.NewFileRepository(database.ConnectToMongoDB())
	newFileRecord := database.File{
		ID:         primitive.ObjectID{},
		UserID:     userId,
		Username:   username,
		FileName:   handler.Filename,
		FilePath:   fileURL,
		SpecialURL: fileUUID,
		ExpiryDate: expirationDateString,
		Password:   string(hashedPassword),
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
	uuidx := r.FormValue("uuid")
	fileRepo := database.NewFileRepository(database.ConnectToMongoDB())
	decodedU, err := file_upload.DecodeUUID(uuidx)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := fileRepo.FindByUUID(context.Background(), decodedU)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(file)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// TODO: implement with front end.
func FileUploadHistory(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	fileRep := database.NewFileRepository(database.ConnectToMongoDB())
	files, err := fileRep.FindAllUserFiles(context.Background(), username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert files to JSON
	response, err := json.Marshal(files)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
