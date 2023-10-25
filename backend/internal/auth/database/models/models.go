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
}
