package service

import (
	"context"
	"errors"

	"gomicro/internal/basket/model"
	"gomicro/internal/basket/repository"
)

type IBasketService interface {
	CreateBasket(ctx context.Context, userID uint) (*model.Basket, error)
	GetBasket(ctx context.Context, basketID uint) (*model.Basket, error)
	AddItemToBasket(ctx context.Context, basketID, productID uint, quantity int) error
	RemoveItemFromBasket(ctx context.Context, basketID, productID uint) error
	ClearBasket(ctx context.Context, basketID uint) error
}

type basketService struct {
	repo repository.BasketRepository
}

func NewBasketService(repo repository.BasketRepository) IBasketService {
	return &basketService{
		repo: repo,
	}
}

func (s *basketService) CreateBasket(ctx context.Context, userID uint) (*model.Basket, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	basket := &model.Basket{
		UserID: userID,
		Items:  []model.BasketItem{},
	}

	err := s.repo.Create(ctx, basket)
	if err != nil {
		return nil, err
	}

	return basket, nil
}

func (s *basketService) GetBasket(ctx context.Context, basketID uint) (*model.Basket, error) {
	return s.repo.GetByID(ctx, basketID)
}

func (s *basketService) AddItemToBasket(ctx context.Context, basketID, productID uint, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	basket, err := s.repo.GetByID(ctx, basketID)
	if err != nil {
		return err
	}

	if basket == nil {
		return errors.New("basket not found")
	}

	// Check if item already exists
	for i, item := range basket.Items {
		if item.ProductID == productID {
			// Update quantity
			basket.Items[i].Quantity = quantity
			return s.repo.Update(ctx, basket)
		}
	}

	// Add new item
	basket.Items = append(basket.Items, model.BasketItem{
		ProductID: productID,
		Quantity:  quantity,
	})

	return s.repo.Update(ctx, basket)
}

func (s *basketService) RemoveItemFromBasket(ctx context.Context, basketID, productID uint) error {
	basket, err := s.repo.GetByID(ctx, basketID)
	if err != nil {
		return err
	}

	if basket == nil {
		return errors.New("basket not found")
	}

	// Find and remove item
	for i, item := range basket.Items {
		if item.ProductID == productID {
			basket.Items = append(basket.Items[:i], basket.Items[i+1:]...)
			return s.repo.Update(ctx, basket)
		}
	}

	return nil
}

func (s *basketService) ClearBasket(ctx context.Context, basketID uint) error {
	basket, err := s.repo.GetByID(ctx, basketID)
	if err != nil {
		return err
	}
	if basket == nil {
		return errors.New("basket not found")
	}
	basket.Items = []model.BasketItem{}
	basket.Total = 0
	return s.repo.Update(ctx, basket)
} 