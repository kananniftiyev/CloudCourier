package repository

import (
	"backend/internal/database"
	"backend/internal/database/models"
	"errors"
	"fmt"
	"time"
)

func CreateNewUser(username, email, hashedPassword string) error {
	db := database.ConnectDatabase()

	var existingUser models.User
	db.Where("email = ? OR username = ?", email, username).First(&existingUser)
	if existingUser.ID != 0 {
		return errors.New("User with this email or username already exists")
	}
	newUser := models.User{Username: username, Email: email, PasswordHash: hashedPassword, RegistrationDate: time.Now()}
	if err := db.Create(&newUser).Error; err != nil {
		fmt.Println("B")
		return err
	}
	return nil
}
