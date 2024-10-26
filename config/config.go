package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig holds database connection settings
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// LoadConfig loads application configuration from .env file
func LoadConfig() *DBConfig {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Return a DBConfig struct filled with environment variable values
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}
