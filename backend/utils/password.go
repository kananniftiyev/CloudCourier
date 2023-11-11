package utils

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(enteredPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
	return err
}
