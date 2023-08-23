package db

import (
	"github.com/omidnasiri/mediana-sms/internal/models"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

// SeedAll handles initial seeding of the data
func SeedAll() []Seed {
	return []Seed{
		{
			Name: "CreateAdminRole",
			Run: func(db *gorm.DB) error {
				if db.Table("roles").Where("title = ?", models.ROLE_ADMIN).First(&models.Role{}).RowsAffected == 0 {
					return db.Create(&models.Role{Title: models.ROLE_ADMIN}).Error
				}
				return nil
			},
		},
		{
			Name: "CreateHeadmasterRole",
			Run: func(db *gorm.DB) error {
				if db.Table("roles").Where("title = ?", models.ROLE_HEADMASTER).First(&models.Role{}).RowsAffected == 0 {
					return db.Create(&models.Role{Title: models.ROLE_HEADMASTER}).Error
				}
				return nil
			},
		},
		{
			Name: "CreateTeacherRole",
			Run: func(db *gorm.DB) error {
				if db.Table("roles").Where("title = ?", models.ROLE_TEACHER).First(&models.Role{}).RowsAffected == 0 {
					return db.Create(&models.Role{Title: models.ROLE_TEACHER}).Error
				}
				return nil
			},
		},
		{
			Name: "CreateStudentRole",
			Run: func(db *gorm.DB) error {
				if db.Table("roles").Where("title = ?", models.ROLE_STUDENT).First(&models.Role{}).RowsAffected == 0 {
					return db.Create(&models.Role{Title: models.ROLE_STUDENT}).Error
				}
				return nil
			},
		},
		{
			Name: "CreateAdmin",
			Run: func(db *gorm.DB) error {
				if db.Table("users").Where("email = ?", "hr@mediana.ir").First(&models.User{}).RowsAffected == 0 {
					return db.Create(&models.User{Name: "Mitra Hesami", Email: "hr@mediana.ir", PasswordHash: "$2a$04$gSwImOq4tHmLq410O4uv/.0j0ZEa571SA.mffMRmpPzFhICnlGWsq", RoleId: 1}).Error
				}
				return nil
			},
		},
	}
}
