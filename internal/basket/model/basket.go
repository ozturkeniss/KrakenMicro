package model

import (
	"encoding/json"
	"time"
)

type BasketItem struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Name      string  `json:"name"`
}

type Basket struct {
	ID        uint         `json:"id"`
	UserID    uint         `json:"user_id"`
	Items     []BasketItem `json:"items"`
	Total     float64      `json:"total"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// Redis için JSON dönüşüm metodları
func (b *Basket) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Basket) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, b)
} 