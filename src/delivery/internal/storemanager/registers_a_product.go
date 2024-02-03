package storemanager

import (
	"context"
	"fmt"

	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/store"
	"github.com/joseluis8906/go-code/src/pkg/repository"
)

func (sm *StoreManager) RegistersAProduct(ctx context.Context, form StoreForm) error {
	criteria := repository.Contains(store.NameField, form.Name)
	aStore, err := sm.Stores.Get(ctx, criteria).ExpectOne()
	if err != nil {
		return fmt.Errorf("getting store: %w", err)
	}

	productsSet := make(map[string]product.Product, len(aStore.Products))
	for _, p := range aStore.Products {
		productsSet[p.Ref.Value] = p
	}

	for _, pf := range form.Products {
		aProduct, err := product.New().
			Ref(pf.Ref).
			Name(pf.Name).
			Price(pf.Price.Amount, pf.Price.Currency).
			Do(ctx)

		if err != nil {
			return fmt.Errorf("creating new product: %w", err)
		}

		productsSet[pf.Ref] = aProduct
	}

	aStore.Products = make([]product.Product, 0, len(productsSet))
	for _, p := range productsSet {
		aStore.Products = append(aStore.Products, p)
	}

	err = sm.Stores.Add(ctx, aStore)
	if err != nil {
		return fmt.Errorf("persisting store: %w", err)
	}

	return nil
}
