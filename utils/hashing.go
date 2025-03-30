package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hashedPassword), nil
	}
}

func CheckPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}
