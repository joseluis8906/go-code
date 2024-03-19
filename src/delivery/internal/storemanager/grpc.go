package storemanager

import (
	"context"
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
	ctx, span := otel.Tracer("").Start(ctx, "storemanager.StoreManager.AddStore")
	defer span.End()

	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		s.Log.Printf("getting x-auth-email: %q", err)

		return nil, err
	}

	sm := StoreManager{
		Email:  email,
		Stores: s.Stores,
	}

	err = sm.AddStore(ctx, req)
	if err != nil {
		s.Log.Printf("adding store: %q", err)
	}

	return &epb.Empty{}, err
}

// AddProduct parses the request and calls the AddProduct method of the
// storemanager.StoreManager domain actor.
// It returns an empty response or an error if the actor fails to execute the method.
// If the gRPC request does not contain a valid x-auth-email header it returns and error.
func (s *GRPCServer) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*epb.Empty, error) {
	ctx, span := otel.Tracer("").Start(ctx, "storemanager.AddProduct")
	defer span.End()

	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		s.Log.Printf("getting x-auth-email: %q", err)

		return nil, err
	}

	theStoremanager := StoreManager{
		Email:  email,
		Stores: s.Stores,
	}

	err = theStoremanager.AddProduct(ctx, req)
	if err != nil {
		s.Log.Printf("adding product: %q", err)
	}

	return &epb.Empty{}, err
}
