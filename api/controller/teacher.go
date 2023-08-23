package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/service"
)

type TeacherController struct {
	teacherService service.TeacherService
}

func NewTeacherController(teacherService service.TeacherService) *TeacherController {
	return &TeacherController{
		teacherService,
	}
}

func (c *TeacherController) Create(ctx *gin.Context) {
	var req presenter.CreateTeacherRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	teacher, err := c.teacherService.Create(req.Name, req.SchoolId)
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	presenter.Success(ctx, teacher)
}
