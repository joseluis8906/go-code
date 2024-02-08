package grpc

import (
	"context"
	"log"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/registry"
	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"

	"github.com/joseluis8906/go-code/src/pkg/grpc"

	"go.uber.org/fx"
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

// RegistersAStore parses the request and calls the RegistersAStore method of the
// storemanager.StoreManager domain actor.
// It returns an empty response or an error if the actor fails to execute the method.
// If the gRPC request does not contain a valid x-auth-email header it returns and error.
func (s *GRPCServer) RegistersAStore(ctx context.Context, req *storemanagerpb.RegistersAStoreRequest) (*emptypb.Empty, error) {
	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		return nil, err
	}

	theStoremanager := storemanager.StoreManager{
		Email:  email,
		Stores: s.registry.Stores,
	}

	form := storemanager.StoreForm{
		Name:    req.GetStore().GetName().GetValue(),
		Country: req.GetStore().GetCountry().GetValue(),
		City:    req.GetStore().GetCity().GetValue(),
		Address: req.GetStore().GetAddress().GetValue(),
	}

	return &emptypb.Empty{}, theStoremanager.RegistersAStore(ctx, form)
}

// RegistersProducts parses the request and calls the RegistersProducts method of the
// storemanager.StoreManager domain actor.
// It returns an empty response or an error if the actor fails to execute the method.
// If the gRPC request does not contain a valid x-auth-email header it returns and error.
func (s *GRPCServer) RegistersProducts(ctx context.Context, req *storemanagerpb.RegistersProductsRequest) (*emptypb.Empty, error) {
	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		return nil, err
	}

	theStoremanager := storemanager.StoreManager{
		Email:  email,
		Stores: s.registry.Stores,
	}

	pform := make([]storemanager.ProductForm, len(req.GetStore().GetProducts()))
	for i, p := range req.GetStore().GetProducts() {
		pform[i] = storemanager.ProductForm{
			Ref:  p.GetRef().GetValue(),
			Name: p.GetName().GetValue(),
			Price: storemanager.MoneyForm{
				Amount:   p.GetPrice().GetAmount().GetValue(),
				Currency: p.GetPrice().GetCurrency().GetValue(),
			},
		}
	}

	form := storemanager.StoreForm{
		Name:     req.GetStore().GetName().GetValue(),
		Products: pform,
	}

	return &emptypb.Empty{}, theStoremanager.RegistersProducts(ctx, form)
}
