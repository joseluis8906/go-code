package waiter

import (
	"context"
	"fmt"

	"github.com/joseluis8906/go-code/src/delivery/internal/product"

	"github.com/joseluis8906/go-code/src/pkg/repository"
)

type (
	Catalog interface {
		Get(ctx context.Context, criteria repository.Criteria) repository.Result[product.Product]
	}

	// Waiter is an extended delivery waiter.
	Waiter struct {
		Catalog  Catalog
		products []product.Product
	}
)

func (w *Waiter) LooksForAProduct(ctx context.Context, productName product.Name) error {
	if w.Catalog == nil {
		return fmt.Errorf("catalog is nil")
	}

	criteria := repository.Contains("", productName.Value)
	result, err := w.Catalog.Get(ctx, criteria).ExpectMany()
	if err != nil {
		return err
	}

	w.products = result
	return nil
}

func (w Waiter) Suggests(ctx context.Context) []product.Product {
	return w.products
}
