package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/internal/minecraft"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/rs/zerolog/log"
)

// ListMinecrafts implements the v1.ServerInterface.
func (a *API) ListMinecrafts(w http.ResponseWriter, r *http.Request, params ListMinecraftsParams) {
	ctx := r.Context()

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Minecraft.List(
		ctx,
		model.ListParams{
			Search: fromSearch(params.Search),
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "ListMinecrafts").
			Msg("Failed to load minecrafts")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load minecrafts"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Minecraft, len(records))
	for id, record := range records {
		payload[id] = a.convertMinecraft(record)
	}

	render.JSON(w, r, MinecraftsResponse{
		Total:    count,
		Versions: payload,
	})
}

// UpdateMinecraft implements the v1.ServerInterface.
func (a *API) UpdateMinecraft(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	versions, err := minecraft.FetchRemote()

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateMinecraft").
			Msg("Failed to fetch versions")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to fetch versions"),
			Status:  ToPtr(http.StatusServiceUnavailable),
		})
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Minecraft.Sync(
		ctx,
		versions,
	); err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateMinecraft").
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

// ListMinecraftBuilds implements the v1.ServerInterface.
func (a *API) ListMinecraftBuilds(w http.ResponseWriter, r *http.Request, _ MinecraftID, params ListMinecraftBuildsParams) {
	ctx := r.Context()
	record := a.MinecraftFromContext(ctx)
	sort, order, limit, offset, search := listMinecraftBuildsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Minecraft.ListBuilds(
		ctx,
		model.MinecraftBuildParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			MinecraftID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("minecraft", record.ID).
			Str("action", "ListMinecraftBuilds").
			Msg("Failed to load minecraft builds")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load minecraft builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuild(record)
	}

	render.JSON(w, r, MinecraftBuildsResponse{
		Total:     count,
		Limit:     limit,
		Offset:    offset,
		Minecraft: ToPtr(a.convertMinecraft(record)),
		Builds:    payload,
	})
}

// AttachMinecraftToBuild implements the v1.ServerInterface.
func (a *API) AttachMinecraftToBuild(w http.ResponseWriter, r *http.Request, _ MinecraftID) {
	ctx := r.Context()
	record := a.MinecraftFromContext(ctx)
	body := &MinecraftBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("minecraft", record.ID).
			Str("action", "AttachMinecraftToBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Minecraft.AttachBuild(
		ctx,
		model.MinecraftBuildParams{
			MinecraftID: record.ID,
			PackID:      body.Pack,
			BuildID:     body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrMinecraftNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find minecraft"),
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
				Message: ToPtr("Minecraft is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("minecraft", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "AttachMinecraftToBuild").
			Msg("Failed to attach minecraft to build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach minecraft to build"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached minecraft to build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteMinecraftFromBuild implements the v1.ServerInterface.
func (a *API) DeleteMinecraftFromBuild(w http.ResponseWriter, r *http.Request, _ MinecraftID) {
	ctx := r.Context()
	record := a.MinecraftFromContext(ctx)
	body := &MinecraftBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "DeleteMinecraftFromBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Minecraft.DropBuild(
		ctx,
		model.MinecraftBuildParams{
			MinecraftID: record.ID,
			PackID:      body.Pack,
			BuildID:     body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrMinecraftNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find minecraft"),
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
				Message: ToPtr("Minecraft is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("minecraft", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "DeleteMinecraftFromBuild").
			Msg("Failed to drop minecraft from build")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop minecraft from build"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped minecraft from build"),
		Status:  ToPtr(http.StatusOK),
	})
}

func (a *API) convertMinecraft(record *model.Minecraft) Minecraft {
	result := Minecraft{
		ID:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Type:      ToPtr(record.Type),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listMinecraftBuildsSorting(request ListMinecraftBuildsParams) (string, string, int64, int64, string) {
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
