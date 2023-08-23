package controller

import (
	"strconv"

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

	student, err := c.studentService.Create(req.Name, req.SchoolId)
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	presenter.Success(ctx, student)
}

func (c *StudentController) BulkAddStudentToTeacher(ctx *gin.Context) {
	var req presenter.BulkAddStudentToTeacherRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	teacherId := ctx.Param("teacher_id")
	teacherIdInt, err := strconv.Atoi(teacherId)
	if err != nil {
		presenter.Failure(ctx, errs.NewValidationError(err.Error()))
		return
	}

	students, err := c.studentService.BulkAddToTeacher(req.StudentIds, uint(teacherIdInt))
	if err != nil {
		presenter.Failure(ctx, err)
		return
	}

	data := presenter.BulkAddStudentToTeacherResponseDTO{
		AffectedStudentsCount: len(students),
		Students:              students,
	}

	presenter.Success(ctx, data)
}
