package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/forge"
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
