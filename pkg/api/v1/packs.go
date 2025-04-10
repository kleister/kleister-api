package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/store"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
)

// ListPacks implements the v1.ServerInterface.
func (a *API) ListPacks(w http.ResponseWriter, r *http.Request, params ListPacksParams) {
	ctx := r.Context()
	sort, order, limit, offset, search := listPacksSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.List(
		ctx,
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
			Str("action", "ListPacks").
			Msg("Failed to load packs")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load packs"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Pack, len(records))
	for id, record := range records {
		payload[id] = a.convertPack(record)
	}

	render.JSON(w, r, PacksResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Packs:  payload,
	})
}

// ShowPack implements the v1.ServerInterface.
func (a *API) ShowPack(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)

	render.JSON(w, r, PackResponse(
		a.convertPack(record),
	))
}

// CreatePack implements the v1.ServerInterface.
func (a *API) CreatePack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := &CreatePackBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("action", "CreatePack").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	record := &model.Pack{}

	if body.Slug != nil {
		record.Slug = FromPtr(body.Slug)
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if body.Website != nil {
		record.Website = FromPtr(body.Website)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.Create(
		ctx,
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
				Message: ToPtr("Failed to validate pack"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("action", "CreatePack").
			Msg("Failed to create pack")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to create pack"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, PackResponse(
		a.convertPack(record),
	))
}

// UpdatePack implements the v1.ServerInterface.
func (a *API) UpdatePack(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	body := &UpdatePackBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "UpdatePack").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if body.Slug != nil {
		record.Slug = FromPtr(body.Slug)
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if body.Website != nil {
		record.Website = FromPtr(body.Website)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.Update(
		ctx,
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
				Message: ToPtr("Failed to validate pack"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "UpdatePack").
			Msg("Failed to update pack")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update pack"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, PackResponse(
		a.convertPack(record),
	))
}

// DeletePack implements the v1.ServerInterface.
func (a *API) DeletePack(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.Delete(
		ctx,
		record.ID,
	); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "DeletePack").
			Msg("Failed to delete pack")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to delete pack"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully deleted pack"),
		Status:  ToPtr(http.StatusOK),
	})
}

// CreatePackAvatar implements the v1.ServerInterface.
func (a *API) CreatePackAvatar(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)

	if err := r.ParseMultipartForm(5 << 20); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "CreatePackAvatar").
			Msg("Failed to parse multipart")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to parse multipart"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	file, meta, err := r.FormFile("file")

	if err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "CreatePackAvatar").
			Msg("Failed to load avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load avatar"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	defer func() { _ = file.Close() }()
	buffer, err := a.resizeAvatar(file, meta)

	if err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "CreatePackAvatar").
			Msg("Failed resize avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed resize avatar"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	avatar, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.CreateAvatar(
		ctx,
		record.ID,
		buffer,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "CreatePackAvatar").
			Msg("Failed to store avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to store avatar"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	render.JSON(w, r, PackAvatarResponse(
		a.convertPackAvatar(avatar),
	))
}

// DeletePackAvatar implements the v1.ServerInterface.
func (a *API) DeletePackAvatar(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)

	avatar, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.DeleteAvatar(
		ctx,
		record.ID,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "DeletePackAvatar").
			Msg("Failed to delete pack avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to delete pack avatar"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	render.JSON(w, r, PackAvatarResponse(
		a.convertPackAvatar(avatar),
	))
}

// ListPackGroups implements the v1.ServerInterface.
func (a *API) ListPackGroups(w http.ResponseWriter, r *http.Request, _ UserID, params ListPackGroupsParams) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	sort, order, limit, offset, search := listPackGroupsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.ListGroups(
		ctx,
		model.GroupPackParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			PackID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "ListPackGroups").
			Msg("Failed to load pack groups")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load pack groups"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]GroupPack, len(records))
	for id, record := range records {
		payload[id] = a.convertPackGroup(record)
	}

	render.JSON(w, r, PackGroupsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Pack:   ToPtr(a.convertPack(record)),
		Groups: payload,
	})
}

// AttachPackToGroup implements the v1.ServerInterface.
func (a *API) AttachPackToGroup(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	body := &PackGroupPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "AttachPackToGroup").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.AttachGroup(
		ctx,
		model.GroupPackParams{
			PackID:  record.ID,
			GroupID: body.Group,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrAlreadyAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Group is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

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
				Message: ToPtr("Failed to validate pack group"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("group", body.Group).
			Str("action", "AttachPackToGroup").
			Msg("Failed to attach user to group")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach pack to group"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached pack to group"),
		Status:  ToPtr(http.StatusOK),
	})
}

