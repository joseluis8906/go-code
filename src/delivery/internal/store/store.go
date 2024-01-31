package store

import "github.com/joseluis8906/go-code/src/delivery/internal/product"

type Store struct {
	Name     Name
	City     City
	Address  Address
	Products []product.Product
}
