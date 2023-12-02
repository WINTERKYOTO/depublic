package entity

import "time"

// Product represents a product in the system
type Product struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	Description string
	Price       int
	Quantity    int
	CreatedAt   time.Time `gorm:"auto_created_at"`
	UpdatedAt   time.Time `gorm:"auto_updated_at"`
}
