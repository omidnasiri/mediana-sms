package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

// Authentication middleware checks for existence of the jwt token in
// the right format in request header, validates it and if successful
// stores user claims in request context
func Authentication(jwtManager jwt.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := authHeader{}

		err := ctx.ShouldBindHeader(&h)
		if err != nil {
			presenter.Failure(ctx, errs.NewUnauthorizedError(err.Error()))
			return
		}

		tokenHeader := strings.Split(h.Token, "Bearer ")

		if len(tokenHeader) < 2 {
			presenter.Failure(ctx, errs.NewUnauthorizedError("authorization header must be in `Bearer {token}` format"))
			return
		}

		userClaims, err := jwtManager.ParseJwtToken(tokenHeader[1])
		if err != nil {
			presenter.Failure(ctx, errs.NewUnauthorizedError("invalid token"))
			return
		}

		ctx.Set("userId", userClaims.UserId)
		ctx.Set("userRole", userClaims.UserRole)

		ctx.Next()
	}
}
