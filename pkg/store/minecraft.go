package store

import (
	"context"
	"database/sql"
	"errors"
	"sort"
	"strings"

	"github.com/kleister/go-minecraft/version"
	"github.com/kleister/kleister-api/pkg/model"
	gover "github.com/mcuadros/go-version"
	"github.com/uptrace/bun"
)

// Minecraft provides all database operations related to minecraft.
type Minecraft struct {
	client *Store
}

// List implements the listing of all minecraft versions.
func (s *Minecraft) List(ctx context.Context, params model.ListParams) ([]*model.Minecraft, int64, error) {
	records := make([]*model.Minecraft, 0)

	q := s.client.handle.NewSelect().
		Model(&records)

	if params.Search != "" {
		q = s.client.SearchQuery(q, params.Search)
	}

	if params.Limit > 0 {
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(0), err
	}

	sort.Slice(records, func(i, j int) bool {
		return gover.Compare(records[i].Name, records[j].Name, "<")
	})

	return records, int64(len(records)), nil
}

// Show implements the details for a specific minecraft.
func (s *Minecraft) Show(ctx context.Context, name string) (*model.Minecraft, error) {
	record := &model.Minecraft{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Where("id = ? OR slug = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrMinecraftNotFound
		}

		return record, err
	}

	return record, nil
}

// Sync implements the sync of the given list of versions.
func (s *Minecraft) Sync(ctx context.Context, versions version.Versions) error {
	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, version := range versions {
			if _, err := tx.NewInsert().
				Model(&model.Minecraft{
					Name: version.ID,
					Type: "release",
				}).
				On("CONFLICT (name) DO UPDATE").
				Set("id = EXCLUDED.id").
				Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	})
}

// ListBuilds implements the listing of all builds for a minecraft.
func (s *Minecraft) ListBuilds(ctx context.Context, params model.MinecraftBuildParams) ([]*model.Build, int64, error) {
	records := make([]*model.Build, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Where("minecraft_id = ?", params.MinecraftID)

	if val, ok := s.client.Builds.ValidSort(params.Sort); ok {
		q = q.Order(strings.Join(
			[]string{
				val,
				sortOrder(params.Order),
			},
			" ",
		))
	}

	if params.Search != "" {
		q = s.client.SearchQuery(q, params.Search)
	}

	counter, err := q.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	if params.Limit > 0 {
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
}

// AttachBuild implements the attachment of a minecraft to a build.
func (s *Minecraft) AttachBuild(ctx context.Context, params model.MinecraftBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	assigned, err := s.isBuildAssigned(ctx, params.MinecraftID, build.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("minecraft_id = ?", params.MinecraftID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DropBuild implements the removal of a minecraft from a build.
func (s *Minecraft) DropBuild(ctx context.Context, params model.MinecraftBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	unassigned, err := s.isBuildUnassigned(ctx, params.MinecraftID, build.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("minecraft_id = NULL").
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Minecraft) isBuildAssigned(ctx context.Context, minecraftID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("minecraft_id = ? AND id = ?", minecraftID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Minecraft) isBuildUnassigned(ctx context.Context, minecraftID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("minecraft_id = ? AND id = ?", minecraftID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}
