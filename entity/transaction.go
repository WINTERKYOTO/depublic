package entity

import (
	"time"
)

type Transaction struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	ProductID uint64    `json:"product_id"`
	UserID    uint64    `json:"user_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
