package scim

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/elimity-com/scim"
	"github.com/elimity-com/scim/errors"
	"github.com/elimity-com/scim/optional"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/secret"
	"github.com/rs/zerolog"
	"github.com/scim2/filter-parser/v2"
	"gorm.io/gorm"
)

var (
	userAttributeMapping = map[string]string{
		"userName":    "username",
		"email":       "email",
		"displayName": "fullname",
		"active":      "active",
	}
)

type userHandlers struct {
	config config.Scim
	store  *gorm.DB
	logger zerolog.Logger
}

// GetAll implements the SCIM v2 server interface for users.
func (us *userHandlers) GetAll(r *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	result := scim.Page{
		TotalResults: 0,
		Resources:    []scim.Resource{},
	}

	q := us.store.WithContext(
		r.Context(),
	).Model(
		&model.User{},
	).Order(
		"username ASC",
	).Where(
		"scim != ''",
	)

	if params.FilterValidator != nil {
		validator := params.FilterValidator

		if err := validator.Validate(); err != nil {
			return result, err
		}

		q = us.filter(
			validator.GetFilter(),
			q,
		)
	}

	counter := int64(0)

	if err := q.Count(
		&counter,
	).Error; err != nil {
		return result, err
	}

	result.TotalResults = int(counter)

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

		records := make([]*model.User, 0)

		if err := q.Find(
			&records,
		).Error; err != nil {
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
						"userName":    record.Username,
						"displayName": record.Fullname,
						"active":      record.Active,
					},
				},
			)
		}
	}

	return result, nil
}

// Get implements the SCIM v2 server interface for users.
func (us *userHandlers) Get(r *http.Request, id string) (scim.Resource, error) {
	record := &model.User{}

	if err := us.store.WithContext(
		r.Context(),
	).Model(
		&model.User{},
	).Where(
		"scim != ''",
	).Where(
		"id = ?",
		id,
	).First(
		record,
	).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return scim.Resource{}, errors.ScimErrorResourceNotFound(id)
		}

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
			"userName":    record.Username,
			"displayName": record.Fullname,
			"active":      record.Active,
		},
	}

	return result, nil
}

// Create implements the SCIM v2 server interface for users.
func (us *userHandlers) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	tx := us.store.WithContext(
		r.Context(),
	).Begin()
	defer tx.Rollback()

	record := &model.User{}

	if val, ok := attributes["externalId"]; ok {
		record.Scim = val.(string)
	}

	if val, ok := attributes["userName"]; ok {
		record.Username = val.(string)
	}

	if val, ok := attributes["displayName"]; ok {
		record.Fullname = val.(string)
	}

	if val, ok := attributes["active"]; ok {
		record.Active = val.(bool)
	}

	if val, ok := attributes["emails"]; ok {
		if is, ok := val.([]interface{}); ok {
			for _, i := range is {
				if vs, ok := i.(map[string]interface{}); ok {
					if p, ok := vs["primary"]; ok && p.(bool) {
						record.Email = vs["value"].(string)
					}
				} else {
					us.logger.Error().
						Str("method", "create").
						Str("path", "emails").
						Msgf("Failed to convert email: %v", i)
				}
			}
		} else {
			us.logger.Error().
				Str("method", "create").
				Str("path", "emails").
				Msgf("Failed to convert interface: %v", val)
		}
	}

	if err := tx.Where(
		model.User{
			Username: record.Username,
		},
	).Attrs(model.User{
		Password: secret.Generate(32),
	}).Assign(
		model.User{
			Scim:     record.Scim,
			Fullname: record.Fullname,
			Active:   record.Active,
			Email:    record.Email,
		},
	).FirstOrCreate(record).Error; err != nil {
		return scim.Resource{}, err
	}

	if err := tx.Commit().Error; err != nil {
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
			"userName":    record.Username,
			"displayName": record.Fullname,
			"active":      record.Active,
		},
	}

	return result, nil
}

