package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/internal/forge"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/rs/zerolog/log"
)

// ListForges implements the v1.ServerInterface.
func (a *API) ListForges(w http.ResponseWriter, r *http.Request, params ListForgesParams) {
	ctx := r.Context()

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Forge.List(
		ctx,
		model.ListParams{
			Search: fromSearch(params.Search),
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "ListForges").
			Msg("Failed to load forges")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load forges"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Forge, len(records))
	for id, record := range records {
		payload[id] = a.convertForge(record)
	}

	render.JSON(w, r, ForgesResponse{
		Total:    count,
		Versions: payload,
	})
}

// UpdateForge implements the v1.ServerInterface.
func (a *API) UpdateForge(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	versions, err := forge.FetchRemote()

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateForge").
			Msg("Failed to fetch versions")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusServiceUnavailable),
		})
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Forge.Sync(
		ctx,
		versions,
	); err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateForge").
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

// ListForgeBuilds implements the v1.ServerInterface.
func (a *API) ListForgeBuilds(w http.ResponseWriter, r *http.Request, _ ForgeID, params ListForgeBuildsParams) {
	ctx := r.Context()
	record := a.ForgeFromContext(ctx)
	sort, order, limit, offset, search := listForgeBuildsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Forge.ListBuilds(
		ctx,
		model.ForgeBuildParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			ForgeID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("forge", record.ID).
			Str("action", "ListForgeBuilds").
			Msg("Failed to load forge builds")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load forge builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuild(record)
	}

	render.JSON(w, r, ForgeBuildsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Forge:  ToPtr(a.convertForge(record)),
		Builds: payload,
	})
}

// AttachForgeToBuild implements the v1.ServerInterface.
func (a *API) AttachForgeToBuild(w http.ResponseWriter, r *http.Request, _ ForgeID) {
	ctx := r.Context()
	record := a.ForgeFromContext(ctx)
	body := &ForgeBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("forge", record.ID).
			Str("action", "AttachForgeToBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Forge.AttachBuild(
		ctx,
		model.ForgeBuildParams{
			ForgeID: record.ID,
			PackID:  body.Pack,
			BuildID: body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrForgeNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find forge"),
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
				Message: ToPtr("Forge is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("forge", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "AttachForgeToBuild").
			Msg("Failed to attach forge to build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach forge to build"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached forge to build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteForgeFromBuild implements the v1.ServerInterface.
func (a *API) DeleteForgeFromBuild(w http.ResponseWriter, r *http.Request, _ ForgeID) {
	ctx := r.Context()
	record := a.ForgeFromContext(ctx)
	body := &ForgeBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "DeleteForgeFromBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Forge.DropBuild(
		ctx,
		model.ForgeBuildParams{
			ForgeID: record.ID,
			PackID:  body.Pack,
			BuildID: body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrForgeNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find forge"),
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
				Message: ToPtr("Forge is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("forge", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "DeleteForgeFromBuild").
			Msg("Failed to drop forge from build")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop forge from build"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped forge from build"),
		Status:  ToPtr(http.StatusOK),
	})
}

func (a *API) convertForge(record *model.Forge) Forge {
	result := Forge{
		ID:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Minecraft: ToPtr(record.Minecraft),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listForgeBuildsSorting(request ListForgeBuildsParams) (string, string, int64, int64, string) {
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
