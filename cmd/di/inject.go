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
	jwtManager := jwt.NewJwtManager()

	// Repositories
	userRepository := repository.NewUserRepository(db)
	schoolRepository := repository.NewSchoolRepository(db)

	// Services
	authService := service.NewAuthService(userRepository, jwtManager)
	schoolService := service.NewSchoolService(schoolRepository, userRepository)

	// Routers
	authController := controller.NewAuthController(authService)
	schoolController := controller.NewSchoolController(schoolService)

	return &api.ControllerContainer{
		JwtManager:       jwtManager,
		AuthController:   authController,
		SchoolController: schoolController,
	}
}
