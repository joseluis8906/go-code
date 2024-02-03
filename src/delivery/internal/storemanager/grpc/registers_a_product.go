package grpc

import (
	"context"

	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"
	"github.com/joseluis8906/go-code/src/pkg/grpc"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) RegistersAProduct(ctx context.Context, req *storemanagerpb.RegistersAProductRequest) (*emptypb.Empty, error) {
	email, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		return nil, err
	}

	theStoremanager := storemanager.StoreManager{
		Email:  email,
		Stores: s.registry.Stores,
	}

	pform := make([]storemanager.ProductForm, len(req.GetStore().GetProducts()))
	for i, p := range req.GetStore().GetProducts() {
		pform[i] = storemanager.ProductForm{
			Ref:  p.GetRef().GetValue(),
			Name: p.GetName().GetValue(),
			Price: storemanager.MoneyForm{
				Amount:   p.GetPrice().GetAmount().GetValue(),
				Currency: p.GetPrice().GetCurrency().GetValue(),
			},
		}
	}

	form := storemanager.StoreForm{
		Name:     req.GetStore().GetName().GetValue(),
		Products: pform,
	}

	return &emptypb.Empty{}, theStoremanager.RegistersAProduct(ctx, form)
}
