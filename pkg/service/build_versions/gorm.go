package buildVersions

import (
	"context"
	"errors"

	"github.com/kleister/kleister-api/pkg/model"
	buildsService "github.com/kleister/kleister-api/pkg/service/builds"
	modsService "github.com/kleister/kleister-api/pkg/service/mods"
	packsService "github.com/kleister/kleister-api/pkg/service/packs"
	versionsService "github.com/kleister/kleister-api/pkg/service/versions"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle   *gorm.DB
	packs    packsService.Service
	builds   buildsService.Service
	mods     modsService.Service
	versions versionsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	packs packsService.Service,
	builds buildsService.Service,
	mods modsService.Service,
	versions versionsService.Service,
) *GormService {
	return &GormService{
		handle:   handle,
		packs:    packs,
		builds:   builds,
		mods:     mods,
		versions: versions,
	}
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, packID, buildID, modID, versionID string) ([]*model.BuildVersion, error) {
	q := s.query(ctx)

	switch {
	case buildID != "" && versionID == "":
		pack, err := s.packID(ctx, packID)
		if err != nil {
			return nil, err
		}

		build, err := s.buildID(ctx, pack, buildID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"build_id = ?",
			build,
		)
	case versionID != "" && buildID == "":
		mod, err := s.modID(ctx, modID)
		if err != nil {
			return nil, err
		}

		version, err := s.versionID(ctx, mod, versionID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"version_id = ?",
			version,
		)
	default:
		return nil, ErrInvalidListParams
	}

	records := make([]*model.BuildVersion, 0)

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, packID, buildID, modID, versionID string) error {
	pack, err := s.packID(ctx, packID)
	if err != nil {
		return err
	}

	build, err := s.buildID(ctx, pack, buildID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	version, err := s.versionID(ctx, mod, versionID)
	if err != nil {
		return err
	}

	if s.isAssigned(ctx, build, version) {
		return ErrAlreadyAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.BuildVersion{
		BuildID:   build,
		VersionID: version,
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Drop implements the Service interface for database persistence.
func (s *GormService) Drop(ctx context.Context, packID, buildID, modID, versionID string) error {
	pack, err := s.packID(ctx, packID)
	if err != nil {
		return err
	}

	build, err := s.buildID(ctx, pack, buildID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	version, err := s.versionID(ctx, mod, versionID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, build, version) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"build_id = ? AND version_id = ?",
		build,
		version,
	).Delete(
		&model.BuildVersion{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (s *GormService) packID(ctx context.Context, id string) (string, error) {
	record, err := s.packs.Show(ctx, id)

	if err != nil {
		if errors.Is(err, packsService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) buildID(ctx context.Context, packID, id string) (string, error) {
	record, err := s.builds.Show(ctx, packID, id)

	if err != nil {
		if errors.Is(err, buildsService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) modID(ctx context.Context, id string) (string, error) {
	record, err := s.mods.Show(ctx, id)

	if err != nil {
		if errors.Is(err, modsService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) versionID(ctx context.Context, modID, id string) (string, error) {
	record, err := s.versions.Show(ctx, modID, id)

	if err != nil {
		if errors.Is(err, versionsService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
}

func (s *GormService) isAssigned(ctx context.Context, buildID, versionID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"build_id = ? AND version_id = ?",
		buildID,
		versionID,
	).Find(
		&model.BuildVersion{},
	)

	return res.RowsAffected != 0
}

func (s *GormService) isUnassigned(ctx context.Context, buildID, versionID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"build_id = ? AND version_id = ?",
		buildID,
		versionID,
	).Find(
		&model.BuildVersion{},
	)

	return res.RowsAffected == 0
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.BuildVersion{},
	).Preload(
		"Build",
	).Preload(
		"Build.Pack",
	).Preload(
		"Version",
	).Preload(
		"Version.Mod",
	)
}
