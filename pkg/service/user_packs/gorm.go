package userpacks

import (
	"context"
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	packsService "github.com/kleister/kleister-api/pkg/service/packs"
	usersService "github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle    *gorm.DB
	config    *config.Config
	principal *model.User
	users     usersService.Service
	packs     packsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	cfg *config.Config,
	users usersService.Service,
	packs packsService.Service,
) *GormService {
	return &GormService{
		handle: handle,
		config: cfg,
		users:  users,
		packs:  packs,
	}
}

// WithPrincipal implements the Service interface for database persistence.
func (s *GormService) WithPrincipal(principal *model.User) Service {
	s.principal = principal
	return s
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, params model.UserPackParams) ([]*model.UserPack, int64, error) {
	counter := int64(0)
	records := make([]*model.UserPack, 0)
	q := s.query(ctx).Debug()

	switch {
	case params.UserID != "" && params.PackID == "":
		user, err := s.userID(ctx, params.UserID)
		if err != nil {
			return nil, counter, err
		}

		q = q.Where(
			"user_id = ?",
			user,
		)

		if val, ok := s.validPackSort(params.Sort); ok {
			q = q.Order(strings.Join(
				[]string{
					val,
					sortOrder(params.Order),
				},
				" ",
			))
		}
	case params.PackID != "" && params.UserID == "":
		pack, err := s.packID(ctx, params.PackID)
		if err != nil {
			return nil, counter, err
		}

		q = q.Where(
			"pack_id = ?",
			pack,
		)

		if val, ok := s.validUserSort(params.Sort); ok {
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
func (s *GormService) Attach(ctx context.Context, params model.UserPackParams) error {
	user, err := s.userID(ctx, params.UserID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, params.PackID)
	if err != nil {
		return err
	}

	if s.isAssigned(ctx, user, pack) {
		return ErrAlreadyAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.UserPack{
		UserID: user,
		PackID: pack,
		Perm:   params.Perm,
	}

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Permit implements the Service interface for database persistence.
func (s *GormService) Permit(ctx context.Context, params model.UserPackParams) error {
	user, err := s.userID(ctx, params.UserID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, params.PackID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, user, pack) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.UserPack{}
	record.Perm = params.Perm

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Where(
		"user_id = ? AND pack_id = ?",
		user,
		pack,
	).Model(
		&model.UserPack{},
	).Updates(
		record,
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Drop implements the Service interface for database persistence.
func (s *GormService) Drop(ctx context.Context, params model.UserPackParams) error {
	user, err := s.userID(ctx, params.UserID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, params.PackID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, user, pack) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"user_id = ? AND pack_id = ?",
		user,
		pack,
	).Delete(
		&model.UserPack{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (s *GormService) userID(ctx context.Context, id string) (string, error) {
	record, err := s.users.Show(ctx, id)

	if err != nil {
		if errors.Is(err, usersService.ErrNotFound) {
			return "", ErrNotFound
		}

		return "", err
	}

	return record.ID, nil
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

func (s *GormService) isAssigned(ctx context.Context, userID, packID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"user_id = ? AND pack_id = ?",
		userID,
		packID,
	).Find(
		&model.UserPack{},
	)

	return res.RowsAffected != 0
}

func (s *GormService) isUnassigned(ctx context.Context, userID, packID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"user_id = ? AND pack_id = ?",
		userID,
		packID,
	).Find(
		&model.UserPack{},
	)

	return res.RowsAffected == 0
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.UserPack{},
	).Preload(
		"Team",
	).Joins(
		"Team",
	).Preload(
		"User",
	).Joins(
		"User",
	)
}

func (s *GormService) validatePerm(perm string) error {
	if err := validation.Validate(
		perm,
		validation.In("user", "admin", "owner"),
	); err != nil {
		return validate.Errors{
			Errors: []validate.Error{
				{
					Field: "perm",
					Error: fmt.Errorf("invalid permission value"),
				},
			},
		}
	}

	return nil
}

func (s *GormService) validUserSort(val string) (string, bool) {
	if val == "" {
		return "User.username", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"username": "User.username",
		"email":    "User.email",
		"fullname": "User.fullname",
		"active":   "User.active",
		"admin":    "User.admin",
	} {
		if val == key {
			return name, true
		}
	}

	return "User.username", true
}

func (s *GormService) validPackSort(val string) (string, bool) {
	if val == "" {
		return "Pack.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"slug": "Pack.slug",
		"name": "Pack.name",
	} {
		if val == key {
			return name, true
		}
	}

	return "Pack.name", true
}
