package service

import (
	"context"
	"errors"

	"gomicro/internal/product/model"
	"gomicro/internal/product/repository"
)

// ProductService defines the interface for product operations
type ProductService interface {
	GetProduct(ctx context.Context, id uint) (*model.Product, error)
	CreateProduct(ctx context.Context, name, description string, price float64, stock int) (*model.Product, error)
	UpdateProduct(ctx context.Context, id uint, name, description string, price float64, stock int) (*model.Product, error)
	DeleteProduct(ctx context.Context, id uint) error
	ListProducts(ctx context.Context) ([]*model.Product, error)
}

// productService implements the ProductService interface
type productService struct {
	repo repository.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

// GetProduct retrieves a product by ID
func (s *productService) GetProduct(ctx context.Context, id uint) (*model.Product, error) {
	return s.repo.GetByID(ctx, id)
}

// CreateProduct creates a new product
func (s *productService) CreateProduct(ctx context.Context, name, description string, price float64, stock int) (*model.Product, error) {
	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}
	if stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}

	product := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	return s.repo.Create(ctx, product)
}

// UpdateProduct updates an existing product
func (s *productService) UpdateProduct(ctx context.Context, id uint, name, description string, price float64, stock int) (*model.Product, error) {
	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}
	if stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Stock = stock

	return s.repo.Update(ctx, product)
}

// DeleteProduct deletes a product by ID
func (s *productService) DeleteProduct(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// ListProducts retrieves all products
func (s *productService) ListProducts(ctx context.Context) ([]*model.Product, error) {
	return s.repo.List(ctx)
} 