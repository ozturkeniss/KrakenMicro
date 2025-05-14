package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	pb "gomicro/api/proto"
	"gomicro/internal/basket/handler"
	"gomicro/internal/basket/repository"
	"gomicro/internal/basket/service"
)

func main() {
	// Redis connection
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test Redis connection
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	// Initialize repository
	repo := repository.NewBasketRepository(rdb)

	// Initialize service
	basketService := service.NewBasketService(repo)

	// Initialize gRPC handler
	basketHandler := handler.NewBasketGRPCHandler(basketService)

	// Create gRPC server
	server := grpc.NewServer()

	// Register service
	pb.RegisterBasketServiceServer(server, basketHandler)

	// Start server
	port := 8082
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Basket service is starting on port %d...\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 