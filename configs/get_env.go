package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetPort() string {
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8000"
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}