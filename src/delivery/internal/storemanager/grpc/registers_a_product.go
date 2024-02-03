package grpc

import (
	"context"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"
	"github.com/joseluis8906/go-code/src/delivery/internal/storemanager"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) RegistersAProduct(ctx context.Context, req *storemanagerpb.RegistersAProductRequest) (*emptypb.Empty, error) {
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
