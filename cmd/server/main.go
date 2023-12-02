package main

import (
	"database/sql"
	"depublic/internal/config"
	"depublic/internal/http/handler" // Pastikan path sesuai dengan struktur proyek Anda

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Connect to database
	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		panic(err)
	}

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	handler.RegisterRoutes(e, db) // Panggil RegisterRoutes dengan parameter yang sesuai

	// Start server
	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
