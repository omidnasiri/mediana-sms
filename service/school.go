package service

import (
	"github.com/omidnasiri/mediana-sms/internal/models"
	"github.com/omidnasiri/mediana-sms/internal/repository"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"
)

type SchoolService interface {
	Create(title string, headmasterId uint) (*models.School, error)
}

type schoolService struct {
	schoolRepo repository.SchoolRepository
	userRepo   repository.UserRepository
}

func NewSchoolService(schoolRepo repository.SchoolRepository, userRepo repository.UserRepository) SchoolService {
	return &schoolService{
		schoolRepo,
		userRepo,
	}
}

func (s *schoolService) Create(title string, headmasterId uint) (*models.School, error) {
	user, err := s.userRepo.GetById(headmasterId)
	if err != nil {
		return nil, err
	}
	if user.Role.Title != models.ROLE_HEADMASTER {
		return nil, errs.NewForbiddenError("provided user_id is doesn't is not a headmaster")
	}
	if user.SchoolId != nil {
		return nil, errs.NewForbiddenError("provided headmaster is already involved with another school")
	}

	school := &models.School{Title: title, HeadmasterId: headmasterId}

	err = s.schoolRepo.Create(school)
	if err != nil {
		return nil, err
	}

	return school, nil
}
