package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/apicustomer"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/bus"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/config"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/logging"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/nosql"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/registry"

	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/waiter"
)

// Module exports the module for app.
var Module = fx.Provide(
	logging.New,
	config.New,
	bus.New,
	nosql.New,

	waiter.NewRepository,
	customer.NewRepository,
	product.NewRepository,

	registry.New,

	apicustomer.NewGRPCServer,
)
