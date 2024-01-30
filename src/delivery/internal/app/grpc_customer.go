package app

import (
	"go.uber.org/fx"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/registry"

	pb "github.com/joseluis8906/go-code/protobuf/delivery/customerpb"
)

type (
	// ServiceDeps represents the dependencies of DeliveryService.
	ServiceDeps struct {
		fx.In
	}

	// DeliveryService represents a delivery service.
	CustomerService struct {
		pb.UnimplementedCustomerServer

		registry *registry.Repository
	}
)

// NewDeliveryService returns a new instance of DeliveryService.
func NewCustomerService(deps ServiceDeps) *CustomerService {
	return &CustomerService{}
}
