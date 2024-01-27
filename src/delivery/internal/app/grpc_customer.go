package app

import (
	"go.uber.org/fx"

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

		registry *RepositoryRegistry
	}
)

// NewDeliveryService returns a new instance of DeliveryService.
func NewCustomerService(deps ServiceDeps) *CustomerService {
	return &CustomerService{}
}