// PermitPackGroup implements the v1.ServerInterface.
func (a *API) PermitPackGroup(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	body := &PackGroupPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "PermitPackGroup").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.PermitGroup(
		ctx,
		model.GroupPackParams{
			PackID:  record.ID,
			GroupID: body.Group,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Group is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

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
				Message: ToPtr("Failed to validate pack group"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("group", body.Group).
			Str("action", "PermitPackGroup").
			Msg("Failed to update pack group perms")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update pack group perms"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully updated pack group perms"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeletePackFromGroup implements the v1.ServerInterface.
func (a *API) DeletePackFromGroup(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	body := &PackGroupPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "DeletePackFromGroup").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.DropGroup(
		ctx,
		model.GroupPackParams{
			PackID:  record.ID,
			GroupID: body.Group,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Group is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("group", body.Group).
			Str("action", "DeletePackFromGroup").
			Msg("Failed to drop pack from group")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop pack from group"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped pack from group"),
		Status:  ToPtr(http.StatusOK),
	})
}

// ListPackUsers implements the v1.ServerInterface.
func (a *API) ListPackUsers(w http.ResponseWriter, r *http.Request, _ UserID, params ListPackUsersParams) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	sort, order, limit, offset, search := listPackUsersSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.ListUsers(
		ctx,
		model.UserPackParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			PackID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "ListPackUsers").
			Msg("Failed to load pack users")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load pack users"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]UserPack, len(records))
	for id, record := range records {
		payload[id] = a.convertPackUser(record)
	}

	render.JSON(w, r, PackUsersResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Pack:   ToPtr(a.convertPack(record)),
		Users:  payload,
	})
}

// AttachPackToUser implements the v1.ServerInterface.
func (a *API) AttachPackToUser(w http.ResponseWriter, r *http.Request, _ PackID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	body := &PackUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "AttachPackToUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.AttachUser(
		ctx,
		model.UserPackParams{
			PackID: record.ID,
			UserID: body.User,
			Perm:   body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrUserNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrAlreadyAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("User is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

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
				Message: ToPtr("Failed to validate pack user"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("user", body.User).
			Str("action", "AttachPackToUser").
			Msg("Failed to attach user to user")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach pack to user"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached pack to user"),
		Status:  ToPtr(http.StatusOK),
	})
}

// PermitPackUser implements the v1.ServerInterface.
func (a *API) PermitPackUser(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	body := &PackUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "PermitPackUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.PermitUser(
		ctx,
		model.UserPackParams{
			PackID: record.ID,
			UserID: body.User,
			Perm:   body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrUserNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

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
				Message: ToPtr("Failed to validate pack user"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("user", body.User).
			Str("action", "PermitPackUser").
			Msg("Failed to update pack user perms")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update pack user perms"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully updated pack user perms"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeletePackFromUser implements the v1.ServerInterface.
func (a *API) DeletePackFromUser(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.PackFromContext(ctx)
	body := &PackUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("action", "DeletePackFromUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Packs.DropUser(
		ctx,
		model.UserPackParams{
			PackID: record.ID,
			UserID: body.User,
		},
	); err != nil {
		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrUserNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("pack", record.ID).
			Str("user", body.User).
			Str("action", "DeletePackFromUser").
			Msg("Failed to drop pack from user")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop pack from user"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped pack from user"),
		Status:  ToPtr(http.StatusOK),
	})
}

// AllowCreatePack defines a middleware to check permissions.
func (a *API) AllowCreatePack(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowShowPack defines a middleware to check permissions.
func (a *API) AllowShowPack(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowManagePack defines a middleware to check permissions.
func (a *API) AllowManagePack(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *API) convertPack(record *model.Pack) Pack {
	result := Pack{
		ID:        ToPtr(record.ID),
		Slug:      ToPtr(record.Slug),
		Name:      ToPtr(record.Name),
		Website:   ToPtr(record.Website),
		Public:    ToPtr(record.Public),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	if record.Avatar != nil {
		result.Avatar = ToPtr(a.convertPackAvatar(record.Avatar))
	}

	return result
}

func (a *API) convertPackAvatar(record *model.PackAvatar) PackAvatar {
	avatar, err := url.JoinPath(
		a.config.Server.Host,
		a.config.Server.Root,
		"api/v1",
		"storage",
		"avatars",
		record.Slug,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("avataer", record.ID).
			Msg("Failed to generate avatar link")
	}

	result := PackAvatar{
		Slug:      ToPtr(record.Slug),
		URL:       ToPtr(avatar),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertPackGroup(record *model.GroupPack) GroupPack {
	result := GroupPack{
		PackID:    record.PackID,
		GroupID:   record.GroupID,
		Group:     ToPtr(a.convertGroup(record.Group)),
		Perm:      ToPtr(GroupPackPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertPackUser(record *model.UserPack) UserPack {
	result := UserPack{
		PackID:    record.PackID,
		UserID:    record.UserID,
		User:      ToPtr(a.convertUser(record.User)),
		Perm:      ToPtr(UserPackPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listPacksSorting(request ListPacksParams) (string, string, int64, int64, string) {
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

func listPackGroupsSorting(request ListPackGroupsParams) (string, string, int64, int64, string) {
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

func listPackUsersSorting(request ListPackUsersParams) (string, string, int64, int64, string) {
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