// Replace implements the SCIM v2 server interface for users.
func (us *userHandlers) Replace(r *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	record := &model.User{}

	if err := us.store.WithContext(
		r.Context(),
	).Model(
		&model.User{},
	).Where(
		"scim != ''",
	).Where(
		"id = ?",
		id,
	).First(
		record,
	).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return scim.Resource{}, errors.ScimErrorResourceNotFound(id)
		}

		return scim.Resource{}, err
	}

	if val, ok := attributes["externalId"]; ok {
		record.Scim = val.(string)
	}

	if val, ok := attributes["userName"]; ok {
		record.Username = val.(string)
	}

	if val, ok := attributes["displayName"]; ok {
		record.Fullname = val.(string)
	}

	if val, ok := attributes["active"]; ok {
		record.Active = val.(bool)
	}

	if val, ok := attributes["emails"]; ok {
		if is, ok := val.([]interface{}); ok {
			for _, i := range is {
				if vs, ok := i.(map[string]interface{}); ok {
					if p, ok := vs["primary"]; ok && p.(bool) {
						record.Email = vs["value"].(string)
					}
				} else {
					us.logger.Error().
						Str("method", "update").
						Str("path", "emails").
						Msgf("Failed to convert email: %v", i)
				}
			}
		} else {
			us.logger.Error().
				Str("method", "update").
				Str("path", "emails").
				Msgf("Failed to convert interface: %v", val)
		}
	}

	tx := us.store.WithContext(
		r.Context(),
	).Begin()
	defer tx.Rollback()

	if err := tx.Save(record).Error; err != nil {
		return scim.Resource{}, err
	}

	if err := tx.Commit().Error; err != nil {
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
			"userName":    record.Username,
			"displayName": record.Fullname,
			"active":      record.Active,
		},
	}

	return result, nil
}

// Patch implements the SCIM v2 server interface for users.
func (us *userHandlers) Patch(r *http.Request, id string, operations []scim.PatchOperation) (scim.Resource, error) {
	record := &model.User{}

	if err := us.store.WithContext(
		r.Context(),
	).Model(
		&model.User{},
	).Where(
		"scim != ''",
	).Where(
		"id = ?",
		id,
	).First(
		record,
	).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return scim.Resource{}, errors.ScimErrorResourceNotFound(id)
		}

		return scim.Resource{}, err
	}

	tx := us.store.WithContext(
		r.Context(),
	).Begin()
	defer tx.Rollback()

	for _, operation := range operations {
		switch op := operation.Op; op {
		default:
			us.logger.Error().
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

	if err := tx.Commit().Error; err != nil {
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
			"userName":    record.Username,
			"displayName": record.Fullname,
			"active":      record.Active,
		},
	}

	return result, nil
}

// Delete implements the SCIM v2 server interface for users.
func (us *userHandlers) Delete(r *http.Request, id string) error {
	tx := us.store.WithContext(
		r.Context(),
	).Begin()
	defer tx.Rollback()

	if err := tx.Where(
		"scim != ''",
	).Where(
		"id = ?",
		id,
	).Delete(
		&model.User{},
	).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (us *userHandlers) filter(expr filter.Expression, db *gorm.DB) *gorm.DB {
	switch e := expr.(type) {
	case *filter.AttributeExpression:
		return us.handleAttributeExpression(e, db)
	default:
		us.logger.Error().
			Str("type", fmt.Sprintf("%T", e)).
			Msg("Unsupported expression type for user filter")
	}

	return db
}

func (us *userHandlers) handleAttributeExpression(e *filter.AttributeExpression, db *gorm.DB) *gorm.DB {
	scimAttr := e.AttributePath.String()
	column, ok := userAttributeMapping[scimAttr]

	if !ok {
		us.logger.Error().
			Str("attribute", scimAttr).
			Msg("Attribute is not mapped for users")

		return db
	}

	value := e.CompareValue

	switch operator := strings.ToLower(string(e.Operator)); operator {
	case "eq":
		return db.Where(fmt.Sprintf("%s = ?", column), value)
	case "ne":
		return db.Where(fmt.Sprintf("%s <> ?", column), value)
	case "co":
		return db.Where(fmt.Sprintf("%s LIKE ?", column), "%"+fmt.Sprintf("%v", value)+"%")
	case "sw":
		return db.Where(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%v", value)+"%")
	case "ew":
		return db.Where(fmt.Sprintf("%s LIKE ?", column), "%"+fmt.Sprintf("%v", value))
	case "gt":
		return db.Where(fmt.Sprintf("%s > ?", column), value)
	case "ge":
		return db.Where(fmt.Sprintf("%s >= ?", column), value)
	case "lt":
		return db.Where(fmt.Sprintf("%s < ?", column), value)
	case "le":
		return db.Where(fmt.Sprintf("%s <= ?", column), value)
	default:
		us.logger.Error().
			Str("operator", operator).
			Msgf("Unsupported attribute operator for user filter")
	}

	return db
}
