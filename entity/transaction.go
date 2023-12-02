package entity

import "time"

// Transaction represents a transaction in the system
type Transaction struct {
	ID        uint64 `gorm:"primary_key"`
	UserID    uint64
	ProductID uint64
	Quantity  int
	CreatedAt time.Time `gorm:"auto_created_at"`
	UpdatedAt time.Time `gorm:"auto_updated_at"`
}
