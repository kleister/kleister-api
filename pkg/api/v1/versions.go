package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
	"github.com/vincent-petithory/dataurl"
)

// ListVersions implements the v1.ServerInterface.
func (a *API) ListVersions(w http.ResponseWriter, r *http.Request, _ ModID, params ListVersionsParams) {
	ctx := r.Context()
	mod := a.ModFromContext(ctx)
	sort, order, limit, offset, search := listVersionsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Versions.List(
		ctx,
		mod,
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
			Str("mod", mod.ID).
			Str("action", "ListVersions").
			Msg("Failed to load mods")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load mods"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Version, len(records))
	for id, record := range records {
		payload[id] = a.convertVersion(record)
	}

	render.JSON(w, r, VersionsResponse{
		Total:    count,
		Limit:    limit,
		Offset:   offset,
		Mod:      ToPtr(a.convertMod(mod)),
		Versions: payload,
	})
}

// ShowVersion implements the v1.ServerInterface.
func (a *API) ShowVersion(w http.ResponseWriter, r *http.Request, _ ModID, _ VersionID) {
	ctx := r.Context()
	record := a.VersionFromContext(ctx)

	render.JSON(w, r, VersionResponse(
		a.convertVersion(record),
	))
}

// CreateVersion implements the v1.ServerInterface.
func (a *API) CreateVersion(w http.ResponseWriter, r *http.Request, _ ModID) {
	ctx := r.Context()
	mod := a.ModFromContext(ctx)
	body := &CreateVersionBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("action", "CreateVersion").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	record := &model.Version{
		ModID: mod.ID,
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if body.Upload != nil {
		data, err := dataurl.DecodeString(
			FromPtr(body.Upload),
		)

		if err != nil {
			log.Error().
				Err(err).
				Str("mod", mod.ID).
				Str("action", "CreateVersion").
				Msg("Failed to decode upload")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to decode upload"),
				Status:  ToPtr(http.StatusBadRequest),
			})

			return
		}

		record.FileUpload = data
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Versions.Create(
		ctx,
		mod,
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
				Message: ToPtr("Failed to validate version"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("action", "CreateVersion").
			Msg("Failed to create version")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to create version"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	fmt.Printf("%+v\n", record.File)

	render.JSON(w, r, VersionResponse(
		a.convertVersion(record),
	))
}

// UpdateVersion implements the v1.ServerInterface.
func (a *API) UpdateVersion(w http.ResponseWriter, r *http.Request, _ ModID, _ VersionID) {
	ctx := r.Context()
	mod := a.ModFromContext(ctx)
	record := a.VersionFromContext(ctx)
	body := &UpdateVersionBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("action", "UpdateVersion").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if body.Upload != nil {
		data, err := dataurl.DecodeString(
			FromPtr(body.Upload),
		)

		if err != nil {
			log.Error().
				Err(err).
				Str("mod", mod.ID).
				Str("action", "CreateVersion").
				Msg("Failed to decode upload")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to decode upload"),
				Status:  ToPtr(http.StatusBadRequest),
			})

			return
		}

		record.FileUpload = data
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Versions.Update(
		ctx,
		mod,
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
				Message: ToPtr("Failed to validate version"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("action", "UpdateVersion").
			Msg("Failed to update version")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update version"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, VersionResponse(
		a.convertVersion(record),
	))
}

// DeleteVersion implements the v1.ServerInterface.
func (a *API) DeleteVersion(w http.ResponseWriter, r *http.Request, _ ModID, _ VersionID) {
	ctx := r.Context()
	mod := a.ModFromContext(ctx)
	record := a.VersionFromContext(ctx)

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Versions.Delete(
		ctx,
		mod,
		record.ID,
	); err != nil {
		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("action", "DeleteVersion").
			Msg("Failed to delete version")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to delete version"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully deleted version"),
		Status:  ToPtr(http.StatusOK),
	})
}

