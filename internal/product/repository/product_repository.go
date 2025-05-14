package repository

import (
	"context"

	"gorm.io/gorm"
	"gomicro/internal/product/model"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	Update(ctx context.Context, product *model.Product) (*model.Product, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*model.Product, error)
}

// productRepository implements the ProductRepository interface
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

// Create creates a new product
func (r *productRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// GetByID retrieves a product by ID
func (r *productRepository) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// Update updates an existing product
func (r *productRepository) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	if err := r.db.WithContext(ctx).Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// Delete deletes a product by ID
func (r *productRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Product{}, id).Error
}

// List retrieves all products
func (r *productRepository) List(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
} 