package waiter

import (
	"context"
	"fmt"

	"github.com/joseluis8906/go-code/src/delivery/internal/product"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"github.com/joseluis8906/go-code/src/pkg/repository"
)

type (
	Catalog interface {
		Matching(ctx context.Context, criteria cmp.Criteria) repository.Result[product.Product]
	}

	// Waiter is an extended delivery waiter.
	Waiter struct {
		catalog  Catalog
		products []product.Product
	}
)

func (w *Waiter) Using(catalog Catalog) *Waiter {
	w.catalog = catalog
	return w
}

func (w *Waiter) LooksForAProduct(ctx context.Context, aProduct fmt.Stringer) error {
	if w.catalog == nil {
		return fmt.Errorf("catalog is nil")
	}

	criteria := cmp.Contains(product.NameField, aProduct.String())
	result, err := w.catalog.Matching(ctx, criteria).ExpectMulti()
	if err != nil {
		return err
	}

	w.products = result
	return nil
}

func (w Waiter) Suggests(ctx context.Context) []product.Product {
	return w.products
}
