package grpc

import (
	"context"

	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *GRPCServer) RegistersAStore(ctx context.Context, req *storemanagerpb.RegistersAStoreRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	}

	email := md.Get("x-auth-email")
	if len(email) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing x-auth-email header")
	}

	theStoremanager := storemanager.StoreManager{
		Email:  email[0],
		Stores: s.registry.Stores,
	}

	form := storemanager.StoreForm{
		Name:    req.GetStore().GetName().GetValue(),
		Country: req.GetStore().GetCountry().GetValue(),
		City:    req.GetStore().GetCity().GetValue(),
		Address: req.GetStore().GetAddress().GetValue(),
	}

	return &emptypb.Empty{}, theStoremanager.RegistersAStore(ctx, form)
}
