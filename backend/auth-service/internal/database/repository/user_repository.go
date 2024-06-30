package repository

import (
	"backend/auth-service/internal/database"
	"backend/auth-service/internal/database/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

// TODO: put erros into one file.
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrDatabaseOperation = errors.New("database operation failed")
)

type UserRepository struct {
	db *gorm.DB
}

// TODO: Singelton Pattern.
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
	ur.db.Where("email = ? OR username = ?", email, username).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func (ur *UserRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	ur.db.Where("id = ?", id).First(&user)

	return user, nil
}
