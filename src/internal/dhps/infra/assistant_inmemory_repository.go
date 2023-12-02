package infra

import (
	"context"
	"sync"

	"github.com/joseluis8906/go-code/src/pkg/delivery"
)

const (
	system string = "uberEats"
)

const (
	city string = "Greensboro"
)

type (
	AssistantInMemoryRepository struct {
		assistants map[string]*delivery.Assistant
		mux        sync.Mutex
	}
)

func NewAssistantInMemoryRepository() *AssistantInMemoryRepository {
	return &AssistantInMemoryRepository{
		assistants: map[string]*delivery.Assistant{
			system: delivery.NewAssistant(
				"Rappi",
				delivery.NewCatalog([]*delivery.Store{
					delivery.NewStore(
						1,
						"McDonald's",
						delivery.NewAddress("3003 SW 34th St", city),
						delivery.NewProduct(1, "Chess Burger"),
						delivery.NewProduct(2, "French Fries"),
						delivery.NewProduct(3, "Coke"),
						delivery.NewProduct(4, "Burger Magnifica"),
					),
					delivery.NewStore(
						2,
						"Burger King",
						delivery.NewAddress("2304 Franklin Ave", city),
						delivery.NewProduct(5, "Chess Burger"),
						delivery.NewProduct(6, "French Fries"),
						delivery.NewProduct(7, "Coke"),
						delivery.NewProduct(8, "Mini Burger"),
					),
				}),
			),
		},
	}
}

func (r *AssistantInMemoryRepository) Get(ctx context.Context) *delivery.Assistant {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.assistants[system]
}
