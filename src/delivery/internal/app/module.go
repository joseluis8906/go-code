package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/waiter"
)

// Module exports the module for app.
var Module = fx.Provide(
	NewCustomerService,

	waiter.NewRepository,
	customer.NewRepository,
	product.NewRepository,
)
