package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/forge"
	"github.com/kleister/kleister-api/pkg/model"
	forgeSvc "github.com/kleister/kleister-api/pkg/service/forge"
)

// ListForgesHandler implements the handler for the ForgeListForges operation.
func ListForgesHandler(forgeService forgeSvc.Service) forge.ListForgesHandlerFunc {
	return func(params forge.ListForgesParams, principal *models.User) middleware.Responder {
		records, err := forgeService.List(params.HTTPRequest.Context())

		if err != nil {
			return forge.NewListForgesDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Forge, len(records))
		for id, record := range records {
			payload[id] = convertForge(record)
		}

		return forge.NewListForgesOK().WithPayload(payload)
	}
}

// UpdateForgeHandler implements the handler for the ForgeUpdateForge operation.
func UpdateForgeHandler(forgeService forgeSvc.Service) forge.UpdateForgeHandlerFunc {
	return func(params forge.UpdateForgeParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return forge.NewUpdateForgeForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		err := forgeService.Update(params.HTTPRequest.Context())

		if err != nil {
			if err == forgeSvc.ErrSyncUnavailable {
				message := "forge version service is unavailable"

				return forge.NewUpdateForgeServiceUnavailable().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return forge.NewUpdateForgeDefault(http.StatusInternalServerError)
		}

		message := "successfully updated forge versions"
		return forge.NewUpdateForgeOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

func convertForge(record *model.Forge) *models.Forge {
	return &models.Forge{
		ID:        strfmt.UUID(record.ID),
		Name:      &record.Name,
		Minecraft: &record.Minecraft,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}
