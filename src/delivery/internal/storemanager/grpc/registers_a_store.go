package grpc

import (
	"context"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) RegistersAProduct(ctx context.Context, req *storemanagerpb.RegistersAProductRequest) (*emptypb.Empty, error) {
	return nil, nil
}
