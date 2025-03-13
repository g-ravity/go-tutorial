package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	port       string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
}

var Config *EnvConfig

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading.env file")
	}

	Config = &EnvConfig{
		port:       os.Getenv("PORT"),
		dbName:     os.Getenv("DB_NAME"),
		dbPort:     os.Getenv("DB_PORT"),
		dbHost:     os.Getenv("DB_HOST"),
		dbUser:     os.Getenv("DB_USER"),
		dbPassword: os.Getenv("DB_PASSWORD"),
	}
}
