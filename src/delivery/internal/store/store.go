package store

import "github.com/joseluis8906/go-code/src/delivery/internal/product"

type Store struct {
	Name     Name
	Products []product.Product
}
