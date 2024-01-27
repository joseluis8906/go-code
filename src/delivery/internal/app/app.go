package app

import (
	"context"
	"fmt"
	"net"

	pb "github.com/joseluis8906/go-code/protobuf/delivery/customerpb"

	"go.uber.org/fx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/nats-io/nats.go"
)

type (
	GRPCServer struct {
		Args []string

		Log interface {
			Printf(format string, v ...interface{})
		}

		Conf interface {
			GetString(key string) string
			GetInt(key string) int
		}
	}

	Params struct {
		fx.In
		CustomerService *CustomerService
	}
)

func (a *GRPCServer) Run(ctx context.Context) error {
	nc, err := nats.Connect(a.Conf.GetString("nats.url"))
	if err != nil {
		return fmt.Errorf("connecting nats: %w", err)
	}

	nc.Subscribe("foo", func(m *nats.Msg) {
		a.Log.Printf("message received: %s", m.Data)
	})

	nc.Publish("foo", []byte("Hello World"))

	ioc := fx.New(
		fx.Options(Module),
		fx.Provide(a.NewGRPCServer),
		fx.Invoke(func(*grpc.Server) {}),
	)

	err = ioc.Start(ctx)
	if err != nil {
		return fmt.Errorf("starting contained app: %w", err)
	}

	<-ctx.Done()

	err = ioc.Stop(context.TODO())
	if err != nil {
		return fmt.Errorf("stoping contained app: %w", err)
	}

	return nil
}

func (a *GRPCServer) NewGRPCServer(lc fx.Lifecycle, params Params) *grpc.Server {
	grpcServer := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", a.Conf.GetInt("grpc.port")))
			if err != nil {
				return err
			}

			pb.RegisterCustomerServer(grpcServer, params.CustomerService)
			reflection.Register(grpcServer)

			go func() {
				err := grpcServer.Serve(lis)
				if err != nil {
					a.Log.Printf("starting grpc server: %w", err)
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
