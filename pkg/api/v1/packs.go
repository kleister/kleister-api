package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/pack"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListPacksHandler implements the handler for the ListPacks operation.
func ListPacksHandler(packsService packs.Service) pack.ListPacksHandlerFunc {
	return func(params pack.ListPacksParams, principal *models.User) middleware.Responder {
		records, err := packsService.List(params.HTTPRequest.Context())

		if err != nil {
			return pack.NewListPacksDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Pack, len(records))
		for id, record := range records {
			payload[id] = convertPack(record)
		}

		return pack.NewListPacksOK().WithPayload(payload)
	}
}

// ShowPackHandler implements the handler for the ShowPack operation.
func ShowPackHandler(packsService packs.Service) pack.ShowPackHandlerFunc {
	return func(params pack.ShowPackParams, principal *models.User) middleware.Responder {
		record, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewShowPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewShowPackDefault(http.StatusInternalServerError)
		}

		return pack.NewShowPackOK().WithPayload(convertPack(record))
	}
}

// CreatePackHandler implements the handler for the CreatePack operation.
func CreatePackHandler(packsService packs.Service) pack.CreatePackHandlerFunc {
	return func(params pack.CreatePackParams, principal *models.User) middleware.Responder {
		record := &model.Pack{}

		if params.Pack.Slug != nil {
			record.Slug = *params.Pack.Slug
		}

		if params.Pack.Name != nil {
			record.Name = *params.Pack.Name
		}

		created, err := packsService.Create(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate pack"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return pack.NewCreatePackUnprocessableEntity().WithPayload(payload)
			}

			return pack.NewCreatePackDefault(http.StatusInternalServerError)
		}

		return pack.NewCreatePackOK().WithPayload(convertPack(created))
	}
}

// UpdatePackHandler implements the handler for the UpdatePack operation.
func UpdatePackHandler(packsService packs.Service) pack.UpdatePackHandlerFunc {
	return func(params pack.UpdatePackParams, principal *models.User) middleware.Responder {
		record, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewUpdatePackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewUpdatePackDefault(http.StatusInternalServerError)
		}

		if params.Pack.Slug != nil {
			record.Slug = *params.Pack.Slug
		}

		if params.Pack.Name != nil {
			record.Name = *params.Pack.Name
		}

		updated, err := packsService.Update(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate pack"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return pack.NewUpdatePackUnprocessableEntity().WithPayload(payload)
			}

			return pack.NewUpdatePackDefault(http.StatusInternalServerError)
		}

		return pack.NewUpdatePackOK().WithPayload(convertPack(updated))
	}
}

// DeletePackHandler implements the handler for the DeletePack operation.
func DeletePackHandler(packsService packs.Service) pack.DeletePackHandlerFunc {
	return func(params pack.DeletePackParams, principal *models.User) middleware.Responder {
		record, err := packsService.Show(params.HTTPRequest.Context(), params.PackID)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return pack.NewDeletePackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return pack.NewDeletePackDefault(http.StatusInternalServerError)
		}

		if err := packsService.Delete(params.HTTPRequest.Context(), record.ID); err != nil {
			message := "failed to delete pack"

			return pack.NewDeletePackBadRequest().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		message := "successfully deleted pack"
		return pack.NewDeletePackOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// convertPack is a simple helper to convert between different model formats.
func convertPack(record *model.Pack) *models.Pack {
	builds := make([]*models.Build, 0)

	for _, build := range record.Builds {
		builds = append(builds, convertBuild(build))
	}

	users := make([]*models.UserPack, 0)

	for _, user := range record.Users {
		users = append(users, convertUserPack(user))
	}

	teams := make([]*models.TeamPack, 0)

	for _, team := range record.Teams {
		teams = append(teams, convertTeamPack(team))
	}

	return &models.Pack{
		ID:            strfmt.UUID(record.ID),
		Icon:          convertPackIcon(record.Icon),
		Logo:          convertPackLogo(record.Logo),
		Background:    convertPackBackground(record.Background),
		Recommended:   convertBuild(record.Recommended),
		RecommendedID: (*strfmt.UUID)(&record.RecommendedID),
		Latest:        convertBuild(record.Latest),
		LatestID:      (*strfmt.UUID)(&record.LatestID),
		Slug:          &record.Slug,
		Name:          &record.Name,
		Website:       &record.Website,
		Published:     &record.Published,
		Hidden:        &record.Hidden,
		Private:       &record.Private,
		Public:        &record.Public,
		CreatedAt:     strfmt.DateTime(record.CreatedAt),
		UpdatedAt:     strfmt.DateTime(record.UpdatedAt),
		Builds:        builds,
		Users:         users,
		Teams:         teams,
	}
}

// convertPackIcon is a simple helper to convert between different model formats.
func convertPackIcon(record *model.PackIcon) *models.PackIcon {
	return &models.PackIcon{}
}

// convertPackLogo is a simple helper to convert between different model formats.
func convertPackLogo(record *model.PackLogo) *models.PackLogo {
	return &models.PackLogo{}
}

// convertPackBackground is a simple helper to convert between different model formats.
func convertPackBackground(record *model.PackBackground) *models.PackBackground {
	return &models.PackBackground{}
}
