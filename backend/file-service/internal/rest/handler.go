package rest

import (
	file_upload "backend/file-service/internal"
	"backend/file-service/internal/database"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kananniftiyev/cloudcourier-lib/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	firebaseBucket = "cloudsharex-b8353.appspot.com"
	fileSizeLimit  = 50 * 1024 * 1024
)


// Todo: Write code to check if there are file with same name if yes then do not let them do it.
// Todo: Last Changes
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*shared.CustomClaims)
	if !ok {
		http.Error(w, "Failed to get user claims", http.StatusUnauthorized)
		return
	}
	userId := claims.UserID
	username := claims.Username
	
	password := r.FormValue("password")
	title := r.FormValue("title")
	app, err := database.InitializeFirebase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a Firebase Storage client
	storageClient, err := app.Storage(context.Background())
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Get a reference to the Firebase Storage bucket
	bucket, err := storageClient.Bucket(firebaseBucket)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Parse the form data to extract the file
	err = r.ParseMultipartForm(fileSizeLimit)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		shared.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			shared.RespondWithError(w, err, http.StatusInternalServerError)
			return
		}
	}(file)

	// Get a reference to the file in Firebase Storage

	userRef := username + "/"
	fileRef := bucket.Object(userRef + handler.Filename)

	// Create a writer for the Firebase Storage object
	writer := fileRef.NewWriter(context.Background())

	expDate := time.Now().Add(6 * time.Hour)
	expirationDateString := expDate.Format(time.RFC3339)
	metadata := map[string]string{
		"expiry_date": expirationDateString,
	}

	writer.ObjectAttrs.Metadata = metadata
	// Copy the uploaded file's content to Firebase Storage
	_, err = io.Copy(writer, file)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	fileUUID := uuid.New()

	fileURL, err := file_upload.GetFileURL(firebaseBucket, userRef+handler.Filename)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Create New Mongo Record
	fileRepo := database.NewFileRepository(database.ConnectToMongoDB())
	newFileRecord := database.File{
		ID:             primitive.ObjectID{},
		UserID:         int(userId),
		Username:       username,
		FileName:       handler.Filename,
		FilePath:       fileURL,
		SpecialURL:     fileUUID,
		ExpiryDate:     expirationDateString,
		Password:       string(hashedPassword),
		Title:          title,
		TotalDownloads: 0,
		UploadDate:     time.Now().Format("January 2, 2006"),
	}

	err = fileRepo.Create(context.Background(), &newFileRecord)

	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func FileRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	uuidx := r.FormValue("uuid")
	password := r.FormValue("password")
	fileRepo := database.NewFileRepository(database.ConnectToMongoDB())
	decodedU, err := file_upload.DecodeUUID(uuidx)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	file, err := fileRepo.FindByUUID(context.Background(), decodedU)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	err = shared.VerifyPassword(password, file.Password)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(file)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// FileUploadHistory TODO: implement with front end.
func FileUploadHistory(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*shared.CustomClaims)
	if !ok {
		http.Error(w, "Failed to get user claims", http.StatusUnauthorized)
		return
	}
	username := claims.Username
	fileRep := database.NewFileRepository(database.ConnectToMongoDB())
	files, err := fileRep.FindAllUserFiles(context.Background(), username)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Convert files to JSON
	response, err := json.Marshal(files)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
