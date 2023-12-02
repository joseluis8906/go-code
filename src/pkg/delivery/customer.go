package delivery

import "math/rand"

// Customer represents a customer.
type Customer struct {
	Email        string
	Name         string
	Address      Address
	Orders       []*Order
	CurrentStore Store
}

// NewCustomer creates a new customer.
func NewCustomer(email string, name string, address Address) Customer {
	return Customer{
		Email:   email,
		Name:    name,
		Address: address,
	}
}

// IsZero returns true if the customer is zero.
func (c *Customer) IsZero() bool {
	return c.Email == ""
}

// AsksFor asks for a product.
func (c *Customer) AsksFor(product Product) *CustomerAsksForProductAction {
	return NewCustomerAsksForProductAction(product)
}

// ChoosesAProductFrom chooses a product from a list of products.
func (c *Customer) ChoosesAProductFrom(theListOfSuggestions []Product) Product {
	min := 0
	max := len(theListOfSuggestions) - 1

	return theListOfSuggestions[rand.Intn(max-min)]
}

// Receives receives an order.
func (c *Customer) Receives(theOrder *Order) {
	theOrder.IsNowDelivered()
}

// Confirms confirms the delivery of an order.
func (c *Customer) Confirms(theOrder *Order) *CustomerConfirmsDeliveryAction {
	return NewCustomerConfirmsDeliveryAction(theOrder)
}

type (
	// CustomerAsksForProductAction represents an action of a customer.
	CustomerAsksForProductAction struct {
		product   Product
		assistant *Assistant
	}

	// CustomerConfirmsDeliveryAction represents an action of a customer.
	CustomerConfirmsDeliveryAction struct {
		order *Order
	}
)

func NewCustomerAsksForProductAction(product Product) *CustomerAsksForProductAction {
	return &CustomerAsksForProductAction{product: product}
}

// NewCustomerConfirmsDeliveryAction creates a new action of a customer.
func NewCustomerConfirmsDeliveryAction(order *Order) *CustomerConfirmsDeliveryAction {
	return &CustomerConfirmsDeliveryAction{order: order}
}

// To sets the assistant of the action.
func (c *CustomerAsksForProductAction) To(theAssistant *Assistant) *CustomerAsksForProductAction {
	c.assistant = theAssistant

	return c
}

// To sets the assistant of the action.
func (c *CustomerAsksForProductAction) Do() []Product {
	return c.assistant.LooksForAProduct(c.product)
}

// WasReceived confirms the delivery of an order.
func (a *CustomerConfirmsDeliveryAction) WasReceived() {
	a.order.IsNowCompleted()
}
