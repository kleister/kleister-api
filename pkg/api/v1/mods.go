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

// ListMods implements the v1.ServerInterface.
func (a *API) ListMods(w http.ResponseWriter, r *http.Request, params ListModsParams) {
	ctx := r.Context()
	sort, order, limit, offset, search := listModsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.List(
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
			Str("action", "ListMods").
			Msg("Failed to load mods")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load mods"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Mod, len(records))
	for id, record := range records {
		payload[id] = a.convertMod(record)
	}

	render.JSON(w, r, ModsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Mods:   payload,
	})
}

// ShowMod implements the v1.ServerInterface.
func (a *API) ShowMod(w http.ResponseWriter, r *http.Request, _ ModID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)

	render.JSON(w, r, ModResponse(
		a.convertMod(record),
	))
}

// CreateMod implements the v1.ServerInterface.
func (a *API) CreateMod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := &CreateModBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("action", "CreateMod").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	record := &model.Mod{}

	if body.Slug != nil {
		record.Slug = FromPtr(body.Slug)
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if body.Side != nil {
		record.Side = FromPtr(body.Side)
	}

	if body.Description != nil {
		record.Description = FromPtr(body.Description)
	}

	if body.Author != nil {
		record.Author = FromPtr(body.Author)
	}

	if body.Website != nil {
		record.Website = FromPtr(body.Website)
	}

	if body.Donate != nil {
		record.Donate = FromPtr(body.Donate)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.Create(
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
				Message: ToPtr("Failed to validate mod"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("action", "CreateMod").
			Msg("Failed to create mod")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to create mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, ModResponse(
		a.convertMod(record),
	))
}

// UpdateMod implements the v1.ServerInterface.
func (a *API) UpdateMod(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	body := &UpdateModBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "UpdateMod").
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

	if body.Side != nil {
		record.Side = FromPtr(body.Side)
	}

	if body.Description != nil {
		record.Description = FromPtr(body.Description)
	}

	if body.Author != nil {
		record.Author = FromPtr(body.Author)
	}

	if body.Website != nil {
		record.Website = FromPtr(body.Website)
	}

	if body.Donate != nil {
		record.Donate = FromPtr(body.Donate)
	}

	if body.Public != nil {
		record.Public = FromPtr(body.Public)
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.Update(
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
				Message: ToPtr("Failed to validate mod"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "UpdateMod").
			Msg("Failed to update mod")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, ModResponse(
		a.convertMod(record),
	))
}

// DeleteMod implements the v1.ServerInterface.
func (a *API) DeleteMod(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.Delete(
		ctx,
		record.ID,
	); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "DeleteMod").
			Msg("Failed to delete mod")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to delete mod"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully deleted mod"),
		Status:  ToPtr(http.StatusOK),
	})
}

// CreateModAvatar implements the v1.ServerInterface.
func (a *API) CreateModAvatar(w http.ResponseWriter, r *http.Request, _ ModID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)

	if err := r.ParseMultipartForm(5 << 20); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "CreateModAvatar").
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
			Str("mod", record.ID).
			Str("action", "CreateModAvatar").
			Msg("Failed to load avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load avatar"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	defer file.Close()
	buffer, err := a.resizeAvatar(file, meta)

	if err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "CreateModAvatar").
			Msg("Failed resize avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed resize avatar"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	avatar, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.CreateAvatar(
		ctx,
		record.ID,
		buffer,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "CreateModAvatar").
			Msg("Failed store avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to store avatar"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	render.JSON(w, r, ModAvatarResponse(
		a.convertModAvatar(avatar),
	))
}

// DeleteModAvatar implements the v1.ServerInterface.
func (a *API) DeleteModAvatar(w http.ResponseWriter, r *http.Request, _ ModID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)

	avatar, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.DeleteAvatar(
		ctx,
		record.ID,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "DeleteModAvatar").
			Msg("Failed to delete mod avatar")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to delete mod avatar"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	render.JSON(w, r, ModAvatarResponse(
		a.convertModAvatar(avatar),
	))
}

// ListModGroups implements the v1.ServerInterface.
func (a *API) ListModGroups(w http.ResponseWriter, r *http.Request, _ UserID, params ListModGroupsParams) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	sort, order, limit, offset, search := listModGroupsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.ListGroups(
		ctx,
		model.GroupModParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			ModID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "ListModGroups").
			Msg("Failed to load mod groups")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load mod groups"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]GroupMod, len(records))
	for id, record := range records {
		payload[id] = a.convertModGroup(record)
	}

	render.JSON(w, r, ModGroupsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Mod:    ToPtr(a.convertMod(record)),
		Groups: payload,
	})
}

// AttachModToGroup implements the v1.ServerInterface.
func (a *API) AttachModToGroup(w http.ResponseWriter, r *http.Request, _ ModID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	body := &ModGroupPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "AttachModToGroup").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.AttachGroup(
		ctx,
		model.GroupModParams{
			ModID:   record.ID,
			GroupID: body.Group,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
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
				Message: ToPtr("Failed to validate mod group"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("group", body.Group).
			Str("action", "AttachModToGroup").
			Msg("Failed to attach user to group")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach mod to group"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached mod to group"),
		Status:  ToPtr(http.StatusOK),
	})
}