// ListVersionBuilds implements the v1.ServerInterface.
func (a *API) ListVersionBuilds(w http.ResponseWriter, r *http.Request, _ ModID, _ VersionID, params ListVersionBuildsParams) {
	ctx := r.Context()
	mod := a.ModFromContext(ctx)
	record := a.VersionFromContext(ctx)
	sort, order, limit, offset, search := listVersionBuildsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Versions.ListBuilds(
		ctx,
		mod,
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
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("action", "ListVersionBuilds").
			Msg("Failed to load version builds")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load version builds"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]BuildVersion, len(records))
	for id, record := range records {
		payload[id] = a.convertVersionBuild(record)
	}

	render.JSON(w, r, VersionBuildsResponse{
		Total:   count,
		Limit:   limit,
		Offset:  offset,
		Mod:     ToPtr(a.convertMod(mod)),
		Version: ToPtr(a.convertVersion(record)),
		Builds:  payload,
	})
}

// AttachVersionToBuild implements the v1.ServerInterface.
func (a *API) AttachVersionToBuild(w http.ResponseWriter, r *http.Request, _ ModID, _ VersionID) {
	ctx := r.Context()
	mod := a.ModFromContext(ctx)
	record := a.VersionFromContext(ctx)
	body := &VersionBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("action", "AttachVersionToBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Versions.AttachBuild(
		ctx,
		mod,
		record,
		model.BuildVersionParams{
			PackID:  body.Pack,
			BuildID: body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
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
				Message: ToPtr("Build is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "AttachVersionToBuild").
			Msg("Failed to attach version to build")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach version to build"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached version to build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteVersionFromBuild implements the v1.ServerInterface.
func (a *API) DeleteVersionFromBuild(w http.ResponseWriter, r *http.Request, _ ModID, _ VersionID) {
	ctx := r.Context()
	mod := a.ModFromContext(ctx)
	record := a.VersionFromContext(ctx)
	body := &VersionBuildBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("action", "DeleteVersionFromBuild").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Versions.DropBuild(
		ctx,
		mod,
		record,
		model.BuildVersionParams{
			PackID:  body.Pack,
			BuildID: body.Build,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
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

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Build is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("mod", mod.ID).
			Str("version", record.ID).
			Str("pack", body.Pack).
			Str("build", body.Build).
			Str("action", "DeleteUserFromGroup").
			Msg("Failed to drop version from build")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop version from build"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped version from build"),
		Status:  ToPtr(http.StatusOK),
	})
}

// AllowCreateVersion defines a middleware to check permissions.
func (a *API) AllowCreateVersion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowShowVersion defines a middleware to check permissions.
func (a *API) AllowShowVersion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowManageVersion defines a middleware to check permissions.
func (a *API) AllowManageVersion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *API) convertVersion(record *model.Version) Version {
	result := Version{
		ID:        ToPtr(record.ID),
		Name:      ToPtr(record.Name),
		Public:    ToPtr(record.Public),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	if record.Mod != nil {
		result.Mod = ToPtr(a.convertMod(record.Mod))
	}

	if record.File != nil {
		result.File = ToPtr(a.convertVersionFile(record.File))
	}

	return result
}

func (a *API) convertVersionFile(record *model.VersionFile) VersionFile {
	upload, err := url.JoinPath(
		a.config.Server.Host,
		a.config.Server.Root,
		"storage",
		"versions",
		record.Version.ModID,
		record.Slug,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("avataer", record.ID).
			Msg("Failed to generate version link")
	}

	result := VersionFile{
		Slug:        ToPtr(record.Slug),
		ContentType: ToPtr(record.ContentType),
		MD5:         ToPtr(record.MD5),
		URL:         ToPtr(upload),
		CreatedAt:   ToPtr(record.CreatedAt),
		UpdatedAt:   ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertVersionBuild(record *model.BuildVersion) BuildVersion {
	result := BuildVersion{
		VersionID: record.VersionID,
		BuildID:   record.BuildID,
		Build:     ToPtr(a.convertBuild(record.Build)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listVersionsSorting(request ListVersionsParams) (string, string, int64, int64, string) {
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

func listVersionBuildsSorting(request ListVersionBuildsParams) (string, string, int64, int64, string) {
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
