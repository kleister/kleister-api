package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/mod"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListModsHandler implements the handler for the ListMods operation.
func ListModsHandler(modsService mods.Service) mod.ListModsHandlerFunc {
	return func(params mod.ListModsParams, principal *models.User) middleware.Responder {
		records, err := modsService.List(params.HTTPRequest.Context())

		if err != nil {
			return mod.NewListModsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Mod, len(records))
		for id, record := range records {
			payload[id] = convertMod(record)
		}

		return mod.NewListModsOK().WithPayload(payload)
	}
}

// ShowModHandler implements the handler for the ShowMod operation.
func ShowModHandler(modsService mods.Service) mod.ShowModHandlerFunc {
	return func(params mod.ShowModParams, principal *models.User) middleware.Responder {
		record, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewShowModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewShowModDefault(http.StatusInternalServerError)
		}

		return mod.NewShowModOK().WithPayload(convertMod(record))
	}
}

// CreateModHandler implements the handler for the CreateMod operation.
func CreateModHandler(modsService mods.Service) mod.CreateModHandlerFunc {
	return func(params mod.CreateModParams, principal *models.User) middleware.Responder {
		record := &model.Mod{}

		if params.Mod.Slug != nil {
			record.Slug = *params.Mod.Slug
		}

		if params.Mod.Name != nil {
			record.Name = *params.Mod.Name
		}

		created, err := modsService.Create(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate mod"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return mod.NewCreateModUnprocessableEntity().WithPayload(payload)
			}

			return mod.NewCreateModDefault(http.StatusInternalServerError)
		}

		return mod.NewCreateModOK().WithPayload(convertMod(created))
	}
}

// UpdateModHandler implements the handler for the UpdateMod operation.
func UpdateModHandler(modsService mods.Service) mod.UpdateModHandlerFunc {
	return func(params mod.UpdateModParams, principal *models.User) middleware.Responder {
		record, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewUpdateModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewUpdateModDefault(http.StatusInternalServerError)
		}

		if params.Mod.Slug != nil {
			record.Slug = *params.Mod.Slug
		}

		if params.Mod.Name != nil {
			record.Name = *params.Mod.Name
		}

		updated, err := modsService.Update(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate mod"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return mod.NewUpdateModUnprocessableEntity().WithPayload(payload)
			}

			return mod.NewUpdateModDefault(http.StatusInternalServerError)
		}

		return mod.NewUpdateModOK().WithPayload(convertMod(updated))
	}
}

// DeleteModHandler implements the handler for the DeleteMod operation.
func DeleteModHandler(modsService mods.Service) mod.DeleteModHandlerFunc {
	return func(params mod.DeleteModParams, principal *models.User) middleware.Responder {
		record, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewDeleteModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewDeleteModDefault(http.StatusInternalServerError)
		}

		if err := modsService.Delete(params.HTTPRequest.Context(), record.ID); err != nil {
			message := "failed to delete mod"

			return mod.NewDeleteModBadRequest().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		message := "successfully deleted mod"
		return mod.NewDeleteModOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// convertMod is a simple helper to convert between different model formats.
func convertMod(record *model.Mod) *models.Mod {
	versions := make([]*models.Version, 0)

	for _, version := range record.Versions {
		versions = append(versions, convertVersion(version))
	}

	users := make([]*models.UserMod, 0)

	for _, user := range record.Users {
		users = append(users, convertUserMod(user))
	}

	teams := make([]*models.TeamMod, 0)

	for _, team := range record.Teams {
		teams = append(teams, convertTeamMod(team))
	}

	return &models.Mod{
		ID:          strfmt.UUID(record.ID),
		Slug:        &record.Slug,
		Name:        &record.Name,
		Side:        &record.Side,
		Description: &record.Description,
		Author:      &record.Author,
		Website:     &record.Website,
		Donate:      &record.Donate,
		CreatedAt:   strfmt.DateTime(record.CreatedAt),
		UpdatedAt:   strfmt.DateTime(record.UpdatedAt),
		Versions:    versions,
		Users:       users,
		Teams:       teams,
	}
}
