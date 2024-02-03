package storemanager

import (
	"context"

	"github.com/joseluis8906/go-code/src/delivery/internal/store"
	"github.com/joseluis8906/go-code/src/pkg/repository"
)

type StoreManager struct {
	Email string

	Stores interface {
		Get(ctx context.Context, criteria repository.Criteria) repository.Result[store.Store]
		Add(ctx context.Context, aStore store.Store) error
	}
}
