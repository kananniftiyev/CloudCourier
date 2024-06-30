package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username         string `gorm:"unique;not null"`
	Email            string `gorm:"unique;not null"`
	PasswordHash     string
	RegistrationDate time.Time
}
