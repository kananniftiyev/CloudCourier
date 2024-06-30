package rest

import (
	"time"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Message struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestedUserData struct {
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}
