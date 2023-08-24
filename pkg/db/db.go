package db

import (
	"fmt"
	"log"
	"os"

	"github.com/omidnasiri/mediana-sms/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: Move dbcon into an interface to separate driver implementations

// Migrate opens the postgres database connection
// and migrates the database schema
func Migrate() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_NAME")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s", username, password, host, port, name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Role{}, &models.User{}, &models.Teacher{}, &models.Student{}, &models.School{})

	for _, seed := range SeedAll() {
		if err := seed.Run(db); err != nil {
			log.Fatal(err)
		}
	}

	return db
}
