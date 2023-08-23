package service

import (
	"github.com/omidnasiri/mediana-sms/internal/repository"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (jwt.Token, string, error)
}

type authService struct {
	userRepo   repository.UserRepository
	jwtManager jwt.JWT
}

func NewAuthService(userRepo repository.UserRepository, jwtManager jwt.JWT) AuthService {
	return &authService{
		userRepo,
		jwtManager,
	}
}

func (s *authService) Login(email, password string) (jwt.Token, string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return jwt.Token{}, "", err
	}

	err = s.verifyPassword(user.PasswordHash, password)
	if err != nil {
		return jwt.Token{}, "", err
	}

	jwtToken, err := s.jwtManager.CreateAccessToken(user.ID, user.Role.Title)
	if err != nil {
		return jwt.Token{}, "", err
	}

	return jwtToken, user.Role.Title, nil
}

func (s *authService) verifyPassword(passwordHash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return errs.NewUnauthorizedError("invalid username or password")
	}
	return nil
}
