package domain

import (
	"github.com/joseluis8906/go-code/src/pkg/delivery"
)

type (
	// CourierAsksForAProduct represents an action of a courier.
	CustomerAsksForAProduct struct {
		theCustomer    *delivery.Customer
		theProductName string
		theAssistant   *delivery.Assistant
	}
)

// NewCustomerAsksForAProduct creates a new CourierAsksForAProduct.
func NewCustomerAsksForAProduct(customer *delivery.Customer, productName string, assistant *delivery.Assistant) *CustomerAsksForAProduct {
	return &CustomerAsksForAProduct{
		theCustomer:    customer,
		theProductName: productName,
		theAssistant:   assistant,
	}
}

// Do returns a list of suggestions.
func (c *CustomerAsksForAProduct) Do() []delivery.Product {
	listOfSuggestions := c.theCustomer.AsksFor(delivery.ProductName(c.theProductName)).To(c.theAssistant).Do()

	return listOfSuggestions
}
