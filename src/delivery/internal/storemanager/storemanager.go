package storemanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/store"
	"go.opentelemetry.io/otel"

	"github.com/joseluis8906/go-code/src/pkg/repository"
)

type (
	// StoreManager represent a store manager.
	// It is responsible for creating and updating stores and products.
	// It is also responsible for reporting that an order has been taken and
	// also if the order is ready to be delivered.
	StoreManager struct {
		Email string

		Stores interface {
			Get(ctx context.Context, criteria repository.Criteria) repository.Result[store.Store]
			Add(ctx context.Context, aStore store.Store) error
		}
	}
)

// RegistersAStore creates a new store with the provided data.
// If the store already exists, it will update the existing store with the new data.
// The store's products are not updated, only the store's information.
// If the store already has products they will be kept.
// The store's products are updated by the RegistersProducts method.
func (sm *StoreManager) RegistersAStore(ctx context.Context, form StoreForm) error {
	ctx, span := otel.Tracer("").Start(ctx, "storemanager.StoreManager.RegistersAStore")
	defer span.End()

	criteria := repository.Contains(store.NameField, form.Name)
	currentStore, err := sm.Stores.Get(ctx, criteria).ExpectOne()
	if err != nil && !errors.Is(err, repository.ErrNoData) {
		return fmt.Errorf("getting store: %w", err)
	}

	aStore, err := store.New().
		Name(form.Name).
		Country(form.Country).
		City(form.City).
		Address(form.Address).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("creating new store: %w", err)
	}

	aStore.Products = currentStore.Products

	err = sm.Stores.Add(ctx, aStore)
	if err != nil {
		return fmt.Errorf("persisting store: %w", err)
	}

	return nil
}

// RegistersProducts adds new products to a store.
// If the store does not exist, it will return an error.
// If products already exist, it will update the existing products with the new data.
func (sm *StoreManager) RegistersProducts(ctx context.Context, form StoreForm) error {
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

// ReportsOrderIsTaken reports that an order has been taken by the store
// and the order is being prepared.
// If the order does not exist, it will return an error.
// If the order is already taken, it will return an error.
func (sm *StoreManager) ReportsOrderIsTaken(ctx context.Context) error {
	return nil
}

// ReportsOrderIsReady reports that an order is ready to be delivered.
// If the order does not exist, it will return an error.
// If the order is not taken, it will return an error.
func (sm *StoreManager) ReportsOrderIsReady(ctx context.Context) error {
	return nil
}
