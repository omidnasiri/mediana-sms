package jwt

import (
	"time"

	goJwt "github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Raw       string
	ExpiresAt time.Time
}

type Config struct {
	secret         string
	accessTokenTTL time.Duration
}

type JwtClaims struct {
	goJwt.RegisteredClaims
	UserId   uint   `json:"user_id"`
	UserRole string `json:"user_role"`
}
