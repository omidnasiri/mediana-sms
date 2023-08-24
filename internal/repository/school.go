package repository

import (
	"errors"

	"github.com/omidnasiri/mediana-sms/internal/models"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	Create(model *models.School) error
	GetById(id uint) (*models.School, error)
}

type schoolRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{
		db,
	}
}

func (r *schoolRepository) Create(model *models.School) error {
	return r.db.Create(model).Error
}

func (r *schoolRepository) GetById(id uint) (*models.School, error) {
	var school models.School
	err := r.db.Where("id = ?", id).First(&school).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("school")
		}
		return nil, err
	}

	return &school, nil
}
