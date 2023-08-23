package di

import (
	"github.com/omidnasiri/mediana-sms/api"
	"github.com/omidnasiri/mediana-sms/api/handler"
	"github.com/omidnasiri/mediana-sms/internal/repository"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
	"github.com/omidnasiri/mediana-sms/service"
	"gorm.io/gorm"
)

// Inject handles dependency injection between application layers
func Inject(db *gorm.DB) *api.HandlerContainer {
	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	jwtService := jwt.NewJwtService()
	authService := service.NewAuthService(userRepository, jwtService)

	// Routers
	authHandler := handler.NewAuthHandler(authService)

	return &api.HandlerContainer{
		AuthHandler: authHandler,
	}
}
