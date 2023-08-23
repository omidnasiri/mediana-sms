package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/omidnasiri/mediana-sms/api"
	"github.com/omidnasiri/mediana-sms/cmd/di"
	"github.com/omidnasiri/mediana-sms/pkg/db"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// database connection
	dbObj := db.Migrate()

	// dependency injection
	handlers := di.Inject(dbObj)

	// start api server
	r := api.SetupRoutes(handlers)
	r.Run()
}
