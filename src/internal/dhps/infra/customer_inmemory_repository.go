package infra

import (
	"context"
	"sync"

	"github.com/joseluis8906/go-code/src/pkg/delivery"
)

// CustomerInMemoryRepository represents an in-memory repository for customers.
type CustomerInMemoryRepository struct {
	customers map[string]delivery.Customer
	mux       sync.Mutex
}

// NewCustomerInMemoryRepository returns a new instance of CustomerInMemoryRepository.
func NewCustomerInMemoryRepository() *CustomerInMemoryRepository {
	return &CustomerInMemoryRepository{
		customers: map[string]delivery.Customer{
			"ellie.hang@example.com": delivery.NewCustomer("ellie.hang@example.com", "Ellie Hang", delivery.NewAddress("211 Southside Square", city)),
		},
	}
}

func (r *CustomerInMemoryRepository) Get(ctx context.Context, email string) delivery.Customer {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.customers[email]
}
