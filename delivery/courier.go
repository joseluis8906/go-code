package delivery

// Courier represents a currier.
type Courier struct {
	Name string
}

// NewCourier creates a new currier.
func NewCourier(name string) *Courier {
	return &Courier{
		Name: name,
	}
}

// GoesUpTo is travel until some point within the city.
func (c *Courier) GoesUpTo(theAddress Address) {
	// Go to store.
}

// PicksUp takes an order.
func (c *Courier) PicksUp(theOrder *Order) {
	theOrder.IsNowTraveling()
}

// Delivers the order.
func (c *Courier) Delivers(theOrder *Order) *CourierDeliversOrderToCustomerAction {
	return NewCourierDeliversOrderToCustomerAction(theOrder)
}

// CourierDeliversOrderToCustomerAction represents an action of a courier.
type CourierDeliversOrderToCustomerAction struct {
	order *Order
}

// NewCourierDeliversOrderToCustomerAction creates a new action of a courier.
func NewCourierDeliversOrderToCustomerAction(order *Order) *CourierDeliversOrderToCustomerAction {
	return &CourierDeliversOrderToCustomerAction{order: order}
}

// To delivers an order to a customer.
func (a *CourierDeliversOrderToCustomerAction) To(theCustomer *Customer) {
	a.order.IsNowDelivered()
}
