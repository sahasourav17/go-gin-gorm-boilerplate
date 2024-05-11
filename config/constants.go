package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST string
	DB_NAME string
	DB_USER string
	DB_PASS string
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
}
