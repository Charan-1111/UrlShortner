package utils

import (
	"goapp/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string) (string, error) {
	issuedAt := time.Now()
	expirationTime := time.Now().Add(models.Config.Jwt.TokenExpiryTime * time.Minute)

	jwtKey := []byte(models.Config.Jwt.JwtSecretKey)

	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			Issuer:    "urlShortner",
			Subject:   "userAuth",
		},
	}

	// create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
