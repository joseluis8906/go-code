package app

import (
	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/waiter"
)

type (
	// RepositoryRegistry represents the repository registry.
	RepositoryRegistry struct {
		Waiters   *waiter.Repository
		Customers *customer.Repository
		Catalog   *product.Repository
	}
)
