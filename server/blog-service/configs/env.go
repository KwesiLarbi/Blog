package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// function for loading mongo environmental variable
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}