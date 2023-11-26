package utils

import (
	"gorest/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(user models.User) (string, error) {

	claims := jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"number":     user.Number,
		"password":   user.Password,
		"created_at": user.CreatedAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("bishnudevkhutiasecretkey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
