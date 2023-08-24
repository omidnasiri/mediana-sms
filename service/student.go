package service

import (
	"github.com/omidnasiri/mediana-sms/internal/models"
	"github.com/omidnasiri/mediana-sms/internal/repository"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
	"github.com/omidnasiri/mediana-sms/pkg/utils"
)

type StudentService interface {
	Create(name, email, password string, schoolId uint) (*models.Student, error)
	BulkAddToTeacher(studentIds []uint, teacherId uint) ([]*models.Student, error)
}

type studentService struct {
	studentRepo repository.StudentRepository
	teacherRepo repository.TeacherRepository
	schoolRepo  repository.SchoolRepository
	userRepo    repository.UserRepository
}

func NewStudentService(studentRepo repository.StudentRepository, teacherRepo repository.TeacherRepository, schoolRepo repository.SchoolRepository, userRepo repository.UserRepository) StudentService {
	return &studentService{
		studentRepo,
		teacherRepo,
		schoolRepo,
		userRepo,
	}
}

func (s *studentService) Create(name, email, password string, schoolId uint) (*models.Student, error) {
	_, err := s.schoolRepo.GetById(schoolId)
	if err != nil {
		return nil, err
	}

	passwordHash, err := utils.HashPassword([]byte(password))
	if err != nil {
		return nil, err
	}

	role := &models.Role{Title: models.ROLE_STUDENT}
	user := &models.User{Name: name, Email: email, PasswordHash: passwordHash, SchoolId: &schoolId, Role: role}
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	student := &models.Student{User: user}
	err = s.studentRepo.Create(student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentService) BulkAddToTeacher(studentIds []uint, teacherId uint) ([]*models.Student, error) {
	teacher, err := s.teacherRepo.GetById(teacherId)
	if err != nil {
		return nil, err
	}

	models, err := s.studentRepo.GetListByIDsExcludeTeacherID(studentIds, teacherId)
	if err != nil {
		return nil, err
	}
	if len(models) < 1 {
		err := errs.NewNotFoundError("no entities with the provided student ids and teacher id needed to be affected")
		return nil, err
	}

	for i := range models {
		models[i].Teachers = append(models[i].Teachers, teacher)
	}

	err = s.studentRepo.BulkUpdate(models)
	if err != nil {
		return nil, err
	}

	return models, nil
}
