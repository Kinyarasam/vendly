package main

import (
	"log"
	"net"

	pb "github.com/kinyarasam/vendly/api/proto"
	"github.com/kinyarasam/vendly/internal/vendor"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVendorServiceServer(s, &vendor.Server{})

	log.Println("Vendor service is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
