package api

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/controller"
	"github.com/omidnasiri/mediana-sms/api/middleware"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
)

const basePath string = "api/v1"

type ControllerContainer struct {
	JwtManager jwt.JWT

	AuthController *controller.AuthController
}

func SetupRoutes(controllers *ControllerContainer) *gin.Engine {
	app := gin.Default()

	app.Use(middleware.Authentication(controllers.JwtManager))

	apiV1 := app.Group(basePath)
	{
		authRouter := apiV1.Group("/auth")
		{
			authRouter.POST("/login", controllers.AuthController.Login)
		}
	}

	return app
}
