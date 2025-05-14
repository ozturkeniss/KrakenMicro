package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"gomicro/internal/product/model"
	"gomicro/internal/product/service"
)

// MockProductRepository implements repository.ProductRepository interface
type MockProductRepository struct {
	products map[uint]*model.Product
}

func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products: make(map[uint]*model.Product),
	}
}

func (m *MockProductRepository) Create(ctx context.Context, product *model.Product) error {
	product.ID = uint(len(m.products) + 1)
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	m.products[product.ID] = product
	return nil
}

func (m *MockProductRepository) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	if product, exists := m.products[id]; exists {
		return product, nil
	}
	return nil, nil
}

func (m *MockProductRepository) Update(ctx context.Context, product *model.Product) error {
	if _, exists := m.products[product.ID]; exists {
		product.UpdatedAt = time.Now()
		m.products[product.ID] = product
		return nil
	}
	return nil
}

func (m *MockProductRepository) Delete(ctx context.Context, id uint) error {
	if _, exists := m.products[id]; exists {
		delete(m.products, id)
		return nil
	}
	return nil
}

func (m *MockProductRepository) List(ctx context.Context) ([]*model.Product, error) {
	products := make([]*model.Product, 0, len(m.products))
	for _, product := range m.products {
		products = append(products, product)
	}
	return products, nil
}

func (m *MockProductRepository) GetAll(ctx context.Context, offset int, limit int) ([]*model.Product, error) {
	var products []*model.Product
	count := 0
	for _, p := range m.products {
		if count >= offset && count < offset+limit {
			products = append(products, p)
		}
		count++
	}
	return products, nil
}

func (m *MockProductRepository) GetByCategory(ctx context.Context, category string, offset int, limit int) ([]*model.Product, error) {
	var products []*model.Product
	count := 0
	for _, p := range m.products {
		if count >= offset && count < offset+limit {
			products = append(products, p)
		}
		count++
	}
	return products, nil
}

func (m *MockProductRepository) UpdateStock(ctx context.Context, id uint, quantity int) error {
	if p, ok := m.products[id]; ok {
		p.Stock += quantity
		if p.Stock < 0 {
			p.Stock = 0
		}
		m.products[id] = p
		return nil
	}
	return errors.New("product not found")
}

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name        string
		product     *model.Product
		wantErr     bool
		checkFields bool
	}{
		{
			name: "valid product",
			product: &model.Product{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       100.0,
				Stock:       10,
			},
			wantErr:     false,
			checkFields: true,
		},
		{
			name: "zero price",
			product: &model.Product{
				Name:        "Zero Price Product",
				Description: "Test Description",
				Price:       0,
				Stock:       10,
			},
			wantErr:     false,
			checkFields: false,
		},
		{
			name: "negative stock",
			product: &model.Product{
				Name:        "Negative Stock Product",
				Description: "Test Description",
				Price:       100.0,
				Stock:       -5,
			},
			wantErr:     true,
			checkFields: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			repo := NewMockProductRepository()
			productService := service.NewProductService(repo)

			// Execute
			err := productService.CreateProduct(context.Background(), tt.product)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("CreateProduct() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("CreateProduct() unexpected error: %v", err)
				return
			}

			if tt.checkFields {
				// Verify product was created
				created, err := repo.GetByID(context.Background(), tt.product.ID)
				if err != nil {
					t.Errorf("Failed to get created product: %v", err)
					return
				}

				if created == nil {
					t.Error("CreateProduct() product was not created")
					return
				}

				// Verify fields
				if created.Name != tt.product.Name {
					t.Errorf("CreateProduct() name = %v, want %v", created.Name, tt.product.Name)
				}
				if created.Description != tt.product.Description {
					t.Errorf("CreateProduct() description = %v, want %v", created.Description, tt.product.Description)
				}
				if created.Price != tt.product.Price {
					t.Errorf("CreateProduct() price = %v, want %v", created.Price, tt.product.Price)
				}
				if created.Stock != tt.product.Stock {
					t.Errorf("CreateProduct() stock = %v, want %v", created.Stock, tt.product.Stock)
				}
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	// Setup
	repo := NewMockProductRepository()
	productService := service.NewProductService(repo)

	// Create a test product
	testProduct := &model.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100.0,
		Stock:       10,
	}
	repo.Create(context.Background(), testProduct)

	tests := []struct {
		name      string
		productID uint
		wantErr   bool
	}{
		{
			name:      "existing product",
			productID: 1,
			wantErr:   false,
		},
		{
			name:      "non-existing product",
			productID: 999,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			product, err := productService.GetProduct(context.Background(), tt.productID)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetProduct() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("GetProduct() unexpected error: %v", err)
				return
			}

			if tt.productID == 1 {
				if product == nil {
					t.Error("GetProduct() returned nil product for existing ID")
					return
				}
				if product.ID != tt.productID {
					t.Errorf("GetProduct() product ID = %v, want %v", product.ID, tt.productID)
				}
			} else {
				if product != nil {
					t.Error("GetProduct() returned product for non-existing ID")
				}
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	// Setup
	repo := NewMockProductRepository()
	productService := service.NewProductService(repo)

	// Create a test product
	testProduct := &model.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100.0,
		Stock:       10,
	}
	repo.Create(context.Background(), testProduct)

	tests := []struct {
		name    string
		product *model.Product
		wantErr bool
	}{
		{
			name: "valid update",
			product: &model.Product{
				ID:          1,
				Name:        "Updated Product",
				Description: "Updated Description",
				Price:       150.0,
				Stock:       20,
			},
			wantErr: false,
		},
		{
			name: "non-existing product",
			product: &model.Product{
				ID:          999,
				Name:        "Non-existing Product",
				Description: "Test Description",
				Price:       100.0,
				Stock:       10,
			},
			wantErr: true,
		},
		{
			name: "invalid price",
			product: &model.Product{
				ID:          1,
				Name:        "Invalid Price Product",
				Description: "Test Description",
				Price:       -50.0,
				Stock:       10,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			err := productService.UpdateProduct(context.Background(), tt.product)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("UpdateProduct() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("UpdateProduct() unexpected error: %v", err)
				return
			}

			if tt.product.ID == 1 {
				// Verify product was updated
				updated, err := repo.GetByID(context.Background(), tt.product.ID)
				if err != nil {
					t.Errorf("Failed to get updated product: %v", err)
					return
				}

				if updated == nil {
					t.Error("UpdateProduct() product was not found")
					return
				}

				// Verify fields
				if updated.Name != tt.product.Name {
					t.Errorf("UpdateProduct() name = %v, want %v", updated.Name, tt.product.Name)
				}
				if updated.Description != tt.product.Description {
					t.Errorf("UpdateProduct() description = %v, want %v", updated.Description, tt.product.Description)
				}
				if updated.Price != tt.product.Price {
					t.Errorf("UpdateProduct() price = %v, want %v", updated.Price, tt.product.Price)
				}
				if updated.Stock != tt.product.Stock {
					t.Errorf("UpdateProduct() stock = %v, want %v", updated.Stock, tt.product.Stock)
				}
			}
		})
	}
} 