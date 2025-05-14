package handler

import (
	"context"
	"time"

	pb "gomicro/api/proto"
	"gomicro/internal/payment/service"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.PaymentResponse, error) {
	payment, err := h.service.ProcessPayment(ctx, uint(req.UserId), req.Amount, req.Currency, req.PaymentMethod)
	if err != nil {
		return nil, err
	}

	return &pb.PaymentResponse{
		PaymentId: uint32(payment.ID),
		UserId:    uint32(payment.UserID),
		Amount:    payment.Amount,
		Currency:  payment.Currency,
		Status:    payment.Status,
		CreatedAt: payment.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (h *PaymentHandler) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.PaymentResponse, error) {
	payment, err := h.service.GetPayment(ctx, uint(req.PaymentId))
	if err != nil {
		return nil, err
	}

	return &pb.PaymentResponse{
		PaymentId: uint32(payment.ID),
		UserId:    uint32(payment.UserID),
		Amount:    payment.Amount,
		Currency:  payment.Currency,
		Status:    payment.Status,
		CreatedAt: payment.CreatedAt.Format(time.RFC3339),
	}, nil
} 