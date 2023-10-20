package rest

import (
	"backend/internal/auth"
	"backend/internal/database/repository"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	userRepo := repository.NewUserRepository()
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
		http.Error(w, err.Error(), http.StatusConflict)
		w.Write(errorJson)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	userRepo := repository.NewUserRepository()
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
	logedUser, err := userRepo.GetUserWithEmail(loginReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(logedUser.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	tokken, err := claims.SigningString()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokken,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	// Todo: Sent Necessary data to user, and look how to use jwt in this.

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
