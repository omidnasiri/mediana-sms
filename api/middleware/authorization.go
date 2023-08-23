package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	"github.com/omidnasiri/mediana-sms/internal/models"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/pkg/utils"
)

// Authorization middleware checks if the provided user role
// can be found within the accepted roles of the endpoint
func Authorization(Roles []models.RoleType) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userRole := ctx.GetString("userRole")

		if !utils.Contains(Roles, userRole) {
			presenter.Failure(ctx, errs.NewForbiddenError(""))
			return
		}

		ctx.Next()
	}
}
