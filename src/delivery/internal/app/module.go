package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/bus"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/config"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/log"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/nosql"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/registry"

	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/store"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/grpc"
	grpccustomer "github.com/joseluis8906/go-code/src/delivery/internal/customer/grpc"
)

// Module exports the module for app.
var Module = fx.Provide(
	config.New,
	log.New,
	bus.New,
	nosql.New,

	customer.NewRepository,
	store.NewRepository,
	registry.New,

	grpc.NewStoreManager,
	grpccustomer.NewGRPCServer,
)