// PermitModGroup implements the v1.ServerInterface.
func (a *API) PermitModGroup(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	body := &ModGroupPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "PermitModGroup").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.PermitGroup(
		ctx,
		model.GroupModParams{
			ModID:   record.ID,
			GroupID: body.Group,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
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
				Message: ToPtr("Failed to validate mod group"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("group", body.Group).
			Str("action", "PermitModGroup").
			Msg("Failed to update mod group perms")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update mod group perms"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully updated mod group perms"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteModFromGroup implements the v1.ServerInterface.
func (a *API) DeleteModFromGroup(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	body := &ModGroupPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "DeleteModFromGroup").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.DropGroup(
		ctx,
		model.GroupModParams{
			ModID:   record.ID,
			GroupID: body.Group,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
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
			Str("mod", record.ID).
			Str("group", body.Group).
			Str("action", "DeleteModFromGroup").
			Msg("Failed to drop mod from group")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop mod from group"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped mod from group"),
		Status:  ToPtr(http.StatusOK),
	})
}

// ListModUsers implements the v1.ServerInterface.
func (a *API) ListModUsers(w http.ResponseWriter, r *http.Request, _ UserID, params ListModUsersParams) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	sort, order, limit, offset, search := listModUsersSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.ListUsers(
		ctx,
		model.UserModParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			ModID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "ListModUsers").
			Msg("Failed to load mod users")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load mod users"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]UserMod, len(records))
	for id, record := range records {
		payload[id] = a.convertModUser(record)
	}

	render.JSON(w, r, ModUsersResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Mod:    ToPtr(a.convertMod(record)),
		Users:  payload,
	})
}

// AttachModToUser implements the v1.ServerInterface.
func (a *API) AttachModToUser(w http.ResponseWriter, r *http.Request, _ ModID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	body := &ModUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "AttachModToUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.AttachUser(
		ctx,
		model.UserModParams{
			ModID:  record.ID,
			UserID: body.User,
			Perm:   body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
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
				Message: ToPtr("Failed to validate mod user"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("user", body.User).
			Str("action", "AttachModToUser").
			Msg("Failed to attach user to user")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to attach mod to user"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached mod to user"),
		Status:  ToPtr(http.StatusOK),
	})
}

// PermitModUser implements the v1.ServerInterface.
func (a *API) PermitModUser(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	body := &ModUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "PermitModUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.PermitUser(
		ctx,
		model.UserModParams{
			ModID:  record.ID,
			UserID: body.User,
			Perm:   body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
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
				Message: ToPtr("Failed to validate mod user"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("user", body.User).
			Str("action", "PermitModUser").
			Msg("Failed to update mod user perms")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update mod user perms"),
			Status:  ToPtr(http.StatusUnprocessableEntity),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully updated mod user perms"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteModFromUser implements the v1.ServerInterface.
func (a *API) DeleteModFromUser(w http.ResponseWriter, r *http.Request, _ UserID) {
	ctx := r.Context()
	record := a.ModFromContext(ctx)
	body := &ModUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("mod", record.ID).
			Str("action", "DeleteModFromUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Mods.DropUser(
		ctx,
		model.UserModParams{
			ModID:  record.ID,
			UserID: body.User,
		},
	); err != nil {
		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
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
			Str("mod", record.ID).
			Str("user", body.User).
			Str("action", "DeleteModFromUser").
			Msg("Failed to drop mod from user")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop mod from user"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped mod from user"),
		Status:  ToPtr(http.StatusOK),
	})
}

// AllowCreateMod defines a middleware to check permissions.
func (a *API) AllowCreateMod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowShowMod defines a middleware to check permissions.
func (a *API) AllowShowMod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AllowManageMod defines a middleware to check permissions.
func (a *API) AllowManageMod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *API) convertMod(record *model.Mod) Mod {
	result := Mod{
		ID:          ToPtr(record.ID),
		Slug:        ToPtr(record.Slug),
		Name:        ToPtr(record.Name),
		Side:        ToPtr(record.Side),
		Description: ToPtr(record.Description),
		Author:      ToPtr(record.Author),
		Website:     ToPtr(record.Website),
		Donate:      ToPtr(record.Donate),
		Public:      ToPtr(record.Public),
		CreatedAt:   ToPtr(record.CreatedAt),
		UpdatedAt:   ToPtr(record.UpdatedAt),
	}

	if record.Avatar != nil {
		result.Avatar = ToPtr(a.convertModAvatar(record.Avatar))
	}

	return result
}

func (a *API) convertModAvatar(record *model.ModAvatar) ModAvatar {
	avatar, err := url.JoinPath(
		a.config.Server.Host,
		a.config.Server.Root,
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

	result := ModAvatar{
		Slug:      ToPtr(record.Slug),
		URL:       ToPtr(avatar),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertModGroup(record *model.GroupMod) GroupMod {
	result := GroupMod{
		ModID:     record.ModID,
		GroupID:   record.GroupID,
		Group:     ToPtr(a.convertGroup(record.Group)),
		Perm:      ToPtr(GroupModPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertModUser(record *model.UserMod) UserMod {
	result := UserMod{
		ModID:     record.ModID,
		UserID:    record.UserID,
		User:      ToPtr(a.convertUser(record.User)),
		Perm:      ToPtr(UserModPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listModsSorting(request ListModsParams) (string, string, int64, int64, string) {
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

func listModGroupsSorting(request ListModGroupsParams) (string, string, int64, int64, string) {
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

func listModUsersSorting(request ListModUsersParams) (string, string, int64, int64, string) {
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
