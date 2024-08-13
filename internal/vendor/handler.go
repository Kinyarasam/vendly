package vendor

import (
	"context"
	"errors"

	pb "github.com/kinyarasam/vendly/api/proto"
)

type Server struct {
	pb.UnimplementedVendorServiceServer
	repo *InMemoryVendorRepository
}

func NewServer(repo *InMemoryVendorRepository) *Server {
	return &Server{repo: repo}
}

func (s *Server) RegisterVendor(ctx context.Context, req *pb.RegisterVendorRequest) (*pb.RegisterVendorResponse, error) {
	vendor := &Vendor{
		Name:  req.Name,
		Email: req.Email,
	}

	err := s.repo.SaveVendor(vendor)
	if err != nil {
		return nil, errors.New("could not save vendor")
	}

	return &pb.RegisterVendorResponse{
		Success:  true,
		VendorId: vendor.ID,
	}, err
}
