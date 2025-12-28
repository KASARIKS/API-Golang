package auth

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/api_golang/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiresAt": config.Envs.JWTExpirationInSeconds,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
