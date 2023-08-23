package repository

import (
	"github.com/omidnasiri/mediana-sms/internal/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(*models.Student) error
	GetListByIDsExcludeTeacherID(ids []uint, teacherId uint) ([]*models.Student, error)
	BulkUpdate(students []*models.Student) error
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

func (r *studentRepository) GetListByIDsExcludeTeacherID(ids []uint, teacherId uint) ([]*models.Student, error) {
	models := make([]*models.Student, 0, len(ids)+1)
	err := r.db.Raw("select * from students where id in ? and id not in (select id from students join teacher_students ts on students.id = ts.student_id where ts.teacher_id = ?)", ids, teacherId).Find(&models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (r *studentRepository) BulkUpdate(students []*models.Student) error {
	err := r.db.Save(&students).Error
	if err != nil {
		return err
	}
	return nil
}
