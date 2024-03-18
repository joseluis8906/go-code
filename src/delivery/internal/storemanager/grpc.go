package storemanager

import (
	"context"
	"fmt"
	"log"

	pb "github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	epb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/joseluis8906/go-code/src/pkg/grpc"
	"go.opentelemetry.io/otel"
)

const authEmail string = "x-auth-email"

type (
	// DeliveryService represents a delivery service.
	GRPCServer struct {
		pb.UnimplementedStoreManagerServiceServer

		Log    *log.Logger
		Stores Stores
	}
)

// AddStore parses the request and calls the RegistersAStore method of the
// storemanager.StoreManager domain actor.
// It returns an empty response or an error if the actor fails to execute the method.
// If the gRPC request does not contain a valid x-auth-email header it returns and error.
func (s *GRPCServer) AddStore(ctx context.Context, req *pb.AddStoreRequest) (*epb.Empty, error) {
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

	return &epb.Empty{}, sm.AddStore(ctx, req)
}

// AddProduct parses the request and calls the AddProduct method of the
// storemanager.StoreManager domain actor.
// It returns an empty response or an error if the actor fails to execute the method.
// If the gRPC request does not contain a valid x-auth-email header it returns and error.
func (s *GRPCServer) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*epb.Empty, error) {
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

	return &epb.Empty{}, theStoremanager.AddProduct(ctx, req)
}
