package repository

import (
	"github.com/omidnasiri/mediana-sms/internal/models"

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
	return nil, nil
}
