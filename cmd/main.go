package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/omidnasiri/mediana-sms/api"
	"github.com/omidnasiri/mediana-sms/cmd/di"
	"github.com/omidnasiri/mediana-sms/pkg/db"
)

func main() {
	// TODO: Setup app environments and load only in dev environment
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// database connection
	dbObj := db.Migrate()

	// dependency injection
	controllers := di.Inject(dbObj)

	// start api server
	r := api.SetupRoutes(controllers)
	r.Run()
}
