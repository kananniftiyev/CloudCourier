// File: rest/handlers.go

package rest

import (
	"backend/auth-service/common"
	auth "backend/auth-service/internal"
	"backend/auth-service/internal/database/repository"
	"net/http"
	"time"
)

// TODO: fix perfomance.
// TODO: create gRPC if needed.
// TODO: Reformat code

var userRepo = repository.NewUserRepository()


func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var registerReq RegisterRequest
	if err := parseRequestBody(r, &registerReq); err != nil {
		common.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(registerReq.Password)
	if err != nil {
		common.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	err = userRepo.CreateUser(registerReq.Username, registerReq.Email, hashedPassword)

	if err == repository.ErrUserAlreadyExists{
		common.RespondWithError(w, err, http.StatusConflict)
		return
	}
	if err != nil {
		common.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	cookieCheck, err := r.Cookie("jwt")
	if err == nil && cookieCheck != nil {
		http.Error(w, "User Already Logged in", http.StatusConflict)
		return
	}

	var loginReq LoginRequest
	if err := parseRequestBody(r, &loginReq); err != nil {
		common.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	hashedPassword, err := userRepo.LoginUserCheck(loginReq.Email)
	if err != nil {
		common.RespondWithError(w, err, http.StatusNotFound)
		return
	}

	err = common.VerifyPassword(loginReq.Password, hashedPassword)
	if err != nil {
		common.RespondWithError(w, err, http.StatusConflict)
		return
	}

	loggedUser, err := userRepo.GetUserWithEmail(loginReq.Email)
	if err != nil {
		common.RespondWithError(w, err, http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateNewJWT(loggedUser.ID, loggedUser.Username)

	if err != nil {
	  common.RespondWithError(w, err, http.StatusInternalServerError)
	  return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{Name: "jwt", Value: "", Expires: time.Unix(0, 0), HttpOnly: true, Domain: "localhost",
		Path: "/"}
	http.SetCookie(w, cookie)

	response := Message{
		Message: "Logged out successfully",
	}
	
	common.RespondWithOkay(w, response)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*common.CustomClaims)
	if !ok {
		http.Error(w, "Failed to get user claims", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID
	user, err := userRepo.GetUserById(userID)
	if err != nil {
		common.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	reqUser := RequestedUserData{Email: user.Email, Username: user.Username, CreatedAt: user.CreatedAt}
	
	common.RespondWithOkay(w, reqUser)
}
