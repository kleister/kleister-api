package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	usermods "github.com/kleister/kleister-api/pkg/service/user_mods"
	userpacks "github.com/kleister/kleister-api/pkg/service/user_packs"
	userteams "github.com/kleister/kleister-api/pkg/service/user_teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListUsers implements the v1.ServerInterface.
func (a *API) ListUsers(ctx context.Context, request ListUsersRequestObject) (ListUsersResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListUsers403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	records, count, err := a.users.WithPrincipal(
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
		return ListUsers500JSONResponse{
			Message: ToPtr("Failed to load users"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]User, len(records))
	for id, record := range records {
		payload[id] = a.convertUser(record)
	}

	return ListUsers200JSONResponse{
		Total: ToPtr(count),
		Users: ToPtr(payload),
	}, nil
}

// ShowUser implements the v1.ServerInterface.
func (a *API) ShowUser(ctx context.Context, request ShowUserRequestObject) (ShowUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ShowUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return ShowUser404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowUser500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowUser200JSONResponse(
		a.convertUser(record),
	), nil
}

// CreateUser implements the v1.ServerInterface.
func (a *API) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return CreateUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record := &model.User{}

	if request.Body.Username != nil {
		record.Username = FromPtr(request.Body.Username)
	}

	if request.Body.Password != nil {
		record.Password = FromPtr(request.Body.Password)
	}

	if request.Body.Email != nil {
		record.Email = FromPtr(request.Body.Email)
	}

	if request.Body.Fullname != nil {
		record.Fullname = FromPtr(request.Body.Fullname)
	}

	if request.Body.Admin != nil {
		record.Admin = FromPtr(request.Body.Admin)
	}

	if request.Body.Active != nil {
		record.Active = FromPtr(request.Body.Active)
	}

	if err := a.users.WithPrincipal(
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

			return CreateUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreateUser500JSONResponse{
			Message: ToPtr("Failed to create user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreateUser200JSONResponse(
		a.convertUser(record),
	), nil
}

// UpdateUser implements the v1.ServerInterface.
func (a *API) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return UpdateUser404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdateUser500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.Username != nil {
		record.Username = FromPtr(request.Body.Username)
	}

	if request.Body.Password != nil {
		record.Password = FromPtr(request.Body.Password)
	}

	if request.Body.Email != nil {
		record.Email = FromPtr(request.Body.Email)
	}

	if request.Body.Fullname != nil {
		record.Fullname = FromPtr(request.Body.Fullname)
	}

	if request.Body.Admin != nil {
		record.Admin = FromPtr(request.Body.Admin)
	}

	if request.Body.Active != nil {
		record.Active = FromPtr(request.Body.Active)
	}

	if err := a.users.WithPrincipal(
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

			return UpdateUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateUser500JSONResponse{
			Message: ToPtr("Failed to update user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateUser200JSONResponse(
		a.convertUser(record),
	), nil
}

// DeleteUser implements the v1.ServerInterface.
func (a *API) DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return DeleteUser404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeleteUser500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.users.WithPrincipal(
		current.GetUser(ctx),
	).Delete(
		ctx,
		record.ID,
	); err != nil {
		return DeleteUser400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete user"),
		}, nil
	}

	return DeleteUser200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted user"),
	}, nil
}

