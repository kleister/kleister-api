package gormdb

import (
	"context"

	"github.com/kleister/kleister-api/pkg/model"
)

// Versions implements versions.Store interface.
type Versions struct {
	client *gormdbStore
}

// List implements List from versions.Store interface.
func (v *Versions) List(ctx context.Context, modID string) ([]*model.Version, error) {
	records := make([]*model.Version, 0)

	err := v.client.handle.Where(
		"mod_id = ?",
		modID,
	).Order(
		"name ASC",
	).Model(
		&model.Version{},
	).Find(
		&records,
	).Error

	return records, err
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
