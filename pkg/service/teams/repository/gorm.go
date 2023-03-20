package repository

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/validate"
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

// GormRepository implements the TeamsRepository interface.
type GormRepository struct {
	handle *gorm.DB
}

// List implements the TeamsRepository interface.
func (r *GormRepository) List(ctx context.Context, _ string) ([]*model.Team, error) {
	records := make([]*model.Team, 0)

	// TODO: use search if given

	if err := r.query(ctx).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Create implements the TeamsRepository interface.
func (r *GormRepository) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if team.Slug == "" {
		team.Slug = store.Slugify(
			tx.Model(&model.Team{}),
			team.Name,
			"",
		)
	}

	team.ID = uuid.New().String()

	if err := r.validate(ctx, team, false); err != nil {
		return nil, err
	}

	if err := tx.Create(team).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return team, nil
}

// Update implements the TeamsRepository interface.
func (r *GormRepository) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if team.Slug == "" {
		team.Slug = store.Slugify(
			tx.Model(&model.Team{}),
			team.Name,
			team.ID,
		)
	}

	if err := r.validate(ctx, team, true); err != nil {
		return nil, err
	}

	if err := tx.Save(team).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return team, nil
}

// Show implements the TeamsRepository interface.
func (r *GormRepository) Show(ctx context.Context, id string) (*model.Team, error) {
	record := &model.Team{}

	err := r.query(ctx).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrTeamNotFound
	}

	return record, err
}

// Delete implements the TeamsRepository interface.
func (r *GormRepository) Delete(ctx context.Context, id string) error {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Delete(
		&model.Team{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the TeamsRepository interface.
func (r *GormRepository) Exists(ctx context.Context, id string) (bool, string, error) {
	record := &model.Team{}

	res := r.query(ctx).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Find(
		record,
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, "", nil
	}

	if res.Error != nil {
		return false, "", res.Error
	}

	return res.RowsAffected > 0, record.ID, nil
}

func (r *GormRepository) validate(ctx context.Context, record *model.Team, existing bool) error {
	errs := validate.Errors{}

	if existing {
		if err := validation.Validate(
			record.ID,
			validation.Required,
			is.UUIDv4,
			validation.By(r.uniqueValueIsPresent(ctx, "id", record.ID)),
		); err != nil {
			errs.Errors = append(errs.Errors, validate.Error{
				Field: "id",
				Error: err,
			})
		}
	}

	if err := validation.Validate(
		record.Slug,
		validation.Required,
		validation.Length(3, 255),
		validation.By(r.uniqueValueIsPresent(ctx, "slug", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
		validation.By(r.uniqueValueIsPresent(ctx, "name", record.ID)),
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

func (r *GormRepository) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		res := r.handle.WithContext(ctx).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		).Not(
			"id = ?",
			id,
		).Find(
			&model.Team{},
		)

		if res.RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (r *GormRepository) query(ctx context.Context) *gorm.DB {
	return r.handle.WithContext(
		ctx,
	).Order(
		"name ASC",
	).Model(
		&model.Team{},
	).Preload(
		"Users",
	).Preload(
		"Users.User",
	)
}
