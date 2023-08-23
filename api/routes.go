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

	AuthController   *controller.AuthController
	SchoolController *controller.SchoolController
}

func SetupRoutes(controllers *ControllerContainer) *gin.Engine {
	app := gin.Default()

	apiV1 := app.Group(basePath)
	{
		authRouter := apiV1.Group("/auth")
		{
			authRouter.POST("/login", controllers.AuthController.Login)
		}
		schoolRouter := apiV1.Group("/school", middleware.Authentication(controllers.JwtManager))
		{
			schoolRouter.POST("/", middleware.Authorization([]string{models.ROLE_ADMIN}), controllers.SchoolController.Create)
		}
	}

	return app
}
