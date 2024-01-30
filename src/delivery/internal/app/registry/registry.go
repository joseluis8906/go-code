package registry

import (
	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/waiter"

	"go.uber.org/fx"
)

type (
	Deps struct {
		fx.In

		Catalog   *product.Repository
		Customers *customer.Repository
		Waiters   *waiter.Repository
	}

	// RepositoryRegistry represents the repository registry.
	Repository struct {
		Customers *customer.Repository
		Catalog   *product.Repository
		Waiters   *waiter.Repository
	}
)

func New(deps Deps) *Repository {
	return &Repository{
		Catalog:   deps.Catalog,
		Customers: deps.Customers,
		Waiters:   deps.Waiters,
	}
}
