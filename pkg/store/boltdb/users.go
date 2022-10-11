package boltdb

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrPasswordEncrypt inditcates that bcrypt failed to create password.
	ErrPasswordEncrypt = errors.New("failed to encrypt password")
)

// Users implements users.Store interface.
type Users struct {
	client *botldbStore
}

// ByBasicAuth implements ByBasicAuth from users.Store interface.
func (u *Users) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	record := &model.User{}

	if err := u.client.handle.Select(
		q.Or(
			q.Eq("Username", username),
			q.Eq("Email", username),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
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

	err := u.client.handle.AllByIndex(
		"Username",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	for _, record := range records {
		teams, err := u.ListTeams(ctx, record.ID)

		if err != nil {
			return records, err
		}

		record.Teams = teams
	}

	return records, err
}

// Show implements Show from users.Store interface.
func (u *Users) Show(ctx context.Context, name string) (*model.User, error) {
	record := &model.User{}

	err := u.client.handle.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record)

	if err == storm.ErrNotFound {
		return record, users.ErrNotFound
	}

	teams, err := u.ListTeams(ctx, record.ID)

	if err != nil {
		return record, err
	}

	record.Teams = teams
	return record, err
}

// Create implements Create from users.Store interface.
func (u *Users) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if user.Password != "" && !strings.HasPrefix(user.Password, "$2a") {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
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

			if err := tx.Select(
				q.Eq("Slug", user.Slug),
			).First(new(model.User)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	user.ID = uuid.New().String()
	user.UpdatedAt = time.Now().UTC()
	user.CreatedAt = time.Now().UTC()

	if err := u.validateCreate(user); err != nil {
		return nil, err
	}

	if err := tx.Save(user); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

// Update implements Update from users.Store interface.
func (u *Users) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if user.Password != "" && !strings.HasPrefix(user.Password, "$2a") {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
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

			if err := tx.Select(
				q.And(
					q.Eq("Slug", user.Slug),
					q.Not(
						q.Eq("ID", user.ID),
					),
				),
			).First(new(model.User)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	user.UpdatedAt = time.Now().UTC()

	if err := u.validateUpdate(user); err != nil {
		return nil, err
	}

	if err := tx.Save(user); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete implements Delete from users.Store interface.
func (u *Users) Delete(ctx context.Context, name string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	record := &model.User{}

	if err := tx.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return users.ErrNotFound
		}

		return err
	}

	if err := tx.Select(
		q.Eq("UserID", record.ID),
	).Delete(new(model.TeamUser)); err != nil {
		return err
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

// ListTeams implements ListTeams from users.Store interface.
func (u *Users) ListTeams(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	if err := u.client.handle.Select(
		q.Eq("UserID", id),
	).Find(&records); err != nil && err != storm.ErrNotFound {
		return records, err
	}

	for _, record := range records {
		user := &model.User{}

		if err := u.client.handle.Select(
			q.Eq("ID", record.UserID),
		).First(user); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.User = user

		team := &model.Team{}

		if err := u.client.handle.Select(
			q.Eq("ID", record.TeamID),
		).First(team); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Team = team
	}

	return records, nil
}

// AppendTeam implements AppendTeam from teams.Store interface.
func (u *Users) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("TeamID", teamID),
		),
	).First(new(model.TeamUser)); err == nil {
		return users.ErrAlreadyAssigned
	}

	record := &model.TeamUser{
		UserID:    userID,
		TeamID:    teamID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := u.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitTeam implements PermitTeam from teams.Store interface.
func (u *Users) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("TeamID", teamID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := u.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropTeam implements DropTeam from teams.Store interface.
func (u *Users) DropTeam(ctx context.Context, userID, teamID string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("TeamID", teamID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

// ListMods implements ListMods from users.Store interface.
func (u *Users) ListMods(ctx context.Context, id string) ([]*model.UserMod, error) {
	records := make([]*model.UserMod, 0)

	if err := u.client.handle.Select(
		q.Eq("UserID", id),
	).Find(&records); err != nil && err != storm.ErrNotFound {
		return records, err
	}

	for _, record := range records {
		user := &model.User{}

		if err := u.client.handle.Select(
			q.Eq("ID", record.UserID),
		).First(user); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.User = user

		mod := &model.Mod{}

		if err := u.client.handle.Select(
			q.Eq("ID", record.ModID),
		).First(mod); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Mod = mod
	}

	return records, nil
}

// AppendMod implements AppendMod from users.Store interface.
func (u *Users) AppendMod(ctx context.Context, userID, modID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("ModID", modID),
		),
	).First(new(model.UserMod)); err == nil {
		return users.ErrAlreadyAssigned
	}

	record := &model.UserMod{
		UserID:    userID,
		ModID:     modID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := u.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitMod implements PermitMod from users.Store interface.
func (u *Users) PermitMod(ctx context.Context, userID, modID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.UserMod{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("ModID", modID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := u.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropMod implements DropMod from users.Store interface.
func (u *Users) DropMod(ctx context.Context, userID, modID string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.UserMod{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("ModID", modID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

// ListPacks implements ListPacks from users.Store interface.
func (u *Users) ListPacks(ctx context.Context, id string) ([]*model.UserPack, error) {
	records := make([]*model.UserPack, 0)

	if err := u.client.handle.Select(
		q.Eq("UserID", id),
	).Find(&records); err != nil && err != storm.ErrNotFound {
		return records, err
	}

	for _, record := range records {
		user := &model.User{}

		if err := u.client.handle.Select(
			q.Eq("ID", record.UserID),
		).First(user); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.User = user

		pack := &model.Pack{}

		if err := u.client.handle.Select(
			q.Eq("ID", record.PackID),
		).First(pack); err != nil && err != storm.ErrNotFound {
			return records, err
		}

		record.Pack = pack
	}

	return records, nil
}

// AppendPack implements AppendPack from users.Store interface.
func (u *Users) AppendPack(ctx context.Context, userID, packID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("PackID", packID),
		),
	).First(new(model.UserPack)); err == nil {
		return users.ErrAlreadyAssigned
	}

	record := &model.UserPack{
		UserID:    userID,
		PackID:    packID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := u.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitPack implements PermitPack from users.Store interface.
func (u *Users) PermitPack(ctx context.Context, userID, packID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.UserPack{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("PackID", packID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := u.validatePerm(record.Perm); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropPack implements DropPack from users.Store interface.
func (u *Users) DropPack(ctx context.Context, userID, packID string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.UserPack{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("PackID", packID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

func (u *Users) validateCreate(record *model.User) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Slug,
		validation.Length(3, 255),
		validation.By(u.uniqueValueIsPresent("slug", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Username,
		validation.Length(3, 255),
		validation.By(u.uniqueValueIsPresent("username", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Email,
		validation.Length(3, 255),
		is.EmailFormat,
		validation.By(u.uniqueValueIsPresent("email", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "email",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Password,
		validation.Required,
		validation.Length(3, 255),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "password",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (u *Users) validateUpdate(record *model.User) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.ID,
		validation.Required,
		is.UUIDv4,
		validation.By(u.uniqueValueIsPresent("id", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "id",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Slug,
		validation.Length(3, 255),
		validation.By(u.uniqueValueIsPresent("slug", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Username,
		validation.Length(3, 255),
		validation.By(u.uniqueValueIsPresent("username", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "username",
			Error: err,
		})
	}

	if err := validation.Validate(
		record.Email,
		validation.Length(3, 255),
		is.EmailFormat,
		validation.By(u.uniqueValueIsPresent("email", record.ID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "email",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (u *Users) validatePerm(perm string) error {
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

func (u *Users) uniqueValueIsPresent(key, id string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		if err := u.client.handle.Select(
			q.And(
				q.Eq(key, val),
				q.Not(
					q.Eq("ID", id),
				),
			),
		).First(new(model.User)); err == storm.ErrNotFound {
			return nil
		}

		return errors.New("taken")
	}
}
