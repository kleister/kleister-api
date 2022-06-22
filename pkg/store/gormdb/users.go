package gormdb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	// ErrPasswordEncrypt inditcates that bcrypt failed to create password.
	ErrPasswordEncrypt = errors.New("failed to encrypt password")
)

// Users implements users.Store interface.
type Users struct {
	client *gormdbStore
}

// ByBasicAuth implements ByBasicAuth from users.Store interface.
func (u *Users) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	record := &model.User{}

	if err := u.client.handle.Where(
		"username = ?",
		username,
	).Or(
		"email = ?",
		username,
	).First(
		record,
	).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, users.ErrNotFound
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(record.Password),
		[]byte(password),
	); err != nil {
		return nil, users.ErrWrongAuth
	}

	return record, nil
}

// List implements List from users.Store interface.
func (u *Users) List(ctx context.Context) ([]*model.User, error) {
	records := make([]*model.User, 0)

	err := u.client.handle.Order(
		"username ASC",
	).Model(
		&model.User{},
	).Preload(
		"Teams",
	).Preload(
		"Teams.User",
	).Preload(
		"Teams.Team",
	).Find(
		&records,
	).Error

	for _, record := range records {
		fmt.Printf("%+v\n", record)
	}

	return records, err
}

// Show implements Show from users.Store interface.
func (u *Users) Show(ctx context.Context, name string) (*model.User, error) {
	record := &model.User{}

	err := u.client.handle.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).Preload(
		"Teams",
	).Preload(
		"Teams.User",
	).Preload(
		"Teams.Team",
	).First(
		record,
	).Error

	if err == gorm.ErrRecordNotFound {
		return record, users.ErrNotFound
	}

	return record, err
}

// Create implements Create from users.Store interface.
func (u *Users) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx := u.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if user.Password != "" && !strings.HasPrefix(user.Password, "$2a") {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			tx.Rollback()
			return nil, ErrPasswordEncrypt
		}

		user.Password = string(encrypt)
	}

	if user.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				user.Slug = slugify.Slugify(user.Username)
			} else {
				user.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", user.Username, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				user.Slug,
			).First(
				&model.User{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	user.ID = uuid.New().String()

	if err := u.validateCreate(user); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Update implements Update from users.Store interface.
func (u *Users) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx := u.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if user.Password != "" && !strings.HasPrefix(user.Password, "$2a") {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			tx.Rollback()
			return nil, ErrPasswordEncrypt
		}

		user.Password = string(encrypt)
	}

	if user.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				user.Slug = slugify.Slugify(user.Username)
			} else {
				user.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", user.Username, i),
				)
			}

			if res := tx.Where(
				"slug = ?",
				user.Slug,
			).Not(
				"id",
				user.ID,
			).First(
				&model.User{},
			); errors.Is(res.Error, gorm.ErrRecordNotFound) {
				break
			}
		}
	}

	if err := u.validateUpdate(user); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete implements Delete from users.Store interface.
func (u *Users) Delete(ctx context.Context, name string) error {
	tx := u.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.User{}

	if err := tx.Where(
		"id = ?",
		name,
	).Or(
		"slug = ?",
		name,
	).First(
		record,
	).Error; err != nil {
		tx.Rollback()

		if err == gorm.ErrRecordNotFound {
			return users.ErrNotFound
		}

		return err
	}

	if err := tx.Where(
		"user_id = ?",
		record.ID,
	).Delete(
		&model.TeamUser{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ListTeams implements ListTeams from users.Store interface.
func (u *Users) ListTeams(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	err := u.client.handle.Where(
		"user_id = ?",
		id,
	).Model(
		&model.TeamUser{},
	).Preload(
		"User",
	).Preload(
		"Team",
	).Find(
		&records,
	).Error

	return records, err
}

// AppendTeam implements AppendTeam from teams.Store interface.
func (u *Users) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	if u.isAssignedToTeam(userID, teamID) {
		return users.ErrAlreadyAssigned
	}

	tx := u.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamUser{
		UserID: userID,
		TeamID: teamID,
		Perm:   perm,
	}

	if err := u.validatePerm(record); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// PermitTeam implements PermitTeam from teams.Store interface.
func (u *Users) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	if u.isUnassignedFromTeam(userID, teamID) {
		return users.ErrNotAssigned
	}

	tx := u.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	record := &model.TeamUser{}
	record.Perm = perm

	if err := u.validatePerm(record); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(
		"user_id = ? AND team_id = ?",
		userID,
		teamID,
	).Model(
		&model.TeamUser{},
	).Updates(
		record,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DropTeam implements DropTeam from teams.Store interface.
func (u *Users) DropTeam(ctx context.Context, userID, teamID string) error {
	if u.isUnassignedFromTeam(userID, teamID) {
		return users.ErrNotAssigned
	}

	tx := u.client.handle.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where(
		"user_id = ? AND team_id = ?",
		userID,
		teamID,
	).Delete(
		&model.TeamUser{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *Users) isAssignedToTeam(userID, teamID string) bool {
	res := u.client.handle.Where(
		"user_id = ? AND team_id = ?",
		userID,
		teamID,
	).Find(
		&model.TeamUser{},
	)

	return res.RowsAffected != 0
}

func (u *Users) isUnassignedFromTeam(userID, teamID string) bool {
	res := u.client.handle.Where(
		"user_id = ? AND team_id = ?",
		userID,
		teamID,
	).Find(
		&model.TeamUser{},
	)

	return res.RowsAffected == 0
}

func (u *Users) validateCreate(record *model.User) error {
	errs := validate.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if u.uniqueValueIsPresent("slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsEmail(record.Email); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "email",
			Error: fmt.Errorf("is not valid"),
		})
	}

	if u.uniqueValueIsPresent("email", record.Email, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "email",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Username, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if u.uniqueValueIsPresent("username", record.Username, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Password, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "password",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (u *Users) validateUpdate(record *model.User) error {
	errs := validate.Errors{}

	if ok := govalidator.IsUUIDv4(record.ID); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "id",
			Error: fmt.Errorf("is not a valid uuid v4"),
		})
	}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if u.uniqueValueIsPresent("slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsEmail(record.Email); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "email",
			Error: fmt.Errorf("is not valid"),
		})
	}

	if u.uniqueValueIsPresent("email", record.Email, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "email",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Username, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if u.uniqueValueIsPresent("username", record.Username, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (u *Users) validatePerm(record *model.TeamUser) error {
	if ok := govalidator.IsIn(record.Perm, "user", "admin", "owner"); !ok {
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

func (u *Users) uniqueValueIsPresent(key, val, id string) bool {
	res := u.client.handle.Where(
		fmt.Sprintf("%s = ?", key),
		val,
	).Not(
		"id = ?",
		id,
	).Find(
		&model.User{},
	)

	return res.RowsAffected != 0
}
