package repository

import (
	"backend/internal/database"
	"backend/internal/database/models"
	"time"
)

func CreateNewUser(username, email, hashedPassword string) {
	db := database.ConnectDatabase()
	newUser := models.User{Username: username, Email: email, PasswordHash: hashedPassword, RegistrationDate: time.Now()}
	db.Create(&newUser)

}
