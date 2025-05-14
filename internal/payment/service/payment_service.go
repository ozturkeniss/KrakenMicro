package service

import (
	"context"
	"errors"
	"time"

	"gomicro/internal/payment/model"
	"gomicro/internal/payment/repository"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, userID uint, amount float64, currency, paymentMethod string) (*model.Payment, error)
	GetPayment(ctx context.Context, paymentID uint) (*model.Payment, error)
}

type paymentService struct {
	repo      repository.PaymentRepository
	publisher IRabbitMQPublisher
}

func NewPaymentService(repo repository.PaymentRepository, publisher IRabbitMQPublisher) PaymentService {
	return &paymentService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *paymentService) ProcessPayment(ctx context.Context, userID uint, amount float64, currency, paymentMethod string) (*model.Payment, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	payment := &model.Payment{
		UserID:        userID,
		Amount:        amount,
		Currency:      currency,
		Status:        "pending",
		PaymentMethod: paymentMethod,
	}

	if err := s.repo.Create(ctx, payment); err != nil {
		return nil, err
	}

	// Simulate payment processing
	time.Sleep(2 * time.Second)

	// Update payment status
	payment.Status = "completed"
	if err := s.repo.Update(ctx, payment); err != nil {
		return nil, err
	}

	// Send stock update event
	event := &model.StockUpdateEvent{
		ProductID: 1, // This should come from the request
		Quantity:  -1, // Decrease stock by 1
	}
	if err := s.publisher.SendStockUpdateEvent(event); err != nil {
		// Log the error but don't fail the payment
		// In a real system, you might want to handle this differently
		return nil, err
	}

	return payment, nil
}

func (s *paymentService) GetPayment(ctx context.Context, paymentID uint) (*model.Payment, error) {
	return s.repo.GetByID(ctx, paymentID)
} 