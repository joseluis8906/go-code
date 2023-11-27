package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/joseluis8906/go-code/idl/dhpspb"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = 8080

func main() {
	fx.New(
		fx.Provide(NewGRPCServer),
		fx.Invoke(func(*grpc.Server) {}),
	).Run()
}

// NewGRPCServer creates a new gRPC server.
func NewGRPCServer(lc fx.Lifecycle) *grpc.Server {
	grpcServer := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
			if err != nil {
				return err
			}

			dhpspb.RegisterDhpsServer(grpcServer, NewServer())
			reflection.Register(grpcServer)

			log.Printf("Starting gRPC server: %v", grpcServer.GetServiceInfo())

			go func() {
				sErr := grpcServer.Serve(lis)
				log.Fatalf("failed to serve: %v", sErr)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()

			return nil
		},
	})

	return grpcServer
}

// NewServer creates a new server.
func NewServer() dhpspb.DhpsServer {
	return &Server{}
}

// Server is the server for the gRPC service.
type Server struct{}

// CustomerLooksForProduct is the implementation of the gRPC method.
func (s *Server) CustomerLooksForProduct(ctx context.Context, req *dhpspb.CustomerLooksForProductRequest) (*dhpspb.CustomerLooksForProductResponse, error) {
	return &dhpspb.CustomerLooksForProductResponse{Products: []*dhpspb.CustomerLooksForProductResponseProduct{{ProductId: 1, ProductName: "Chess Burger", StoreName: "McDonald's"}}}, nil
}
