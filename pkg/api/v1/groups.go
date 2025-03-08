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

// ListGroups implements the v1.ServerInterface.
func (a *API) ListGroups(w http.ResponseWriter, r *http.Request, params ListGroupsParams) {
	ctx := r.Context()
	sort, order, limit, offset, search := listGroupsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.List(
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
			Str("action", "ListGroups").
			Msg("Failed to load groups")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load groups"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]Group, len(records))
	for id, record := range records {
		payload[id] = a.convertGroup(record)
	}

	render.JSON(w, r, GroupsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Groups: payload,
	})
}

// ShowGroup implements the v1.ServerInterface.
func (a *API) ShowGroup(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)

	render.JSON(w, r, GroupResponse(
		a.convertGroup(record),
	))
}

// CreateGroup implements the v1.ServerInterface.
func (a *API) CreateGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body := &CreateGroupBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("action", "CreateGroup").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	record := &model.Group{}

	if body.Slug != nil {
		record.Slug = FromPtr(body.Slug)
	}

	if body.Name != nil {
		record.Name = FromPtr(body.Name)
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.Create(
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("action", "CreateGroup").
			Msg("Failed to create group")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to create group"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, GroupResponse(
		a.convertGroup(record),
	))
}

// UpdateGroup implements the v1.ServerInterface.
func (a *API) UpdateGroup(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &CreateGroupBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "UpdateGroup").
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

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.Update(
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "UpdateGroup").
			Msg("Failed to update group")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update group"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, GroupResponse(
		a.convertGroup(record),
	))
}

// DeleteGroup implements the v1.ServerInterface.
func (a *API) DeleteGroup(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.Delete(
		ctx,
		record.ID,
	); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "DeleteGroup").
			Msg("Failed to delete group")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete group"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted group"),
	})
}

// ListGroupUsers implements the v1.ServerInterface.
func (a *API) ListGroupUsers(w http.ResponseWriter, r *http.Request, _ GroupID, params ListGroupUsersParams) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	sort, order, limit, offset, search := listGroupUsersSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.ListUsers(
		ctx,
		model.UserGroupParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			GroupID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "ListGroupUsers").
			Msg("Failed to load group users")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load group users"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]UserGroup, len(records))
	for id, record := range records {
		payload[id] = a.convertGroupUser(record)
	}

	render.JSON(w, r, GroupUsersResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Group:  ToPtr(a.convertGroup(record)),
		Users:  payload,
	})
}

// AttachGroupToUser implements the v1.ServerInterface.
func (a *API) AttachGroupToUser(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "AttachGroupToUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.AttachUser(
		ctx,
		model.UserGroupParams{
			GroupID: record.ID,
			UserID:  body.User,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group user"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("user", body.User).
			Str("action", "AttachGroupToUser").
			Msg("Failed to attach group to user")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach group to user"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached group to user"),
		Status:  ToPtr(http.StatusOK),
	})
}

// PermitGroupUser implements the v1.ServerInterface.
func (a *API) PermitGroupUser(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "PermitGroupUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.PermitUser(
		ctx,
		model.UserGroupParams{
			GroupID: record.ID,
			UserID:  body.User,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group user"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("user", body.User).
			Str("action", "PermitGroupUser").
			Msg("Failed to update group user perms")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update group user perms"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully updated group user perms"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteGroupFromUser implements the v1.ServerInterface.
func (a *API) DeleteGroupFromUser(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupUserPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "DeleteGroupFromUser").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.DropUser(
		ctx,
		model.UserGroupParams{
			GroupID: record.ID,
			UserID:  body.User,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
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
			Str("group", record.ID).
			Str("user", body.User).
			Str("action", "DeleteGroupFromUser").
			Msg("Failed to drop group from user")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop group from user"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped group from user"),
		Status:  ToPtr(http.StatusOK),
	})
}

// ListGroupMods implements the v1.ServerInterface.
func (a *API) ListGroupMods(w http.ResponseWriter, r *http.Request, _ GroupID, params ListGroupModsParams) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	sort, order, limit, offset, search := listGroupModsSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.ListMods(
		ctx,
		model.GroupModParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			GroupID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "ListGroupMods").
			Msg("Failed to load group mods")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load group mods"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]GroupMod, len(records))
	for id, record := range records {
		payload[id] = a.convertGroupMod(record)
	}

	render.JSON(w, r, GroupModsResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Group:  ToPtr(a.convertGroup(record)),
		Mods:   payload,
	})
}

// AttachGroupToMod implements the v1.ServerInterface.
func (a *API) AttachGroupToMod(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupModPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "AttachGroupToMod").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.AttachMod(
		ctx,
		model.GroupModParams{
			GroupID: record.ID,
			ModID:   body.Mod,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrAlreadyAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Mod is already attached"),
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group mod"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("mod", body.Mod).
			Str("action", "AttachGroupToMod").
			Msg("Failed to attach group to mod")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach group to mod"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached group to mod"),
		Status:  ToPtr(http.StatusOK),
	})
}

