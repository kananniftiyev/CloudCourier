package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

const SECRET_KEY = "secret"

func HashPassword(enteredPassword string) (string, error) {
	cost := 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(enteredPassword), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(enteredPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
	return err
}

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the JWT token from the cookie
		cookie, err := r.Cookie("jwt") // "jwt" is the cookie name, change it to match your setup
		if err != nil {
			http.Error(w, "No token provided", http.StatusForbidden)
			return
		}

		// Verify the JWT token
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})

		if err != nil {
			http.Error(w, "Failed to authenticate token", http.StatusUnauthorized)
			return
		}

		_ = token

		// Token is valid; you can access user information in the token.Claims field if needed
		next.ServeHTTP(w, r)
	})
}

func CreateNewJWT(ID uint) (string, error) {
	// Convert the SECRET_KEY string to a byte array
	key := []byte(SECRET_KEY)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}
