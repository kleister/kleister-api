package fabric

import (
	"context"
	"errors"
	"strings"

	"github.com/kleister/kleister-api/pkg/config"
	fabricClient "github.com/kleister/kleister-api/pkg/internal/fabric"
	"github.com/kleister/kleister-api/pkg/model"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle    *gorm.DB
	config    *config.Config
	principal *model.User
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	cfg *config.Config,
) *GormService {
	return &GormService{
		handle: handle,
		config: cfg,
	}
}

// WithPrincipal implements the Service interface for database persistence.
func (s *GormService) WithPrincipal(principal *model.User) Service {
	s.principal = principal
	return s
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, params model.ListParams) ([]*model.Fabric, int64, error) {
	records := make([]*model.Fabric, 0)
	q := s.query(ctx)

	if params.Search != "" {
		term := strings.Join(
			[]string{
				"%",
				params.Search,
				"%",
			},
			"",
		)

		q = q.Or(
			"name LIKE ?",
			term,
		)
	}

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, 0, err
	}

	return records, int64(len(records)), nil
}

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, name string) (*model.Fabric, error) {
	record := &model.Fabric{}

	err := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"name = ?",
		name,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrNotFound
	}

	return record, err
}

// Sync implements the Store interface for database persistence.
func (s *GormService) Sync(ctx context.Context, versions fabricClient.Versions) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	for _, v := range versions {
		if err := tx.Where(
			&model.Fabric{
				Name: v.Value,
			},
		).FirstOrCreate(
			&model.Fabric{},
		).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

// ListBuilds implements the Service interface for database persistence.
func (s *GormService) ListBuilds(ctx context.Context, params model.FabricBuildParams) ([]*model.Build, int64, error) {
	parent, err := s.Show(ctx, params.FabricID)

	if err != nil {
		return nil, 0, err
	}

	counter := int64(0)
	records := make([]*model.Build, 0)

	q := s.handle.WithContext(
		ctx,
	).Model(
		&model.Build{},
	).Joins(
		"Pack",
	).Preload(
		"Pack",
	).Where(
		"fabric_id = ?",
		parent.ID,
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

	// if params.Search != "" {
	// 	opts := queryparser.Options{
	// 		CutFn: searchCut,
	// 		Allowed: []string{
	// 			"slug",
	// 			"name",
	// 		},
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

// AttachBuild implements the Service interface for database persistence.
func (s *GormService) AttachBuild(ctx context.Context, params model.FabricBuildParams) error {
	parent, err := s.Show(ctx, params.FabricID)

	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, params.PackID)
	if err != nil {
		return err
	}

	build, current, err := s.buildID(ctx, pack, params.BuildID)
	if err != nil {
		return err
	}

	if current == parent.ID {
		return ErrAlreadyAssigned
	}

	if err := s.handle.WithContext(
		ctx,
	).Table(
		"builds",
	).Where(
		"id = ?",
		build,
	).Update(
		"fabric_id",
		parent.ID,
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	return nil
}

// DropBuild implements the Service interface for database persistence.
func (s *GormService) DropBuild(ctx context.Context, params model.FabricBuildParams) error {
	parent, err := s.Show(ctx, params.FabricID)

	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, params.PackID)
	if err != nil {
		return err
	}

	build, current, err := s.buildID(ctx, pack, params.BuildID)
	if err != nil {
		return err
	}

	if current != parent.ID {
		return ErrNotAssigned
	}

	if err := s.handle.WithContext(
		ctx,
	).Table(
		"builds",
	).Where(
		"id = ?",
		build,
	).Update(
		"fabric_id",
		gorm.Expr("NULL"),
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	return nil
}

func (s *GormService) packID(ctx context.Context, id string) (string, error) {
	var (
		result string
	)

	if err := s.handle.WithContext(
		ctx,
	).Table(
		"packs",
	).Select(
		"id",
	).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		&result,
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return result, nil
}

func (s *GormService) buildID(ctx context.Context, packID, id string) (string, string, error) {
	type idAndFabric struct {
		ID       string
		FabricID string
	}

	result := idAndFabric{}

	if err := s.handle.WithContext(
		ctx,
	).Table(
		"builds",
	).Select(
		"id",
		"fabric_id",
	).Where(
		"pack_id = ?",
		packID,
	).Where(
		"id = ? OR name = ?",
		id,
		id,
	).First(
		&result,
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", ErrNotFound
		}

		return "", "", err
	}

	return result.ID, result.FabricID, nil
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Fabric{},
	)
}

func (s *GormService) validBuildSort(val string) (string, bool) {
	if val == "" {
		return "Build.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"build_name":   "Build.name",
		"build_public": "Build.public",
		"pack_slug":    "Pack.slug",
		"pack_name":    "Pack.name",
	} {
		if val == key {
			return name, true
		}
	}

	return "Build.name", true
}
