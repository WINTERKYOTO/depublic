package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// UserRepository is an interface that defines methods for interacting with users in a database.
type UserRepository interface {
	GetUser(ctx context.Context, userID uint64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userID uint64) error
}

// UserRepositoryImpl is a concrete implementation of the UserRepository interface.
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepositoryImpl returns a new UserRepositoryImpl instance.
func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// GetUser gets a user by ID.
func (repo *UserRepositoryImpl) GetUser(ctx context.Context, userID uint64) (*User, error) {
	query := `
		SELECT id, email, full_name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user User
	err := repo.db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// GetUserByEmail gets a user by email.
func (repo *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, full_name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user User
	err := repo.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil
}

// CreateUser creates a new user.
func (repo *UserRepositoryImpl) CreateUser(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (email, password, full_name)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	var userID uint64
	err = repo.db.QueryRowContext(ctx, query, user.Email, hashedPassword, user.FullName).Scan(&userID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = userID
	return nil
}

// UpdateUser updates a user.
func (repo *UserRepositoryImpl) UpdateUser(ctx context.Context, user *User) error {
	query := `
		UPDATE users
		SET email = $1, full_name = $2
		WHERE id = $3
	`

	_, err := repo.db.ExecContext(ctx, query, user.Email, user.FullName, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// DeleteUser deletes a user.
func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, userID uint64) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := repo.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
