package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWT struct {
	JwtSecretKey    string        `json:"jwt_secret_key"`
	TokenExpiryTime time.Duration `json:"token_expiry_time"`
}
