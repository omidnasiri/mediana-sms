package service

import (
	"github.com/omidnasiri/mediana-sms/internal/models"
	"github.com/omidnasiri/mediana-sms/internal/repository"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
)

type StudentService interface {
	Create(name string, schoolId uint) (*models.Student, error)
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

func (s *studentService) Create(name string, schoolId uint) (*models.Student, error) {
	school, err := s.schoolRepo.GetById(schoolId)
	if err != nil {
		return nil, err
	}

	user := &models.User{Name: name}
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	student := &models.Student{UserId: user.ID}
	err = s.studentRepo.Create(student)
	if err != nil {
		return nil, err
	}

	err = s.schoolRepo.Create(school)
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
		err := errs.NewNotFoundError("no entities with the provided ids and seo_id needed to be affected")
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
