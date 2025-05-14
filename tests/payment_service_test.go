package tests

import (
	"context"
	"testing"
	"time"

	"gomicro/internal/payment/model"
	"gomicro/internal/payment/service"
)

// MockPaymentRepository implements repository.PaymentRepository interface
type MockPaymentRepository struct {
	payments map[uint]*model.Payment
}

func NewMockPaymentRepository() *MockPaymentRepository {
	return &MockPaymentRepository{
		payments: make(map[uint]*model.Payment),
	}
}

func (m *MockPaymentRepository) Create(ctx context.Context, payment *model.Payment) error {
	payment.ID = uint(len(m.payments) + 1)
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	m.payments[payment.ID] = payment
	return nil
}

func (m *MockPaymentRepository) GetByID(ctx context.Context, id uint) (*model.Payment, error) {
	if payment, exists := m.payments[id]; exists {
		return payment, nil
	}
	return nil, nil
}

func (m *MockPaymentRepository) Update(ctx context.Context, payment *model.Payment) error {
	if _, exists := m.payments[payment.ID]; exists {
		payment.UpdatedAt = time.Now()
		m.payments[payment.ID] = payment
		return nil
	}
	return nil
}

// MockRabbitMQPublisher implements RabbitMQ publisher interface
type MockRabbitMQPublisher struct {
	messages []*model.StockUpdateEvent
}

func NewMockRabbitMQPublisher() *MockRabbitMQPublisher {
	return &MockRabbitMQPublisher{
		messages: make([]*model.StockUpdateEvent, 0),
	}
}

func (m *MockRabbitMQPublisher) SendStockUpdateEvent(event *model.StockUpdateEvent) error {
	m.messages = append(m.messages, event)
	return nil
}

func (m *MockRabbitMQPublisher) Close() {}

func TestProcessPayment(t *testing.T) {
	tests := []struct {
		name          string
		userID        uint
		amount        float64
		currency      string
		paymentMethod string
		wantErr       bool
	}{
		{
			name:          "successful payment",
			userID:        1,
			amount:        100.0,
			currency:      "TRY",
			paymentMethod: "credit_card",
			wantErr:       false,
		},
		{
			name:          "invalid amount",
			userID:        1,
			amount:        0,
			currency:      "TRY",
			paymentMethod: "credit_card",
			wantErr:       true,
		},
		{
			name:          "negative amount",
			userID:        1,
			amount:        -50.0,
			currency:      "TRY",
			paymentMethod: "credit_card",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			repo := NewMockPaymentRepository()
			publisher := NewMockRabbitMQPublisher()
			paymentService := service.NewPaymentService(repo, publisher)

			// Execute
			payment, err := paymentService.ProcessPayment(context.Background(), tt.userID, tt.amount, tt.currency, tt.paymentMethod)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("ProcessPayment() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("ProcessPayment() unexpected error: %v", err)
				return
			}

			if payment == nil {
				t.Error("ProcessPayment() returned nil payment")
				return
			}

			// Verify payment details
			if payment.UserID != tt.userID {
				t.Errorf("ProcessPayment() userID = %v, want %v", payment.UserID, tt.userID)
			}
			if payment.Amount != tt.amount {
				t.Errorf("ProcessPayment() amount = %v, want %v", payment.Amount, tt.amount)
			}
			if payment.Currency != tt.currency {
				t.Errorf("ProcessPayment() currency = %v, want %v", payment.Currency, tt.currency)
			}
			if payment.PaymentMethod != tt.paymentMethod {
				t.Errorf("ProcessPayment() paymentMethod = %v, want %v", payment.PaymentMethod, tt.paymentMethod)
			}
			if payment.Status != "completed" {
				t.Errorf("ProcessPayment() status = %v, want %v", payment.Status, "completed")
			}

			// Verify RabbitMQ message
			if len(publisher.messages) != 1 {
				t.Errorf("Expected 1 RabbitMQ message, got %d", len(publisher.messages))
				return
			}

			event := publisher.messages[0]
			if event.ProductID != 1 {
				t.Errorf("Expected ProductID = 1, got %d", event.ProductID)
			}
			if event.Quantity != -1 {
				t.Errorf("Expected Quantity = -1, got %d", event.Quantity)
			}
		})
	}
}

func TestGetPayment(t *testing.T) {
	// Setup
	repo := NewMockPaymentRepository()
	publisher := NewMockRabbitMQPublisher()
	paymentService := service.NewPaymentService(repo, publisher)

	// Create a test payment
	testPayment := &model.Payment{
		UserID:        1,
		Amount:        100.0,
		Currency:      "TRY",
		Status:        "completed",
		PaymentMethod: "credit_card",
	}
	repo.Create(context.Background(), testPayment)

	tests := []struct {
		name      string
		paymentID uint
		wantErr   bool
	}{
		{
			name:      "existing payment",
			paymentID: 1,
			wantErr:   false,
		},
		{
			name:      "non-existing payment",
			paymentID: 999,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			payment, err := paymentService.GetPayment(context.Background(), tt.paymentID)

			// Assert
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetPayment() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("GetPayment() unexpected error: %v", err)
				return
			}

			if tt.paymentID == 1 {
				if payment == nil {
					t.Error("GetPayment() returned nil payment for existing ID")
					return
				}
				if payment.ID != tt.paymentID {
					t.Errorf("GetPayment() payment ID = %v, want %v", payment.ID, tt.paymentID)
				}
			} else {
				if payment != nil {
					t.Error("GetPayment() returned payment for non-existing ID")
				}
			}
		})
	}
} 