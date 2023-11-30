package service

import (
	"context"

	"github.com/depublic/depublic/internal/repository"
)

// UserService is a struct that holds the handlers for user-related requests.
type UserService struct {
	repo repository.UserRepository
}

// NewUserService returns a new UserService instance.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetUserHandler handles the `GET /users/:id` request.
func (s *UserService) GetUserHandler(ctx context.Context, userID uint64) (*User, error) {
	user, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserHandler handles the `PUT /users/:id` request.
func (s *UserService) UpdateUserHandler(ctx context.Context, userID uint64, updateRequest *UpdateUserRequest) (*User, error) {
	user, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	user.FullName = updateRequest.FullName
	err = s.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUserHandler handles the `DELETE /users/:id` request.
func (s *UserService) DeleteUserHandler(ctx context.Context, userID uint64) error {
	if err := s.repo.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}
