// Package auth provides authentication and authorization functionalities.
package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

// JWTTokenVerifyMiddleware is a middleware function for verifying JWT tokens
func JWTTokenVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the JWT token from the cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			shared.RespondWithError(w, err, http.StatusForbidden)
			return
		}

		// Verify the JWT token
		token, err := jwt.ParseWithClaims(cookie.Value, &shared.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(shared.SECRET_KEY), nil
		})

		if err != nil {
			shared.RespondWithError(w, err, http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			shared.RespondWithError(w, err, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*shared.CustomClaims)

		if !ok {
			shared.RespondWithError(w, errors.New("failed to get token claims"), http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CreateNewJWT generates a new JWT token with custom claims including user ID,
func CreateNewJWT(ID uint, username string) (string, error) {
	// Convert the SECRET_KEY string to a byte array
	key := []byte(shared.SECRET_KEY)

	claimsS := shared.CustomClaims{
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
