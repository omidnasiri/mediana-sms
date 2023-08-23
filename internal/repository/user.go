package repository

import (
	"errors"

	"github.com/omidnasiri/mediana-sms/internal/models"
	errs "github.com/omidnasiri/mediana-sms/pkg/err"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("user")
		}
		return nil, err
	}

	return &user, nil
}
