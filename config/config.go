package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // for .env file support (optional)
)

// Config struct represents the configuration structure
type Config struct {
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	APIPort string
}

// LoadConfig loads configuration from the .env file and/or config file
func LoadConfig() (*Config, error) {
	// Load .env variables (optional)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var config Config
	config.APIPort = os.Getenv("API_PORT")
	return &config, nil
}
