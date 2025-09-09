package validator

import (
	"fmt"
	"goapp/models"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) (*models.Claims, error) {
	jwtKey := []byte(models.Config.Jwt.JwtSecretKey)

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token provided")
	}

	return claims, nil
}
