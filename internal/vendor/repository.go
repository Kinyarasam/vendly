package vendor

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type InMemoryVendorRepository struct {
	vendors map[string]*Vendor
	mu      sync.RWMutex
}

func NewInMemoryVendorRepository() *InMemoryVendorRepository {
	return &InMemoryVendorRepository{
		vendors: make(map[string]*Vendor),
	}
}

func (r *InMemoryVendorRepository) SaveVendor(vendor *Vendor) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	vendor.ID = generateID()
	vendor.CreatedAt = time.Now()
	vendor.UpdatedAt = time.Now()
	r.vendors[vendor.ID] = vendor
	return nil
}

func (r *InMemoryVendorRepository) GetVendor(id string) (*Vendor, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	vendor, exists := r.vendors[id]
	if !exists {
		return nil, ErrVendorNotFound
	}
	return vendor, nil
}

var ErrVendorNotFound = errors.New("vendor not found")

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
