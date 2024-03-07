package userPacks

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	packsService "github.com/kleister/kleister-api/pkg/service/packs"
	usersService "github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
	"gorm.io/gorm"
)

// GormService defines the service to store content within a database based on Gorm.
type GormService struct {
	handle *gorm.DB
	users  usersService.Service
	packs  packsService.Service
}

// NewGormService initializes the service to store content within a database based on Gorm.
func NewGormService(
	handle *gorm.DB,
	users usersService.Service,
	packs packsService.Service,
) *GormService {
	return &GormService{
		handle: handle,
		users:  users,
		packs:  packs,
	}
}

// List implements the Service interface for database persistence.
func (s *GormService) List(ctx context.Context, userID, packID string) ([]*model.UserPack, error) {
	q := s.query(ctx)

	switch {
	case userID != "" && packID == "":
		user, err := s.userID(ctx, userID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"user_id = ?",
			user,
		)
	case packID != "" && userID == "":
		pack, err := s.packID(ctx, packID)
		if err != nil {
			return nil, err
		}

		q = q.Where(
			"pack_id = ?",
			pack,
		)
	default:
		return nil, ErrInvalidListParams
	}

	records := make([]*model.UserPack, 0)

	if err := q.Find(
		&records,
	).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// Attach implements the Service interface for database persistence.
func (s *GormService) Attach(ctx context.Context, userID, packID, perm string) error {
	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, packID)
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
func (s *GormService) Permit(ctx context.Context, userID, packID, perm string) error {
	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, packID)
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
	record.Perm = perm

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
func (s *GormService) Drop(ctx context.Context, userID, packID string) error {
	user, err := s.userID(ctx, userID)
	if err != nil {
		return err
	}

	pack, err := s.packID(ctx, packID)
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
	).Preload(
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
