package boltdb

import (
	"context"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/kleister/kleister-api/pkg/model"
)

// Builds implements builds.Store interface.
type Builds struct {
	client *botldbStore
}

// List implements List from builds.Store interface.
func (b *Builds) List(ctx context.Context, packID string) ([]*model.Build, error) {
	records := make([]*model.Build, 0)

	err := b.client.handle.Select(
		q.Eq("PackID", packID),
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

// Show implements Show from builds.Store interface.
func (b *Builds) Show(ctx context.Context, packID string, name string) (*model.Build, error) {
	return nil, nil
}

// Create implements Create from builds.Store interface.
func (b *Builds) Create(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	return nil, nil
}

// Update implements Update from builds.Store interface.
func (b *Builds) Update(ctx context.Context, packID string, build *model.Build) (*model.Build, error) {
	return nil, nil
}

// Delete implements Delete from builds.Store interface.
func (b *Builds) Delete(ctx context.Context, packID string, name string) error {
	return nil
}
