package main

import (
	"log"

	"github.com/gsh519/suggestion/suggestion"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	suggestion.Suggest()
}
