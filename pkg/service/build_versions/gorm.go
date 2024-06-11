package buildversions

import (
	"context"
	"errors"
	"strings"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	buildsService "github.com/kleister/kleister-api/pkg/service/builds"
	modsService "github.com/kleister/kleister-api/pkg/service/mods"
	packsService "github.com/kleister/kleister-api/pkg/service/packs"
	versionsService "github.com/kleister/kleister-api/pkg/service/versions"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle    *gorm.DB
	config    *config.Config
	principal *model.User
	packs     packsService.Service
	builds    buildsService.Service
	mods      modsService.Service
	versions  versionsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	cfg *config.Config,
	packs packsService.Service,
	builds buildsService.Service,
	mods modsService.Service,
	versions versionsService.Service,
) *GormService {
	return &GormService{
		handle:   handle,
		config:   cfg,
		packs:    packs,
		builds:   builds,
		mods:     mods,
		versions: versions,
	}
}

// WithPrincipal implements the Service interface for database persistence.
func (s *GormService) WithPrincipal(principal *model.User) Service {
	s.principal = principal
	return s
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, params model.BuildVersionParams) ([]*model.BuildVersion, int64, error) {
	counter := int64(0)
	records := make([]*model.BuildVersion, 0)
	q := s.query(ctx)

	switch {
	case params.BuildID != "" && params.VersionID == "":
		pack, err := s.packID(ctx, params.PackID)
		if err != nil {
			return nil, counter, err
		}

		build, err := s.buildID(ctx, pack, params.BuildID)
		if err != nil {
			return nil, counter, err
		}

		q = q.Where(
			"build_id = ?",
			build,
		)

		if val, ok := s.validBuildSort(params.Sort); ok {
			q = q.Order(strings.Join(
				[]string{
					val,
					sortOrder(params.Order),
				},
				" ",
			))
		}
	case params.VersionID != "" && params.BuildID == "":
		mod, err := s.modID(ctx, params.ModID)
		if err != nil {
			return nil, counter, err
		}

		version, err := s.versionID(ctx, mod, params.VersionID)
		if err != nil {
			return nil, counter, err
		}

		q = q.Where(
			"version_id = ?",
			version,
		)

		if val, ok := s.validVersionSort(params.Sort); ok {
			q = q.Order(strings.Join(
				[]string{
					val,
					sortOrder(params.Order),
				},
				" ",
			))
		}
	default:
		return nil, counter, ErrInvalidListParams
	}

	// if params.Search != "" {
	// 	opts := queryparser.Options{
	// 		CutFn: searchCut,
	// 		Allowed: []string{},
	// 	}

	// 	parser := queryparser.New(
	// 		params.Search,
	// 		opts,
	// 	).Parse()

	// 	for _, name := range opts.Allowed {
	// 		if parser.Has(name) {

	// 			q = q.Where(
	// 				fmt.Sprintf(
	// 					"%s LIKE ?",
	// 					name,
	// 				),
	// 				strings.ReplaceAll(
	// 					parser.GetOne(name),
	// 					"*",
	// 					"%",
	// 				),
	// 			)
	// 		}
	// 	}
	// }

	if err := q.Count(
		&counter,
	).Error; err != nil {
		return nil, counter, err
	}

	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, counter, err
	}

	return records, counter, nil
}

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, params model.BuildVersionParams) error {
	pack, err := s.packID(ctx, params.PackID)
	if err != nil {
		return err
	}

	build, err := s.buildID(ctx, pack, params.BuildID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, params.ModID)
	if err != nil {
		return err
	}

	version, err := s.versionID(ctx, mod, params.VersionID)
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
func (s *GormService) Drop(ctx context.Context, params model.BuildVersionParams) error {
	pack, err := s.packID(ctx, params.PackID)
	if err != nil {
		return err
	}

	build, err := s.buildID(ctx, pack, params.BuildID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, params.ModID)
	if err != nil {
		return err
	}

	version, err := s.versionID(ctx, mod, params.VersionID)
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
	record, err := s.builds.Show(ctx, model.BuildParams{
		PackID:  packID,
		BuildID: id,
	})

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
	record, err := s.versions.Show(ctx, model.VersionParams{
		ModID:     modID,
		VersionID: id,
	})

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
	).Joins(
		"Build",
	).Preload(
		"Version",
	).Joins(
		"Version",
	)
}

func (s *GormService) validBuildSort(val string) (string, bool) {
	if val == "" {
		return "Build.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug": "Build.slug",
		"name": "Build.name",
	} {
		if val == key {
			return name, true
		}
	}

	return "Build.name", true
}

func (s *GormService) validVersionSort(val string) (string, bool) {
	if val == "" {
		return "Version.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug": "Version.slug",
		"name": "Version.name",
	} {
		if val == key {
			return name, true
		}
	}

	return "Version.name", true
}
