package delivery

// Cashier represents a cashier.
type Cashier struct {
	Name string
}

// NewCashier creates a new cashier.
func NewCashier(name string) Cashier {
	return Cashier{
		Name: name,
	}
}

// CreatesAnOrder creates an order.
func (c *Cashier) CreatesAnOrder() *CashierCreatesOrderAction {
	return NewCashierCreatesOrderAction()
}

// Adds adds a product to an order.
func (c *Cashier) Adds(aProduct Product) *CashierAddsAProductToOrderAction {
	return NewCashierAddsAProductToOrderAction(aProduct)
}

// LocatesTheStore locates a store.
func (c *Cashier) LocatesTheStore() *CashierLocatesTheStoreAction {
	return NewCashierLocatesStoreAction()
}

// Sends sends an order to a store.
func (c *Cashier) Sends(theOrder *Order) *CashierSendsOrderToStoreAction {
	return NewCashierSendsOrderToStoreAction()
}

// CashierCreatesOrderAction represents an action of a cashier.
type CashierCreatesOrderAction struct {
	store    *Store
	customer *Customer
	catalog  *Catalog
}

// NewCashierCreatesOrderAction creates a new action of a cashier.
func NewCashierCreatesOrderAction() *CashierCreatesOrderAction {
	return &CashierCreatesOrderAction{}
}

// Using sets the catalog.
func (a *CashierCreatesOrderAction) Using(theCatalog *Catalog) *CashierCreatesOrderAction {
	a.catalog = theCatalog

	return a
}

// On sets the store.
func (a *CashierCreatesOrderAction) On(theStore *Store) *CashierCreatesOrderAction {
	a.store = theStore

	return a
}

// For sets the customer.
func (a *CashierCreatesOrderAction) For(theCustomer *Customer) *CashierCreatesOrderAction {
	a.customer = theCustomer

	return a
}

// With add a product to a new order.
func (a *CashierCreatesOrderAction) With(thisProduct Product) *Order {
	theStore := a.catalog.LocateTheStoreFor(thisProduct)
	newOrder := NewOrder(*a.customer, *theStore, thisProduct)

	return &newOrder
}

// CashierLocatesTheStoreAction represents an action of a cashier.
type CashierLocatesTheStoreAction struct {
	product Product
}

// NewCashierLocatesStoreAction creates a new action of a cashier.
func NewCashierLocatesStoreAction() *CashierLocatesTheStoreAction {
	return &CashierLocatesTheStoreAction{}
}

// Of sets the product.
func (a *CashierLocatesTheStoreAction) Of(theFirstProduct Product) *CashierLocatesTheStoreAction {
	a.product = theFirstProduct

	return a
}

// Using uses a catalog to find a product.
func (a *CashierLocatesTheStoreAction) Using(theCatalog *Catalog) *Store {
	return theCatalog.LocateTheStoreFor(a.product)
}

// CashierAddsAProductToOrderAction represents an action of a cashier.
type CashierAddsAProductToOrderAction struct {
	product Product
}

// NewCashierAddsAProductToOrderAction creates a new action of a cashier.
func NewCashierAddsAProductToOrderAction(product Product) *CashierAddsAProductToOrderAction {
	return &CashierAddsAProductToOrderAction{product: product}
}

// To adds a product to an order.
func (c *CashierAddsAProductToOrderAction) To(theOrder *Order) {
	theOrder.Append(c.product)
}

// CashierSendsOrderToStoreAction represents an action of a cashier.
type CashierSendsOrderToStoreAction struct {
	order *Order
}

// NewCashierSendsOrderToStoreAction creates a new action of a cashier.
func NewCashierSendsOrderToStoreAction() *CashierSendsOrderToStoreAction {
	return &CashierSendsOrderToStoreAction{}
}

// To sends an order to a store.
func (a *CashierSendsOrderToStoreAction) To(theStore *Store) {
	theStore.Receives(a.order)
}
