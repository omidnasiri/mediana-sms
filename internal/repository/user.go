package repository

import (
	"errors"

	"github.com/omidnasiri/mediana-sms/internal/models"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*models.User) error
	GetByEmail(email string) (*models.User, error)
	GetById(id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) Create(model *models.User) error {
	return r.db.Create(model).Error
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("user")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("user")
		}
		return nil, err
	}

	return &user, nil
}
