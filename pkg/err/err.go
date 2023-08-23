package errs

import (
	"errors"

	goJwt "github.com/golang-jwt/jwt/v4"
)

var (
	// Empty secret key error
	ErrEmptyJWTSecretKey = errors.New("empty jwt secret key")
	// Unexpected signing method error
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	// Invalid token error
	ErrInvalidToken = errors.New("invalid token")
	// Validation Errors which could be returned by "github.com/golang-jwt/jwt/v4" package
	GoJwtValidationErrors = []error{
		goJwt.ErrTokenMalformed,
		goJwt.ErrTokenExpired,
		goJwt.ErrTokenNotValidYet,
		goJwt.ErrTokenInvalidClaims,
	}
)
