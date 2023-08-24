package controller

import (
	"strconv"

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

	teacher, err := c.teacherService.Create(req.Name, req.Email, req.Password, req.SchoolId)
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	presenter.Success(ctx, teacher)
}

func (c *TeacherController) GetStudents(ctx *gin.Context) {
	teacherId := ctx.Param("teacher_id")
	teacherIdInt, err := strconv.Atoi(teacherId)
	if err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	students, err := c.teacherService.GetStudents(uint(teacherIdInt))
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	presenter.Success(ctx, students)
}
