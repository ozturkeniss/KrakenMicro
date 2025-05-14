package handler

import (
	"context"
	"errors"

	pb "gomicro/api/proto"
	"gomicro/internal/product/service"
)

// ProductGRPCHandler handles gRPC requests for products
type ProductGRPCHandler struct {
	pb.UnimplementedProductServiceServer
	productService service.ProductService
}

// NewProductGRPCHandler creates a new gRPC handler for products
func NewProductGRPCHandler(productService service.ProductService) *ProductGRPCHandler {
	return &ProductGRPCHandler{
		productService: productService,
	}
}

// GetProduct implements the GetProduct gRPC method
func (h *ProductGRPCHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	product, err := h.productService.GetProduct(ctx, uint(req.ProductId))
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          uint32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
	}, nil
}

// GetProducts implements the GetProducts gRPC method
func (h *ProductGRPCHandler) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	var products []*pb.Product
	for _, id := range req.ProductIds {
		product, err := h.productService.GetProduct(ctx, uint(id))
		if err != nil {
			continue
		}
		products = append(products, &pb.Product{
			Id:          uint32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       int32(product.Stock),
		})
	}

	return &pb.GetProductsResponse{
		Products: products,
	}, nil
}

// CreateProduct implements the CreateProduct gRPC method
func (h *ProductGRPCHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	product, err := h.productService.CreateProduct(ctx, req.Name, req.Description, req.Price, int(req.Stock))
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          uint32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
	}, nil
}

// UpdateProduct implements the UpdateProduct gRPC method
func (h *ProductGRPCHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	product, err := h.productService.UpdateProduct(ctx, uint(req.Id), req.Name, req.Description, req.Price, int(req.Stock))
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          uint32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
	}, nil
}

// DeleteProduct implements the DeleteProduct gRPC method
func (h *ProductGRPCHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	err := h.productService.DeleteProduct(ctx, uint(req.Id))
	if err != nil {
		return &pb.DeleteProductResponse{Success: false}, err
	}

	return &pb.DeleteProductResponse{Success: true}, nil
}

// ListProducts implements the ListProducts gRPC method
func (h *ProductGRPCHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	products, err := h.productService.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	var pbProducts []*pb.Product
	for _, p := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       int32(p.Stock),
		})
	}

	return &pb.ListProductsResponse{
		Products: pbProducts,
	}, nil
} 