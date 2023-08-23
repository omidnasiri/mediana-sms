package repository

import (
	"github.com/omidnasiri/mediana-sms/internal/models"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	Create(model *models.School) error
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
