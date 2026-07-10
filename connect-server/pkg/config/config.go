package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConnectionString string
	Port                     string
	Environment              string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading env file")
	}

	DatabaseConnectionString := getEnv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	Port := getEnv("PORT", "6969")
	// prod
	Environment := getEnv("ENVIRONMENT", "dev")

	return &Config{
		DatabaseConnectionString: DatabaseConnectionString,
		Port:                     Port,
		Environment:              Environment,
	}
}

func getEnv(key string, defaultValue string) string {
	res := os.Getenv(key)
	if res == "" {
		return defaultValue
	}
	return res
}
