package store

import (
	"context"
	"database/sql"
	"errors"
	"sort"
	"strings"

	neoforgeClient "github.com/kleister/kleister-api/pkg/internal/neoforge"
	"github.com/kleister/kleister-api/pkg/model"
	gover "github.com/mcuadros/go-version"
	"github.com/uptrace/bun"
)

// Neoforge provides all database operations related to neoforge.
type Neoforge struct {
	client *Store
}

// List implements the listing of all neoforge versions.
func (s *Neoforge) List(ctx context.Context, params model.ListParams) ([]*model.Neoforge, int64, error) {
	records := make([]*model.Neoforge, 0)

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

// Show implements the details for a specific neoforge.
func (s *Neoforge) Show(ctx context.Context, name string) (*model.Neoforge, error) {
	record := &model.Neoforge{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Where("id = ? OR slug = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrNeoforgeNotFound
		}

		return record, err
	}

	return record, nil
}

// Sync implements the sync of the given list of versions.
func (s *Neoforge) Sync(ctx context.Context, versions neoforgeClient.Versions) error {
	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, version := range versions {
			if _, err := tx.NewInsert().
				Model(&model.Neoforge{
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

// ListBuilds implements the listing of all builds for a neoforge.
func (s *Neoforge) ListBuilds(ctx context.Context, params model.NeoforgeBuildParams) ([]*model.Build, int64, error) {
	records := make([]*model.Build, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Where("neoforge_id = ?", params.NeoforgeID)

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

// AttachBuild implements the attachment of a neoforge to a build.
func (s *Neoforge) AttachBuild(ctx context.Context, params model.NeoforgeBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	assigned, err := s.isBuildAssigned(ctx, params.NeoforgeID, build.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("neoforge_id = ?", params.NeoforgeID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DropBuild implements the removal of a neoforge from a build.
func (s *Neoforge) DropBuild(ctx context.Context, params model.NeoforgeBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	unassigned, err := s.isBuildUnassigned(ctx, params.NeoforgeID, build.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("neoforge_id = NULL").
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Neoforge) isBuildAssigned(ctx context.Context, neoforgeID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("neoforge_id = ? AND id = ?", neoforgeID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Neoforge) isBuildUnassigned(ctx context.Context, neoforgeID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("neoforge_id = ? AND id = ?", neoforgeID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}
