package delivery

import "math/rand"

// Customer represents a customer.
type Customer struct {
	Name         string
	Address      Address
	Orders       []*Order
	CurrentStore Store
}

// NewCustomer creates a new customer.
func NewCustomer(name string, address Address) *Customer {
	return &Customer{
		Name:    name,
		Address: address,
	}
}

// Confirms confirms the delivery of an order.
func (c *Customer) Confirms(theOrder *Order) *CustomerConfirmsDeliveryAction {
	return NewCustomerConfirmsDeliveryAction(theOrder)
}

// LooksFor looks for a product in the catalog.
func (c *Customer) LooksFor(criteria Product) *CustomerLooksForProductAction {
	return NewCustomerLooksForProductAction(criteria)
}

// CustomerLooksForProductAction represents an action of a customer.
type CustomerLooksForProductAction struct {
	product Product
}

// NewCustomerLooksForProductAction creates a new action of a customer.
func NewCustomerLooksForProductAction(product Product) *CustomerLooksForProductAction {
	return &CustomerLooksForProductAction{product: product}
}

// Using uses a catalog to find a product.
func (c *CustomerLooksForProductAction) Using(theCatalog *Catalog) Product {
	products := theCatalog.ProductsByName(c.product.Name)
	min := 0
	max := len(products) - 1

	return products[rand.Intn(max-min)]
}

// CustomerConfirmsDeliveryAction represents an action of a customer.
type CustomerConfirmsDeliveryAction struct {
	order *Order
}

// NewCustomerConfirmsDeliveryAction creates a new action of a customer.
func NewCustomerConfirmsDeliveryAction(order *Order) *CustomerConfirmsDeliveryAction {
	return &CustomerConfirmsDeliveryAction{order: order}
}

// WasReceived confirms the delivery of an order.
func (a *CustomerConfirmsDeliveryAction) WasReceived() {
	a.order.IsNowCompleted()
}
