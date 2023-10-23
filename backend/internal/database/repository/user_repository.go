package repository

import (
	"backend/internal/database"
	"backend/internal/database/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserNotFound      = errors.New("User not found")
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrDatabaseOperation = errors.New("Database operation failed")
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.ConnectDatabase(),
	}
}

func (ur *UserRepository) CreateUser(username, email, hashedPassword string) error {
	if ur.userExists(email, username) {
		return ErrUserAlreadyExists
	}

	newUser := models.User{Username: username, Email: email, PasswordHash: hashedPassword, RegistrationDate: time.Now()}
	if err := ur.db.Create(&newUser).Error; err != nil {
		return ErrDatabaseOperation
	}
	return nil
}

func (ur *UserRepository) LoginUserCheck(email string) (string, error) {
	var existingUser models.User
	ur.db.Where("email = ?", email).First(&existingUser)
	if existingUser.ID == 0 {
		return "", ErrUserNotFound
	}
	return existingUser.PasswordHash, nil
}

func (ur *UserRepository) GetUserWithEmail(email string) (models.User, error) {
	var existingUser models.User
	ur.db.Where(ur.db.Where("email = ?", email).First(&existingUser))
	if existingUser.ID == 0 {
		return models.User{}, ErrUserNotFound
	}
	return existingUser, nil
}

func (ur *UserRepository) userExists(email, username string) bool {
	var user models.User
	err := ur.db.Where("email = ? AND username = ?", email, username).First(&user)
	if err != nil {
		return true
	}
	return false
}

func (ur *UserRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	ur.db.Where("id = ?", id).First(&user)

	return user, nil
}
