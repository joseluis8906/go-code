package delivery

// Product represents a product.
type Product struct {
	Ref  string
	Name string
}

// NewProduct creates a new product.
func NewProduct(ref string, name string) Product {
	return Product{
		Ref:  ref,
		Name: name,
	}
}

// ProductName creates a new product with name.
func ProductName(name string) Product {
	return Product{Name: name}
}
