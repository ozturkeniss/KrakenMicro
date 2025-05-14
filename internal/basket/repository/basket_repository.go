package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gomicro/internal/basket/model"
)

type BasketRepository interface {
	GetBasket(ctx context.Context, userID uint) (*model.Basket, error)
	SaveBasket(ctx context.Context, basket *model.Basket) error
	DeleteBasket(ctx context.Context, userID uint) error

	// New methods for test/service compatibility
	Create(ctx context.Context, basket *model.Basket) error
	GetByID(ctx context.Context, basketID uint) (*model.Basket, error)
	Update(ctx context.Context, basket *model.Basket) error
}

type basketRepository struct {
	client *redis.Client
}

func NewBasketRepository(client *redis.Client) BasketRepository {
	return &basketRepository{client: client}
}

func (r *basketRepository) GetBasket(ctx context.Context, userID uint) (*model.Basket, error) {
	key := fmt.Sprintf("basket:%d", userID)
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			// Return empty basket if not found
			return &model.Basket{
				UserID:    userID,
				Items:     []model.BasketItem{},
				Total:     0,
				UpdatedAt: time.Now(),
			}, nil
		}
		return nil, err
	}

	var basket model.Basket
	if err := json.Unmarshal(data, &basket); err != nil {
		return nil, err
	}

	return &basket, nil
}

func (r *basketRepository) SaveBasket(ctx context.Context, basket *model.Basket) error {
	key := fmt.Sprintf("basket:%d", basket.UserID)
	data, err := json.Marshal(basket)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, 24*time.Hour).Err()
}

func (r *basketRepository) DeleteBasket(ctx context.Context, userID uint) error {
	key := fmt.Sprintf("basket:%d", userID)
	return r.client.Del(ctx, key).Err()
}

// New methods for test/service compatibility
func (r *basketRepository) Create(ctx context.Context, basket *model.Basket) error {
	return r.SaveBasket(ctx, basket)
}

func (r *basketRepository) GetByID(ctx context.Context, basketID uint) (*model.Basket, error) {
	return r.GetBasket(ctx, basketID)
}

func (r *basketRepository) Update(ctx context.Context, basket *model.Basket) error {
	return r.SaveBasket(ctx, basket)
} 