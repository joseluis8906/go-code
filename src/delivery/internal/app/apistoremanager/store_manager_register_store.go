package apistoremanager

import (
	"context"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) RegisterStore(ctx context.Context, req *storemanagerpb.RegisterStoreRequest) (*emptypb.Empty, error) {
	return nil, nil
}
