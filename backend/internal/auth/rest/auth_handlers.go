package rest

import (
	"backend/internal/auth"
	"backend/internal/auth/database/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var userRepo = repository.NewUserRepository()

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var registerReq RegisterRequest
	err = json.Unmarshal(requestBody, &registerReq)
	if err != nil {
		http.Error(w, "Failed to parse JSON request", http.StatusBadRequest)
	}
	hashedPassword, err := auth.HashPassword(registerReq.Password)
	if err != nil {
		log.Fatal("Could not hash The password")
	}

	err = userRepo.CreateUser(registerReq.Username, registerReq.Email, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		log.Println(err)

		newError := Error{ErrorN: err.Error()}
		errorJson, _ := json.Marshal(newError)
		http.Error(w, "", http.StatusConflict)
		w.Write(errorJson)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cookieCheck, err := r.Cookie("jwt")
	if err == nil && cookieCheck != nil {
		http.Error(w, "User is already authenticated", http.StatusUnauthorized)
		return

	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var loginReq LoginRequest
	err = json.Unmarshal(requestBody, &loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPassword, err := userRepo.LoginUserCheck(loginReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	err = auth.VerifyPassword(loginReq.Password, hashedPassword)
	if err != nil {
		http.Error(w, "Email or Password is wrong", http.StatusConflict)
		return
	}
	loggedUser, err := userRepo.GetUserWithEmail(loginReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateNewJWT(loggedUser.ID, loggedUser.Username)

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{Name: "jwt", Value: "", Expires: time.Unix(0, 0), HttpOnly: true, Domain: "localhost",
		Path: "/"}
	http.SetCookie(w, cookie)

	// Prepare a response message.
	response := Message{
		Message: "Logged out successfully",
	}

	// Marshal the response to JSON.
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	// Set the response headers.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*auth.CustomClaims)
	if !ok {
		http.Error(w, "Failed to get user claims", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID
	user, err := userRepo.GetUserById(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reqUser := RequestedUserData{Email: user.Email, Username: user.Username, CreatedAt: user.CreatedAt}
	userJSON, err := json.Marshal(reqUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
