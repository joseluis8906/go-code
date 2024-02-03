package store

import (
	"context"

	"github.com/joseluis8906/go-code/src/delivery/internal/product"
)

type (
	Store struct {
		Name     Name
		Country  Country
		City     City
		Address  Address
		Products []product.Product
	}

	storeBuilder struct {
		store Store
		err   error
	}
)

func New() *storeBuilder {
	return &storeBuilder{}
}

func (sb *storeBuilder) Name(name string) *storeBuilder {
	if sb.err != nil {
		return sb
	}

	sb.store.Name, sb.err = NewName(name)
	return sb
}

func (sb *storeBuilder) Country(country string) *storeBuilder {
	if sb.err != nil {
		return sb
	}

	sb.store.Country, sb.err = NewCountry(country)
	return sb
}

func (sb *storeBuilder) City(city string) *storeBuilder {
	if sb.err != nil {
		return sb
	}

	sb.store.City, sb.err = NewCity(city)
	return sb
}

func (sb *storeBuilder) Address(address string) *storeBuilder {
	if sb.err != nil {
		return sb
	}

	sb.store.Address, sb.err = NewAddress(address)
	return sb
}

func (sb *storeBuilder) Do(ctx context.Context) (Store, error) {
	return sb.store, sb.err
}
