package utils

import (
	"log"
	
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}