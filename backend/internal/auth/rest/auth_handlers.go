package rest

import (
	"backend/internal/auth"
	"backend/internal/database/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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
	err = repository.CreateNewUser(registerReq.Username, registerReq.Email, hashedPassword)
	if err != nil {
		log.Println(err)

		newError := Error{ErrorN: err.Error()}
		errorJson, _ := json.Marshal(newError)
		http.Error(w, "", http.StatusConflict)
		w.Write(errorJson)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}
