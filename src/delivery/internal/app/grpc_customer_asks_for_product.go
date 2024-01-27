package app

import (
	"context"
	"fmt"

	"github.com/joseluis8906/go-code/src/delivery/internal/customer"
	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/waiter"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"github.com/joseluis8906/go-code/src/pkg/gglpb"

	pb "github.com/joseluis8906/go-code/protobuf/delivery/customerpb"
)

const appName = "delivery"

// CustomerAsksForAProduct returns a list of suggestions.
func (s *CustomerService) AsksForAProduct(ctx context.Context, req *pb.AsksForAProductReq) (*pb.AsksForAProductRes, error) {
	email, err := customer.NewEmail(req.GetCustomerEmail().GetValue())
	if err != nil {
		return nil, fmt.Errorf("validating customer email: %w", err)
	}

	aProduct, err := product.NewName(req.GetProductName().GetValue())
	if err != nil {
		return nil, fmt.Errorf("validating product name: %w", err)
	}

	criteria := cmp.Equals(customer.EmailField, email.String())
	theCustomer, err := s.registry.Customers.Matching(ctx, criteria).ExpectOne()
	if err != nil {
		return nil, fmt.Errorf("getting customer: %w", err)
	}

	criteria = cmp.Equals(waiter.NameField, appName)
	theWaiter, err := s.registry.Waiters.Matching(ctx, criteria).ExpectOne()
	if err != nil {
		return nil, fmt.Errorf("getting waiter: %w", err)
	}

	err = theCustomer.AsksFor(aProduct).
		To(theWaiter.Using(s.registry.Catalog)).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("asking for a product: %w", err)
	}

	suggestions := theWaiter.Suggests(ctx)

	products := make([]*pb.Product, len(suggestions))
	for i, suggestion := range suggestions {
		product := pb.Product{
			Ref:  gglpb.String(suggestion.Ref.String()),
			Name: gglpb.String(suggestion.Name.String()),
		}

		products[i] = &product
	}

	res := &pb.AsksForAProductRes{
		Products: products,
	}

	return res, nil
}
