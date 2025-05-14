package handler

import (
	"context"
	"errors"
	"time"

	pb "gomicro/api/proto"
	"gomicro/internal/basket/model"
	"gomicro/internal/basket/service"
)

type BasketGRPCHandler struct {
	pb.UnimplementedBasketServiceServer
	basketService service.IBasketService
}

func NewBasketGRPCHandler(basketService service.IBasketService) *BasketGRPCHandler {
	return &BasketGRPCHandler{
		basketService: basketService,
	}
}

func (h *BasketGRPCHandler) GetBasket(ctx context.Context, req *pb.GetBasketRequest) (*pb.Basket, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	basket, err := h.basketService.GetBasket(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return convertToProtoBasket(basket), nil
}

func (h *BasketGRPCHandler) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.Basket, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	err := h.basketService.AddItemToBasket(ctx, uint(req.UserId), uint(req.ProductId), int(req.Quantity))
	if err != nil {
		return nil, err
	}
	basket, err := h.basketService.GetBasket(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return convertToProtoBasket(basket), nil
}

func (h *BasketGRPCHandler) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.Basket, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	err := h.basketService.AddItemToBasket(ctx, uint(req.UserId), uint(req.ProductId), int(req.Quantity))
	if err != nil {
		return nil, err
	}
	basket, err := h.basketService.GetBasket(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return convertToProtoBasket(basket), nil
}

func (h *BasketGRPCHandler) RemoveItem(ctx context.Context, req *pb.RemoveItemRequest) (*pb.Basket, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	err := h.basketService.RemoveItemFromBasket(ctx, uint(req.UserId), uint(req.ProductId))
	if err != nil {
		return nil, err
	}
	basket, err := h.basketService.GetBasket(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return convertToProtoBasket(basket), nil
}

func (h *BasketGRPCHandler) ClearBasket(ctx context.Context, req *pb.ClearBasketRequest) (*pb.ClearBasketResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	err := h.basketService.ClearBasket(ctx, uint(req.UserId))
	if err != nil {
		return &pb.ClearBasketResponse{Success: false}, err
	}
	return &pb.ClearBasketResponse{Success: true}, nil
}

func convertToProtoBasket(basket *model.Basket) *pb.Basket {
	protoBasket := &pb.Basket{
		UserId:    uint32(basket.UserID),
		Total:     basket.Total,
		UpdatedAt: basket.UpdatedAt.Format(time.RFC3339),
		Items:     make([]*pb.BasketItem, len(basket.Items)),
	}

	for i, item := range basket.Items {
		protoBasket.Items[i] = &pb.BasketItem{
			ProductId: uint32(item.ProductID),
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
			Name:      item.Name,
		}
	}

	return protoBasket
} 