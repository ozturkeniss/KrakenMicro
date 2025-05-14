package tests

import (
	"context"
	"testing"
	"time"

	"gomicro/internal/basket/model"
	"gomicro/internal/basket/service"
)

// MockBasketRepository implements repository.BasketRepository interface
type MockBasketRepository struct {
	baskets map[uint]*model.Basket
}

func NewMockBasketRepository() *MockBasketRepository {
	return &MockBasketRepository{
		baskets: make(map[uint]*model.Basket),
	}
}

func (m *MockBasketRepository) Create(ctx context.Context, basket *model.Basket) error {
	basket.ID = uint(len(m.baskets) + 1)
	basket.CreatedAt = time.Now()
	basket.UpdatedAt = time.Now()
	m.baskets[basket.ID] = basket
	return nil
}

func (m *MockBasketRepository) GetByID(ctx context.Context, id uint) (*model.Basket, error) {
	if basket, exists := m.baskets[id]; exists {
		return basket, nil
	}
	return nil, nil
}

func (m *MockBasketRepository) GetByUserID(ctx context.Context, userID uint) (*model.Basket, error) {
	for _, basket := range m.baskets {
		if basket.UserID == userID {
			return basket, nil
		}
	}
	return nil, nil
}

func (m *MockBasketRepository) Update(ctx context.Context, basket *model.Basket) error {
	if _, exists := m.baskets[basket.ID]; exists {
		basket.UpdatedAt = time.Now()
		m.baskets[basket.ID] = basket
		return nil
	}
	return nil
}

func (m *MockBasketRepository) Delete(ctx context.Context, id uint) error {
	if _, exists := m.baskets[id]; exists {
		delete(m.baskets, id)
		return nil
	}
	return nil
}

func (m *MockBasketRepository) DeleteBasket(ctx context.Context, userID uint) error {
	delete(m.baskets, userID)
	return nil
}

func (m *MockBasketRepository) GetBasket(ctx context.Context, userID uint) (*model.Basket, error) {
	return m.GetByID(ctx, userID)
}

func (m *MockBasketRepository) SaveBasket(ctx context.Context, basket *model.Basket) error {
	m.baskets[basket.ID] = basket
	return nil
}

func TestCreateBasket(t *testing.T) {
	tests := []struct {
		name     string
		userID   uint
		wantErr  bool
	}{
		{
			name:     "valid basket",
			userID:   1,
			wantErr:  false,
		},
		{
			name:     "zero user ID",
			userID:   0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			repo := NewMockBasketRepository()
			basketService := service.NewBasketService(repo)

			// Execute
			basket, err := basketService.CreateBasket(context.Background(), tt.userID)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("CreateBasket() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("CreateBasket() unexpected error: %v", err)
				return
			}

			if basket == nil {
				t.Error("CreateBasket() returned nil basket")
				return
			}

			// Verify basket details
			if basket.UserID != tt.userID {
				t.Errorf("CreateBasket() userID = %v, want %v", basket.UserID, tt.userID)
			}
			if len(basket.Items) != 0 {
				t.Errorf("CreateBasket() items = %v, want empty slice", basket.Items)
			}
		})
	}
}

func TestGetBasket(t *testing.T) {
	// Setup
	repo := NewMockBasketRepository()
	basketService := service.NewBasketService(repo)

	// Create a test basket
	testBasket := &model.Basket{
		UserID: 1,
		Items:  []model.BasketItem{},
	}
	repo.Create(context.Background(), testBasket)

	tests := []struct {
		name      string
		basketID  uint
		wantErr   bool
	}{
		{
			name:      "existing basket",
			basketID:  1,
			wantErr:   false,
		},
		{
			name:      "non-existing basket",
			basketID:  999,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			basket, err := basketService.GetBasket(context.Background(), tt.basketID)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetBasket() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("GetBasket() unexpected error: %v", err)
				return
			}

			if tt.basketID == 1 {
				if basket == nil {
					t.Error("GetBasket() returned nil basket for existing ID")
					return
				}
				if basket.ID != tt.basketID {
					t.Errorf("GetBasket() basket ID = %v, want %v", basket.ID, tt.basketID)
				}
			} else {
				if basket != nil {
					t.Error("GetBasket() returned basket for non-existing ID")
				}
			}
		})
	}
}

func TestAddItemToBasket(t *testing.T) {
	// Setup
	repo := NewMockBasketRepository()
	basketService := service.NewBasketService(repo)

	// Create a test basket
	testBasket := &model.Basket{
		UserID: 1,
		Items:  []model.BasketItem{},
	}
	repo.Create(context.Background(), testBasket)

	tests := []struct {
		name      string
		basketID  uint
		productID uint
		quantity  int
		wantErr   bool
	}{
		{
			name:      "valid item",
			basketID:  1,
			productID: 1,
			quantity:  2,
			wantErr:   false,
		},
		{
			name:      "non-existing basket",
			basketID:  999,
			productID: 1,
			quantity:  1,
			wantErr:   true,
		},
		{
			name:      "zero quantity",
			basketID:  1,
			productID: 1,
			quantity:  0,
			wantErr:   true,
		},
		{
			name:      "negative quantity",
			basketID:  1,
			productID: 1,
			quantity:  -1,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			err := basketService.AddItemToBasket(context.Background(), tt.basketID, tt.productID, tt.quantity)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("AddItemToBasket() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("AddItemToBasket() unexpected error: %v", err)
				return
			}

			// Verify item was added
			basket, err := repo.GetByID(context.Background(), tt.basketID)
			if err != nil {
				t.Errorf("Failed to get basket: %v", err)
				return
			}

			if basket == nil {
				t.Error("AddItemToBasket() basket was not found")
				return
			}

			// Find the added item
			var found bool
			for _, item := range basket.Items {
				if item.ProductID == tt.productID {
					found = true
					if item.Quantity != tt.quantity {
						t.Errorf("AddItemToBasket() quantity = %v, want %v", item.Quantity, tt.quantity)
					}
					break
				}
			}

			if !found {
				t.Error("AddItemToBasket() item was not added")
			}
		})
	}
}

func TestRemoveItemFromBasket(t *testing.T) {
	// Setup
	repo := NewMockBasketRepository()
	basketService := service.NewBasketService(repo)

	// Create a test basket with an item
	testBasket := &model.Basket{
		UserID: 1,
		Items: []model.BasketItem{
			{
				ProductID: 1,
				Quantity:  2,
			},
		},
	}
	repo.Create(context.Background(), testBasket)

	tests := []struct {
		name      string
		basketID  uint
		productID uint
		wantErr   bool
	}{
		{
			name:      "existing item",
			basketID:  1,
			productID: 1,
			wantErr:   false,
		},
		{
			name:      "non-existing basket",
			basketID:  999,
			productID: 1,
			wantErr:   true,
		},
		{
			name:      "non-existing item",
			basketID:  1,
			productID: 999,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			err := basketService.RemoveItemFromBasket(context.Background(), tt.basketID, tt.productID)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("RemoveItemFromBasket() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("RemoveItemFromBasket() unexpected error: %v", err)
				return
			}

			if tt.basketID == 1 {
				// Verify item was removed
				basket, err := repo.GetByID(context.Background(), tt.basketID)
				if err != nil {
					t.Errorf("Failed to get basket: %v", err)
					return
				}

				if basket == nil {
					t.Error("RemoveItemFromBasket() basket was not found")
					return
				}

				// Check if item still exists
				for _, item := range basket.Items {
					if item.ProductID == tt.productID {
						t.Error("RemoveItemFromBasket() item was not removed")
						break
					}
				}
			}
		})
	}
} 