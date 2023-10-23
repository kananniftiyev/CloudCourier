package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username         string
	Email            string
	PasswordHash     string
	RegistrationDate time.Time
	Files            []Files
}

type Files struct {
	gorm.Model
	UserID     uint
	User       User
	FileName   string
	FilePath   string
	FileSize   float64
	ExpiryDate time.Time
	IsDeleted  bool
}
