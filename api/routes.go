package api

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/controller"
	"github.com/omidnasiri/mediana-sms/api/middleware"
	"github.com/omidnasiri/mediana-sms/internal/models"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
)

const basePath string = "api/v1"

type ControllerContainer struct {
	JwtManager jwt.JWT

	AuthController    *controller.AuthController
	SchoolController  *controller.SchoolController
	TeacherController *controller.TeacherController
	StudentController *controller.StudentController
}

func SetupRoutes(controllers *ControllerContainer) *gin.Engine {
	app := gin.Default()

	apiV1 := app.Group(basePath)
	{
		authRouter := apiV1.Group("/auth")
		{
			authRouter.POST("/login", controllers.AuthController.Login)
		}
		// TODO: implement admin create headmaster api
		schoolRouter := apiV1.Group("/school", middleware.Authentication(controllers.JwtManager))
		{
			schoolRouter.POST("/", middleware.Authorization([]string{models.ROLE_ADMIN}), controllers.SchoolController.Create)
		}
		teacherRouter := apiV1.Group("/teacher", middleware.Authentication(controllers.JwtManager))
		{
			teacherRouter.POST("/", middleware.Authorization([]string{models.ROLE_HEADMASTER}), controllers.TeacherController.Create)
			teacherRouter.GET("/students", middleware.Authorization([]string{models.ROLE_TEACHER}), controllers.TeacherController.GetStudents)
		}
		studentRouter := apiV1.Group("/student", middleware.Authentication(controllers.JwtManager))
		{
			studentRouter.POST("/", middleware.Authorization([]string{models.ROLE_HEADMASTER}), controllers.StudentController.Create)
			studentRouter.POST("/:teacher_id", middleware.Authorization([]string{models.ROLE_HEADMASTER}), controllers.StudentController.BulkAddStudentToTeacher)
		}
	}

	return app
}
