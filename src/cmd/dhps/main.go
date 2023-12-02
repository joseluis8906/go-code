package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/joseluis8906/go-code/idl/dhpspb"
	"github.com/joseluis8906/go-code/src/internal/dhps/app"
	"github.com/joseluis8906/go-code/src/internal/dhps/infra"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	fx.In

	DeliveryService *app.DeliveryService
}

const port = 80

func main() {
	fx.New(
		fx.Options(
			app.Module,
			infra.Module,
		),
		fx.Provide(
			NewGRPCServer,
		),
		fx.Invoke(func(*grpc.Server) {}),
	).Run()
}

// NewGRPCServer creates a new gRPC server.
func NewGRPCServer(lc fx.Lifecycle, params Params) *grpc.Server {
	grpcServer := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
			if err != nil {
				return err
			}

			dhpspb.RegisterDeliveryServer(grpcServer, params.DeliveryService)
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
