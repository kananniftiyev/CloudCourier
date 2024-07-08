// Package auth provides authentication and authorization functionalities.
package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kananniftiyev/cloudcourier-lib/shared"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash from the provided password string.
func HashPassword(enteredPassword string) (string, error) {
	cost := 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(enteredPassword), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CreateNewJWT generates a new JWT token with custom claims including user ID,
func CreateNewJWT(ID uint, username string) (string, error) {
	// Convert the SECRET_KEY string to a byte array
	key := []byte(shared.SECRET_KEY)

	// Create custom claim.
	claims := shared.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		UserID:   ID,
		Username: username,
	}

	// Create token struct with signing method and claims we gave.
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and turn into string
	tokenString, err := tokenStruct.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
