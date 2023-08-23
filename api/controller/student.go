package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omidnasiri/mediana-sms/api/presenter"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/service"
)

type StudentController struct {
	studentService service.StudentService
}

func NewStudentController(studentService service.StudentService) *StudentController {
	return &StudentController{
		studentService,
	}
}

func (c *StudentController) Create(ctx *gin.Context) {
	var req presenter.CreateStudentRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	Student, err := c.studentService.Create(req.Name, req.SchoolId)
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	presenter.Success(ctx, Student)
}
