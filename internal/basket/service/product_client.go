package service

import (
	"context"
	"log"

	pb "gomicro/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(address string) (*ProductClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to product service: %v", err)
		return nil, err
	}

	client := pb.NewProductServiceClient(conn)
	return &ProductClient{client: client}, nil
}

func (c *ProductClient) GetProduct(ctx context.Context, productID uint32) (*pb.Product, error) {
	resp, err := c.client.GetProduct(ctx, &pb.GetProductRequest{
		ProductId: productID,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProductClient) GetProducts(ctx context.Context, productIDs []uint32) ([]*pb.Product, error) {
	resp, err := c.client.GetProducts(ctx, &pb.GetProductsRequest{
		ProductIds: productIDs,
	})
	if err != nil {
		return nil, err
	}
	return resp.Products, nil
} 