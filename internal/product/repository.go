package product

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type InMemoryProductRepository struct {
	products map[string]*Product
	mu       sync.RWMutex
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*Product),
	}
}

func (r *InMemoryProductRepository) SaveProduct(product *Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = generateID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	r.products[product.ID] = product
	return nil
}

func (r *InMemoryProductRepository) GetProduct(id string) (*Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, ErrProductNotFound
	}
	return product, nil
}

var ErrProductNotFound = errors.New("product not found")

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
