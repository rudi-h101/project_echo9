package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetHost() (string, []string) {
	app_host := os.Getenv("APP_URL")
	temp := strings.Split(app_host, "://")
	schema := []string{temp[0]}
	host := temp[1]

	return host, schema
}

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