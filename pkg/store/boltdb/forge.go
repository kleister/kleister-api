package boltdb

import (
	"context"

	"github.com/asdine/storm/v3"
	"github.com/kleister/kleister-api/pkg/model"
)

// Forge implements forge.Store interface.
type Forge struct {
	client *botldbStore
}

// List implements List from forge.Store interface.
func (t *Forge) List(ctx context.Context) ([]*model.Forge, error) {
	records := make([]*model.Forge, 0)

	err := t.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	return records, nil
}
