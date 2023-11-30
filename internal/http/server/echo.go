package server

import (
	"context"
	"depublic-app/internal/http/router"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/depublic/depublic/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	// Load the configuration
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	// Create the Echo instance
	e := echo.New()

	// Set the Logger middleware
	e.Use(middleware.Logger())

	// Setup the routes
	router.SetupRoutes(e, config)

	// Start the server
	go func() {
		err := e.Start(fmt.Sprintf(":%s", config.Port))
		if err != nil {
			fmt.Println("Failed to start the server: %v", err)
			os.Exit(1)
		}
	}()

	// Wait for SIGINT or SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	// Gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = e.Shutdown(ctx)
	if err != nil {
		fmt.Println("Failed to shutdown the server: %v", err)
		os.Exit(1)
	}

	fmt.Println("Server stopped")
}
