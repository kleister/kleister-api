package boltdb

import (
	"context"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/google/uuid"
	"github.com/kleister/go-forge/version"
	"github.com/kleister/kleister-api/pkg/model"
)

// Forge implements forge.Store interface.
type Forge struct {
	client *botldbStore
}

// List implements List from forge.Store interface.
func (f *Forge) List(ctx context.Context) ([]*model.Forge, error) {
	records := make([]*model.Forge, 0)

	err := f.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	return records, nil
}

// Sync implements Sync from forge.Store interface.
func (f *Forge) Sync(ctx context.Context, versions version.Versions) error {
	tx, err := f.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, v := range versions {
		record := &model.Forge{}

		err := f.client.handle.Select(
			q.Eq("Name", v.ID),
		).First(record)

		if err == storm.ErrNotFound {
			if err := f.create(tx, &model.Forge{
				Name:      v.ID,
				Minecraft: v.Minecraft,
			}); err != nil {
				return err
			}
		} else {
			record.Name = v.ID
			record.Minecraft = v.Minecraft

			if err := f.update(tx, record); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

func (f *Forge) create(tx storm.Node, record *model.Forge) error {
	record.ID = uuid.New().String()
	record.UpdatedAt = time.Now().UTC()
	record.CreatedAt = time.Now().UTC()
	return tx.Save(record)
}

func (f *Forge) update(tx storm.Node, record *model.Forge) error {
	record.UpdatedAt = time.Now().UTC()
	return tx.Save(record)
}
