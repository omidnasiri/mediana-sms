package service

import (
	"github.com/omidnasiri/mediana-sms/internal/models"
	"github.com/omidnasiri/mediana-sms/internal/repository"
	"github.com/omidnasiri/mediana-sms/pkg/utils"
)

type TeacherService interface {
	Create(name, email, password string, schoolId uint) (*models.Teacher, error)
	GetStudents(teacherUserId uint) ([]*models.Student, error)
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

func (s *teacherService) Create(name, email, password string, schoolId uint) (*models.Teacher, error) {
	_, err := s.schoolRepo.GetById(schoolId)
	if err != nil {
		return nil, err
	}

	passwordHash, err := utils.HashPassword([]byte(password))
	if err != nil {
		return nil, err
	}

	role := &models.Role{Title: models.ROLE_TEACHER}
	user := &models.User{Name: name, Email: email, PasswordHash: passwordHash, SchoolId: &schoolId, Role: role}
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	teacher := &models.Teacher{User: user}
	err = s.teacherRepo.Create(teacher)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (s *teacherService) GetStudents(teacherUserId uint) ([]*models.Student, error) {
	students, err := s.teacherRepo.GetStudentsByTeacherUserId(teacherUserId)
	if err != nil {
		return nil, err
	}
	return students, nil
}
