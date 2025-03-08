package scim

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/dchest/uniuri"
	"github.com/elimity-com/scim"
	serrors "github.com/elimity-com/scim/errors"
	"github.com/elimity-com/scim/optional"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/rs/zerolog"
	"github.com/scim2/filter-parser/v2"
	"github.com/uptrace/bun"
)

var (
	groupAttributeMapping = map[string]string{
		"displayName": "name",
	}
)

type groupHandlers struct {
	config config.Scim
	store  *bun.DB
	logger zerolog.Logger
}

// GetAll implements the SCIM v2 server interface for groups.
func (gs *groupHandlers) GetAll(r *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	result := scim.Page{
		TotalResults: 0,
		Resources:    []scim.Resource{},
	}

	records := make([]*model.Group, 0)

	q := gs.store.NewSelect().
		Model(&records).
		Order("group.name ASC")

	if params.FilterValidator != nil {
		validator := params.FilterValidator

		if err := validator.Validate(); err != nil {
			return result, err
		}

		q = gs.filter(
			validator.GetFilter(),
			q,
		)
	}

	counter, err := q.Count(r.Context())

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, nil
		}

		gs.logger.Error().
			Err(err).
			Msg("Failed to count all")

		return result, err
	}

	result.TotalResults = counter

	if params.Count > 0 {
		q = q.Limit(
			params.Count,
		)

		if params.StartIndex < 1 {
			params.StartIndex = 1
		}

		if params.StartIndex > 1 {
			q = q.Offset(
				params.StartIndex * params.Count,
			)
		}

		if err := q.Scan(r.Context()); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return result, nil
			}

			gs.logger.Error().
				Err(err).
				Msg("Failed to fetch all")

			return result, err
		}

		for _, record := range records {
			result.Resources = append(
				result.Resources,
				scim.Resource{
					ID:         record.ID,
					ExternalID: optional.NewString(record.Scim),
					Meta: scim.Meta{
						Created:      &record.CreatedAt,
						LastModified: &record.UpdatedAt,
					},
					Attributes: scim.ResourceAttributes{
						"displayName": record.Name,
					},
				},
			)
		}
	}

	return result, nil
}

// Get implements the SCIM v2 server interface for groups.
func (gs *groupHandlers) Get(r *http.Request, id string) (scim.Resource, error) {
	record := &model.Group{}

	if err := gs.store.NewSelect().
		Model(record).
		Where("id = ?", id).
		Scan(r.Context()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return scim.Resource{}, serrors.ScimErrorResourceNotFound(id)
		}

		gs.logger.Error().
			Err(err).
			Str("id", id).
			Msg("Failed to fetch group")

		return scim.Resource{}, err
	}

	result := scim.Resource{
		ID:         record.ID,
		ExternalID: optional.NewString(record.Scim),
		Meta: scim.Meta{
			Created:      &record.CreatedAt,
			LastModified: &record.UpdatedAt,
		},
		Attributes: scim.ResourceAttributes{
			"displayName": record.Name,
		},
	}

	return result, nil
}

// Create implements the SCIM v2 server interface for groups.
func (gs *groupHandlers) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	externalID := ""
	if val, ok := attributes["externalId"]; ok {
		externalID = val.(string)
	}

	displayName := ""
	if val, ok := attributes["displayName"]; ok {
		displayName = val.(string)
	}

	record := &model.Group{}

	if err := gs.store.NewSelect().
		Model(record).
		Where("name = ? OR scim = ?", displayName, externalID).
		Scan(r.Context()); err != nil && err != sql.ErrNoRows {
		gs.logger.Error().
			Err(err).
			Str("group", displayName).
			Msg("Failed to check if group exists")

		return scim.Resource{}, err
	}

	record.Scim = externalID
	record.Name = displayName

	if record.ID == "" {
		record.Slug = gs.slugify(r.Context(), displayName, record.ID)

		gs.logger.Debug().
			Str("group", record.Name).
			Msg("Creating new group")

		if _, err := gs.store.NewInsert().
			Model(record).
			Exec(r.Context()); err != nil {
			gs.logger.Error().
				Err(err).
				Str("group", record.Name).
				Msg("Failed to create group")

			return scim.Resource{}, err
		}
	} else {
		if _, err := gs.store.NewUpdate().
			Model(record).
			Where("id = ?", record.ID).
			Exec(r.Context()); err != nil {
			gs.logger.Error().
				Err(err).
				Str("group", record.Name).
				Msg("Failed to update group")

			return scim.Resource{}, err
		}
	}

	result := scim.Resource{
		ID:         record.ID,
		ExternalID: optional.NewString(record.Scim),
		Meta: scim.Meta{
			Created:      &record.CreatedAt,
			LastModified: &record.UpdatedAt,
		},
		Attributes: scim.ResourceAttributes{
			"displayName": record.Name,
		},
	}

	return result, nil
}

