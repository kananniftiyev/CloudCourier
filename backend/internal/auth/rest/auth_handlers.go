package rest

import (
	"backend/internal/auth"
	"backend/internal/database/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	const SECRET_KEY = "secret"
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
	loggedUser, err := userRepo.GetUserWithEmail(loginReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateNewJWT(loggedUser.ID)
	fmt.Println(token + "a")

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	// Todo: Sent Necessary data to user, and look how to use jwt in this.

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
