package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Stock       int           `gorm:"not null" json:"stock"`
	Category    string         `gorm:"not null" json:"category"`
	ImageURL    string         `json:"image_url"`
	IsActive    bool          `gorm:"default:true" json:"is_active"`
} 