// Replace implements the SCIM v2 server interface for groups.
func (gs *groupHandlers) Replace(r *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	externalID := ""
	if val, ok := attributes["externalId"]; ok {
		externalID = val.(string)
	}

	displayName := ""
	if val, ok := attributes["displayName"]; ok {
		displayName = val.(string)
	}

	record := &model.Group{}

	if err := gs.store.NewSelect().
		Model(record).
		Where("id = ? OR scim = ?", id, externalID).
		Scan(r.Context()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return scim.Resource{}, serrors.ScimErrorResourceNotFound(id)
		}

		gs.logger.Error().
			Err(err).
			Str("id", id).
			Msg("Failed to fetch group")

		return scim.Resource{}, err
	}

	record.Scim = externalID
	record.Name = displayName

	if _, err := gs.store.NewUpdate().
		Model(record).
		Where("id = ?", record.ID).
		Exec(r.Context()); err != nil {
		gs.logger.Error().
			Err(err).
			Str("id", id).
			Msg("Failed to update group")

		return scim.Resource{}, err
	}

	result := scim.Resource{
		ID:         record.ID,
		ExternalID: optional.NewString(record.Scim),
		Meta: scim.Meta{
			Created:      &record.CreatedAt,
			LastModified: &record.UpdatedAt,
		},
		Attributes: scim.ResourceAttributes{
			"displayName": record.Name,
		},
	}

	return result, nil
}

// Patch implements the SCIM v2 server interface for groups.
func (gs *groupHandlers) Patch(r *http.Request, id string, operations []scim.PatchOperation) (scim.Resource, error) {
	record := &model.Group{}

	if err := gs.store.NewSelect().
		Model(record).
		Where("id = ?", id).
		Scan(r.Context()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return scim.Resource{}, serrors.ScimErrorResourceNotFound(id)
		}

		gs.logger.Error().
			Err(err).
			Str("id", id).
			Msg("Failed to fetch group")

		return scim.Resource{}, err
	}

	for _, operation := range operations {
		switch op := operation.Op; op {
		case "remove":
			switch {
			case operation.Path.String() == "members":
				if is, ok := operation.Value.([]interface{}); ok {
					for _, i := range is {
						if vs, ok := i.(map[string]interface{}); ok {
							if v, ok := vs["value"]; ok {
								if _, err := gs.store.NewDelete().
									Model((*model.UserGroup)(nil)).
									Where("group_id = ? AND user_id = ?", record.ID, v.(string)).
									Exec(r.Context()); err != nil {
									gs.logger.Error().
										Err(err).
										Str("group", record.Name).
										Str("user", v.(string)).
										Msg("Failed to delete member")

									return scim.Resource{}, err
								}
							} else {
								gs.logger.Error().
									Str("method", "patch").
									Str("id", id).
									Str("operation", op).
									Str("path", "members").
									Msgf("Failed to convert member: %v", vs)
							}
						} else {
							gs.logger.Error().
								Str("method", "patch").
								Str("id", id).
								Str("operation", op).
								Str("path", "members").
								Msgf("Failed to convert values: %v", i)
						}
					}
				} else {
					gs.logger.Error().
						Str("method", "patch").
						Str("id", id).
						Str("operation", op).
						Str("path", "members").
						Msgf("Failed to convert interface: %v", operation.Value)
				}
			default:
				gs.logger.Error().
					Str("method", "patch").
					Str("id", id).
					Str("operation", op).
					Str("path", operation.Path.String()).
					Msg("Unknown path")

				return scim.Resource{}, fmt.Errorf(
					"unknown path: %s",
					operation.Path.String(),
				)
			}
		case "add":
			switch {
			case operation.Path.String() == "members":
				if is, ok := operation.Value.([]interface{}); ok {
					for _, i := range is {
						if vs, ok := i.(map[string]interface{}); ok {
							if v, ok := vs["value"]; ok {
								user := &model.User{}

								if err := gs.store.NewSelect().
									Model(user).
									Where("id = ?", v.(string)).
									Scan(r.Context()); err != nil {
									if errors.Is(err, sql.ErrNoRows) {
										continue
									}

									gs.logger.Error().
										Err(err).
										Str("group", record.Name).
										Str("user", v.(string)).
										Msg("Failed to fetch user")

									return scim.Resource{}, err
								}

								exists, err := gs.store.NewSelect().
									Model((*model.UserGroup)(nil)).
									Where("group_id = ? AND user_id = ?", record.ID, user.ID).
									Exists(r.Context())

								if err != nil {
									gs.logger.Error().
										Err(err).
										Str("group", record.Name).
										Str("user", user.Username).
										Msg("Failed to check member")

									return scim.Resource{}, err
								}

								if exists {
									gs.logger.Debug().
										Str("group", record.Name).
										Str("user", user.Username).
										Msg("Member already exists")

									continue
								}

								if _, err := gs.store.NewInsert().
									Model(&model.UserGroup{
										GroupID: record.ID,
										UserID:  user.ID,
										Perm:    "owner",
									}).Exec(r.Context()); err != nil {
									gs.logger.Error().
										Err(err).
										Str("group", record.Name).
										Str("user", user.Username).
										Msg("Failed to append member")

									return scim.Resource{}, err
								}
							} else {
								gs.logger.Error().
									Str("method", "patch").
									Str("id", id).
									Str("operation", op).
									Str("path", "members").
									Msgf("Failed to convert member: %v", vs)
							}
						} else {
							gs.logger.Error().
								Str("method", "patch").
								Str("id", id).
								Str("operation", op).
								Str("path", "members").
								Msgf("Failed to convert values: %v", i)
						}
					}
				} else {
					gs.logger.Error().
						Str("method", "patch").
						Str("id", id).
						Str("operation", op).
						Str("path", "members").
						Msgf("Failed to convert interface: %v", operation.Value)
				}
			default:
				gs.logger.Error().
					Str("method", "patch").
					Str("id", id).
					Str("operation", op).
					Str("path", operation.Path.String()).
					Msg("Unknown path")

				return scim.Resource{}, fmt.Errorf(
					"unknown path: %s",
					operation.Path.String(),
				)
			}
		default:
			gs.logger.Error().
				Str("method", "patch").
				Str("id", id).
				Str("operation", op).
				Msg("Unknown operation")

			return scim.Resource{}, fmt.Errorf(
				"unknown operation: %s",
				op,
			)
		}
	}

	result := scim.Resource{
		ID:         record.ID,
		ExternalID: optional.NewString(record.Scim),
		Meta: scim.Meta{
			Created:      &record.CreatedAt,
			LastModified: &record.UpdatedAt,
		},
		Attributes: scim.ResourceAttributes{
			"displayName": record.Name,
		},
	}

	return result, nil
}

