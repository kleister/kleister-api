package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/mod"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/versions"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListVersionsHandler implements the handler for the ListVersions operation.
func ListVersionsHandler(modsService mods.Service, versionsService versions.Service) mod.ListVersionsHandlerFunc {
	return func(params mod.ListVersionsParams, principal *models.User) middleware.Responder {
		parent, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewListVersionsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewListVersionsDefault(http.StatusInternalServerError)
		}

		records, err := versionsService.List(params.HTTPRequest.Context(), parent)

		if err != nil {
			return mod.NewListVersionsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Version, len(records))
		for id, record := range records {
			payload[id] = convertVersion(record)
		}

		return mod.NewListVersionsOK().WithPayload(payload)
	}
}

// ShowVersionHandler implements the handler for the ShowVersion operation.
func ShowVersionHandler(modsService mods.Service, versionsService versions.Service) mod.ShowVersionHandlerFunc {
	return func(params mod.ShowVersionParams, principal *models.User) middleware.Responder {
		parent, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewListVersionsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewListVersionsDefault(http.StatusInternalServerError)
		}

		record, err := versionsService.Show(params.HTTPRequest.Context(), parent, params.VersionID)

		if err != nil {
			if err == versions.ErrNotFound {
				message := "version not found"

				return mod.NewShowVersionNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewShowVersionDefault(http.StatusInternalServerError)
		}

		return mod.NewShowVersionOK().WithPayload(convertVersion(record))
	}
}

// CreateVersionHandler implements the handler for the CreateVersion operation.
func CreateVersionHandler(modsService mods.Service, versionsService versions.Service) mod.CreateVersionHandlerFunc {
	return func(params mod.CreateVersionParams, principal *models.User) middleware.Responder {
		parent, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewListVersionsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewListVersionsDefault(http.StatusInternalServerError)
		}

		record := &model.Version{
			ModID: parent.ID,
		}

		if params.Version.Slug != nil {
			record.Slug = *params.Version.Slug
		}

		if params.Version.Name != nil {
			record.Name = *params.Version.Name
		}

		created, err := versionsService.Create(params.HTTPRequest.Context(), parent, record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate version"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return mod.NewCreateVersionUnprocessableEntity().WithPayload(payload)
			}

			return mod.NewCreateVersionDefault(http.StatusInternalServerError)
		}

		return mod.NewCreateVersionOK().WithPayload(convertVersion(created))
	}
}

// UpdateVersionHandler implements the handler for the UpdateVersion operation.
func UpdateVersionHandler(modsService mods.Service, versionsService versions.Service) mod.UpdateVersionHandlerFunc {
	return func(params mod.UpdateVersionParams, principal *models.User) middleware.Responder {
		parent, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewListVersionsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewListVersionsDefault(http.StatusInternalServerError)
		}

		record, err := versionsService.Show(params.HTTPRequest.Context(), parent, params.VersionID)

		if err != nil {
			if err == versions.ErrNotFound {
				message := "version not found"

				return mod.NewUpdateVersionNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewUpdateVersionDefault(http.StatusInternalServerError)
		}

		if params.Version.Slug != nil {
			record.Slug = *params.Version.Slug
		}

		if params.Version.Name != nil {
			record.Name = *params.Version.Name
		}

		updated, err := versionsService.Update(params.HTTPRequest.Context(), parent, record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate version"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return mod.NewUpdateVersionUnprocessableEntity().WithPayload(payload)
			}

			return mod.NewUpdateVersionDefault(http.StatusInternalServerError)
		}

		return mod.NewUpdateVersionOK().WithPayload(convertVersion(updated))
	}
}

// DeleteVersionHandler implements the handler for the DeleteVersion operation.
func DeleteVersionHandler(modsService mods.Service, versionsService versions.Service) mod.DeleteVersionHandlerFunc {
	return func(params mod.DeleteVersionParams, principal *models.User) middleware.Responder {
		parent, err := modsService.Show(params.HTTPRequest.Context(), params.ModID)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return mod.NewListVersionsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewListVersionsDefault(http.StatusInternalServerError)
		}

		record, err := versionsService.Show(params.HTTPRequest.Context(), parent, params.VersionID)

		if err != nil {
			if err == versions.ErrNotFound {
				message := "version not found"

				return mod.NewDeleteVersionNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return mod.NewDeleteVersionDefault(http.StatusInternalServerError)
		}

		if err := versionsService.Delete(params.HTTPRequest.Context(), parent, record.ID); err != nil {
			message := "failed to delete version"

			return mod.NewDeleteVersionBadRequest().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		message := "successfully deleted version"
		return mod.NewDeleteVersionOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// convertVersion is a simple helper to convert between different model formats.
func convertVersion(record *model.Version) *models.Version {
	modID := strfmt.UUID(record.ModID)

	builds := make([]*models.BuildVersion, 0)

	for _, build := range record.Builds {
		builds = append(builds, convertBuildVersion(build))
	}

	return &models.Version{
		ID:        strfmt.UUID(record.ID),
		File:      convertVersionFile(record.File),
		ModID:     modID,
		Mod:       convertMod(record.Mod),
		Slug:      &record.Slug,
		Name:      &record.Name,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
		Builds:    builds,
	}
}

// convertVersionFile is a simple helper to convert between different model formats.
func convertVersionFile(record *model.VersionFile) *models.VersionFile {
	return &models.VersionFile{}
}
