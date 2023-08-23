package jwt

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	goJwt "github.com/golang-jwt/jwt/v4"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
)

type JWT interface {
	CreateAccessToken(userId uint, userRole string) (Token, error)
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

func (jwt *jwtService) CreateAccessToken(userId uint, userRole string) (Token, error) {
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

func (jwt *jwtService) ParseJwtToken(tokenString string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	token, err := goJwt.ParseWithClaims(tokenString, claims,
		func(token *goJwt.Token) (any, error) {
			if _, ok := token.Method.(*goJwt.SigningMethodHMAC); !ok {
				return nil, errs.ErrUnexpectedSigningMethod
			}
			return []byte(jwt.cfg.secret), nil
		})

	if err != nil {
		ve, errHasValidationError := err.(*goJwt.ValidationError)
		if errHasValidationError {
			for _, goJwtValidationErr := range errs.GoJwtValidationErrors {
				if ve.Is(goJwtValidationErr) {
					return nil, goJwtValidationErr
				}
			}
			return nil, ve.Inner
		} else {
			return nil, fmt.Errorf("failled to parse token >>> %s", err.Error())
		}
	}

	if !token.Valid {
		return nil, errs.ErrInvalidToken
	}

	return claims, nil
}
