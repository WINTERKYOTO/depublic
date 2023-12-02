package router

import (
	"depublic/internal/http/handler/auth"
	"depublic/internal/http/handler/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterRoutes registers all the HTTP routes
func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	// Auth
	e.POST("/auth/login", auth.Login(db))

	// Users
	e.GET("/users", user.GetAll(db))
	e.GET("/users/:id", user.GetByID(db))
	e.POST("/users", user.Create(db))
	e.PUT("/users/:id", user.Update(db))
}
