package gormdb

import (
	"context"

	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/model"
)

// Minecraft implements minecraft.Store interface.
type Minecraft struct {
	client *gormdbStore
}

// List implements List from minecraft.Store interface.
func (m *Minecraft) List(ctx context.Context) ([]*model.Minecraft, error) {
	records := make([]*model.Minecraft, 0)

	err := m.client.handle.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}

// Sync implements Sync from minecraft.Store interface.
func (m *Minecraft) Sync(ctx context.Context, versions version.Versions) error {
	tx := m.client.handle.Begin()
	defer tx.Rollback()

	for _, v := range versions {
		record := &model.Minecraft{}

		if err := tx.Where(
			&model.Minecraft{
				Name: v.ID,
			},
		).Attrs(
			&model.Minecraft{
				Type: "release",
			},
		).FirstOrCreate(
			record,
		).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}
