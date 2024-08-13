package vendor

import (
	"context"
	"testing"

	pb "github.com/kinyarasam/vendly/api/proto"
	"github.com/stretchr/testify/assert"
)

func TestRegisterVendor(t *testing.T) {
	repo := NewInMemoryVendorRepository()
	server := NewServer(repo)

	req := &pb.RegisterVendorRequest{
		Name:  "Test Vendor",
		Email: "test@vendor.com",
	}

	res, err := server.RegisterVendor(context.Background(), req)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.NotEmpty(t, res.VendorId)

	savedVendor, err := repo.GetVendor((res.VendorId))
	assert.NoError(t, err)
	assert.Equal(t, "Test Vendor", savedVendor.Name)
	assert.Equal(t, "test@vendor.com", savedVendor.Email)
}
