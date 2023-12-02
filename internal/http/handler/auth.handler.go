package handler

import (
	"depublic/common"
	"depublic/service"

	"github.com/labstack/echo/v4"
)

// AuthController is the handler for the `/auth` endpoints
type AuthController struct {
	service service.AuthService
}

// NewAuthController creates a new `AuthController`
func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}

// Define the BindUser function
func BindUser(ctx echo.Context) (*common.User, error) {
	user := new(common.User)
	if err := ctx.Bind(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login handles the `POST /auth/login` endpoint
func (c *AuthController) Login(ctx echo.Context) error {
	// Bind the request
	user := new(common.User)
	if err := ctx.Bind(user); err != nil {
		return err
	}

	// Authenticate the user
	claims, err := c.service.Authenticate(user.Username, user.Password)
	if err != nil {
		return err
	}

	// Generate a token
	token, err := common.GenerateToken(claims)
	if err != nil {
		return err
	}

	// Respond with the token
	return ctx.JSON(200, echo.Map{"token": token})
}
