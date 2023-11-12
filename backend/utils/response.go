package utils

import (
	"encoding/json"
	"log"
)

type Response struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func NewResponse(error error, code int) []byte {
	newResponse := Response{
		Status:     "error",
		Message:    error.Error(),
		StatusCode: code,
	}
	js, err := json.Marshal(newResponse)
	if err != nil {
		log.Fatal(err.Error())
	}
	return js
}
