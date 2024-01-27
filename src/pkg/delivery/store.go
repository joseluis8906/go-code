package delivery

// Store represents a store.
type Store struct {
	Name     string
	Address  Address
	Products []Product
	Orders   []*Order
}

// NewStore creates a new store.
func NewStore(name string, address Address, products ...Product) *Store {
	return &Store{
		Name:     name,
		Address:  address,
		Products: products,
	}
}

// Receives receives an order.
func (s *Store) Receives(theOrder *Order) {
	s.Orders = append(s.Orders, theOrder)
}

// BeginsToPrepare prepares an order.
func (s *Store) BeginsToPrepare(theOrder *Order) {
	theOrder.IsNowInPreparation()
}

// FinishesPreparing completes an order.
func (s *Store) FinishesPreparing(theOrder *Order) {
	theOrder.IsNowPrepared()
}

// Delivers delivers an order to a courier.
func (s *Store) Delivers(theOrder *Order) *StoreDeliversTheOrderToCourierAction {
	return NewStoreDeliversTheOrderToCourierAction(theOrder)
}

// StoreDeliversTheOrderToCourierAction represents an action of a store.
type StoreDeliversTheOrderToCourierAction struct {
	order   *Order
	courier *Courier
}

// NewStoreDeliversTheOrderToCourierAction creates a new action of a store.
func NewStoreDeliversTheOrderToCourierAction(order *Order) *StoreDeliversTheOrderToCourierAction {
	return &StoreDeliversTheOrderToCourierAction{order: order}
}

// To delivers an order to a courier.
func (a *StoreDeliversTheOrderToCourierAction) To(theCourier *Courier) *StoreDeliversTheOrderToCourierAction {
	a.courier = theCourier

	return a
}

func (a *StoreDeliversTheOrderToCourierAction) Do() {
	a.courier.PicksUp(a.order)
}
