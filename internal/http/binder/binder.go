package binder

import (
	"depublic/entity"

	"github.com/labstack/echo/v4"
)

// BindUser binds a `User` from an HTTP request
func BindUser(c echo.Context) (*entity.User, error) {
	user := &entity.User{}
	err := c.Bind(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
