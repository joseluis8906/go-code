package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/joseluis8906/go-code/protobuf/delivery/customerpb"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/apicustomer"
	"github.com/joseluis8906/go-code/src/delivery/internal/app/apistoremanager"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Params struct {
		fx.In

		Config *viper.Viper
		Logger *log.Logger

		CustomerServer     *apicustomer.GRPCServer
		StoreManagerServer *apistoremanager.GRPCServer
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

			customerpb.RegisterCustomerServer(grpcServer, params.CustomerServer)
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
