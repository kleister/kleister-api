package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
)

// ListBuilds implements the v1.ServerInterface.
func (a *API) ListBuilds(w http.ResponseWriter, r *http.Request, _ PackID, params ListBuildsParams) {
	ctx := r.Context()
	pack := a.PackFromContext(ctx)
	sort, order, limit, offset, search := listBuildsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Builds.List(
		ctx,
		pack,
		model.ListParams{
			Sort:   sort,
			Order:  order,
			Limit:  limit,
			Offset: offset,
			Search: search,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("action", "ListBuilds").
			Msg("Failed to load builds")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Build, len(records))
	for id, record := range records {
		payload[id] = a.convertBuild(record)
	}

	render.JSON(w, r, BuildsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Pack:   ToPtr(a.convertPack(pack)),
		Builds: payload,
	})
}

// ShowBuild implements the v1.ServerInterface.
func (a *API) ShowBuild(w http.ResponseWriter, r *http.Request, _ PackID, _ BuildID) {
	ctx := r.Context()
	record := a.BuildFromContext(ctx)

	render.JSON(w, r, BuildResponse(
		a.convertBuild(record),
	))
}

// CreateBuild implements the v1.ServerInterface.
func (a *API) CreateBuild(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	pack := a.PackFromContext(ctx)
	body := &CreateBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("action", "CreateBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	record := &model.Build{
		PackID: pack.ID,
	}

	if body.MinecraftID != nil {
		record.MinecraftID = body.MinecraftID
	}

	if body.ForgeID != nil {
		record.ForgeID = body.ForgeID
	}

	if body.NeoforgeID != nil {
		record.NeoforgeID = body.NeoforgeID
	}

	if body.QuiltID != nil {
		record.QuiltID = body.QuiltID
	}

	if body.FabricID != nil {
		record.FabricID = body.FabricID
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if body.Java != nil {
		record.Java = FromPtr(body.Java)
	}

	if body.Memory != nil {
		record.Memory = FromPtr(body.Memory)
	}

	if body.Latest != nil {
		record.Latest = FromPtr(body.Latest)
	}

	if body.Recommended != nil {
		record.Recommended = FromPtr(body.Recommended)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Builds.Create(
		ctx,
		pack,
		record,
	); err != nil {
		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to validate build"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("action", "CreateBuild").
			Msg("Failed to create build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to create build"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, BuildResponse(
		a.convertBuild(record),
	))
}

// UpdateBuild implements the v1.ServerInterface.
func (a *API) UpdateBuild(w http.ResponseWriter, r *http.Request, _ PackID, _ BuildID) {
	ctx := r.Context()
	pack := a.PackFromContext(ctx)
	record := a.BuildFromContext(ctx)
	body := &UpdateBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("action", "UpdateBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if body.MinecraftID != nil {
		record.MinecraftID = body.MinecraftID
	}

	if body.ForgeID != nil {
		record.ForgeID = body.ForgeID
	}

	if body.NeoforgeID != nil {
		record.NeoforgeID = body.NeoforgeID
	}

	if body.QuiltID != nil {
		record.QuiltID = body.QuiltID
	}

	if body.FabricID != nil {
		record.FabricID = body.FabricID
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if body.Java != nil {
		record.Java = FromPtr(body.Java)
	}

	if body.Memory != nil {
		record.Memory = FromPtr(body.Memory)
	}

	if body.Latest != nil {
		record.Latest = FromPtr(body.Latest)
	}

	if body.Recommended != nil {
		record.Recommended = FromPtr(body.Recommended)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Builds.Update(
		ctx,
		pack,
		record,
	); err != nil {
		if v, ok := err.(validate.Errors); ok {
			errors := make([]Validation, 0)

			for _, verr := range v.Errors {
				errors = append(
					errors,
					Validation{
						Field:   ToPtr(verr.Field),
						Message: ToPtr(verr.Error.Error()),
					},
				)
			}

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to validate build"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("action", "UpdateBuild").
			Msg("Failed to update build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update build"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, BuildResponse(
		a.convertBuild(record),
	))
}

// DeleteBuild implements the v1.ServerInterface.
func (a *API) DeleteBuild(w http.ResponseWriter, r *http.Request, _ PackID, _ BuildID) {
	ctx := r.Context()
	pack := a.PackFromContext(ctx)
	record := a.BuildFromContext(ctx)

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Builds.Delete(
		ctx,
		pack,
		record.ID,
	); err != nil {
		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("action", "DeleteBuild").
			Msg("Failed to delete build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to delete build"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully deleted build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// ListBuildVersions implements the v1.ServerInterface.
func (a *API) ListBuildVersions(w http.ResponseWriter, r *http.Request, _ PackID, _ BuildID, params ListBuildVersionsParams) {
	ctx := r.Context()
	pack := a.PackFromContext(ctx)
	record := a.BuildFromContext(ctx)
	sort, order, limit, offset, search := listBuildVersionsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Builds.ListVersions(
		ctx,
		pack,
		record,
		model.ListParams{
			Sort:   sort,
			Order:  order,
			Limit:  limit,
			Offset: offset,
			Search: search,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("action", "ListBuildVersions").
			Msg("Failed to load build versions")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load build versions"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]BuildVersion, len(records))
	for id, record := range records {
		payload[id] = a.convertBuildVersion(record)
	}

	render.JSON(w, r, BuildVersionsResponse{
		Total:    count,
		Limit:    limit,
		Offset:   offset,
		Pack:     ToPtr(a.convertPack(pack)),
		Build:    ToPtr(a.convertBuild(record)),
		Versions: payload,
	})
}

// AttachBuildToVersion implements the v1.ServerInterface.
func (a *API) AttachBuildToVersion(w http.ResponseWriter, r *http.Request, _ PackID, _ BuildID) {
	ctx := r.Context()
	pack := a.PackFromContext(ctx)
	record := a.BuildFromContext(ctx)
	body := &BuildVersionBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("action", "AttachBuildToVersion").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Builds.AttachVersion(
		ctx,
		pack,
		record,
		model.BuildVersionParams{
			ModID:     body.Mod,
			VersionID: body.Version,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrVersionNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find version"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrAlreadyAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Version is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("mod", body.Mod).
			Str("version", body.Version).
			Str("action", "AttachBuildToVersion").
			Msg("Failed to attach build to version")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach build to version"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached build to version"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteBuildFromVersion implements the v1.ServerInterface.
func (a *API) DeleteBuildFromVersion(w http.ResponseWriter, r *http.Request, _ PackID, _ BuildID) {
	ctx := r.Context()
	pack := a.PackFromContext(ctx)
	record := a.BuildFromContext(ctx)
	body := &BuildVersionBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("action", "DeleteBuildFromVersion").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Builds.DropVersion(
		ctx,
		pack,
		record,
		model.BuildVersionParams{
			ModID:     body.Mod,
			VersionID: body.Version,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrVersionNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find version"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Version is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", pack.ID).
			Str("build", record.ID).
			Str("mod", body.Mod).
			Str("version", body.Version).
			Str("action", "DeleteBuildFromVersion").
			Msg("Failed to drop build from version")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop build from version"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped build from version"),
		Status:  ToPtr(http.StatusOK),
	})
}

// AllowCreateBuild defines a middleware to check permissions.
func (a *API) AllowCreateBuild(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowShowBuild defines a middleware to check permissions.
func (a *API) AllowShowBuild(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowManageBuild defines a middleware to check permissions.
func (a *API) AllowManageBuild(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *API) convertBuild(record *model.Build) Build {
	result := Build{
		ID:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Java:      ToPtr(record.Java),
		Memory:    ToPtr(record.Memory),
		Public:    ToPtr(record.Public),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	if record.Pack != nil {
		result.Pack = ToPtr(a.convertPack(record.Pack))
	}

	if record.MinecraftID != nil {
		result.MinecraftID = record.MinecraftID
		result.Minecraft = ToPtr(a.convertMinecraft(record.Minecraft))
	}

	if record.ForgeID != nil {
		result.ForgeID = record.ForgeID
		result.Forge = ToPtr(a.convertForge(record.Forge))
	}

	if record.NeoforgeID != nil {
		result.NeoforgeID = record.NeoforgeID
		result.Neoforge = ToPtr(a.convertNeoforge(record.Neoforge))
	}

	if record.QuiltID != nil {
		result.QuiltID = record.QuiltID
		result.Quilt = ToPtr(a.convertQuilt(record.Quilt))
	}

	if record.FabricID != nil {
		result.FabricID = record.FabricID
		result.Fabric = ToPtr(a.convertFabric(record.Fabric))
	}

	return result
}

func (a *API) convertBuildVersion(record *model.BuildVersion) BuildVersion {
	result := BuildVersion{
		BuildID:   record.BuildID,
		VersionID: record.VersionID,
		Version:   ToPtr(a.convertVersion(record.Version)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listBuildsSorting(request ListBuildsParams) (string, string, int64, int64, string) {
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

func listBuildVersionsSorting(request ListBuildVersionsParams) (string, string, int64, int64, string) {
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
