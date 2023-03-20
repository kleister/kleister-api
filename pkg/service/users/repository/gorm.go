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
func NewGormRepository(handle *gorm.DB) *GormRepository {
	return &GormRepository{
		handle: handle,
	}
}

// GormRepository implements the UsersRepository interface.
type GormRepository struct {
	handle *gorm.DB
}

// List implements the UsersRepository interface.
func (r *GormRepository) List(ctx context.Context, _ string) ([]*model.User, error) {
	records := make([]*model.User, 0)

	// TODO: use search if given

	if err := r.query(ctx).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Create implements the UsersRepository interface.
func (r *GormRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if user.Slug == "" {
		user.Slug = store.Slugify(
			tx.Model(&model.User{}),
			user.Username,
			"",
		)
	}

	user.ID = uuid.New().String()

	if err := r.validate(ctx, user, false); err != nil {
		return nil, err
	}

	if err := tx.Create(user).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Update implements the UsersRepository interface.
func (r *GormRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx := r.handle.WithContext(ctx).Begin()
	defer tx.Rollback()

	if user.Slug == "" {
		user.Slug = store.Slugify(
			tx.Model(&model.User{}),
			user.Username,
			user.ID,
		)
	}

	if err := r.validate(ctx, user, true); err != nil {
		return nil, err
	}

	if err := tx.Save(user).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Show implements the UsersRepository interface.
func (r *GormRepository) Show(ctx context.Context, id string) (*model.User, error) {
	record := &model.User{}

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
		return record, ErrUserNotFound
	}

	return record, err
}

// Delete implements the UsersRepository interface.
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
		&model.User{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the UsersRepository interface.
func (r *GormRepository) Exists(ctx context.Context, id string) (bool, string, error) {
	record := &model.User{}

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

func (r *GormRepository) validate(ctx context.Context, record *model.User, existing bool) error {
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
		record.Username,
		validation.Required,
		validation.Length(3, 255),
		validation.By(r.uniqueValueIsPresent(ctx, "username", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
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
			&model.User{},
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
		"username ASC",
	).Model(
		&model.User{},
	).Preload(
		"Teams",
	).Preload(
		"Teams.Team",
	)
}
