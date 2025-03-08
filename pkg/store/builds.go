package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/uptrace/bun"
)

// Builds provides all database operations related to builds.
type Builds struct {
	client *Store
}

// List implements the listing of all builds.
func (s *Builds) List(ctx context.Context, pack *model.Pack, params model.ListParams) ([]*model.Build, int64, error) {
	records := make([]*model.Build, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Where("build.pack_id = ?", pack.ID)

	if val, ok := s.ValidSort(params.Sort); ok {
		q = q.Order(strings.Join(
			[]string{
				val,
				sortOrder(params.Order),
			},
			" ",
		))
	}

	if params.Search != "" {
		q = s.client.SearchQuery(q, params.Search)
	}

	counter, err := q.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	if params.Limit > 0 {
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
}

// Show implements the details for a specific build.
func (s *Builds) Show(ctx context.Context, pack *model.Pack, name string) (*model.Build, error) {
	record := &model.Build{}

	q := s.client.handle.NewSelect().
		Model(record).
		Where("build.pack_id = ?", pack.ID).
		Where("build.id = ? OR name = ?", name, name)

	if err := q.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record, ErrBuildNotFound
		}

		return record, err
	}

	return record, nil
}

// Create implements the create of a new credential.
func (s *Builds) Create(ctx context.Context, pack *model.Pack, record *model.Build) error {
	if record.Name == "" {
		record.Name = s.slugify(
			ctx,
			"name",
			record.Name,
			"",
			pack.ID,
		)
	}

	if err := s.validate(ctx, record, false); err != nil {
		return err
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// Update implements the update of an existing credential.
func (s *Builds) Update(ctx context.Context, pack *model.Pack, record *model.Build) error {
	if record.Name == "" {
		record.Name = s.slugify(
			ctx,
			"name",
			record.Name,
			record.ID,
			pack.ID,
		)
	}

	if err := s.validate(ctx, record, true); err != nil {
		return err
	}

	q := s.client.handle.NewUpdate().
		Model(record).
		Where("pack_id = ?", pack.ID).
		Where("id = ?", record.ID)

	if _, err := q.Exec(ctx); err != nil {
		return err
	}

	return nil
}

// Delete implements the deletion of a credential.
func (s *Builds) Delete(ctx context.Context, pack *model.Pack, name string) error {
	record, err := s.Show(ctx, pack, name)

	if err != nil {
		return err
	}

	q := s.client.handle.NewDelete().
		Model((*model.Build)(nil)).
		Where("pack_id = ?", pack.ID).
		Where("id = ?", record.ID)

	if _, err := q.Exec(ctx); err != nil {
		return err
	}

	return nil
}

// ListVersions implements the listing of all versions for a build.
func (s *Builds) ListVersions(ctx context.Context, _ *model.Pack, build *model.Build, params model.ListParams) ([]*model.BuildVersion, int64, error) {
	records := make([]*model.BuildVersion, 0)

	q := s.client.handle.NewSelect().
		Model(&records).
		Relation("Version").
		Relation("Build").
		Where("build_id = ?", build.ID)

	if val, ok := s.client.Versions.ValidSort(params.Sort); ok {
		q = q.Order(strings.Join(
			[]string{
				val,
				sortOrder(params.Order),
			},
			" ",
		))
	}

	if params.Search != "" {
		q = s.client.SearchQuery(q, params.Search)
	}

	counter, err := q.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	if params.Limit > 0 {
		q = q.Limit(int(params.Limit))
	}

	if params.Offset > 0 {
		q = q.Offset(int(params.Offset))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, int64(counter), err
	}

	return records, int64(counter), nil
}

// AttachVersion implements the attachment of an build to a version.
func (s *Builds) AttachVersion(ctx context.Context, _ *model.Pack, build *model.Build, params model.BuildVersionParams) error {
	mod, err := s.client.Mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	version, err := s.client.Versions.Show(ctx, mod, params.VersionID)

	if err != nil {
		return err
	}

	assigned, err := s.isVersionAssigned(ctx, version.ID, build.ID)

	if err != nil {
		return err
	}

	if assigned {
		return ErrAlreadyAssigned
	}

	record := &model.BuildVersion{
		BuildID:   build.ID,
		VersionID: version.ID,
	}

	if _, err := s.client.handle.NewInsert().
		Model(record).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DropVersion implements the removal of an build from a version.
func (s *Builds) DropVersion(ctx context.Context, _ *model.Pack, build *model.Build, params model.BuildVersionParams) error {
	mod, err := s.client.Mods.Show(ctx, params.ModID)

	if err != nil {
		return err
	}

	version, err := s.client.Versions.Show(ctx, mod, params.VersionID)

	if err != nil {
		return err
	}

	unassigned, err := s.isVersionUnassigned(ctx, version.ID, build.ID)

	if err != nil {
		return err
	}

	if unassigned {
		return ErrNotAssigned
	}

	if _, err := s.client.handle.NewDelete().
		Model((*model.BuildVersion)(nil)).
		Where("build_id = ? AND version_id = ?", build.ID, version.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Builds) isVersionAssigned(ctx context.Context, versionID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.BuildVersion)(nil)).
		Where("version_id = ? AND build_id = ?", versionID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Builds) isVersionUnassigned(ctx context.Context, versionID, buildID string) (bool, error) {
	count, err := s.client.handle.NewSelect().
		Model((*model.BuildVersion)(nil)).
		Where("version_id = ? AND build_id = ?", versionID, buildID).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count < 1, nil
}

func (s *Builds) validate(ctx context.Context, record *model.Build, _ bool) error {
	errs := validate.Errors{}

	if err := validation.Validate(
		record.Name,
		validation.Required,
		validation.Length(3, 255),
		validation.By(s.uniqueValueIsPresent(ctx, "name", record.ID, record.PackID)),
	); err != nil {
		errs.Errors = append(errs.Errors, validate.Error{
			Field: "name",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (s *Builds) uniqueValueIsPresent(ctx context.Context, key, id, packID string) func(value interface{}) error {
	return func(value interface{}) error {
		val, _ := value.(string)

		q := s.client.handle.NewSelect().
			Model((*model.Build)(nil)).
			Where("pack_id = ?", packID).
			Where("? = ?", bun.Ident(key), val)

		if id != "" {
			q = q.Where(
				"id != ?",
				id,
			)
		}

		exists, err := q.Exists(ctx)

		if err != nil {
			return err
		}

		if exists {
			return errors.New("is already taken")
		}

		return nil
	}
}

func (s *Builds) slugify(ctx context.Context, column, value, id, packID string) string {
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

		query := s.client.handle.NewSelect().
			Model((*model.Version)(nil)).
			Where("pack_id = ?", packID).
			Where("? = ?", bun.Ident(column), slug)

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

// ValidSort validates the given sorting column.
func (s *Builds) ValidSort(val string) (string, bool) {
	if val == "" {
		return "build.name", true
	}

	val = strings.ToLower(val)

	for key, name := range map[string]string{
		"name":        "build.name",
		"latest":      "build.latest",
		"recommended": "build.recommended",
		"created":     "build.created_at",
		"updated":     "build.updated_at",
	} {
		if val == key {
			return name, true
		}
	}

	return "build.name", true
}
