package boltdb

import (
	"context"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/kleister/kleister-api/pkg/model"
)

// Versions implements versions.Store interface.
type Versions struct {
	client *botldbStore
}

// List implements List from versions.Store interface.
func (v *Versions) List(ctx context.Context, modID string) ([]*model.Version, error) {
	records := make([]*model.Version, 0)

	err := v.client.handle.Select(
		q.Eq("ModID", modID),
	).OrderBy(
		"Name",
	).Find(
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	return records, nil
}

// Show implements Show from versions.Store interface.
func (v *Versions) Show(ctx context.Context, modID string, name string) (*model.Version, error) {
	return nil, nil
}

// Create implements Create from versions.Store interface.
func (v *Versions) Create(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	return nil, nil
}

// Update implements Update from versions.Store interface.
func (v *Versions) Update(ctx context.Context, modID string, version *model.Version) (*model.Version, error) {
	return nil, nil
}

// Delete implements Delete from versions.Store interface.
func (v *Versions) Delete(ctx context.Context, modID string, name string) error {
	return nil
}
