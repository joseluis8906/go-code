package delivery

// Product represents a product.
type Product struct {
	Code int
	Name string
}

// NewProduct creates a new product.
func NewProduct(code int, name string) Product {
	return Product{
		Code: code,
		Name: name,
	}
}

// ProductName creates a new product with name.
func ProductName(name string) Product {
	return Product{Name: name}
}
