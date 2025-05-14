package model

type StockUpdateEvent struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
} 