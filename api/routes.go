package api

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/controller"
)

const basePath string = "api/v1"

type ControllerContainer struct {
	AuthController *controller.AuthController
}

func SetupRoutes(controllers *ControllerContainer) *gin.Engine {
	app := gin.Default()

	apiV1 := app.Group(basePath)
	{
		authRouter := apiV1.Group("/auth")
		{
			authRouter.POST("/login", controllers.AuthController.Login)
		}
	}

	return app
}
