package entity

import "time"

// User represents a user in the system
type User struct {
	ID        uint64 `gorm:"primary_key"`
	Username  string `gorm:"unique"`
	Password  string
	Role      string
	CreatedAt time.Time `gorm:"auto_created_at"`
	UpdatedAt time.Time `gorm:"auto_updated_at"`
}
