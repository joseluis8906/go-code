package grpc

import (
	"log"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/registry"
	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"
	"go.uber.org/fx"
)

type Deps struct {
	fx.In

	Log      *log.Logger
	Registry *registry.Repository
}

// NewDeliveryService returns a new instance of DeliveryService.
func NewStoreManager(deps Deps) *storemanager.GRPCServer {
	return &storemanager.GRPCServer{
		Log:    deps.Log,
		Stores: deps.Registry.Stores,
	}
}
