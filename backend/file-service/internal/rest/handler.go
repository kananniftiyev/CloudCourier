package rest

import (
	file_upload "backend/file-service/internal"
	"backend/file-service/internal/database"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kananniftiyev/cloudcourier-lib/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Refactor Code.
const (
	fileSizeLimit = 50 * 1024 * 1024
)

var firebaseBucket = os.Getenv("FIREBASE_BUCKET_LINK")

// TODO: Write code to check if there are file with same name if yes then do not let them do it.
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
		shared.RespondWithError(w, err, http.StatusInternalServerError)
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

	if err = json.NewEncoder(w).Encode(map[string]string{
		"status":  "201",
		"message": "File Uploaded Successfuly",
	}); err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

}

// TODO: Add code to increase amount of download number of file.
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

	// TODO: Check this
	shared.RespondWithOkay(w, file)
}

// TODO: FileUploadHistory  implement with front end.
func FileUploadHistory(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*shared.CustomClaims)
	if !ok {
		shared.RespondWithError(w, errors.New("failed to get user claims"), http.StatusUnauthorized)
		return
	}
	username := claims.Username
	fileRep := database.NewFileRepository(database.ConnectToMongoDB())
	files, err := fileRep.FindAllUserFiles(context.Background(), username)
	if err != nil {
		shared.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	shared.RespondWithOkay(w, files)

}
