package customer

import (
	"context"
	"fmt"
)

type (
	AsksForAProduct struct {
		product fmt.Stringer

		waiter interface {
			LooksForAProduct(ctx context.Context, product fmt.Stringer) error
		}
	}

	Waiter interface {
		LooksForAProduct(ctx context.Context, product fmt.Stringer) error
	}
)

func (a *AsksForAProduct) To(waiter Waiter) *AsksForAProduct {
	a.waiter = waiter
	return a
}

func (a *AsksForAProduct) Do(ctx context.Context) error {
	return a.waiter.LooksForAProduct(ctx, a.product)
}
