package builder

import (
	"depublic/entity"

	"gorm.io/gorm"
)

// UserBuilder is a builder for the `User` entity
type UserBuilder struct {
	db *gorm.DB
}

// NewUserBuilder creates a new `UserBuilder`
func NewUserBuilder(db *gorm.DB) *UserBuilder {
	return &UserBuilder{db}
}

// WithUsername sets the `Username` field
func (b *UserBuilder) WithUsername(username string) *UserBuilder {
	b.db.Where("username = ?", username)
	return b
}

// WithPassword sets the `Password` field
func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.db.Where("password = ?", password)
	return b
}

// WithRole sets the `Role` field
func (b *UserBuilder) WithRole(role string) *UserBuilder {
	b.db.Where("role = ?", role)
	return b
}

// FindAll finds all users that match the given criteria
func (b *UserBuilder) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	result := b.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindOne finds the first user that matches the given criteria
func (b *UserBuilder) FindOne() (*entity.User, error) {
	var user *entity.User
	result := b.db.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
