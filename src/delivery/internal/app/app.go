package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	ll "github.com/joseluis8906/go-code/src/delivery/internal/app/log"
	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"

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

			storemanagerpb.RegisterStoreManagerServiceServer(grpcServer, deps.StoreManagerServer)

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

func NewHTTPServer(lc fx.Lifecycle, deps Deps) *http.Server {
	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    ":9090",
		Handler: handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			fmt.Println("Starting HTTP server at", srv.Addr)

			go srv.Serve(ln)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
