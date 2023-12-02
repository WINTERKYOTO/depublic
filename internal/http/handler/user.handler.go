package handler

import (
	"depublic/entity"
	"depublic/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserController adalah handler untuk endpoint '/users'
type UserController struct {
	service service.UserService
}

// NewUserController membuat `UserController` baru
func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

// GetAll menangani endpoint 'GET /users'
func (c *UserController) GetAll(ctx echo.Context) error {
	// Dapatkan semua pengguna
	users, err := c.service.GetAll()
	if err != nil {
		return err
	}

	// Respon dengan daftar pengguna
	return ctx.JSON(200, users)
}

// GetByID menangani endpoint 'GET /users/:id'
func (c *UserController) GetByID(ctx echo.Context) error {
	// Dapatkan ID pengguna dari path sebagai string
	userID := ctx.Param("id")

	// Konversi userID dari string ke uint64 (jika dibutuhkan)
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err // Tindakan yang sesuai jika konversi gagal
	}

	// Dapatkan pengguna berdasarkan ID yang diharapkan dalam tipe string
	user, err := c.service.GetByID(strconv.FormatUint(userIDUint, 10))
	if err != nil {
		return err
	}

	// Respon dengan pengguna yang didapatkan
	return ctx.JSON(200, user)
}

// Create menangani endpoint 'POST /users'
func (c *UserController) Create(ctx echo.Context) error {
	// Bind request
	user := new(entity.User)
	if err := ctx.Bind(user); err != nil {
		return err
	}

	// Buat pengguna baru
	err := c.service.Create(user)
	if err != nil {
		return err
	}

	// Respon dengan pengguna yang telah dibuat
	return ctx.JSON(201, user)
}

// Update menangani endpoint 'PUT /users/:id'
func (c *UserController) Update(ctx echo.Context) error {
	// Dapatkan ID pengguna dari path sebagai string
	userID := ctx.Param("id")

	// Konversi userID dari string ke uint64 (jika dibutuhkan)
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err // Tindakan yang sesuai jika konversi gagal
	}

	// Bind request
	user := new(entity.User)
	if err := ctx.Bind(user); err != nil {
		return err
	}

	// Setel ID pengguna ke dalam entitas user
	user.ID = userIDUint

	// Lakukan pembaruan pengguna menggunakan entitas user
	err = c.service.Update(user)
	if err != nil {
		return err
	}

	// Respon dengan pengguna yang telah diperbarui
	return ctx.JSON(200, user)
}
