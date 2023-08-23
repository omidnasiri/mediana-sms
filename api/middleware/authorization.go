package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	"github.com/omidnasiri/mediana-sms/internal/models"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/pkg/jwt"
	"github.com/omidnasiri/mediana-sms/pkg/utils"
)

// Authentication middleware checks for existence of the jwt token in
// the right format in request header, validates it and if successful
// stores user claims in request context
func Authorization(Roles []models.RoleType) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		claims, ok := ctx.Get("userClaims")
		if !ok {
			presenter.Failure(ctx, errs.NewForbiddenError(""))
			return
		}

		if !utils.Contains(Roles, *jwt.JwtClaims(claims).Role) {
			presenter.Failure(ctx, errs.NewForbiddenError(""))
			return
		}

		ctx.Next()
	}
}
