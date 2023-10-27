package auth

import (
	"backend/utils"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HashPassword(enteredPassword string) (string, error) {
	cost := 14
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
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "No token provided", http.StatusForbidden)
			return
		}

		// Verify the JWT token
		token, err := jwt.ParseWithClaims(cookie.Value, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.SECRET_KEY), nil
		})

		if err != nil {
			http.Error(w, "Failed to authenticate token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			log.Printf("JWT token is not valid")
			return
		}

		claims, ok := token.Claims.(*utils.CustomClaims)

		if !ok {
			http.Error(w, "Failed to get token claims", http.StatusUnauthorized)
			log.Printf("Failed to get custom claims from JWT token", err)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CreateNewJWT(ID uint, username string) (string, error) {
	// Convert the SECRET_KEY string to a byte array
	key := []byte(utils.SECRET_KEY)

	claimsS := utils.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		UserID:   ID,
		Username: username,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsS)

	token, err := claims.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}
