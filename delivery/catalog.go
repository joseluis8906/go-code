package delivery

// Catalog represents a catalog.
type Catalog struct {
	Stores []*Store
}

// NewCatalog creates a new catalog.
func NewCatalog(stores []*Store) *Catalog {
	return &Catalog{stores}
}

// ProductsByName returns the products by name.
func (c *Catalog) ProductsByName(name string) []Product {
	result := []Product{}

	for _, store := range c.Stores {
		for _, product := range store.Products {
			if product.Name == name {
				result = append(result, product)
			}
		}
	}

	return result
}

func (c *Catalog) LocateTheStoreFor(thisProduct Product) *Store {
	for _, store := range c.Stores {
		for _, product := range store.Products {
			if thisProduct.Code == product.Code {
				return store
			}
		}
	}

	return nil
}
