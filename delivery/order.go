package delivery

const (
	orderStatusCreated       OrderStatus = "created"
	orderStatusInPreparation OrderStatus = "in_preparation"
	orderStatusPrepared      OrderStatus = "prepared"
	orderStatusTraveling     OrderStatus = "traveling"
	orderStatusDelivered     OrderStatus = "delivered"
	orderStatusCompleted     OrderStatus = "completed"
	orderStatusCanceled      OrderStatus = "canceled"
)

type (
	// Order represents an order.
	Order struct {
		Customer Customer
		Store    Store
		Products []Product
		Status   OrderStatus
	}

	// OrderStatus represents an order status.
	OrderStatus string
)

// NewOrder creates a new order.
func NewOrder(customer Customer, store Store, products ...Product) Order {
	return Order{
		Customer: customer,
		Store:    store,
		Products: products,
		Status:   orderStatusCreated,
	}
}

// IsNowInPreparation means the order is being prepared.
func (o *Order) IsNowInPreparation() {
	o.Status = orderStatusInPreparation
}

// IsNowPrepared means the order is ready to be delivered.
func (o *Order) IsNowPrepared() {
	o.Status = orderStatusPrepared
}

// IsNowTraveling means the order is going up to the Customer.
func (o *Order) IsNowTraveling() {
	o.Status = orderStatusTraveling
}

// IsNowDelivered means the order was delivered.
func (o *Order) IsNowDelivered() {
	o.Status = orderStatusDelivered
}

// IsNowCanceled means the order was canceled.
func (o *Order) IsNowCanceled() {
	o.Status = orderStatusCanceled
}

// IsNowCompleted means the order was completed.
func (o *Order) IsNowCompleted() {
	o.Status = orderStatusCompleted
}

// IsCompleted returns true if the order is completed.
func (o *Order) IsCompleted() bool {
	return o.Status == orderStatusCompleted
}

// Append adds a product to an order.
func (o *Order) Append(product Product) {
	o.Products = append(o.Products, product)
}
