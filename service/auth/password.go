package auth

import (
	"crypto/sha512"
)

func HashPassword(password string) (string, error) {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hash := hasher.Sum(nil)

	return string(hash), nil
}
