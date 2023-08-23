package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req presenter.LoginRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	token, role, err := c.authService.Login(req.Email, req.Password)
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
