package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService,
	}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	err := handler.authService.Login()
	if err != nil {
		return
	}
}
