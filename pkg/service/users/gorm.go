package users

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/validate"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(handle *gorm.DB) *GormService {
	return &GormService{
		handle: handle,
	}
}

// ByBasicAuth implements the Service interface for database persistence.
func (s *GormService) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	record := &model.User{}

	if err := s.handle.WithContext(
		ctx,
	).Where(
		"username = ?",
		username,
	).Or(
		"email = ?",
		username,
	).First(
		record,
	).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(record.Hashword),
		[]byte(password),
	); err != nil {
		return nil, ErrWrongAuth
	}

	return record, nil
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context) ([]*model.User, error) {
	records := make([]*model.User, 0)

	if err := s.query(ctx).Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Show implements the Service interface for database persistence.
func (s *GormService) Show(ctx context.Context, name string) (*model.User, error) {
	record := &model.User{}

	err := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).First(
		record,
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, ErrNotFound
	}

	return record, err
}

// Create implements the Service interface for database persistence.
func (s *GormService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if user.Slug == "" {
		user.Slug = store.Slugify(
			tx.Model(&model.User{}),
			user.Username,
			"",
		)
	}

	if err := s.validate(ctx, user, false); err != nil {
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

// Update implements the Service interface for database persistence.
func (s *GormService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if user.Slug == "" {
		user.Slug = store.Slugify(
			tx.Model(&model.User{}),
			user.Username,
			user.ID,
		)
	}

	if err := s.validate(ctx, user, true); err != nil {
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

// Delete implements the Service interface for database persistence.
func (s *GormService) Delete(ctx context.Context, name string) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Delete(
		&model.User{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Exists implements the Service interface for database persistence.
func (s *GormService) Exists(ctx context.Context, name string) (bool, error) {
	res := s.query(ctx).Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Find(
		&model.User{},
	)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

// External implements the Service interface for database persistence.
func (s *GormService) External(ctx context.Context, username, email, fullname string, admin bool) (*model.User, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.User{}

	if err := tx.Where(
		&model.User{
			Username: username,
		},
	).First(
		record,
	).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	record.Username = username
	record.Email = email
	record.Fullname = fullname

	if record.ID == "" {
		if record.Slug == "" {
			record.Slug = store.Slugify(
				tx.Model(&model.User{}),
				record.Username,
				"",
			)
		}

		record.Password = secret.Generate(32)
		record.Active = true
		record.Admin = admin

		if err := tx.Create(record).Error; err != nil {
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
	} else {
		if record.Slug == "" {
			record.Slug = store.Slugify(
				tx.Model(&model.User{}),
				record.Username,
				record.ID,
			)
		}

		if err := tx.Save(record).Error; err != nil {
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
	}

	return record, nil
}

func (s *GormService) validate(ctx context.Context, record *model.User, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Slug,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "slug", record.ID)),
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
		validation.By(s.uniqueValueIsPresent(ctx, "username", record.ID)),
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

func (s *GormService) uniqueValueIsPresent(ctx context.Context, key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.handle.WithContext(
			ctx,
		).Where(
			fmt.Sprintf("%s = ?", key),
			val,
		)

		if id != "" {
			q = q.Not(
				"id = ?",
				id,
			)
		}

		if q.Find(
			&model.User{},
		).RowsAffected != 0 {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
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
