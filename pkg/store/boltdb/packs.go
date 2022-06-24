package boltdb

import (
	"context"
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/google/uuid"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/validate"
)

// Packs implements packs.Store interface.
type Packs struct {
	client *botldbStore
}

// List implements List from packs.Store interface.
func (p *Packs) List(ctx context.Context) ([]*model.Pack, error) {
	records := make([]*model.Pack, 0)

	err := p.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	// for _, record := range records {
	// 	users, err := t.ListUsers(ctx, record.ID)

	// 	if err != nil {
	// 		return records, err
	// 	}

	// 	record.Users = users
	// }

	return records, nil
}

// Show implements Show from packs.Store interface.
func (p *Packs) Show(ctx context.Context, name string) (*model.Pack, error) {
	record := &model.Pack{}

	if err := p.client.handle.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return record, packs.ErrNotFound
		}

		return nil, err
	}

	// users, err := t.ListUsers(ctx, record.ID)

	// if err != nil {
	// 	return record, err
	// }

	// record.Users = users
	return record, nil
}

// Create implements Create from packs.Store interface.
func (p *Packs) Create(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	tx, err := p.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if pack.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				pack.Slug = slugify.Slugify(pack.Name)
			} else {
				pack.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", pack.Name, i),
				)
			}

			if err := tx.Select(
				q.Eq("Slug", pack.Slug),
			).First(new(model.Pack)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	pack.ID = uuid.New().String()
	pack.UpdatedAt = time.Now().UTC()
	pack.CreatedAt = time.Now().UTC()

	if err := p.validateCreate(pack); err != nil {
		return nil, err
	}

	if err := tx.Save(pack); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return pack, nil
}

// Update implements Update from packs.Store interface.
func (p *Packs) Update(ctx context.Context, pack *model.Pack) (*model.Pack, error) {
	tx, err := p.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if pack.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				pack.Slug = slugify.Slugify(pack.Name)
			} else {
				pack.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", pack.Name, i),
				)
			}

			if err := tx.Select(
				q.And(
					q.Eq("Slug", pack.Slug),
					q.Not(
						q.Eq("ID", pack.ID),
					),
				),
			).First(new(model.Pack)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	pack.UpdatedAt = time.Now().UTC()

	if err := p.validateUpdate(pack); err != nil {
		return nil, err
	}

	if err := tx.Save(pack); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return pack, nil
}

// Delete implements Delete from packs.Store interface.
func (p *Packs) Delete(ctx context.Context, name string) error {
	tx, err := p.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	record := &model.Pack{}
	if err := tx.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return packs.ErrNotFound
		}

		return err
	}

	// if err := tx.Select(
	// 	q.Eq("PackID", record.ID),
	// ).Delete(new(model.UserPack)); err != nil {
	// 	return err
	// }

	// if err := tx.DeleteStruct(record); err != nil {
	// 	return err
	// }

	return tx.Commit()
}

func (p *Packs) validateCreate(record *model.Pack) error {
	errs := validate.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if p.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Name, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if p.uniqueValueIsPresent("Name", record.Name, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (p *Packs) validateUpdate(record *model.Pack) error {
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

	if p.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Name, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if p.uniqueValueIsPresent("Name", record.Name, record.ID) {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (p *Packs) uniqueValueIsPresent(key, val, id string) bool {
	if err := p.client.handle.Select(
		q.And(
			q.Eq(key, val),
			q.Not(
				q.Eq("ID", id),
			),
		),
	).First(new(model.Pack)); err == storm.ErrNotFound {
		return false
	}

	return true
}
