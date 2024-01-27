package delivery

import "strings"

// Assistant represents a assistant.
type Assistant struct {
	Name    string
	Catalog *Catalog
}

// NewAssistant creates a new assistant.
func NewAssistant(name string, catalog *Catalog) *Assistant {
	return &Assistant{
		Name:    name,
		Catalog: catalog,
	}
}

// Using sets the catalog.
func (c *Assistant) Using(theCatalog *Catalog) {
	c.Catalog = theCatalog
}

// CreatesAnOrder creates an order.
func (c *Assistant) LooksForAProduct(query Product) []Product {
	var suggestedProducts []Product

	for _, store := range c.Catalog.Stores {
		for _, product := range store.Products {
			if strings.Contains(product.Name, query.Name) {
				suggestedProducts = append(suggestedProducts, product)
			}
		}
	}

	return suggestedProducts
}

// CreatesAnOrder creates an order.
func (c *Assistant) CreatesAnOrder() *AssistantCreatesOrderAction {
	return NewAssistantCreatesOrderAction()
}

// Adds adds a product to an order.
func (c *Assistant) Adds(aProduct Product) *AssistantAddsAProductToOrderAction {
	return NewAssistantAddsAProductToOrderAction(aProduct)
}

// LocatesTheStoreOf locates a store for a product.
func (c *Assistant) LocatesTheStoreOf(thisProduct Product) *Store {
	for _, store := range c.Catalog.Stores {
		for _, product := range store.Products {
			if thisProduct.Ref == product.Ref {
				return store
			}
		}
	}

	return nil
}

// Sends sends an order to a store.
func (c *Assistant) Sends(theOrder *Order) *AssistantSendsOrderToStoreAction {
	return NewAssistantSendsOrderToStoreAction()
}

// AssistantCreatesOrderAction represents an action of a assistant.
type AssistantCreatesOrderAction struct {
	store    *Store
	customer *Customer
	catalog  *Catalog
	product  Product
}

// NewAssistantCreatesOrderAction creates a new action of a assistant.
func NewAssistantCreatesOrderAction() *AssistantCreatesOrderAction {
	return &AssistantCreatesOrderAction{}
}

// On sets the store.
func (a *AssistantCreatesOrderAction) On(theStore *Store) *AssistantCreatesOrderAction {
	a.store = theStore

	return a
}

// For sets the customer.
func (a *AssistantCreatesOrderAction) For(theCustomer *Customer) *AssistantCreatesOrderAction {
	a.customer = theCustomer

	return a
}

// Adding add a product to a new order.
func (a *AssistantCreatesOrderAction) Adding(thisProduct Product) *AssistantCreatesOrderAction {
	a.product = thisProduct

	return a
}

// Do creates a new order
func (a *AssistantCreatesOrderAction) Do() *Order {
	newOrder := NewOrder(*a.customer, *a.store, a.product)

	return &newOrder
}

// AssistantAddsAProductToOrderAction represents an action of a assistant.
type AssistantAddsAProductToOrderAction struct {
	order   *Order
	product Product
}

// NewAssistantAddsAProductToOrderAction creates a new action of a assistant.
func NewAssistantAddsAProductToOrderAction(product Product) *AssistantAddsAProductToOrderAction {
	return &AssistantAddsAProductToOrderAction{product: product}
}

// To adds a product to an order.
func (c *AssistantAddsAProductToOrderAction) To(theOrder *Order) *AssistantAddsAProductToOrderAction {
	c.order = theOrder

	return c
}

func (c *AssistantAddsAProductToOrderAction) Do() {
	c.order.Append(c.product)
}

// AssistantSendsOrderToStoreAction represents an action of a assistant.
type AssistantSendsOrderToStoreAction struct {
	store *Store
	order *Order
}

// NewAssistantSendsOrderToStoreAction creates a new action of a assistant.
func NewAssistantSendsOrderToStoreAction() *AssistantSendsOrderToStoreAction {
	return &AssistantSendsOrderToStoreAction{}
}

// To sends an order to a store.
func (a *AssistantSendsOrderToStoreAction) To(theStore *Store) *AssistantSendsOrderToStoreAction {
	a.store = theStore

	return a
}

// Do sends an order to a store.
func (a *AssistantSendsOrderToStoreAction) Do() {
	a.store.Receives(a.order)
}
