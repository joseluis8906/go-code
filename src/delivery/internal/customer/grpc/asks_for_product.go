package grpc

import (
	"context"
	"fmt"

	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/waiter"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"github.com/joseluis8906/go-code/src/pkg/gglpb"

	"github.com/joseluis8906/go-code/protobuf/delivery/customerpb"
	"github.com/joseluis8906/go-code/protobuf/delivery/deliverypb"
)

const appName = "delivery"

// CustomerAsksForAProduct returns a list of suggestions.
func (s *GRPCServer) AsksForAProduct(ctx context.Context, req *customerpb.AsksForAProductRequest) (*customerpb.AsksForAProductResponse, error) {
	email, err := customer.NewEmail(req.GetCustomer().GetEmail().GetValue())
	if err != nil {
		return nil, fmt.Errorf("validating customer email: %w", err)
	}

	productName, err := product.NewName(req.GetProduct().GetName().GetValue())
	if err != nil {
		return nil, fmt.Errorf("validating product name: %w", err)
	}

	criteria := cmp.Equals(customer.EmailField, email.Value)
	theCustomer, err := s.registry.Customers.Matching(ctx, criteria).ExpectOne()
	if err != nil {
		return nil, fmt.Errorf("getting customer: %w", err)
	}

	criteria = cmp.Equals(waiter.NameField, appName)
	theWaiter, err := s.registry.Waiters.Matching(ctx, criteria).ExpectOne()
	if err != nil {
		return nil, fmt.Errorf("getting waiter: %w", err)
	}

	theWaiter.TakesACatalog(s.registry.Catalog)

	err = theCustomer.AsksForAProduct(ctx, productName, &theWaiter)
	if err != nil {
		return nil, fmt.Errorf("asking for a product: %w", err)
	}

	suggestions := theWaiter.Suggests(ctx)

	products := make([]*deliverypb.Product, len(suggestions))
	for i, suggestion := range suggestions {
		products[i] = &deliverypb.Product{
			Ref:  gglpb.String(suggestion.Ref.Value),
			Name: gglpb.String(suggestion.Name.Value),
		}
	}

	res := &customerpb.AsksForAProductResponse{
		// Products: products,
	}

	return res, nil
}
