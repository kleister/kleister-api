package boltdb

import (
	"context"

	"github.com/asdine/storm/v3"
	"github.com/kleister/kleister-api/pkg/model"
)

// Minecraft implements minecraft.Store interface.
type Minecraft struct {
	client *botldbStore
}

// List implements List from minecraft.Store interface.
func (t *Minecraft) List(ctx context.Context) ([]*model.Minecraft, error) {
	records := make([]*model.Minecraft, 0)

	err := t.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	return records, nil
}
