package store

import "context"

type Manager struct {
	Name string
}

func (m *Manager) RegisterStore(ctx context.Context, store Store) error {
	return nil
}
