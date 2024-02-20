package storemanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"

	"github.com/joseluis8906/go-code/src/delivery/internal/product"
	"github.com/joseluis8906/go-code/src/delivery/internal/store"

	"github.com/joseluis8906/go-code/src/pkg/repository"
)

type (
	// StoreManager represent a store manager.
	// It is responsible for creating and updating stores and products.
	// It is also responsible for reporting that an order has been taken and
	// also if the order is ready to be delivered.
	StoreManager struct {
		Email  string
		Stores Stores
	}

	Stores interface {
		Get(ctx context.Context, criteria repository.Criteria) repository.Result[store.Store]
		Add(ctx context.Context, aStore store.Store) error
	}
)

// RegistersAStore creates a new store with the provided data.
// If the store already exists, it will update the existing store with the new data.
// The store's products are not updated, only the store's information.
// If the store already has products they will be kept.
// The store's products are updated by the RegistersProducts method.
func (sm *StoreManager) RegistersAStore(ctx context.Context, form *storemanagerpb.RegistersAStoreRequest) error {
	reqStore := form.GetStore()
	name := reqStore.GetName().Value
	country := reqStore.GetCountry().Value
	city := reqStore.GetCity().Value
	address := reqStore.GetAddress().Value

	criteria := repository.Contains(store.NameField, name)
	currentStore, err := sm.Stores.Get(ctx, criteria).ExpectOne()
	if err != nil && !errors.Is(err, repository.ErrNoData) {
		return fmt.Errorf("getting store: %w", err)
	}

	var sb store.Builder
	sb.Name(name)
	sb.Country(country)
	sb.City(city)
	sb.Address(address)

	newStore, err := sb.Build()
	if err != nil {
		return fmt.Errorf("creating new store: %w", err)
	}

	newStore.Products = currentStore.Products
	err = sm.Stores.Add(ctx, newStore)
	if err != nil {
		return fmt.Errorf("persisting store: %w", err)
	}

	return nil
}

// RegistersProducts adds new products to a store.
// If the store does not exist, it will return an error.
// If products already exist, it will update the existing products with the new data.
func (sm *StoreManager) RegistersProducts(ctx context.Context, form *storemanagerpb.RegistersProductsRequest) error {
	reqStore := form.GetStore()
	name := reqStore.GetName().Value

	criteria := repository.Contains(store.NameField, name)
	theStore, err := sm.Stores.Get(ctx, criteria).ExpectOne()
	if err != nil {
		return fmt.Errorf("getting store: %w", err)
	}

	currentMenu := make(map[string]product.Product, len(theStore.Products))
	for _, p := range theStore.Products {
		currentMenu[p.Ref.Value] = p
	}

	newProducts := form.GetStore().GetProducts()
	for _, pf := range newProducts {
		price := pf.GetPrice()

		var pb product.Builder
		pb.Ref(pf.GetRef().Value)
		pb.Name(pf.GetName().Value)
		pb.Price(price.GetAmount().Value, price.GetCurrency().Value)

		newProduct, err := pb.Build()
		if err != nil {
			return fmt.Errorf("creating new product: %w", err)
		}

		currentMenu[newProduct.Name.Value] = newProduct
	}

	theStore.Products = make([]product.Product, 0, len(currentMenu))
	for _, p := range currentMenu {
		theStore.Products = append(theStore.Products, p)
	}

	err = sm.Stores.Add(ctx, theStore)
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
