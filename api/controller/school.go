package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/service"
)

type SchoolController struct {
	schoolService service.SchoolService
}

func NewSchoolController(schoolService service.SchoolService) *SchoolController {
	return &SchoolController{
		schoolService,
	}
}

func (c *SchoolController) Create(ctx *gin.Context) {
	var req presenter.CreateSchoolRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	school, err := c.schoolService.Create(req.Title, req.HeadmasterId)
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	presenter.Success(ctx, school)
}
