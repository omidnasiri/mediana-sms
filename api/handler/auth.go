package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
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
	var req presenter.LoginRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	token, role, err := handler.authService.Login(req.Email, req.Password)
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	data := presenter.LoginResultDTO{
		AccessToken:           token.Raw,
		AccessTokenExpiration: token.ExpiresAt,
		Role:                  role,
	}

	presenter.Success(ctx, data)
}
