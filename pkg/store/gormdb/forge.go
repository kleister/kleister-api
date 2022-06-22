package gormdb

import (
	"context"

	"github.com/kleister/kleister-api/pkg/model"
)

// Forge implements forge.Store interface.
type Forge struct {
	client *gormdbStore
}

// List implements List from forge.Store interface.
func (t *Forge) List(ctx context.Context) ([]*model.Forge, error) {
	records := make([]*model.Forge, 0)

	err := t.client.handle.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}
