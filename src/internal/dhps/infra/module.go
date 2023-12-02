package infra

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAssistantInMemoryRepository,
	NewCustomerInMemoryRepository,
)
