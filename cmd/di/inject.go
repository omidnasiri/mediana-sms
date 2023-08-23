package di

import (
	"github.com/omidnasiri/mediana-sms/api"
	"github.com/omidnasiri/mediana-sms/api/controller"
	"github.com/omidnasiri/mediana-sms/internal/repository"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
	"github.com/omidnasiri/mediana-sms/service"
	"gorm.io/gorm"
)

// Inject handles dependency injection between application layers
func Inject(db *gorm.DB) *api.ControllerContainer {
	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	jwtService := jwt.NewJwtService()
	authService := service.NewAuthService(userRepository, jwtService)

	// Routers
	authController := controller.NewAuthController(authService)

	return &api.ControllerContainer{
		AuthController: authController,
	}
}
