package store

import (
	"github.com/joseluis8906/go-code/protobuf/deliverypb"
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

	Builder struct {
		store Store
		err   error
	}
)

func (sb *Builder) Name(name string) *Builder {
	if sb.err != nil {
		return sb
	}

	sb.store.Name, sb.err = NewName(name)
	return sb
}

func (sb *Builder) Country(country string) *Builder {
	if sb.err != nil {
		return sb
	}

	sb.store.Country, sb.err = NewCountry(country)
	return sb
}

func (sb *Builder) City(city string) *Builder {
	if sb.err != nil {
		return sb
	}

	sb.store.City, sb.err = NewCity(city)
	return sb
}

func (sb *Builder) Address(address string) *Builder {
	if sb.err != nil {
		return sb
	}

	sb.store.Address, sb.err = NewAddress(address)
	return sb
}

func (sb *Builder) Build() (Store, error) {
	return sb.store, sb.err
}

func (s Store) IsZero() bool {
	return s.Name.IsZero()
}

func FromPB(data *deliverypb.Store) (Store, error) {
	name := data.GetName().GetValue()
	country := data.GetCountry().GetValue()
	city := data.GetCity().GetValue()
	address := data.GetAddress().GetValue()

	var sb Builder
	sb.Name(name)
	sb.Country(country)
	sb.City(city)
	sb.Address(address)

	return sb.Build()
}
