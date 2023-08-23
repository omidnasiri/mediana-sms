package repository

import (
	"github.com/omidnasiri/mediana-sms/internal/models"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Create(*models.Teacher) error
}

type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{
		db,
	}
}

func (r *teacherRepository) Create(model *models.Teacher) error {
	return r.db.Create(model).Error
}
