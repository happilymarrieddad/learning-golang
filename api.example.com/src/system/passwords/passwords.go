package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = 10
)

func IsValid(hash string, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return false
	}
	return true
}

func Encrypt(password string) (string, error) {
	str, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(str), err
}
