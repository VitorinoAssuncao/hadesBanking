package types

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) Hash() Password {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	hash := string(bytes)
	return Password(hash)
}

func (p Password) CompareSecret(input string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(input))
	return err
}

func (p Password) ToString() string {
	return string(p)
}
