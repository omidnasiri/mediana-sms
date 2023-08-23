package errs

import (
	"errors"
	"fmt"
	"net/http"

	goJwt "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var (
	ErrValidation              = errors.New("validation error")
	ErrNotFound                = errors.New("not found error")
	ErrDuplicateEntity         = errors.New("duplicate entity")
	ErrForbidden               = errors.New("forbidden error")
	ErrUnauthorized            = errors.New("unauthorized error")
	ErrInvalidToken            = errors.New("invalid token")
	ErrEmptyJWTSecretKey       = errors.New("empty jwt secret key")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

	// Validation Errors which could be returned by "github.com/golang-jwt/jwt/v4" package
	GoJwtValidationErrors = []error{
		goJwt.ErrTokenMalformed,
		goJwt.ErrTokenExpired,
		goJwt.ErrTokenNotValidYet,
		goJwt.ErrTokenInvalidClaims,
	}
)

func NewValidationError(errMsg string) error {
	return fmt.Errorf("%w: %s", ErrValidation, errMsg)
}

func NewNotFoundError(errMsg string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, errMsg)
}

func NewDuplicateEntity(errMsg string) error {
	return fmt.Errorf("%w: %s", ErrDuplicateEntity, errMsg)
}

func NewForbiddenError(errMsg string) error {
	return fmt.Errorf("%w: %s", ErrForbidden, errMsg)
}

func NewUnauthorizedError(errMsg string) error {
	return fmt.Errorf("%w: %s", ErrUnauthorized, errMsg)
}

func GetHttpStatusCodeFromError(err error) int {
	if errors.Is(err, ErrValidation) {
		return http.StatusBadRequest
	} else if errors.Is(err, ErrNotFound) || errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	} else if errors.Is(err, ErrDuplicateEntity) {
		return http.StatusConflict
	} else if errors.Is(err, ErrForbidden) {
		return http.StatusForbidden
	} else if errors.Is(err, ErrUnauthorized) {
		return http.StatusUnauthorized
	} else {
		return http.StatusInternalServerError
	}
}
