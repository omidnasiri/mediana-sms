package service

import "github.com/omidnasiri/mediana-sms/internal/repository"

type AuthService interface {
	Login() error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo,
	}
}

func (s *authService) Login() error {
	return nil
}
