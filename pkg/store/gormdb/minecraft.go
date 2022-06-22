package gormdb

import (
	"context"

	"github.com/kleister/kleister-api/pkg/model"
)

// Minecraft implements minecraft.Store interface.
type Minecraft struct {
	client *gormdbStore
}

// List implements List from minecraft.Store interface.
func (t *Minecraft) List(ctx context.Context) ([]*model.Minecraft, error) {
	records := make([]*model.Minecraft, 0)

	err := t.client.handle.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}
