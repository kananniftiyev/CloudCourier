package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func newResponse(error error, code int) *Response {
	newResponse := &Response{
		Status:     "error",
		Message:    error.Error(),
		StatusCode: code,
	}
	return newResponse
}

func RespondWithError(w http.ResponseWriter, err error, statusCode int) {
	response := newResponse(err, statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		return
	}
}

func RespondWithOkay(w http.ResponseWriter, m any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Failed to encode response: %v", err)
		return
	}
}
