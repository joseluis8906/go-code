package grpc

import (
	"log"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/registry"
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
)

const authEmail string = "x-auth-email"

type (
	Deps struct {
		fx.In

		Logs     *log.Logger
		Registry *registry.Repository
	}

	// DeliveryService represents a delivery service.
	GRPCServer struct {
		storemanagerpb.UnimplementedStoreManagerServer

		logs     *log.Logger
		registry *registry.Repository
	}
)

// NewDeliveryService returns a new instance of DeliveryService.
func NewGRPCServer(deps Deps) *GRPCServer {
	return &GRPCServer{
		logs:     deps.Logs,
		registry: deps.Registry,
	}
}
