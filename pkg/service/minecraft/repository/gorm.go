package repository

import (
	"context"

	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/model"
	"gorm.io/gorm"
)

// NewGormRepository initializes a new repository for GormDB.
func NewGormRepository(
	handle *gorm.DB,
) *GormRepository {
	return &GormRepository{
		handle: handle,
	}
}

// GormRepository implements the MinecraftRepository interface.
type GormRepository struct {
	handle *gorm.DB
}

// Search implements the MinecraftRepository interface.
func (r *GormRepository) Search(ctx context.Context, search string) ([]*model.Minecraft, error) {
	records := make([]*model.Minecraft, 0)
	q := r.query(ctx)

	if search != "" {
		q = q.Or(
			"name LIKE ?",
			"%"+search+"%",
		)
	}

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Update implements the MinecraftRepository interface.
func (r *GormRepository) Update(ctx context.Context) error {
	available, err := version.FromDefault()

	if err != nil {
		return err
	}

	version.ByVersion(
		available.Releases,
	).Sort()

	f := &version.Filter{
		Version: ">=1.7.10",
	}

	for _, row := range available.Releases.Filter(f) {
		if err := r.handle.WithContext(
			ctx,
		).Where(
			model.Minecraft{
				Name: row.ID,
			},
		).Assign(
			model.Minecraft{
				Type: "release",
			},
		).FirstOrCreate(&model.Minecraft{}).Error; err != nil {
			return err
		}
	}

	return nil
}

// ListBuilds implements the MinecraftRepository interface.
func (r *GormRepository) ListBuilds(ctx context.Context, id, _ string) ([]*model.Build, error) {
	records := make([]*model.Build, 0)

	// TODO: use search if given

	if err := r.handle.WithContext(
		ctx,
	).Where(
		"forge_id = ?",
		id,
	).Order(
		"name ASC",
	).Model(
		&model.Build{},
	).Preload(
		"Pack",
	).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *GormRepository) query(ctx context.Context) *gorm.DB {
	return r.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Minecraft{},
	)
}
