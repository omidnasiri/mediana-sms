package presenter

import (
	"github.com/gin-gonic/gin"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
)

func Failure(ctx *gin.Context, err error) {
	ctx.JSON(errs.GetHttpStatusCodeFromError(err), newGenericResponse(nil, err.Error(), false))
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(200, newGenericResponse(data, "", true))
}

func newGenericResponse(data any, err string, result bool) GenericResponse {
	return GenericResponse{
		Data:   data,
		Errors: err,
		Result: result,
	}
}

type GenericResponse struct {
	Data   any    `json:"data,omitempty"`
	Errors string `json:"errors,omitempty"`
	Result bool   `json:"result"`
}
