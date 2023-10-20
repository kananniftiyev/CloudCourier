package repository

import (
	"backend/internal/database"
	"backend/internal/database/models"
	"errors"
	"time"
)

func CreateNewUser(username, email, hashedPassword string) error {
	db := database.ConnectDatabase()

	err := FindUser(email, username)
	if err != nil {
		return err
	}

	newUser := models.User{Username: username, Email: email, PasswordHash: hashedPassword, RegistrationDate: time.Now()}
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func FindUser(email, username string) error {
	db := database.ConnectDatabase()

	var existingUser models.User
	db.Where("email = ? OR username = ?", email, username).First(&existingUser)
	if existingUser.ID != 0 {
		return errors.New("User with this email or username already exists")
	}
	return nil
}

func LoginUserCheck(email string) (string, error) {
	db := database.ConnectDatabase()

	var existingUser models.User
	db.Where("email = ?", email).First(&existingUser)
	if existingUser.ID == 0 {
		return "", errors.New("User Does Not Exist")
	}
	return existingUser.PasswordHash, nil
}

func GetUserWithEmail(email string) (models.User, error) {
	db := database.ConnectDatabase()

	var existingUser models.User
	db.Where(db.Where("email = ?", email).First(&existingUser))
	if existingUser.ID == 0 {
		return models.User{}, errors.New("User Does Not Exist")
	}
	return existingUser, nil
}
