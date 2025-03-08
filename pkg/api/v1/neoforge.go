package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/internal/neoforge"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/rs/zerolog/log"
)

// ListNeoforges implements the v1.ServerInterface.
func (a *API) ListNeoforges(w http.ResponseWriter, r *http.Request, params ListNeoforgesParams) {
	ctx := r.Context()

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Neoforge.List(
		ctx,
		model.ListParams{
			Search: fromSearch(params.Search),
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "ListNeoforges").
			Msg("Failed to load neoforges")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load neoforges"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Neoforge, len(records))
	for id, record := range records {
		payload[id] = a.convertNeoforge(record)
	}

	render.JSON(w, r, NeoforgesResponse{
		Total:    count,
		Versions: payload,
	})
}

// UpdateNeoforge implements the v1.ServerInterface.
func (a *API) UpdateNeoforge(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	versions, err := neoforge.FetchRemote()

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateNeoforge").
			Msg("Failed to fetch versions")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusServiceUnavailable),
		})
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Neoforge.Sync(
		ctx,
		versions,
	); err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateNeoforge").
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

// ListNeoforgeBuilds implements the v1.ServerInterface.
func (a *API) ListNeoforgeBuilds(w http.ResponseWriter, r *http.Request, _ NeoforgeID, params ListNeoforgeBuildsParams) {
	ctx := r.Context()
	record := a.NeoforgeFromContext(ctx)
	sort, order, limit, offset, search := listNeoforgeBuildsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Neoforge.ListBuilds(
		ctx,
		model.NeoforgeBuildParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			NeoforgeID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("neoforge", record.ID).
			Str("action", "ListNeoforgeBuilds").
			Msg("Failed to load neoforge builds")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load neoforge builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuild(record)
	}

	render.JSON(w, r, NeoforgeBuildsResponse{
		Total:    count,
		Limit:    limit,
		Offset:   offset,
		Neoforge: ToPtr(a.convertNeoforge(record)),
		Builds:   payload,
	})
}

// AttachNeoforgeToBuild implements the v1.ServerInterface.
func (a *API) AttachNeoforgeToBuild(w http.ResponseWriter, r *http.Request, _ NeoforgeID) {
	ctx := r.Context()
	record := a.NeoforgeFromContext(ctx)
	body := &NeoforgeBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("neoforge", record.ID).
			Str("action", "AttachNeoforgeToBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Neoforge.AttachBuild(
		ctx,
		model.NeoforgeBuildParams{
			NeoforgeID: record.ID,
			PackID:     body.Pack,
			BuildID:    body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrNeoforgeNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find neoforge"),
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
				Message: ToPtr("Neoforge is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("neoforge", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "AttachNeoforgeToBuild").
			Msg("Failed to attach neoforge to build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach neoforge to build"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached neoforge to build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteNeoforgeFromBuild implements the v1.ServerInterface.
func (a *API) DeleteNeoforgeFromBuild(w http.ResponseWriter, r *http.Request, _ NeoforgeID) {
	ctx := r.Context()
	record := a.NeoforgeFromContext(ctx)
	body := &NeoforgeBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "DeleteNeoforgeFromBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Neoforge.DropBuild(
		ctx,
		model.NeoforgeBuildParams{
			NeoforgeID: record.ID,
			PackID:     body.Pack,
			BuildID:    body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrNeoforgeNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find neoforge"),
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
				Message: ToPtr("Neoforge is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("neoforge", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "DeleteNeoforgeFromBuild").
			Msg("Failed to drop neoforge from build")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop neoforge from build"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped neoforge from build"),
		Status:  ToPtr(http.StatusOK),
	})
}

func (a *API) convertNeoforge(record *model.Neoforge) Neoforge {
	result := Neoforge{
		ID:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listNeoforgeBuildsSorting(request ListNeoforgeBuildsParams) (string, string, int64, int64, string) {
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
