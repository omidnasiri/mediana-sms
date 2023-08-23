package repository

import (
	"github.com/omidnasiri/mediana-sms/internal/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(*models.Student) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewstudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db,
	}
}

func (r *studentRepository) Create(model *models.Student) error {
	return r.db.Create(model).Error
}
