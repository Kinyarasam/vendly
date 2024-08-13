package vendor

import (
	"context"

	pb "github.com/kinyarasam/vendly/api/proto"
)

type Server struct {
	pb.UnimplementedVendorServiceServer
}

func (s *Server) RegisterVendor(ctx context.Context, req *pb.RegisterVendorRequest) (*pb.RegisterVendorResponse, error) {
	return &pb.RegisterVendorResponse{Success: true, VendorId: "1234"}, nil
}
