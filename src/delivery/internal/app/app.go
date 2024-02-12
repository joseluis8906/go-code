package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/joseluis8906/go-code/protobuf/delivery/customerpb"
	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"

	"github.com/joseluis8906/go-code/src/delivery/internal/app/logging"
	customer "github.com/joseluis8906/go-code/src/delivery/internal/customer/grpc"
	storemanager "github.com/joseluis8906/go-code/src/delivery/internal/storemanager/grpc"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Deps struct {
		fx.In

		Config *viper.Viper
		Log    *log.Logger

		CustomerServer     *customer.GRPCServer
		StoreManagerServer *storemanager.GRPCServer
	}
)

func NewGRPCServer(lc fx.Lifecycle, deps Deps) *grpc.Server {
	grpcServer := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", deps.Config.GetInt("grpc.port")))
			if err != nil {
				return err
			}

			customerpb.RegisterCustomerServer(grpcServer, deps.CustomerServer)
			storemanagerpb.RegisterStoreManagerServer(grpcServer, deps.StoreManagerServer)

			reflection.Register(grpcServer)

			go func() {
				err := grpcServer.Serve(lis)
				if err != nil {
					deps.Log.Printf("%v starting grpc server: %v", logging.Error, err)
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