// PermitGroupMod implements the v1.ServerInterface.
func (a *API) PermitGroupMod(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupModPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "PermitGroupMod").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.PermitMod(
		ctx,
		model.GroupModParams{
			GroupID: record.ID,
			ModID:   body.Mod,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Mod is not attached"),
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group mod"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("mod", body.Mod).
			Str("action", "PermitGroupMod").
			Msg("Failed to update group mod perms")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update group mod perms"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully updated group mod perms"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteGroupFromMod implements the v1.ServerInterface.
func (a *API) DeleteGroupFromMod(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupModPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "DeleteGroupFromMod").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.DropMod(
		ctx,
		model.GroupModParams{
			GroupID: record.ID,
			ModID:   body.Mod,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrModNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Mod is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("mod", body.Mod).
			Str("action", "DeleteGroupFromMod").
			Msg("Failed to drop group from mod")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop group from mod"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped group from mod"),
		Status:  ToPtr(http.StatusOK),
	})
}

// ListGroupPacks implements the v1.ServerInterface.
func (a *API) ListGroupPacks(w http.ResponseWriter, r *http.Request, _ GroupID, params ListGroupPacksParams) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	sort, order, limit, offset, search := listGroupPacksSorting(params)

	records, count, err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.ListPacks(
		ctx,
		model.GroupPackParams{
			ListParams: model.ListParams{
				Sort:   sort,
				Order:  order,
				Limit:  limit,
				Offset: offset,
				Search: search,
			},
			GroupID: record.ID,
		},
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "ListGroupPacks").
			Msg("Failed to load group packs")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to load group packs"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	payload := make([]GroupPack, len(records))
	for id, record := range records {
		payload[id] = a.convertGroupPack(record)
	}

	render.JSON(w, r, GroupPacksResponse{
		Total:  count,
		Limit:  limit,
		Offset: offset,
		Group:  ToPtr(a.convertGroup(record)),
		Packs:  payload,
	})
}

// AttachGroupToPack implements the v1.ServerInterface.
func (a *API) AttachGroupToPack(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupPackPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "AttachGroupToPack").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.AttachPack(
		ctx,
		model.GroupPackParams{
			GroupID: record.ID,
			PackID:  body.Pack,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrAlreadyAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Pack is already attached"),
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group pack"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("pack", body.Pack).
			Str("action", "AttachGroupToPack").
			Msg("Failed to attach group to pack")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach group to pack"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully attached group to pack"),
		Status:  ToPtr(http.StatusOK),
	})
}

// PermitGroupPack implements the v1.ServerInterface.
func (a *API) PermitGroupPack(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupPackPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "PermitGroupPack").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.PermitPack(
		ctx,
		model.GroupPackParams{
			GroupID: record.ID,
			PackID:  body.Pack,
			Perm:    body.Perm,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Pack is not attached"),
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
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate group pack"),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("pack", body.Pack).
			Str("action", "PermitGroupPack").
			Msg("Failed to update group pack perms")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update group pack perms"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully updated group pack perms"),
		Status:  ToPtr(http.StatusOK),
	})
}

// DeleteGroupFromPack implements the v1.ServerInterface.
func (a *API) DeleteGroupFromPack(w http.ResponseWriter, r *http.Request, _ GroupID) {
	ctx := r.Context()
	record := a.GroupFromContext(ctx)
	body := &GroupPackPermBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("action", "DeleteGroupFromPack").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	if err := a.storage.WithPrincipal(
		current.GetUser(ctx),
	).Groups.DropPack(
		ctx,
		model.GroupPackParams{
			GroupID: record.ID,
			PackID:  body.Pack,
		},
	); err != nil {
		if errors.Is(err, store.ErrGroupNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find group"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrPackNotFound) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		if errors.Is(err, store.ErrNotAssigned) {
			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Pack is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			})

			return
		}

		log.Error().
			Err(err).
			Str("group", record.ID).
			Str("pack", body.Pack).
			Str("action", "DeleteGroupFromPack").
			Msg("Failed to drop group from pack")

		a.RenderNotify(w, r, Notification{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop group from pack"),
		})

		return
	}

	a.RenderNotify(w, r, Notification{
		Message: ToPtr("Successfully dropped group from pack"),
		Status:  ToPtr(http.StatusOK),
	})
}

func (a *API) convertGroup(record *model.Group) Group {
	result := Group{
		ID:        ToPtr(record.ID),
		Slug:      ToPtr(record.Slug),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertGroupUser(record *model.UserGroup) UserGroup {
	result := UserGroup{
		UserID:    record.UserID,
		User:      ToPtr(a.convertUser(record.User)),
		GroupID:   record.GroupID,
		Perm:      ToPtr(UserGroupPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertGroupMod(record *model.GroupMod) GroupMod {
	result := GroupMod{
		GroupID:   record.GroupID,
		ModID:     record.ModID,
		Mod:       ToPtr(a.convertMod(record.Mod)),
		Perm:      ToPtr(GroupModPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertGroupPack(record *model.GroupPack) GroupPack {
	result := GroupPack{
		GroupID:   record.GroupID,
		PackID:    record.PackID,
		Pack:      ToPtr(a.convertPack(record.Pack)),
		Perm:      ToPtr(GroupPackPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func listGroupsSorting(request ListGroupsParams) (string, string, int64, int64, string) {
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

func listGroupUsersSorting(request ListGroupUsersParams) (string, string, int64, int64, string) {
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

func listGroupModsSorting(request ListGroupModsParams) (string, string, int64, int64, string) {
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

func listGroupPacksSorting(request ListGroupPacksParams) (string, string, int64, int64, string) {
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
