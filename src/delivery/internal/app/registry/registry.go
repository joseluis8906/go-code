package registry

import (
	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/store"

	"go.uber.org/fx"
)

type (
	Deps struct {
		fx.In

		Customers *customer.Repository
		Stores    *store.Repository
	}

	// RepositoryRegistry represents the repository registry.
	Repository struct {
		Customers *customer.Repository
		Stores    *store.Repository
	}
)

func New(deps Deps) *Repository {
	return &Repository{
		Customers: deps.Customers,
		Stores:    deps.Stores,
	}
}
