package api

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/handler"
)

const basePath string = "api/v1"

type HandlerContainer struct {
	AuthHandler *handler.AuthHandler
}

func SetupRoutes(handlers *HandlerContainer) *gin.Engine {
	app := gin.Default()

	apiV1 := app.Group(basePath)
	{
		authRouter := apiV1.Group("/auth")
		{
			authRouter.POST("/login", handlers.AuthHandler.Login)
		}
	}

	return app
}
