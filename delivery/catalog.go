package delivery

// Catalog represents a catalog.
type Catalog struct {
	Stores []*Store
}

// NewCatalog creates a new catalog.
func NewCatalog(stores []*Store) *Catalog {
	return &Catalog{stores}
}
