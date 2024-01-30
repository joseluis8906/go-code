package apicustomer

import (
	"github.com/joseluis8906/go-code/src/delivery/internal/app/registry"
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/protobuf/delivery/customerpb"
)

type (
	Deps struct {
		fx.In

		Registry *registry.Repository
	}

	// DeliveryService represents a delivery service.
	GRPCServer struct {
		customerpb.UnimplementedCustomerServer

		registry *registry.Repository
	}
)

// NewDeliveryService returns a new instance of DeliveryService.
func NewGRPCServer(deps Deps) *GRPCServer {
	return &GRPCServer{
		registry: deps.Registry,
	}
}
