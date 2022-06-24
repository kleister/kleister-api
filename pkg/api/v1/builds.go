package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/pack"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/builds"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListBuildsHandler implements the handler for the ListBuilds operation.
func ListBuildsHandler(packsService packs.Service, buildsService builds.Service) pack.ListBuildsHandlerFunc {
	return func(params pack.ListBuildsParams, principal *models.User) middleware.Responder {
		parent, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewListBuildsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewListBuildsDefault(http.StatusInternalServerError)
		}

		records, err := buildsService.List(params.HTTPRequest.Context(), parent)

		if err != nil {
			return pack.NewListBuildsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Build, len(records))
		for id, record := range records {
			payload[id] = convertBuild(record)
		}

		return pack.NewListBuildsOK().WithPayload(payload)
	}
}

// ShowBuildHandler implements the handler for the ShowBuild operation.
func ShowBuildHandler(packsService packs.Service, buildsService builds.Service) pack.ShowBuildHandlerFunc {
	return func(params pack.ShowBuildParams, principal *models.User) middleware.Responder {
		parent, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewListBuildsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewListBuildsDefault(http.StatusInternalServerError)
		}

		record, err := buildsService.Show(params.HTTPRequest.Context(), parent, params.BuildID)

		if err != nil {
			if err == builds.ErrNotFound {
				message := "build not found"

				return pack.NewShowBuildNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewShowBuildDefault(http.StatusInternalServerError)
		}

		return pack.NewShowBuildOK().WithPayload(convertBuild(record))
	}
}

// CreateBuildHandler implements the handler for the CreateBuild operation.
func CreateBuildHandler(packsService packs.Service, buildsService builds.Service) pack.CreateBuildHandlerFunc {
	return func(params pack.CreateBuildParams, principal *models.User) middleware.Responder {
		parent, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewListBuildsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewListBuildsDefault(http.StatusInternalServerError)
		}

		record := &model.Build{
			PackID: parent.ID,
		}

		if params.Build.Slug != nil {
			record.Slug = *params.Build.Slug
		}

		if params.Build.Name != nil {
			record.Name = *params.Build.Name
		}

		created, err := buildsService.Create(params.HTTPRequest.Context(), parent, record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate build"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return pack.NewCreateBuildUnprocessableEntity().WithPayload(payload)
			}

			return pack.NewCreateBuildDefault(http.StatusInternalServerError)
		}

		return pack.NewCreateBuildOK().WithPayload(convertBuild(created))
	}
}

// UpdateBuildHandler implements the handler for the UpdateBuild operation.
func UpdateBuildHandler(packsService packs.Service, buildsService builds.Service) pack.UpdateBuildHandlerFunc {
	return func(params pack.UpdateBuildParams, principal *models.User) middleware.Responder {
		parent, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewListBuildsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewListBuildsDefault(http.StatusInternalServerError)
		}

		record, err := buildsService.Show(params.HTTPRequest.Context(), parent, params.BuildID)

		if err != nil {
			if err == builds.ErrNotFound {
				message := "build not found"

				return pack.NewUpdateBuildNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewUpdateBuildDefault(http.StatusInternalServerError)
		}

		if params.Build.Slug != nil {
			record.Slug = *params.Build.Slug
		}

		if params.Build.Name != nil {
			record.Name = *params.Build.Name
		}

		updated, err := buildsService.Update(params.HTTPRequest.Context(), parent, record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate build"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return pack.NewUpdateBuildUnprocessableEntity().WithPayload(payload)
			}

			return pack.NewUpdateBuildDefault(http.StatusInternalServerError)
		}

		return pack.NewUpdateBuildOK().WithPayload(convertBuild(updated))
	}
}

// DeleteBuildHandler implements the handler for the DeleteBuild operation.
func DeleteBuildHandler(packsService packs.Service, buildsService builds.Service) pack.DeleteBuildHandlerFunc {
	return func(params pack.DeleteBuildParams, principal *models.User) middleware.Responder {
		parent, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewListBuildsNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewListBuildsDefault(http.StatusInternalServerError)
		}

		record, err := buildsService.Show(params.HTTPRequest.Context(), parent, params.BuildID)

		if err != nil {
			if err == builds.ErrNotFound {
				message := "build not found"

				return pack.NewDeleteBuildNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewDeleteBuildDefault(http.StatusInternalServerError)
		}

		if err := buildsService.Delete(params.HTTPRequest.Context(), parent, record.ID); err != nil {
			message := "failed to delete build"

			return pack.NewDeleteBuildBadRequest().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		message := "successfully deleted build"
		return pack.NewDeleteBuildOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// convertBuild is a simple helper to convert between different model formats.
func convertBuild(record *model.Build) *models.Build {
	packID := strfmt.UUID(record.PackID)
	minecraftID := strfmt.UUID(record.MinecraftID)
	forgeID := strfmt.UUID(record.ForgeID)

	versions := make([]*models.BuildVersion, 0)

	for _, version := range record.Versions {
		versions = append(versions, convertBuildVersion(version))
	}

	return &models.Build{
		ID:          strfmt.UUID(record.ID),
		PackID:      &packID,
		Pack:        convertPack(record.Pack),
		MinecraftID: &minecraftID,
		Minecraft:   convertMinecraft(record.Minecraft),
		ForgeID:     &forgeID,
		Forge:       convertForge(record.Forge),
		Slug:        &record.Slug,
		Name:        &record.Name,
		MinJava:     &record.MinJava,
		MinMemory:   &record.MinMemory,
		Published:   &record.Published,
		Hidden:      &record.Hidden,
		Private:     &record.Private,
		Public:      &record.Public,
		CreatedAt:   strfmt.DateTime(record.CreatedAt),
		UpdatedAt:   strfmt.DateTime(record.UpdatedAt),
		Versions:    versions,
	}
}

// convertBuildVersion is a simple helper to convert between different model formats.
func convertBuildVersion(record *model.BuildVersion) *models.BuildVersion {
	buildID := strfmt.UUID(record.BuildID)
	versionID := strfmt.UUID(record.VersionID)

	return &models.BuildVersion{
		BuildID:   &buildID,
		Build:     convertBuild(record.Build),
		VersionID: &versionID,
		Version:   convertVersion(record.Version),
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}
