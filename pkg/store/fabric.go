package store

import (
	"context"
	"database/sql"
	"errors"
	"sort"
	"strings"

	fabricClient "github.com/kleister/kleister-api/pkg/internal/fabric"
	"github.com/kleister/kleister-api/pkg/model"
	gover "github.com/mcuadros/go-version"
	"github.com/uptrace/bun"
)

// Fabric provides all database operations related to fabric.
type Fabric struct {
	client *Store
}

// List implements the listing of all fabric versions.
func (s *Fabric) List(ctx context.Context, params model.ListParams) ([]*model.Fabric, int64, error) {
	records := make([]*model.Fabric, 0)

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

// Show implements the details for a specific fabric.
func (s *Fabric) Show(ctx context.Context, name string) (*model.Fabric, error) {
	record := &model.Fabric{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Where("id = ? OR slug = ?", name, name).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrFabricNotFound
		}

		return record, err
	}

	return record, nil
}

// Sync implements the sync of the given list of versions.
func (s *Fabric) Sync(ctx context.Context, versions fabricClient.Versions) error {
	return s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, version := range versions {
			if _, err := tx.NewInsert().
				Model(&model.Fabric{
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

// ListBuilds implements the listing of all builds for a fabric.
func (s *Fabric) ListBuilds(ctx context.Context, params model.FabricBuildParams) ([]*model.Build, int64, error) {
	records := make([]*model.Build, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Where("fabric_id = ?", params.FabricID)

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

// AttachBuild implements the attachment of a fabric to a build.
func (s *Fabric) AttachBuild(ctx context.Context, params model.FabricBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	assigned, err := s.isBuildAssigned(ctx, params.FabricID, build.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("fabric_id = ?", params.FabricID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DropBuild implements the removal of a fabric from a build.
func (s *Fabric) DropBuild(ctx context.Context, params model.FabricBuildParams) error {
	pack, err := s.client.Packs.Show(ctx, params.PackID)

	if err != nil {
		return err
	}

	build, err := s.client.Builds.Show(ctx, pack, params.BuildID)

	if err != nil {
		return err
	}

	unassigned, err := s.isBuildUnassigned(ctx, params.FabricID, build.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewInsert().
		Model((*model.Build)(nil)).
		Set("fabric_id = NULL").
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Fabric) isBuildAssigned(ctx context.Context, fabricID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("fabric_id = ? AND id = ?", fabricID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Fabric) isBuildUnassigned(ctx context.Context, fabricID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.Build)(nil)).
		Where("fabric_id = ? AND id = ?", fabricID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}
