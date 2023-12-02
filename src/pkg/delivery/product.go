package delivery

// Product represents a product.
type Product struct {
	Ref  uint64
	Name string
}

// NewProduct creates a new product.
func NewProduct(ref uint64, name string) Product {
	return Product{
		Ref:  ref,
		Name: name,
	}
}

// ProductName creates a new product with name.
func ProductName(name string) Product {
	return Product{Name: name}
}
