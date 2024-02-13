package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/joseluis8906/go-code/protobuf/delivery/customerpb"
	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"

	ll "github.com/joseluis8906/go-code/src/delivery/internal/app/log"
	customer "github.com/joseluis8906/go-code/src/delivery/internal/customer/grpc"
	storemanager "github.com/joseluis8906/go-code/src/delivery/internal/storemanager/grpc"

	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
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
	var tracerProvider *trace.TracerProvider
	var grpcServer *grpc.Server

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpointURL(deps.Config.GetString("otel.endpoint")))
			if err != nil {
				deps.Log.Fatalf("initializing otel trace grpc: %v", err)
			}

			tracerProvider = trace.NewTracerProvider(trace.WithBatcher(exp))
			otel.SetTracerProvider(tracerProvider)

			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", deps.Config.GetInt("grpc.port")))
			if err != nil {
				return err
			}

			grpcServer = grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))

			customerpb.RegisterCustomerServer(grpcServer, deps.CustomerServer)
			storemanagerpb.RegisterStoreManagerServer(grpcServer, deps.StoreManagerServer)

			reflection.Register(grpcServer)

			go func() {
				err := grpcServer.Serve(lis)
				if err != nil {
					deps.Log.Printf(ll.Error("starting grpc server: %v"), err)
				}
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()

			if err := tracerProvider.Shutdown(ctx); err != nil {
				deps.Log.Printf("shtdown tracer provider: %v", err)
			}

			return nil
		},
	})

	return grpcServer
}
