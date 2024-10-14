package users

import (
	"context"
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/kleister/kleister-api/pkg/validate"
	"golang.org/x/crypto/bcrypt"
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

// External implements the Service interface for database persistence.
func (s *GormService) External(ctx context.Context, provider, ref, username, email, fullname string, admin bool) (*model.User, error) {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.UserAuth{}

	if err := tx.Where(
		&model.UserAuth{
			Provider: provider,
			Ref:      ref,
		},
	).Preload(
		"User",
	).First(
		record,
	).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if record.Provider == "" {
		record.Provider = provider
	}

	if record.Ref == "" {
		record.Ref = ref
	}

	if record.User == nil {
		record.User = &model.User{}
	}

	record.User.Email = email
	record.User.Fullname = fullname
	record.User.Admin = admin

	if record.User.ID == "" {
		record.User.Username = username
		record.User.Password = secret.Generate(32)
		record.User.Active = true

		if err := tx.Create(record).Error; err != nil {
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
	} else {
		if err := tx.Save(record).Error; err != nil {
			return nil, err
		}

		if err := tx.Save(record.User).Error; err != nil {
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
	}

	return record.User, nil
}

// AuthByID implements the Service interface for database persistence.
func (s *GormService) AuthByID(ctx context.Context, userID string) (*model.User, error) {
	return s.Show(ctx, userID)
}

// AuthByCreds implements the Service interface.
func (s *GormService) AuthByCreds(ctx context.Context, username, password string) (*model.User, error) {
	record := &model.User{}

	err := s.query(ctx).Where(
		"username = ?",
		username,
	).Or(
		"email = ?",
		username,
	).First(
		record,
	).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return record, ErrNotFound
		}

		return record, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(record.Hashword),
		[]byte(password),
	); err != nil {
		return nil, ErrBadCredentials
	}

	return record, nil
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, params model.ListParams) ([]*model.User, int64, error) {
	counter := int64(0)
	records := make([]*model.User, 0)
	q := s.query(ctx)

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
	// 			"username",
	// 			"email",
	// 			"fullname",
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
func (s *GormService) Show(ctx context.Context, name string) (*model.User, error) {
	record := &model.User{}

	err := s.query(ctx).Debug().Where(
		"id = ?",
		name,
	).Or(
		"username = ?",
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
func (s *GormService) Create(ctx context.Context, record *model.User) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := s.validate(ctx, record, false); err != nil {
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Update implements the Service interface for database persistence.
func (s *GormService) Update(ctx context.Context, record *model.User) error {
	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := s.validate(ctx, record, true); err != nil {
		return err
	}

	if err := tx.Save(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
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
		"username = ?",
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
		"username = ?",
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

// Column implements the Service interface for database persistence.
func (s *GormService) Column(ctx context.Context, name, col string, val any) error {
	if err := s.handle.WithContext(
		ctx,
	).Table(
		"users",
	).Where(
		"id = ? OR name = ?",
		name,
		name,
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

func (s *GormService) validate(ctx context.Context, record *model.User, _ bool) error {
	errs := validate.Errors{}

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
	).Model(
		&model.User{},
	).Preload(
		"Auths",
	).Preload(
		"Teams",
	).Preload(
		"Teams.Team",
	)
}

func (s *GormService) validSort(val string) (string, bool) {
	if val == "" {
		return "username", true
	}

	val = strings.ToLower(val)

	for _, name := range []string{
		"username",
		"email",
		"fullname",
		"admin",
		"active",
	} {
		if val == name {
			return val, true
		}
	}

	return "username", true
}
