package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/joseluis8906/go-code/protobuf/delivery/customerpb"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Params struct {
		fx.In

		Config          *viper.Viper
		Logger          *log.Logger
		CustomerService *CustomerService
	}
)

func NewGRPCServer(lc fx.Lifecycle, params Params) *grpc.Server {
	grpcServer := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", params.Config.GetInt("grpc.port")))
			if err != nil {
				return err
			}

			customerpb.RegisterCustomerServer(grpcServer, params.CustomerService)
			reflection.Register(grpcServer)

			go func() {
				err := grpcServer.Serve(lis)
				if err != nil {
					params.Logger.Printf("starting grpc server: %w", err)
				}
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
