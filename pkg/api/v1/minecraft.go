package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/minecraft"
	"github.com/kleister/kleister-api/pkg/model"
	minecraftSvc "github.com/kleister/kleister-api/pkg/service/minecraft"
)

// ListMinecraftsHandler implements the handler for the MinecraftListMinecrafts operation.
func ListMinecraftsHandler(minecraftService minecraftSvc.Service) minecraft.ListMinecraftsHandlerFunc {
	return func(params minecraft.ListMinecraftsParams, principal *models.User) middleware.Responder {
		records, err := minecraftService.List(params.HTTPRequest.Context())

		if err != nil {
			return minecraft.NewListMinecraftsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Minecraft, len(records))
		for id, record := range records {
			payload[id] = convertMinecraft(record)
		}

		return minecraft.NewListMinecraftsOK().WithPayload(payload)
	}
}

// UpdateMinecraftHandler implements the handler for the MinecraftUpdateMinecraft operation.
func UpdateMinecraftHandler(minecraftService minecraftSvc.Service) minecraft.UpdateMinecraftHandlerFunc {
	return func(params minecraft.UpdateMinecraftParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return minecraft.NewUpdateMinecraftForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		err := minecraftService.Update(params.HTTPRequest.Context())

		if err != nil {
			if err == minecraftSvc.ErrSyncUnavailable {
				message := "minecraft version service is unavailable"

				return minecraft.NewUpdateMinecraftServiceUnavailable().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return minecraft.NewUpdateMinecraftDefault(http.StatusInternalServerError)
		}

		message := "successfully updated minecraft versions"
		return minecraft.NewUpdateMinecraftOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

func convertMinecraft(record *model.Minecraft) *models.Minecraft {
	return &models.Minecraft{
		ID:        strfmt.UUID(record.ID),
		Name:      &record.Name,
		Type:      &record.Type,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}