// Delete implements the SCIM v2 server interface for groups.
func (gs *groupHandlers) Delete(r *http.Request, id string) error {
	if _, err := gs.store.NewDelete().
		Model((*model.Group)(nil)).
		Where("id = ?", id).
		Exec(r.Context()); err != nil {
		gs.logger.Error().
			Err(err).
			Str("id", id).
			Msg("Failed to delete group")

		return err
	}

	return nil
}

func (gs *groupHandlers) slugify(ctx context.Context, value, id string) string {
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

		query := gs.store.NewSelect().
			Model((*model.Group)(nil)).
			Where("slug = ?", slug)

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

func (gs *groupHandlers) filter(expr filter.Expression, db *bun.SelectQuery) *bun.SelectQuery {
	switch e := expr.(type) {
	case *filter.AttributeExpression:
		return gs.handleAttributeExpression(e, db)
	default:
		gs.logger.Error().
			Str("type", fmt.Sprintf("%T", e)).
			Msg("Unsupported expression type for group filter")
	}

	return db
}

func (gs *groupHandlers) handleAttributeExpression(e *filter.AttributeExpression, db *bun.SelectQuery) *bun.SelectQuery {
	scimAttr := e.AttributePath.String()
	column, ok := groupAttributeMapping[scimAttr]

	if !ok {
		gs.logger.Error().
			Str("attribute", scimAttr).
			Msg("Attribute is not mapped for groups")

		return db
	}

	value := e.CompareValue

	switch operator := strings.ToLower(string(e.Operator)); operator {
	case "eq":
		return db.Where("? = ?", bun.Ident(column), value)
	case "ne":
		return db.Where("? <> ?", bun.Ident(column), value)
	case "co":
		return db.Where("? LIKE ?", bun.Ident(column), "%"+fmt.Sprintf("%v", value)+"%")
	case "sw":
		return db.Where("? LIKE ?", bun.Ident(column), fmt.Sprintf("%v", value)+"%")
	case "ew":
		return db.Where("? LIKE ?", bun.Ident(column), "%"+fmt.Sprintf("%v", value))
	case "gt":
		return db.Where("? > ?", bun.Ident(column), value)
	case "ge":
		return db.Where("? >= ?", bun.Ident(column), value)
	case "lt":
		return db.Where("? < ?", bun.Ident(column), value)
	case "le":
		return db.Where("? <= ?", bun.Ident(column), value)
	default:
		gs.logger.Error().
			Str("operator", operator).
			Msgf("Unsupported attribute operator for group filter")
	}

	return db
}
