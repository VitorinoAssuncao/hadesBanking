package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(text string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(text), 14)
	return string(bytes)
}

func ValidateHash(accountSecret, loginSecret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(accountSecret), []byte(loginSecret))
	return err == nil
}
