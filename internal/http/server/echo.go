package server

import (
	"depublic/internal/config"
	"depublic/internal/http/router" // Import paket router yang berisi RegisterRoutes

	"github.com/labstack/echo/v4"
)

// Start starts the HTTP server
func Start(cfg *config.Config) {
	// Create Echo instance
	e := echo.New()

	// Register routes dari paket router
	router.RegisterRoutes(e, nil) // Panggil RegisterRoutes dari paket router

	// Start server
	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
