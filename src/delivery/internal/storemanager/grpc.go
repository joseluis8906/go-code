package storemanager

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/joseluis8906/go-code/src/pkg/grpc"
)

const authEmail string = "x-auth-email"

type (
	// DeliveryService represents a delivery service.
	GRPCServer struct {
		storemanagerpb.UnimplementedStoreManagerServer

		Log    *log.Logger
		Stores Stores
	}
)

// RegistersAStore parses the request and calls the RegistersAStore method of the
// storemanager.StoreManager domain actor.
// It returns an empty response or an error if the actor fails to execute the method.
// If the gRPC request does not contain a valid x-auth-email header it returns and error.
func (s *GRPCServer) RegistersAStore(ctx context.Context, req *storemanagerpb.RegistersAStoreRequest) (*emptypb.Empty, error) {
	ctx, span := otel.Tracer("").Start(ctx, "storemanager.StoreManager.RegistersAStore")
	defer span.End()

	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		return nil, fmt.Errorf("getting grpc x-auth-email header: %q", err)
	}

	sm := StoreManager{
		Email:  email,
		Stores: s.Stores,
	}

	return &emptypb.Empty{}, sm.RegistersAStore(ctx, req)
}

// RegistersProducts parses the request and calls the RegistersProducts method of the
// storemanager.StoreManager domain actor.
// It returns an empty response or an error if the actor fails to execute the method.
// If the gRPC request does not contain a valid x-auth-email header it returns and error.
func (s *GRPCServer) RegistersProducts(ctx context.Context, req *storemanagerpb.RegistersProductsRequest) (*emptypb.Empty, error) {
	ctx, span := otel.Tracer("").Start(ctx, "storemanager.RegistersProducts")
	defer span.End()

	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		return nil, err
	}

	theStoremanager := StoreManager{
		Email:  email,
		Stores: s.Stores,
	}

	return &emptypb.Empty{}, theStoremanager.RegistersProducts(ctx, req)
}
