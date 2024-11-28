package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     getEnvAsInt("DB_PORT", 5432), // Default to 5432 if not set
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)

	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Printf("Error converting %s to int: %v. Using default value: %d", key, err, defaultValue)
		return defaultValue
	}

	return value
}
