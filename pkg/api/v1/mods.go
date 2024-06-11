package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods"
	teammods "github.com/kleister/kleister-api/pkg/service/team_mods"
	usermods "github.com/kleister/kleister-api/pkg/service/user_mods"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListMods implements the v1.ServerInterface.
func (a *API) ListMods(ctx context.Context, request ListModsRequestObject) (ListModsResponseObject, error) {
	records, count, err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		toListParams(
			string(FromPtr(request.Params.Sort)),
			string(FromPtr(request.Params.Order)),
			request.Params.Limit,
			request.Params.Offset,
			request.Params.Search,
		),
	)

	if err != nil {
		return ListMods500JSONResponse{
			Message: ToPtr("Failed to load mods"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Mod, len(records))
	for id, record := range records {
		payload[id] = a.convertMod(record)
	}

	return ListMods200JSONResponse{
		Total: ToPtr(count),
		Mods:  ToPtr(payload),
	}, nil
}

// ShowMod implements the v1.ServerInterface.
func (a *API) ShowMod(ctx context.Context, request ShowModRequestObject) (ShowModResponseObject, error) {
	record, err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.ModId,
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return ShowMod404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowMod500JSONResponse{
			Message: ToPtr("Failed to load mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowMod200JSONResponse(
		a.convertMod(record),
	), nil
}

// CreateMod implements the v1.ServerInterface.
func (a *API) CreateMod(ctx context.Context, request CreateModRequestObject) (CreateModResponseObject, error) {
	record := &model.Mod{}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Side != nil {
		record.Side = string(ModSide(FromPtr(request.Body.Side)))
	}

	if request.Body.Description != nil {
		record.Description = FromPtr(request.Body.Description)
	}

	if request.Body.Author != nil {
		record.Author = FromPtr(request.Body.Author)
	}

	if request.Body.Website != nil {
		record.Website = FromPtr(request.Body.Website)
	}

	if request.Body.Donate != nil {
		record.Donate = FromPtr(request.Body.Donate)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	if err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Create(
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

			return CreateMod422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate mod"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreateMod500JSONResponse{
			Message: ToPtr("Failed to create mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreateMod200JSONResponse(
		a.convertMod(record),
	), nil
}

// UpdateMod implements the v1.ServerInterface.
func (a *API) UpdateMod(ctx context.Context, request UpdateModRequestObject) (UpdateModResponseObject, error) {
	record, err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.ModId,
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return UpdateMod404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdateMod500JSONResponse{
			Message: ToPtr("Failed to load mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Side != nil {
		record.Side = string(ModSide(FromPtr(request.Body.Side)))
	}

	if request.Body.Description != nil {
		record.Description = FromPtr(request.Body.Description)
	}

	if request.Body.Author != nil {
		record.Author = FromPtr(request.Body.Author)
	}

	if request.Body.Website != nil {
		record.Website = FromPtr(request.Body.Website)
	}

	if request.Body.Donate != nil {
		record.Donate = FromPtr(request.Body.Donate)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	if err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Update(
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

			return UpdateMod422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate mod"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateMod500JSONResponse{
			Message: ToPtr("Failed to update mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateMod200JSONResponse(
		a.convertMod(record),
	), nil
}

// DeleteMod implements the v1.ServerInterface.
func (a *API) DeleteMod(ctx context.Context, request DeleteModRequestObject) (DeleteModResponseObject, error) {
	record, err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.ModId,
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return DeleteMod404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeleteMod500JSONResponse{
			Message: ToPtr("Failed to load mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Delete(
		ctx,
		record.ID,
	); err != nil {
		return DeleteMod400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete mod"),
		}, nil
	}

	return DeleteMod200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted mod"),
	}, nil
}

// ListModTeams implements the v1.ServerInterface.
func (a *API) ListModTeams(ctx context.Context, request ListModTeamsRequestObject) (ListModTeamsResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListModTeams403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.ModId,
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return ListModTeams404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListModTeams500JSONResponse{
			Message: ToPtr("Failed to load mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.teammods.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.TeamModParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			ModID: record.ID,
		},
	)

	if err != nil {
		return ListModTeams500JSONResponse{
			Message: ToPtr("Failed to load teams"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]TeamMod, len(records))
	for id, record := range records {
		payload[id] = a.convertModTeam(record)
	}

	return ListModTeams200JSONResponse{
		Total: ToPtr(count),
		Mod:   ToPtr(a.convertMod(record)),
		Teams: ToPtr(payload),
	}, nil
}

// AttachModToTeam implements the v1.ServerInterface.
func (a *API) AttachModToTeam(ctx context.Context, request AttachModToTeamRequestObject) (AttachModToTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachModToTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teammods.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.TeamModParams{
			ModID:  request.ModId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teammods.ErrNotFound) {
			return AttachModToTeam404JSONResponse{
				Message: ToPtr("Failed to find mod or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teammods.ErrAlreadyAssigned) {
			return AttachModToTeam412JSONResponse{
				Message: ToPtr("Team is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
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

			return AttachModToTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate mod team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachModToTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach mod to team"),
		}, nil
	}

	return AttachModToTeam200JSONResponse{
		Message: ToPtr("Successfully attached mod to team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitModTeam implements the v1.ServerInterface.
func (a *API) PermitModTeam(ctx context.Context, request PermitModTeamRequestObject) (PermitModTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitModTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teammods.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.TeamModParams{
			ModID:  request.ModId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teammods.ErrNotFound) {
			return PermitModTeam404JSONResponse{
				Message: ToPtr("Failed to find mod or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teammods.ErrNotAssigned) {
			return PermitModTeam412JSONResponse{
				Message: ToPtr("Team is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
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

			return PermitModTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate mod team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitModTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update mod team perms"),
		}, nil
	}

	return PermitModTeam200JSONResponse{
		Message: ToPtr("Successfully updated mod team perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteModFromTeam implements the v1.ServerInterface.
func (a *API) DeleteModFromTeam(ctx context.Context, request DeleteModFromTeamRequestObject) (DeleteModFromTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteModFromTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teammods.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.TeamModParams{
			ModID:  request.ModId,
			TeamID: request.Body.Team,
		},
	); err != nil {
		if errors.Is(err, teammods.ErrNotFound) {
			return DeleteModFromTeam404JSONResponse{
				Message: ToPtr("Failed to find mod or team"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, teammods.ErrNotAssigned) {
			return DeleteModFromTeam412JSONResponse{
				Message: ToPtr("Team is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteModFromTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop mod from team"),
		}, nil
	}

	return DeleteModFromTeam200JSONResponse{
		Message: ToPtr("Successfully dropped mod from team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListModUsers implements the v1.ServerInterface.
func (a *API) ListModUsers(ctx context.Context, request ListModUsersRequestObject) (ListModUsersResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListModUsers403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.mods.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.ModId,
	)

	if err != nil {
		if errors.Is(err, mods.ErrNotFound) {
			return ListModUsers404JSONResponse{
				Message: ToPtr("Failed to find mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListModUsers500JSONResponse{
			Message: ToPtr("Failed to load mod"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.usermods.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.UserModParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			ModID: record.ID,
		},
	)

	if err != nil {
		return ListModUsers500JSONResponse{
			Message: ToPtr("Failed to load users"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserMod, len(records))
	for id, record := range records {
		payload[id] = a.convertModUser(record)
	}

	return ListModUsers200JSONResponse{
		Total: ToPtr(count),
		Mod:   ToPtr(a.convertMod(record)),
		Users: ToPtr(payload),
	}, nil
}

// AttachModToUser implements the v1.ServerInterface.
func (a *API) AttachModToUser(ctx context.Context, request AttachModToUserRequestObject) (AttachModToUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachModToUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.usermods.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.UserModParams{
			ModID:  request.ModId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, usermods.ErrNotFound) {
			return AttachModToUser404JSONResponse{
				Message: ToPtr("Failed to find mod or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, usermods.ErrAlreadyAssigned) {
			return AttachModToUser412JSONResponse{
				Message: ToPtr("User is already attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
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

			return AttachModToUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate mod user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachModToUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach mod to user"),
		}, nil
	}

	return AttachModToUser200JSONResponse{
		Message: ToPtr("Successfully attached mod to user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitModUser implements the v1.ServerInterface.
func (a *API) PermitModUser(ctx context.Context, request PermitModUserRequestObject) (PermitModUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitModUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.usermods.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.UserModParams{
			ModID:  request.ModId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, usermods.ErrNotFound) {
			return PermitModUser404JSONResponse{
				Message: ToPtr("Failed to find mod or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, usermods.ErrNotAssigned) {
			return PermitModUser412JSONResponse{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
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

			return PermitModUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate mod user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitModUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update mod user perms"),
		}, nil
	}

	return PermitModUser200JSONResponse{
		Message: ToPtr("Successfully updated mod user perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteModFromUser implements the v1.ServerInterface.
func (a *API) DeleteModFromUser(ctx context.Context, request DeleteModFromUserRequestObject) (DeleteModFromUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteModFromUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.usermods.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.UserModParams{
			ModID:  request.ModId,
			UserID: request.Body.User,
		},
	); err != nil {
		if errors.Is(err, usermods.ErrNotFound) {
			return DeleteModFromUser404JSONResponse{
				Message: ToPtr("Failed to find mod or user"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, usermods.ErrNotAssigned) {
			return DeleteModFromUser412JSONResponse{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteModFromUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop mod from user"),
		}, nil
	}

	return DeleteModFromUser200JSONResponse{
		Message: ToPtr("Successfully dropped mod from user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertMod(record *model.Mod) Mod {
	result := Mod{
		Id:          ToPtr(record.ID),
		Slug:        ToPtr(record.Slug),
		Name:        ToPtr(record.Name),
		Side:        ToPtr(ModSide(record.Side)),
		Description: ToPtr(record.Description),
		Author:      ToPtr(record.Author),
		Website:     ToPtr(record.Website),
		Donate:      ToPtr(record.Donate),
		Public:      ToPtr(record.Public),
		CreatedAt:   ToPtr(record.CreatedAt),
		UpdatedAt:   ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertModTeam(record *model.TeamMod) TeamMod {
	result := TeamMod{
		ModId:     record.ModID,
		TeamId:    record.TeamID,
		Team:      ToPtr(a.convertTeam(record.Team)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertModUser(record *model.UserMod) UserMod {
	result := UserMod{
		ModId:     record.ModID,
		UserId:    record.UserID,
		User:      ToPtr(a.convertUser(record.User)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
