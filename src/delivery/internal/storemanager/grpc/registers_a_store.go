package grpc

import (
	"context"

	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"
	"github.com/joseluis8906/go-code/src/pkg/grpc"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) RegistersAStore(ctx context.Context, req *storemanagerpb.RegistersAStoreRequest) (*emptypb.Empty, error) {
	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		return nil, err
	}

	theStoremanager := storemanager.StoreManager{
		Email:  email,
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
