package delivery_test

import (
	"fmt"

	"github.com/joseluis8906/yummies/go-code/src/pkg/delivery"
)

func ExampleOrder_IsCompleted() {
	country := "USA"
	city := "Greensboro"

	theAssistant := delivery.NewAssistant(
		"Rappi",
		delivery.NewCatalog([]*delivery.Store{
			delivery.NewStore(
				"McDonald's",
				delivery.NewAddress("3003 SW 34th St", city, country),
				delivery.NewProduct("CHB-001", "Chess Burger"),
				delivery.NewProduct("FRF-001", "French Fries"),
				delivery.NewProduct("COK-001", "Coke"),
			),
			delivery.NewStore(
				"Burger King",
				delivery.NewAddress("2304 Franklin Ave", city, country),
				delivery.NewProduct("CHB-002", "Chess Burger"),
				delivery.NewProduct("FRF-002", "French Fries"),
				delivery.NewProduct("COK-002", "Coke"),
			),
		}),
	)

	theCustomer := delivery.NewCustomer("ellie.hang@example.com", "Ellie Hang", delivery.NewAddress("211 Southside Square", city, country))
	theCourier := delivery.NewCourier("John Doe")

	// everything starts when
	listOfSuggestions := theCustomer.AsksFor(delivery.ProductName("Chess Burger")).To(theAssistant).Do()
	// then
	theFirstProduct := theCustomer.ChoosesAProductFrom(listOfSuggestions)
	// then
	theStore := theAssistant.LocatesTheStoreOf(theFirstProduct)
	// then
	theOrder := theAssistant.CreatesAnOrder().On(theStore).For(&theCustomer).Adding(theFirstProduct).Do()
	// and
	listOfSuggestions = theCustomer.AsksFor(delivery.ProductName("Coke")).To(theAssistant).Do()
	// when
	anotherProduct := theCustomer.ChoosesAProductFrom(listOfSuggestions)
	// then
	theAssistant.Adds(anotherProduct).To(theOrder).Do()
	// once all products are added to the order
	theAssistant.Sends(theOrder).To(theStore).Do()
	// and after a while
	theStore.BeginsToPrepare(theOrder)
	// then
	theCourier.GoesUpTo(theStore.Address)
	// and when
	theStore.FinishesPreparing(theOrder)
	// then
	theStore.Delivers(theOrder).To(theCourier).Do()
	// then
	theCourier.GoesUpTo(theCustomer.Address)
	// and once there
	theCourier.Delivers(theOrder).To(&theCustomer).Do()
	// finally
	theCustomer.Confirms(theOrder).WasReceived()
	// and if everything is ok then that's it!

	fmt.Println(theOrder.IsCompleted())

	// Output: true
}
