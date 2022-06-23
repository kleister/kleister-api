package gormdb

import (
	"context"

	"github.com/kleister/go-forge/version"
	"github.com/kleister/kleister-api/pkg/model"
)

// Forge implements forge.Store interface.
type Forge struct {
	client *gormdbStore
}

// List implements List from forge.Store interface.
func (f *Forge) List(ctx context.Context) ([]*model.Forge, error) {
	records := make([]*model.Forge, 0)

	err := f.client.handle.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}

// Sync implements Sync from forge.Store interface.
func (f *Forge) Sync(ctx context.Context, versions version.Versions) error {
	tx := f.client.handle.Begin()
	defer tx.Rollback()

	for _, v := range versions {
		record := &model.Forge{}

		if err := tx.Where(
			&model.Forge{
				Name: v.ID,
			},
		).Attrs(
			&model.Forge{
				Minecraft: v.Minecraft,
			},
		).FirstOrCreate(
			record,
		).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}
