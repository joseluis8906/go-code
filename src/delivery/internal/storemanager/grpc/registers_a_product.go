package grpc

import (
	"context"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) RegistersAStore(ctx context.Context, req *storemanagerpb.RegistersAStoreRequest) (*emptypb.Empty, error) {
	return nil, nil
}
