package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the configuration for the application
type Config struct {
	Database struct {
		DSN string
	}
	Server struct {
		Port string
	}
}

// Load loads the configuration from the `.env` file
func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	// Database configuration
	cfg.Database.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	// Server configuration
	cfg.Server.Port = os.Getenv("PORT")

	log.Println("Loaded configuration")
	return cfg, nil
}
