package delivery_test

import (
	"fmt"

	"github.com/joseluis8906/go-code/delivery"
)

func ExampleOrder_IsCompleted() {
	city := "Greensboro"

	theCatalog := delivery.NewCatalog([]*delivery.Store{
		delivery.NewStore(
			"McDonald's",
			delivery.NewAddress("3003 SW 34th St", city),
			delivery.NewProduct(1, "Chess Burger"),
			delivery.NewProduct(2, "French Fries"),
			delivery.NewProduct(3, "Coke"),
		),

		delivery.NewStore(
			"Burger King",
			delivery.NewAddress("2304 Franklin Ave", city),
			delivery.NewProduct(4, "Chess Burger"),
			delivery.NewProduct(5, "French Fries"),
			delivery.NewProduct(6, "Coke"),
		),
	})

	theCustomer := delivery.NewCustomer("Ellie Hang", delivery.NewAddress("211 Southside Square", city))
	theCashier := delivery.NewCashier("Uber Eats")
	theCourier := delivery.NewCourier("John Doe")

	// everything starts when
	theFirstProduct := theCustomer.LooksFor(delivery.ProductName("Chess Burger")).Using(theCatalog)
	// then
	theStore := theCashier.LocatesTheStore().Of(theFirstProduct).Using(theCatalog)
	// then
	theOrder := theCashier.CreatesAnOrder().Using(theCatalog).On(theStore).For(theCustomer).With(theFirstProduct)
	// and
	anotherProduct := theCustomer.LooksFor(delivery.ProductName("Coke")).Using(theCatalog)
	// then
	theCashier.Adds(anotherProduct).To(theOrder)
	// once all products are added to the order
	theCashier.Sends(theOrder).To(theStore)
	// after a while
	theStore.BeginsToPrepare(theOrder)
	// then
	theCourier.GoesUpTo(theStore.Address)
	// and when
	theStore.FinishesPreparing(theOrder)
	// then
	theStore.Delivers(theOrder).To(theCourier)
	// then
	theCourier.GoesUpTo(theCustomer.Address)
	// and once there
	theCourier.Delivers(theOrder).To(theCustomer)
	// finally
	theCustomer.Confirms(theOrder).WasReceived()
	// and if everything is ok then that's it!

	fmt.Println(theOrder.IsCompleted())

	// Output: true
}
