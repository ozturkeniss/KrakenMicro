package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	UserID        uint          `gorm:"not null" json:"user_id"`
	Amount        float64       `gorm:"not null" json:"amount"`
	Currency      string        `gorm:"not null" json:"currency"`
	Status        string        `gorm:"not null" json:"status"`
	PaymentMethod string        `gorm:"not null" json:"payment_method"`
} 