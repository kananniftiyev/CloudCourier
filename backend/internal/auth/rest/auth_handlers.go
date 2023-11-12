// File: rest/handlers.go

package rest

import (
	"backend/internal/auth"
	"backend/internal/auth/database/repository"
	"backend/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var userRepo = repository.NewUserRepository()

func respondWithError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, "", statusCode)
	w.Write(utils.NewResponse(err, statusCode))
}

func parseRequestBody(r *http.Request, data interface{}) error {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(requestBody, data)
}

func registerUser(w http.ResponseWriter, registerReq RegisterRequest) {
	hashedPassword, err := auth.HashPassword(registerReq.Password)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	err = userRepo.CreateUser(registerReq.Username, registerReq.Email, hashedPassword)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var registerReq RegisterRequest
	if err := parseRequestBody(r, &registerReq); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}

	registerUser(w, registerReq)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cookieCheck, err := r.Cookie("jwt")
	if err == nil && cookieCheck != nil {
		return
	}

	var loginReq LoginRequest
	if err := parseRequestBody(r, &loginReq); err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	hashedPassword, err := userRepo.LoginUserCheck(loginReq.Email)
	if err != nil {
		respondWithError(w, err, http.StatusNotFound)
		return
	}

	err = utils.VerifyPassword(loginReq.Password, hashedPassword)
	if err != nil {
		respondWithError(w, err, http.StatusConflict)
		return
	}

	loggedUser, err := userRepo.GetUserWithEmail(loginReq.Email)
	if err != nil {
		respondWithError(w, err, http.StatusUnauthorized)
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

	response := Message{
		Message: "Logged out successfully",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*utils.CustomClaims)
	if !ok {
		http.Error(w, "Failed to get user claims", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID
	user, err := userRepo.GetUserById(userID)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	reqUser := RequestedUserData{Email: user.Email, Username: user.Username, CreatedAt: user.CreatedAt}
	userJSON, err := json.Marshal(reqUser)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
