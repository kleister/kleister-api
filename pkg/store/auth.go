package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Machiel/slugify"
	"github.com/dchest/uniuri"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

// Auth provides all database operations related to auth.
type Auth struct {
	client *Store
}

// External creates and authenticates external users.
func (s *Auth) External(ctx context.Context, provider, ref, username, email, fullname string, admin bool) (*model.User, error) {
	auth := &model.UserAuth{}
	record := &model.User{}

	if err := s.client.handle.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if s.client.scim.Enabled {
			if err := tx.NewSelect().
				Model(record).
				Relation("Auths").
				Where("user.scim = ?", ref).
				Scan(ctx); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return err
			}

			for _, row := range record.Auths {
				if row.Provider == provider && row.Ref == ref {
					auth = row
				}
			}

			if record.Scim != "" {
				record.Admin = admin

				if _, err := tx.NewUpdate().
					Model(record).
					Where("id = ?", record.ID).
					Exec(ctx); err != nil {
					return err
				}

				auth.UserID = record.ID
				auth.Provider = provider
				auth.Ref = ref
				auth.Login = username
				auth.Email = email
				auth.Name = fullname

				if auth.ID != "" {
					if _, err := tx.NewUpdate().
						Model(auth).
						Where("id = ?", auth.ID).
						Exec(ctx); err != nil {
						return err
					}
				} else {
					if _, err := tx.NewInsert().
						Model(auth).
						Exec(ctx); err != nil {
						return err
					}
				}

				record.Auths = append(
					record.Auths,
					auth,
				)

				return nil
			}
		}

		if err := tx.NewSelect().
			Model(record).
			Relation("Auths").
			Join("JOIN user_auths AS user_auth ON user.id = user_auth.user_id").
			Where("user_auth.provider = ?", provider).
			Where("user_auth.ref = ?", ref).
			Scan(ctx); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		for _, row := range record.Auths {
			if row.Provider == provider && row.Ref == ref {
				auth = row
			}
		}

		record.Email = email
		record.Fullname = fullname
		record.Admin = admin

		if record.ID != "" {
			if _, err := tx.NewUpdate().
				Model(record).
				Where("user.id = ?", record.ID).
				Exec(ctx); err != nil {
				return err
			}
		} else {
			record.Active = true
			record.Username = s.slugify(ctx, tx, username, "")
			record.Password = secret.Generate(32)

			if _, err := tx.NewInsert().
				Model(record).
				Exec(ctx); err != nil {
				return err
			}
		}

		auth.UserID = record.ID
		auth.Provider = provider
		auth.Ref = ref
		auth.Login = username
		auth.Email = email
		auth.Name = fullname

		if auth.ID != "" {
			if _, err := tx.NewUpdate().
				Model(auth).
				Where("id = ?", auth.ID).
				Exec(ctx); err != nil {
				return err
			}
		} else {
			if _, err := tx.NewInsert().
				Model(auth).
				Exec(ctx); err != nil {
				return err
			}

			record.Auths = append(
				record.Auths,
				auth,
			)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return record, nil
}

func (s *Auth) slugify(ctx context.Context, db bun.Tx, value, id string) string {
	var (
		slug string
	)

	for i := 0; true; i++ {
		if i == 0 {
			slug = slugify.Slugify(value)
		} else {
			slug = slugify.Slugify(
				fmt.Sprintf("%s-%s", value, uniuri.NewLen(6)),
			)
		}

		query := db.NewSelect().
			Model((*model.User)(nil)).
			Where("username = ?", slug)

		if id != "" {
			query = query.Where(
				"id != ?",
				id,
			)
		}

		if count, err := query.Count(
			ctx,
		); err == nil && count == 0 {
			break
		}
	}

	return slug
}

// ByID tries to authenticate a user based on identifier.
func (s *Auth) ByID(ctx context.Context, userID string) (*model.User, error) {
	record := &model.User{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Relation("Auths").
		Relation("Groups").
		Relation("Groups.Group").
		Relation("Groups.Group.Mods").
		Relation("Groups.Group.Packs").
		Relation("Mods").
		Relation("Mods.Mod").
		Relation("Packs").
		Relation("Packs.Pack").
		Where("id = ?", userID).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrUserNotFound
		}

		return record, err
	}

	return record, nil
}

// ByCreds tries to authenticate a user based on credentials.
func (s *Auth) ByCreds(ctx context.Context, username, password string) (*model.User, error) {
	record := &model.User{}

	if err := s.client.handle.NewSelect().
		Model(record).
		Relation("Auths").
		Relation("Groups").
		Relation("Groups.Group").
		Relation("Groups.Group.Mods").
		Relation("Groups.Group.Packs").
		Relation("Mods").
		Relation("Mods.Mod").
		Relation("Packs").
		Relation("Packs.Pack").
		Where("username = ?", username).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrUserNotFound
		}

		return record, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(record.Hashword),
		[]byte(password),
	); err != nil {
		return nil, ErrWrongCredentials
	}

	return record, nil
}
