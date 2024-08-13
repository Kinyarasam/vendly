package main

import (
	"log"
	"net"

	pb "github.com/kinyarasam/vendly/api/proto"
	"github.com/kinyarasam/vendly/internal/product"
	"google.golang.org/grpc"
)

func main() {
	repo := product.NewInMemoryProductRepository()
	productServer := product.NewServer(repo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, productServer)

	log.Println("Product service is running on port 50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
