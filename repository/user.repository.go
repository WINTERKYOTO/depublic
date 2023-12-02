package repository

import (
	"depublic/entity"
	"fmt"

	"github.com/jinzhu/gorm"
)

// UserRepository is the interface for interacting with users in the repository
type UserRepository interface {
	GetAll() ([]*entity.User, error)
	GetByID(userID string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(userID string) error
}

// UserRepositoryImpl is the implementation of the `UserRepository` interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepositoryImpl creates a new `UserRepositoryImpl`
func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

// GetAll gets all users from the database
func (r *UserRepositoryImpl) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID gets a user by ID from the database
func (r *UserRepositoryImpl) GetByID(userID string) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %s not found", userID)
		}
		return nil, err
	}
	return &user, nil
}

// Create creates a new user in the database
func (r *UserRepositoryImpl) Create(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Update updates an existing user in the database
func (r *UserRepositoryImpl) Update(user *entity.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a user by ID from the database
func (r *UserRepositoryImpl) Delete(userID string) error {
	if err := r.db.Delete(&entity.User{}, "id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}
