package repository

import (
	"errors"

	"github.com/omidnasiri/mediana-sms/internal/models"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Create(*models.Teacher) error
	GetById(id uint) (*models.Teacher, error)
	GetStudentsById(id uint) ([]*models.Student, error)
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

func (r *teacherRepository) GetById(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.Where("id = ?", id).First(&teacher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("teacher")
		}
		return nil, err
	}

	return &teacher, nil
}

func (r *teacherRepository) GetStudentsById(id uint) ([]*models.Student, error) {
	var teacher models.Teacher
	err := r.db.Preload("Students").Where("id = ?", id).First(&teacher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("teacher")
		}
		return nil, err
	}

	return teacher.Students, nil
}
