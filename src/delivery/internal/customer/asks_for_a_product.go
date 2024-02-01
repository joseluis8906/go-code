package customer

import (
	"context"

	"github.com/joseluis8906/go-code/src/delivery/internal/product"
)

type (
	Waiter interface {
		LooksForAProduct(ctx context.Context, productName product.Name) error
	}
)

func (c *Customer) AsksForAProduct(ctx context.Context, productName product.Name, aWaiter Waiter) error {
	return aWaiter.LooksForAProduct(ctx, productName)
}
