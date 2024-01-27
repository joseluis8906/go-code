package main

import (
  "go.uber.org/fx"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  
  "github.com/joseluis8906/go-code/idl/pongpb"
)

const port = 50051

type Params struct {
  fx.In

  PongService pongpb.PongService
}

func main() {
  fx.New(
    fx.Options(pong.Module),
    fx.Provide(NewGRPCServer),
    fx.Invoke(func(*grpc.Server) {}),
  ).Run()

}
func NewGRPCServer(lc fx.Lifecycle, params Params) *grpc.Server {
  grpcServer := grpc.NewServer()

  lc.Append(fx.Hook{
    OnStart: func(context.Context) error {
      lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
      if err != nil {
        return err
      }

      pongpb.RegisterDeliveryServer(grpcServer, params.PongService)
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
