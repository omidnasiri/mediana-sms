package service

import (
	"github.com/omidnasiri/mediana-sms/internal/models"
	"github.com/omidnasiri/mediana-sms/internal/repository"
)

type TeacherService interface {
	Create(name string, schoolId uint) (*models.Teacher, error)
	GetStudents(id uint) ([]*models.Student, error)
}

type teacherService struct {
	teacherRepo repository.TeacherRepository
	schoolRepo  repository.SchoolRepository
	userRepo    repository.UserRepository
}

func NewTeacherService(teacherRepo repository.TeacherRepository, schoolRepo repository.SchoolRepository, userRepo repository.UserRepository) TeacherService {
	return &teacherService{
		teacherRepo,
		schoolRepo,
		userRepo,
	}
}

func (s *teacherService) Create(name string, schoolId uint) (*models.Teacher, error) {
	school, err := s.schoolRepo.GetById(schoolId)
	if err != nil {
		return nil, err
	}

	user := &models.User{Name: name}
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	teacher := &models.Teacher{UserId: user.ID}
	err = s.teacherRepo.Create(teacher)
	if err != nil {
		return nil, err
	}

	err = s.schoolRepo.Create(school)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (s *teacherService) GetStudents(id uint) ([]*models.Student, error) {
	students, err := s.teacherRepo.GetStudentsById(id)
	if err != nil {
		return nil, err
	}
	return students, nil
}
