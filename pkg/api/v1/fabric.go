package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/internal/fabric"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/rs/zerolog/log"
)

// ListFabrics implements the v1.ServerInterface.
func (a *API) ListFabrics(w http.ResponseWriter, r *http.Request, params ListFabricsParams) {
	ctx := r.Context()

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Fabric.List(
		ctx,
		model.ListParams{
			Search: fromSearch(params.Search),
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "ListFabrics").
			Msg("Failed to load fabrics")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load fabrics"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Fabric, len(records))
	for id, record := range records {
		payload[id] = a.convertFabric(record)
	}

	render.JSON(w, r, FabricsResponse{
		Total:    count,
		Versions: payload,
	})
}

// UpdateFabric implements the v1.ServerInterface.
func (a *API) UpdateFabric(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	versions, err := fabric.FetchRemote()

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateFabric").
			Msg("Failed to fetch versions")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusServiceUnavailable),
		})
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Fabric.Sync(
		ctx,
		versions,
	); err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateFabric").
			Msg("Failed to sync versions")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to sync versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		})
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully synced versions"),
		Status:  ToPtr(http.StatusOK),
	})
}

// ListFabricBuilds implements the v1.ServerInterface.
func (a *API) ListFabricBuilds(w http.ResponseWriter, r *http.Request, _ FabricID, params ListFabricBuildsParams) {
	ctx := r.Context()
	record := a.FabricFromContext(ctx)
	sort, order, limit, offset, search := listFabricBuildsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Fabric.ListBuilds(
		ctx,
		model.FabricBuildParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			FabricID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("fabric", record.ID).
			Str("action", "ListFabricBuilds").
			Msg("Failed to load fabric builds")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load fabric builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuild(record)
	}

	render.JSON(w, r, FabricBuildsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Fabric: ToPtr(a.convertFabric(record)),
		Builds: payload,
	})
}

// AttachFabricToBuild implements the v1.ServerInterface.
func (a *API) AttachFabricToBuild(w http.ResponseWriter, r *http.Request, _ FabricID) {
	ctx := r.Context()
	record := a.FabricFromContext(ctx)
	body := &FabricBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("fabric", record.ID).
			Str("action", "AttachFabricToBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Fabric.AttachBuild(
		ctx,
		model.FabricBuildParams{
			FabricID: record.ID,
			PackID:   body.Pack,
			BuildID:  body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrFabricNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find fabric"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrBuildNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find build"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrAlreadyAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Fabric is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("fabric", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "AttachFabricToBuild").
			Msg("Failed to attach fabric to build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach fabric to build"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached fabric to build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteFabricFromBuild implements the v1.ServerInterface.
func (a *API) DeleteFabricFromBuild(w http.ResponseWriter, r *http.Request, _ FabricID) {
	ctx := r.Context()
	record := a.FabricFromContext(ctx)
	body := &FabricBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "DeleteFabricFromBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Fabric.DropBuild(
		ctx,
		model.FabricBuildParams{
			FabricID: record.ID,
			PackID:   body.Pack,
			BuildID:  body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrFabricNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find fabric"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrBuildNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find build"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Fabric is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("fabric", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "DeleteFabricFromBuild").
			Msg("Failed to drop fabric from build")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop fabric from build"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped fabric from build"),
		Status:  ToPtr(http.StatusOK),
	})
}

func (a *API) convertFabric(record *model.Fabric) Fabric {
	result := Fabric{
		ID:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listFabricBuildsSorting(request ListFabricBuildsParams) (string, string, int64, int64, string) {
	sort, limit, offset, search := toPageParams(
		request.Sort,
		request.Limit,
		request.Offset,
		request.Search,
	)

	order := ""

	if request.Order != nil {
		order = string(FromPtr(request.Order))
	}

	return sort, order, limit, offset, search
}
