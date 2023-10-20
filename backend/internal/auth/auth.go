package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(enteredPassword string) (string, error) {
	cost := 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(enteredPassword), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(enteredPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
	return err
}
