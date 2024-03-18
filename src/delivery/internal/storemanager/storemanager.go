package storemanager

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/joseluis8906/go-code/protobuf/delivery/storemanagerpb"

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

// AddStore creates a new store with the provided data.
// If the store already exists, it will update the existing store with the new data.
// The store's products are not updated, only the store's information.
// If the store already has products they will be kept.
// The store's products are updated by the RegistersProducts method.
func (sm *StoreManager) AddStore(ctx context.Context, req *pb.AddStoreRequest) error {
	newStore, err := store.FromPB(req.GetStore())
	if err != nil {
		return fmt.Errorf("creating new store: %w", err)
	}

	criteria := repository.Equals(store.Fields().Name, fmt.Sprintf("%s", newStore.Name))

	oldStore, err := sm.Stores.Get(ctx, criteria).ExpectOne()
	if err != nil && !errors.Is(err, repository.ErrNoData) {
		return fmt.Errorf("verifing exisiting store: %w", err)
	}

	if !oldStore.IsZero() {
		return errors.New("store already exists")
	}

	err = sm.Stores.Add(ctx, newStore)
	if err != nil {
		return fmt.Errorf("persisting new store: %w", err)
	}

	return nil
}

// AddProduct adds a new product to a store.
// If the store does not exist, it will return an error.
// If products already exist, it will update the existing products with the new data.
func (sm *StoreManager) AddProduct(ctx context.Context, form *pb.AddProductRequest) error {
	name, err := store.NewName(form.GetStore().GetName().GetValue())
	if err != nil {
		return fmt.Errorf("casting store name: %w", err)
	}

	criteria := repository.Contains(store.Fields().Name, fmt.Sprintf("%s", name))

	theStore, err := sm.Stores.Get(ctx, criteria).ExpectOne()
	if err != nil {
		return fmt.Errorf("getting store: %w", err)
	}

	theCatalog := make(map[product.Ref]product.Product, len(theStore.Products))
	for _, p := range theStore.Products {
		theCatalog[p.Ref] = p
	}

	newProduct, err := product.FromPB(form.GetProduct())
	if err != nil {
		return fmt.Errorf("creating new product: %w", err)
	}

	theCatalog[newProduct.Ref] = newProduct

	theStore.Products = make([]product.Product, 0, len(theCatalog))
	for _, p := range theCatalog {
		theStore.Products = append(theStore.Products, p)
	}

	err = sm.Stores.Add(ctx, theStore)
	if err != nil {
		return fmt.Errorf("persisting store: %w", err)
	}

	return nil
}

// OrderIsTaken reports that an order has been taken by the store
// and the order is being prepared.
// If the order does not exist, it will return an error.
// If the order is already taken, it will return an error.
func (sm *StoreManager) OrderIsTaken(ctx context.Context) error {
	return nil
}

// OrderIsReady reports that an order is ready to be delivered.
// If the order does not exist, it will return an error.
// If the order is not taken, it will return an error.
func (sm *StoreManager) OrderIsReady(ctx context.Context) error {
	return nil
}
