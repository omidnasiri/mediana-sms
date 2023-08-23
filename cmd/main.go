package main

import "github.com/joho/godotenv"

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
