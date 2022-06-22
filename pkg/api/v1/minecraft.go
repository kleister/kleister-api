package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/minecraft"
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
