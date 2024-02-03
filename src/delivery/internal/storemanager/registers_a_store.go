package storemanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/joseluis8906/go-code/src/delivery/internal/store"
	"github.com/joseluis8906/go-code/src/pkg/repository"
)

type (
	StoreForm struct {
		Name     string
		Country  string
		City     string
		Address  string
		Products []ProductForm
	}

	ProductForm struct {
		Ref   string
		Name  string
		Price MoneyForm
	}

	MoneyForm struct {
		Amount   int64
		Currency string
	}
)

func (sm *StoreManager) RegistersAStore(ctx context.Context, form StoreForm) error {
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
