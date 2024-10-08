package main

import (
	"log"
	"net"

	pb "github.com/kinyarasam/vendly/api/proto"
	"github.com/kinyarasam/vendly/internal/vendor"
	"google.golang.org/grpc"
)

func main() {
	repo := vendor.NewInMemoryVendorRepository()
	vendorServer := vendor.NewServer(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVendorServiceServer(s, vendorServer)

	log.Println("Vendor service is running on port 50051...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
