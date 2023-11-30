package builder

import (
	"fmt"

	"github.com/depublic/depublic/internal/config"
	"github.com/depublic/depublic/internal/repository"
	"github.com/jinzhu/gorm"
	"log"
)

// Build builds the entire database, including schema and seed data.
func Build(config *config.Config) error {
	log.Println("Building database...")

	// Create the database
	if err := createDatabase(config); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	// Create the tables
	if err := createTables(config); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	// Seed the data
	if err := seedData(config); err != nil {
		return fmt.Errorf("failed to seed data: %w", err)
	}

	log.Println("Database build successfully completed.")
	return nil
}

// createDatabase creates the database if it does not already exist.
func createDatabase(config *config.Config) error {
	// Get the database configuration
	config := config.Database

	// Connect to the database
	db, err := repository.NewDatabase(config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Create the database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `depublic`")
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	return nil
}

// createTables creates the tables in the database.
func createTables(config *config.Config) error {
	// Get the database configuration
	config := config.Database

	// Connect to the database
	db, err := repository.NewDatabase(config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Create the tables
	err = db.AutoMigrate(&entity.Product{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return fmt.Errorf("failed to migrate database schema: %w", err)
	}

	return nil
}

// seedData seeds initial data into the database.
func seedData(config *config.Config) error {
	// Get the database configuration
	config := config.Database

	// Connect to the database
	db, err := repository.NewDatabase(config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Seed products