// ListUserTeams implements the v1.ServerInterface.
func (a *API) ListUserTeams(ctx context.Context, request ListUserTeamsRequestObject) (ListUserTeamsResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListUserTeams403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return ListUserTeams404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListUserTeams500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.userteams.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.UserTeamParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			UserID: record.ID,
		},
	)

	if err != nil {
		return ListUserTeams500JSONResponse{
			Message: ToPtr("Failed to load teams"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserTeam, len(records))
	for id, record := range records {
		payload[id] = a.convertUserTeam(record)
	}

	return ListUserTeams200JSONResponse{
		Total: ToPtr(count),
		User:  ToPtr(a.convertUser(record)),
		Teams: ToPtr(payload),
	}, nil
}

// AttachUserToTeam implements the v1.ServerInterface.
func (a *API) AttachUserToTeam(ctx context.Context, request AttachUserToTeamRequestObject) (AttachUserToTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachUserToTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userteams.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.UserTeamParams{
			UserID: request.UserId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userteams.ErrNotFound) {
			return AttachUserToTeam404JSONResponse{
				Message: ToPtr("Failed to find user or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userteams.ErrAlreadyAssigned) {
			return AttachUserToTeam412JSONResponse{
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

			return AttachUserToTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachUserToTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach user to team"),
		}, nil
	}

	return AttachUserToTeam200JSONResponse{
		Message: ToPtr("Successfully attached user to team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitUserTeam implements the v1.ServerInterface.
func (a *API) PermitUserTeam(ctx context.Context, request PermitUserTeamRequestObject) (PermitUserTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitUserTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userteams.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.UserTeamParams{
			UserID: request.UserId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userteams.ErrNotFound) {
			return PermitUserTeam404JSONResponse{
				Message: ToPtr("Failed to find user or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userteams.ErrNotAssigned) {
			return PermitUserTeam412JSONResponse{
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

			return PermitUserTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitUserTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update user team perms"),
		}, nil
	}

	return PermitUserTeam200JSONResponse{
		Message: ToPtr("Successfully updated user team perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteUserFromTeam implements the v1.ServerInterface.
func (a *API) DeleteUserFromTeam(ctx context.Context, request DeleteUserFromTeamRequestObject) (DeleteUserFromTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteUserFromTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userteams.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.UserTeamParams{
			UserID: request.UserId,
			TeamID: request.Body.Team,
		},
	); err != nil {
		if errors.Is(err, userteams.ErrNotFound) {
			return DeleteUserFromTeam404JSONResponse{
				Message: ToPtr("Failed to find user or team"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, userteams.ErrNotAssigned) {
			return DeleteUserFromTeam412JSONResponse{
				Message: ToPtr("Team is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteUserFromTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop user from team"),
		}, nil
	}

	return DeleteUserFromTeam200JSONResponse{
		Message: ToPtr("Successfully dropped user from team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListUserPacks implements the v1.ServerInterface.
func (a *API) ListUserPacks(ctx context.Context, request ListUserPacksRequestObject) (ListUserPacksResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListUserPacks403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return ListUserPacks404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListUserPacks500JSONResponse{
			Message: ToPtr("Failed to load user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.userpacks.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.UserPackParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			UserID: record.ID,
		},
	)

	if err != nil {
		return ListUserPacks500JSONResponse{
			Message: ToPtr("Failed to load packs"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserPack, len(records))
	for id, record := range records {
		payload[id] = a.convertUserPack(record)
	}

	return ListUserPacks200JSONResponse{
		Total: ToPtr(count),
		User:  ToPtr(a.convertUser(record)),
		Packs: ToPtr(payload),
	}, nil
}

// AttachUserToPack implements the v1.ServerInterface.
func (a *API) AttachUserToPack(ctx context.Context, request AttachUserToPackRequestObject) (AttachUserToPackResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachUserToPack403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userpacks.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.UserPackParams{
			UserID: request.UserId,
			PackID: request.Body.Pack,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userpacks.ErrNotFound) {
			return AttachUserToPack404JSONResponse{
				Message: ToPtr("Failed to find user or pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userpacks.ErrAlreadyAssigned) {
			return AttachUserToPack412JSONResponse{
				Message: ToPtr("Pack is already attached"),
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

			return AttachUserToPack422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user pack"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachUserToPack500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach user to pack"),
		}, nil
	}

	return AttachUserToPack200JSONResponse{
		Message: ToPtr("Successfully attached user to pack"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitUserPack implements the v1.ServerInterface.
func (a *API) PermitUserPack(ctx context.Context, request PermitUserPackRequestObject) (PermitUserPackResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitUserPack403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userpacks.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.UserPackParams{
			UserID: request.UserId,
			PackID: request.Body.Pack,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userpacks.ErrNotFound) {
			return PermitUserPack404JSONResponse{
				Message: ToPtr("Failed to find user or pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userpacks.ErrNotAssigned) {
			return PermitUserPack412JSONResponse{
				Message: ToPtr("Pack is not attached"),
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

			return PermitUserPack422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user pack"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitUserPack500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update user pack perms"),
		}, nil
	}

	return PermitUserPack200JSONResponse{
		Message: ToPtr("Successfully updated user pack perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteUserFromPack implements the v1.ServerInterface.
func (a *API) DeleteUserFromPack(ctx context.Context, request DeleteUserFromPackRequestObject) (DeleteUserFromPackResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteUserFromPack403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userpacks.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.UserPackParams{
			UserID: request.UserId,
			PackID: request.Body.Pack,
		},
	); err != nil {
		if errors.Is(err, userpacks.ErrNotFound) {
			return DeleteUserFromPack404JSONResponse{
				Message: ToPtr("Failed to find user or pack"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, userpacks.ErrNotAssigned) {
			return DeleteUserFromPack412JSONResponse{
				Message: ToPtr("Pack is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteUserFromPack500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop user from pack"),
		}, nil
	}

	return DeleteUserFromPack200JSONResponse{
		Message: ToPtr("Successfully dropped user from pack"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListUserMods implements the v1.ServerInterface.
func (a *API) ListUserMods(ctx context.Context, request ListUserModsRequestObject) (ListUserModsResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListUserMods403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.users.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.UserId,
	)

	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			return ListUserMods404JSONResponse{
				Message: ToPtr("Failed to find user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListUserMods500JSONResponse{
			Message: ToPtr("Failed to load user"),
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
			UserID: record.ID,
		},
	)

	if err != nil {
		return ListUserMods500JSONResponse{
			Message: ToPtr("Failed to load mods"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserMod, len(records))
	for id, record := range records {
		payload[id] = a.convertUserMod(record)
	}

	return ListUserMods200JSONResponse{
		Total: ToPtr(count),
		User:  ToPtr(a.convertUser(record)),
		Mods:  ToPtr(payload),
	}, nil
}

// AttachUserToMod implements the v1.ServerInterface.
func (a *API) AttachUserToMod(ctx context.Context, request AttachUserToModRequestObject) (AttachUserToModResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachUserToMod403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.usermods.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.UserModParams{
			UserID: request.UserId,
			ModID:  request.Body.Mod,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, usermods.ErrNotFound) {
			return AttachUserToMod404JSONResponse{
				Message: ToPtr("Failed to find user or mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, usermods.ErrAlreadyAssigned) {
			return AttachUserToMod412JSONResponse{
				Message: ToPtr("Mod is already attached"),
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

			return AttachUserToMod422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user mod"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachUserToMod500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach user to mod"),
		}, nil
	}

	return AttachUserToMod200JSONResponse{
		Message: ToPtr("Successfully attached user to mod"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitUserMod implements the v1.ServerInterface.
func (a *API) PermitUserMod(ctx context.Context, request PermitUserModRequestObject) (PermitUserModResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitUserMod403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.usermods.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.UserModParams{
			UserID: request.UserId,
			ModID:  request.Body.Mod,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, usermods.ErrNotFound) {
			return PermitUserMod404JSONResponse{
				Message: ToPtr("Failed to find user or mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, usermods.ErrNotAssigned) {
			return PermitUserMod412JSONResponse{
				Message: ToPtr("Mod is not attached"),
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

			return PermitUserMod422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate user mod"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitUserMod500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update user mod perms"),
		}, nil
	}

	return PermitUserMod200JSONResponse{
		Message: ToPtr("Successfully updated user mod perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteUserFromMod implements the v1.ServerInterface.
func (a *API) DeleteUserFromMod(ctx context.Context, request DeleteUserFromModRequestObject) (DeleteUserFromModResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteUserFromMod403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.usermods.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.UserModParams{
			UserID: request.UserId,
			ModID:  request.Body.Mod,
		},
	); err != nil {
		if errors.Is(err, usermods.ErrNotFound) {
			return DeleteUserFromMod404JSONResponse{
				Message: ToPtr("Failed to find user or mod"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, usermods.ErrNotAssigned) {
			return DeleteUserFromMod412JSONResponse{
				Message: ToPtr("Mod is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteUserFromMod500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop user from mod"),
		}, nil
	}

	return DeleteUserFromMod200JSONResponse{
		Message: ToPtr("Successfully dropped user from mod"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertUser(record *model.User) User {
	result := User{
		Id:        ToPtr(record.ID),
		Username:  ToPtr(record.Username),
		Email:     ToPtr(record.Email),
		Fullname:  ToPtr(record.Fullname),
		Profile:   ToPtr(gravatarFor(record.Email)),
		Active:    ToPtr(record.Active),
		Admin:     ToPtr(record.Admin),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	auths := make([]UserAuth, 0)

	for _, auth := range record.Auths {
		auths = append(
			auths,
			a.convertUserAuth(auth),
		)
	}

	result.Auths = ToPtr(auths)

	return result
}

func (a *API) convertUserAuth(record *model.UserAuth) UserAuth {
	result := UserAuth{
		Provider:  ToPtr(record.Provider),
		Ref:       ToPtr(record.Ref),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertUserTeam(record *model.UserTeam) UserTeam {
	result := UserTeam{
		TeamId:    record.TeamID,
		Team:      ToPtr(a.convertTeam(record.Team)),
		UserId:    record.UserID,
		Perm:      ToPtr(UserTeamPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertUserPack(record *model.UserPack) UserPack {
	result := UserPack{
		PackId:    record.PackID,
		Pack:      ToPtr(a.convertPack(record.Pack)),
		UserId:    record.UserID,
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertUserMod(record *model.UserMod) UserMod {
	result := UserMod{
		ModId:     record.ModID,
		Mod:       ToPtr(a.convertMod(record.Mod)),
		UserId:    record.UserID,
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
