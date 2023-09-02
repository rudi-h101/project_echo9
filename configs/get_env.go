package config

import (
	"os"
)

func GetPort() string {
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8000"
}

