package product

import (
	"context"
	"errors"

	pb "github.com/kinyarasam/vendly/api/proto"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	repo *InMemoryProductRepository
}

func NewServer(repo *InMemoryProductRepository) *Server {
	return &Server{repo: repo}
}

func (s *Server) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	product := &Product{
		VendorID:    req.VendorId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	err := s.repo.SaveProduct(product)
	if err != nil {
		return nil, errors.New("could not save product")
	}

	return &pb.AddProductResponse{
		Success:   true,
		ProductId: product.ID,
	}, nil
}
