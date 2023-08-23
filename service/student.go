package service

import (
	"github.com/omidnasiri/mediana-sms/internal/models"
	"github.com/omidnasiri/mediana-sms/internal/repository"
)

type StudentService interface {
	Create(name string, schoolId uint) (*models.Student, error)
}

type studentService struct {
	StudentRepo repository.StudentRepository
	schoolRepo  repository.SchoolRepository
	userRepo    repository.UserRepository
}

func NewStudentService(studentRepo repository.StudentRepository, schoolRepo repository.SchoolRepository, userRepo repository.UserRepository) StudentService {
	return &studentService{
		studentRepo,
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
	err = s.StudentRepo.Create(student)
	if err != nil {
		return nil, err
	}

	err = s.schoolRepo.Create(school)
	if err != nil {
		return nil, err
	}

	return student, nil
}
