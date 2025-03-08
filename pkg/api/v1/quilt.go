package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/internal/quilt"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/rs/zerolog/log"
)

// ListQuilts implements the v1.ServerInterface.
func (a *API) ListQuilts(w http.ResponseWriter, r *http.Request, params ListQuiltsParams) {
	ctx := r.Context()

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Quilt.List(
		ctx,
		model.ListParams{
			Search: fromSearch(params.Search),
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "ListQuilts").
			Msg("Failed to load quilts")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load quilts"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Quilt, len(records))
	for id, record := range records {
		payload[id] = a.convertQuilt(record)
	}

	render.JSON(w, r, QuiltsResponse{
		Total:    count,
		Versions: payload,
	})
}

// UpdateQuilt implements the v1.ServerInterface.
func (a *API) UpdateQuilt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	versions, err := quilt.FetchRemote()

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateQuilt").
			Msg("Failed to fetch versions")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusServiceUnavailable),
		})
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Quilt.Sync(
		ctx,
		versions,
	); err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateQuilt").
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

// ListQuiltBuilds implements the v1.ServerInterface.
func (a *API) ListQuiltBuilds(w http.ResponseWriter, r *http.Request, _ QuiltID, params ListQuiltBuildsParams) {
	ctx := r.Context()
	record := a.QuiltFromContext(ctx)
	sort, order, limit, offset, search := listQuiltBuildsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Quilt.ListBuilds(
		ctx,
		model.QuiltBuildParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			QuiltID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("quilt", record.ID).
			Str("action", "ListQuiltBuilds").
			Msg("Failed to load quilt builds")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load quilt builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuild(record)
	}

	render.JSON(w, r, QuiltBuildsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Quilt:  ToPtr(a.convertQuilt(record)),
		Builds: payload,
	})
}

// AttachQuiltToBuild implements the v1.ServerInterface.
func (a *API) AttachQuiltToBuild(w http.ResponseWriter, r *http.Request, _ QuiltID) {
	ctx := r.Context()
	record := a.QuiltFromContext(ctx)
	body := &QuiltBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("quilt", record.ID).
			Str("action", "AttachQuiltToBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Quilt.AttachBuild(
		ctx,
		model.QuiltBuildParams{
			QuiltID: record.ID,
			PackID:  body.Pack,
			BuildID: body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrQuiltNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find quilt"),
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
				Message: ToPtr("Quilt is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("quilt", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "AttachQuiltToBuild").
			Msg("Failed to attach quilt to build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach quilt to build"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached quilt to build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteQuiltFromBuild implements the v1.ServerInterface.
func (a *API) DeleteQuiltFromBuild(w http.ResponseWriter, r *http.Request, _ QuiltID) {
	ctx := r.Context()
	record := a.QuiltFromContext(ctx)
	body := &QuiltBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "DeleteQuiltFromBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Quilt.DropBuild(
		ctx,
		model.QuiltBuildParams{
			QuiltID: record.ID,
			PackID:  body.Pack,
			BuildID: body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrQuiltNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find quilt"),
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
				Message: ToPtr("Quilt is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("quilt", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "DeleteQuiltFromBuild").
			Msg("Failed to drop quilt from build")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop quilt from build"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped quilt from build"),
		Status:  ToPtr(http.StatusOK),
	})
}

func (a *API) convertQuilt(record *model.Quilt) Quilt {
	result := Quilt{
		ID:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listQuiltBuildsSorting(request ListQuiltBuildsParams) (string, string, int64, int64, string) {
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
