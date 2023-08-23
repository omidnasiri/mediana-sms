package service

import (
	"github.com/omidnasiri/mediana-sms/internal/repository"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
)

type AuthService interface {
	Login() error
}

type authService struct {
	userRepo   repository.UserRepository
	jwtService jwt.JWT
}

func NewAuthService(userRepo repository.UserRepository, jwtService jwt.JWT) AuthService {
	return &authService{
		userRepo,
		jwtService,
	}
}

func (s *authService) Login() error {
	return nil
}
