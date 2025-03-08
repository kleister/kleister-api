package store

import (
	"context"
	"database/sql"
	"errors"
	"sort"
	"strings"

	quiltClient "github.com/kleister/kleister-api/pkg/internal/quilt"
	"github.com/kleister/kleister-api/pkg/model"
	gover "github.com/mcuadros/go-version"
	"github.com/uptrace/bun"
)

// Quilt provides all database operations related to quilt.
type Quilt struct {
	client *Store
}

// List implements the listing of all quilt versions.
func (s *Quilt) List(ctx context.Context, params model.ListParams) ([]*model.Quilt, int64, error) {
	records := make([]*model.Quilt, 0)

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

// Show implements the details for a specific quilt.
func (s *Quilt) Show(ctx context.Context, name string) (*model.Quilt, error) {
	record := &model.Quilt{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Where("id = ? OR slug = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrQuiltNotFound
		}

		return record, err
	}

	return record, nil
}

// Sync implements the sync of the given list of versions.
func (s *Quilt) Sync(ctx context.Context, versions quiltClient.Versions) error {
	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, version := range versions {
			if _, err := tx.NewInsert().
				Model(&model.Quilt{
					Name: version.Value,
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

// ListBuilds implements the listing of all builds for a quilt.
func (s *Quilt) ListBuilds(ctx context.Context, params model.QuiltBuildParams) ([]*model.Build, int64, error) {
	records := make([]*model.Build, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Where("quilt_id = ?", params.QuiltID)

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

// AttachBuild implements the attachment of a quilt to a build.
func (s *Quilt) AttachBuild(ctx context.Context, params model.QuiltBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	assigned, err := s.isBuildAssigned(ctx, params.QuiltID, build.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("quilt_id = ?", params.QuiltID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DropBuild implements the removal of a quilt from a build.
func (s *Quilt) DropBuild(ctx context.Context, params model.QuiltBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	unassigned, err := s.isBuildUnassigned(ctx, params.QuiltID, build.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("quilt_id = NULL").
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Quilt) isBuildAssigned(ctx context.Context, quiltID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("quilt_id = ? AND id = ?", quiltID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Quilt) isBuildUnassigned(ctx context.Context, quiltID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("quilt_id = ? AND id = ?", quiltID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}
