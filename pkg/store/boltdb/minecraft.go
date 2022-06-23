package boltdb

import (
	"context"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/google/uuid"
	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/model"
)

// Minecraft implements minecraft.Store interface.
type Minecraft struct {
	client *botldbStore
}

// List implements List from minecraft.Store interface.
func (m *Minecraft) List(ctx context.Context) ([]*model.Minecraft, error) {
	records := make([]*model.Minecraft, 0)

	err := m.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	return records, nil
}

// Sync implements Sync from minecraft.Store interface.
func (m *Minecraft) Sync(ctx context.Context, versions version.Versions) error {
	tx, err := m.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, v := range versions {
		record := &model.Minecraft{}

		err := m.client.handle.Select(
			q.Eq("Name", v.ID),
		).First(record)

		if err == storm.ErrNotFound {
			if err := m.create(tx, &model.Minecraft{
				Name: v.ID,
				Type: "release",
			}); err != nil {
				return err
			}
		} else {
			record.Name = v.ID
			record.Type = "release"

			if err := m.update(tx, record); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

func (m *Minecraft) create(tx storm.Node, record *model.Minecraft) error {
	record.ID = uuid.New().String()
	record.UpdatedAt = time.Now().UTC()
	record.CreatedAt = time.Now().UTC()
	return tx.Save(record)
}

func (m *Minecraft) update(tx storm.Node, record *model.Minecraft) error {
	record.UpdatedAt = time.Now().UTC()
	return tx.Save(record)
}
