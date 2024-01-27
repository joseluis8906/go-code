package customer

import "fmt"

type (
	// Customer is an extended delivery customer.
	Customer struct {
		Email Email
	}
)

func (c *Customer) AsksFor(aProduct fmt.Stringer) *AsksForAProduct {
	return &AsksForAProduct{product: aProduct}
}
