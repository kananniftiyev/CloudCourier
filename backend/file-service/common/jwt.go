package common

import "github.com/dgrijalva/jwt-go"

const SECRET_KEY = "secret"

type CustomClaims struct {
	jwt.StandardClaims
	UserID   uint
	Username string

	// Add other custom claims as needed
}
