package router

import (
	"depublic-app/internal/config"
	"net/http"

	"github.com/depublic/depublic/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// SetupRoutes sets up the routes for the application.
func SetupRoutes(r *gin.Engine, config *config.Config) {
	// Auth routes
	r.POST("/auth/login", handler.AuthHandler{}.LoginHandler)
	r.POST("/auth/register", handler.AuthHandler{}.RegisterHandler)

	// User routes
	r.GET("/users/:id", handler.UserHandler{}.GetUserHandler)
	r.PUT("/users/:id", handler.UserHandler{}.UpdateUserHandler)
	r.DELETE("/users/:id", handler.UserHandler{}.DeleteUserHandler)

	// Protected routes
	r.Use(jwtAuthMiddleware(config))

	// Example route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})
}

// jwtAuthMiddleware is a middleware that checks for a valid JWT token in the request header.
func jwtAuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request header
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set the user ID in the context
		c.Set("userID", token.Claims.Subject)
	}
}
