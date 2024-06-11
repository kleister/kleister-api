package versions

import (
	"context"
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle    *gorm.DB
	config    *config.Config
	principal *model.User
	mods      mods.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	cfg *config.Config,
	modsService mods.Service,
) *GormService {
	return &GormService{
		handle: handle,
		config: cfg,
		mods:   modsService,
	}
}

// WithPrincipal implements the Service interface for database persistence.
func (s *GormService) WithPrincipal(principal *model.User) Service {
	s.principal = principal
	return s
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, params model.VersionParams) ([]*model.Version, int64, error) {
	counter := int64(0)
	records := make([]*model.Version, 0)

	parent, err := s.mods.Show(ctx, params.ModID)

	if err != nil {
		return nil, counter, err
	}

	q := s.query(ctx, parent.ID)

	if val, ok := s.validSort(params.Sort); ok {
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

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, params model.VersionParams) (*model.Version, error) {
	record := &model.Version{}

	parent, err := s.mods.Show(ctx, params.ModID)

	if err != nil {
		return nil, err
	}

	if err := s.query(ctx, parent.ID).Where(
		"id = ? OR name = ?",
		params.VersionID,
		params.VersionID,
	).First(
		record,
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return record, ErrNotFound
		}

		return nil, err
	}

	return record, nil
}

// Create implements the Service interface for database persistence.
func (s *GormService) Create(ctx context.Context, params model.VersionParams, record *model.Version) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	parent, err := s.mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	record.ModID = parent.ID

	if err := s.validate(ctx, record, false); err != nil {
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Update implements the Service interface for database persistence.
func (s *GormService) Update(ctx context.Context, params model.VersionParams, record *model.Version) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	parent, err := s.mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	record.ModID = parent.ID
	record.ID = params.VersionID

	if err := s.validate(ctx, record, true); err != nil {
		return err
	}

	if err := tx.Save(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Delete implements the Service interface for database persistence.
func (s *GormService) Delete(ctx context.Context, params model.VersionParams) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	parent, err := s.mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	if err := tx.Where(
		"mod_id = ?",
		parent.ID,
	).Where(
		"id = ? OR name = ?",
		params.VersionID,
		params.VersionID,
	).Delete(
		&model.Version{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the Service interface for database persistence.
func (s *GormService) Exists(ctx context.Context, params model.VersionParams) (bool, error) {
	parent, err := s.mods.Show(ctx, params.ModID)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return false, nil
		}

		return false, err
	}

	res := s.query(ctx, parent.ID).Where(
		"id = ? OR name = ?",
		params.VersionID,
		params.VersionID,
	).Find(
		&model.Version{},
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

// Column implements the Service interface for database persistence.
func (s *GormService) Column(ctx context.Context, params model.VersionParams, col string, val any) error {
	parent, err := s.mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	if err := s.handle.WithContext(
		ctx,
	).Table(
		"versions",
	).Where(
		"mod_id = ?",
		parent.ID,
	).Where(
		"id = ? OR name = ?",
		params.VersionID,
		params.VersionID,
	).Update(
		col,
		val,
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return err
	}

	return nil
}

func (s *GormService) validate(ctx context.Context, record *model.Version, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "name", record.ID, record.ModID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (s *GormService) uniqueValueIsPresent(ctx context.Context, key, id, modID string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.handle.WithContext(
			ctx,
		).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		).Where(
			"mod_id = ?",
			modID,
		)

		if id != "" {
			q = q.Not(
				"id = ?",
				id,
			)
		}

		if q.Find(
			&model.Version{},
		).RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (s *GormService) query(ctx context.Context, modID string) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.Version{},
	).Where(
		"mod_id = ?",
		modID,
	).Preload(
		"Mod",
	)
}

func (s *GormService) validSort(val string) (string, bool) {
	if val == "" {
		return "name", true
	}

	val = strings.ToLower(val)

	for _, name := range []string{
		"name",
	} {
		if val == name {
			return val, true
		}
	}

	return "name", true
}
