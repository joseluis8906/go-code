package app

import (
	"github.com/joseluis8906/go-code/src/internal/dhps/infra"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewDeliveryService)

// Params represents the input parameters for the app module.
type Params struct {
	fx.In

	AssistantRepository *infra.AssistantInMemoryRepository
	CustomerRepository  *infra.CustomerInMemoryRepository
}
