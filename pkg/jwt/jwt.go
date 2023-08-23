package jwt

import (
	"log"
	"os"
	"strconv"
	"time"

	goJwt "github.com/golang-jwt/jwt/v4"
)

type JWT interface {
	CreateAccessToken(userId, userRole string) (Token, error)
}

type jwtService struct {
	cfg Config
}

func NewJwtService() JWT {
	accessTokenTtlInt, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_TTL"))
	if err != nil {
		log.Fatal(err)
	}
	cfg := Config{
		secret:         os.Getenv("JWT_SECRET"),
		accessTokenTTL: time.Duration(accessTokenTtlInt),
	}
	return &jwtService{
		cfg: cfg,
	}
}

func (jwt *jwtService) CreateAccessToken(userId, userRole string) (Token, error) {
	now := time.Now()
	var jwtKey = []byte(jwt.cfg.secret)

	accessExpirationTime := now.Add(jwt.cfg.accessTokenTTL)
	claims := &JwtClaims{
		UserId:   userId,
		UserRole: userRole,
		RegisteredClaims: goJwt.RegisteredClaims{
			IssuedAt:  goJwt.NewNumericDate(now),
			ExpiresAt: goJwt.NewNumericDate(accessExpirationTime),
		},
	}
	accessToken := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return Token{}, err
	}

	return Token{accessTokenString, accessExpirationTime}, nil
}
