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
	teacherRepository := repository.NewTeacherRepository(db)
	studentRepository := repository.NewstudentRepository(db)

	// Services
	authService := service.NewAuthService(userRepository, jwtManager)
	schoolService := service.NewSchoolService(schoolRepository, userRepository)
	teacherService := service.NewTeacherService(teacherRepository, schoolRepository, userRepository)
	studentService := service.NewStudentService(studentRepository, schoolRepository, userRepository)

	// Routers
	authController := controller.NewAuthController(authService)
	schoolController := controller.NewSchoolController(schoolService)
	teacherController := controller.NewTeacherController(teacherService)
	studentController := controller.NewStudentController(studentService)

	return &api.ControllerContainer{
		JwtManager:        jwtManager,
		AuthController:    authController,
		SchoolController:  schoolController,
		TeacherController: teacherController,
		StudentController: studentController,
	}
}
