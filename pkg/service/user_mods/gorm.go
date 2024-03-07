package userMods

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	modsService "github.com/kleister/kleister-api/pkg/service/mods"
	usersService "github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
	users  usersService.Service
	mods   modsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	users usersService.Service,
	mods modsService.Service,
) *GormService {
	return &GormService{
		handle: handle,
		users:  users,
		mods:   mods,
	}
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, userID, modID string) ([]*model.UserMod, error) {
	q := s.query(ctx)

	switch {
	case userID != "" && modID == "":
		user, err := s.userID(ctx, userID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"user_id = ?",
			user,
		)
	case modID != "" && userID == "":
		mod, err := s.modID(ctx, modID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"mod_id = ?",
			mod,
		)
	default:
		return nil, ErrInvalidListParams
	}

	records := make([]*model.UserMod, 0)

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, userID, modID, perm string) error {
	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	if s.isAssigned(ctx, user, mod) {
		return ErrAlreadyAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.UserMod{
		UserID: user,
		ModID:  mod,
		Perm:   perm,
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
func (s *GormService) Permit(ctx context.Context, userID, modID, perm string) error {
	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, user, mod) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	record := &model.UserMod{}
	record.Perm = perm

	if err := s.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Where(
		"user_id = ? AND mod_id = ?",
		user,
		mod,
	).Model(
		&model.UserMod{},
	).Updates(
		record,
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

// Drop implements the Service interface for database persistence.
func (s *GormService) Drop(ctx context.Context, userID, modID string) error {
	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	mod, err := s.modID(ctx, modID)
	if err != nil {
		return err
	}

	if s.isUnassigned(ctx, user, mod) {
		return ErrNotAssigned
	}

	tx := s.handle.WithContext(
		ctx,
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"user_id = ? AND mod_id = ?",
		user,
		mod,
	).Delete(
		&model.UserMod{},
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

func (s *GormService) isAssigned(ctx context.Context, userID, modID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"user_id = ? AND mod_id = ?",
		userID,
		modID,
	).Find(
		&model.UserMod{},
	)

	return res.RowsAffected != 0
}

func (s *GormService) isUnassigned(ctx context.Context, userID, modID string) bool {
	res := s.handle.WithContext(
		ctx,
	).Where(
		"user_id = ? AND mod_id = ?",
		userID,
		modID,
	).Find(
		&model.UserMod{},
	)

	return res.RowsAffected == 0
}

func (s *GormService) query(ctx context.Context) *gorm.DB {
	return s.handle.WithContext(
		ctx,
	).Model(
		&model.UserMod{},
	).Preload(
		"User",
	).Preload(
		"Mod",
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
