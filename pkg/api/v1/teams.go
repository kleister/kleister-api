package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	teammods "github.com/kleister/kleister-api/pkg/service/team_mods"
	teampacks "github.com/kleister/kleister-api/pkg/service/team_packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	userteams "github.com/kleister/kleister-api/pkg/service/user_teams"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListTeams implements the v1.ServerInterface.
func (a *API) ListTeams(ctx context.Context, request ListTeamsRequestObject) (ListTeamsResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListTeams403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	records, count, err := a.teams.WithPrincipal(
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
		return ListTeams500JSONResponse{
			Message: ToPtr("Failed to load teams"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Team, len(records))
	for id, record := range records {
		payload[id] = a.convertTeam(record)
	}

	return ListTeams200JSONResponse{
		Total: ToPtr(count),
		Teams: ToPtr(payload),
	}, nil
}

// ShowTeam implements the v1.ServerInterface.
func (a *API) ShowTeam(ctx context.Context, request ShowTeamRequestObject) (ShowTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ShowTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.teams.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if errors.Is(err, teams.ErrNotFound) {
			return ShowTeam404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowTeam500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowTeam200JSONResponse(
		a.convertTeam(record),
	), nil
}

// CreateTeam implements the v1.ServerInterface.
func (a *API) CreateTeam(ctx context.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return CreateTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record := &model.Team{}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if err := a.teams.WithPrincipal(
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

			return CreateTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreateTeam500JSONResponse{
			Message: ToPtr("Failed to create team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreateTeam200JSONResponse(
		a.convertTeam(record),
	), nil
}

// UpdateTeam implements the v1.ServerInterface.
func (a *API) UpdateTeam(ctx context.Context, request UpdateTeamRequestObject) (UpdateTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return UpdateTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.teams.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if errors.Is(err, teams.ErrNotFound) {
			return UpdateTeam404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdateTeam500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if err := a.teams.WithPrincipal(
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

			return UpdateTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateTeam500JSONResponse{
			Message: ToPtr("Failed to update team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateTeam200JSONResponse(
		a.convertTeam(record),
	), nil
}

// DeleteTeam implements the v1.ServerInterface.
func (a *API) DeleteTeam(ctx context.Context, request DeleteTeamRequestObject) (DeleteTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.teams.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if errors.Is(err, teams.ErrNotFound) {
			return DeleteTeam404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeleteTeam500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.teams.WithPrincipal(
		current.GetUser(ctx),
	).Delete(
		ctx,
		record.ID,
	); err != nil {
		return DeleteTeam400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete team"),
		}, nil
	}

	return DeleteTeam200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted team"),
	}, nil
}

// ListTeamUsers implements the v1.ServerInterface.
func (a *API) ListTeamUsers(ctx context.Context, request ListTeamUsersRequestObject) (ListTeamUsersResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListTeamUsers403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.teams.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if errors.Is(err, teams.ErrNotFound) {
			return ListTeamUsers404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListTeamUsers500JSONResponse{
			Message: ToPtr("Failed to load team"),
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
			TeamID: record.ID,
		},
	)

	if err != nil {
		return ListTeamUsers500JSONResponse{
			Message: ToPtr("Failed to load users"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserTeam, len(records))
	for id, record := range records {
		payload[id] = a.convertTeamUser(record)
	}

	return ListTeamUsers200JSONResponse{
		Total: ToPtr(count),
		Team:  ToPtr(a.convertTeam(record)),
		Users: ToPtr(payload),
	}, nil
}

// AttachTeamToUser implements the v1.ServerInterface.
func (a *API) AttachTeamToUser(ctx context.Context, request AttachTeamToUserRequestObject) (AttachTeamToUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachTeamToUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userteams.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.UserTeamParams{
			TeamID: request.TeamId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userteams.ErrNotFound) {
			return AttachTeamToUser404JSONResponse{
				Message: ToPtr("Failed to find team or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userteams.ErrAlreadyAssigned) {
			return AttachTeamToUser412JSONResponse{
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

			return AttachTeamToUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachTeamToUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach team to user"),
		}, nil
	}

	return AttachTeamToUser200JSONResponse{
		Message: ToPtr("Successfully attached team to user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitTeamUser implements the v1.ServerInterface.
func (a *API) PermitTeamUser(ctx context.Context, request PermitTeamUserRequestObject) (PermitTeamUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitTeamUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userteams.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.UserTeamParams{
			TeamID: request.TeamId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userteams.ErrNotFound) {
			return PermitTeamUser404JSONResponse{
				Message: ToPtr("Failed to find team or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userteams.ErrNotAssigned) {
			return PermitTeamUser412JSONResponse{
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

			return PermitTeamUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitTeamUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update team user perms"),
		}, nil
	}

	return PermitTeamUser200JSONResponse{
		Message: ToPtr("Successfully updated team user perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteTeamFromUser implements the v1.ServerInterface.
func (a *API) DeleteTeamFromUser(ctx context.Context, request DeleteTeamFromUserRequestObject) (DeleteTeamFromUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteTeamFromUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userteams.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.UserTeamParams{
			TeamID: request.TeamId,
			UserID: request.Body.User,
		},
	); err != nil {
		if errors.Is(err, userteams.ErrNotFound) {
			return DeleteTeamFromUser404JSONResponse{
				Message: ToPtr("Failed to find team or user"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, userteams.ErrNotAssigned) {
			return DeleteTeamFromUser412JSONResponse{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteTeamFromUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop team from user"),
		}, nil
	}

	return DeleteTeamFromUser200JSONResponse{
		Message: ToPtr("Successfully dropped team from user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListTeamPacks implements the v1.ServerInterface.
func (a *API) ListTeamPacks(ctx context.Context, request ListTeamPacksRequestObject) (ListTeamPacksResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListTeamPacks403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.teams.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if errors.Is(err, teams.ErrNotFound) {
			return ListTeamPacks404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListTeamPacks500JSONResponse{
			Message: ToPtr("Failed to load team"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	records, count, err := a.teampacks.WithPrincipal(
		current.GetUser(ctx),
	).List(
		ctx,
		model.TeamPackParams{
			ListParams: toListParams(
				string(FromPtr(request.Params.Sort)),
				string(FromPtr(request.Params.Order)),
				request.Params.Limit,
				request.Params.Offset,
				request.Params.Search,
			),
			TeamID: record.ID,
		},
	)

	if err != nil {
		return ListTeamPacks500JSONResponse{
			Message: ToPtr("Failed to load packs"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]TeamPack, len(records))
	for id, record := range records {
		payload[id] = a.convertTeamPack(record)
	}

	return ListTeamPacks200JSONResponse{
		Total: ToPtr(count),
		Team:  ToPtr(a.convertTeam(record)),
		Packs: ToPtr(payload),
	}, nil
}

// AttachTeamToPack implements the v1.ServerInterface.
func (a *API) AttachTeamToPack(ctx context.Context, request AttachTeamToPackRequestObject) (AttachTeamToPackResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachTeamToPack403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teampacks.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.TeamPackParams{
			TeamID: request.TeamId,
			PackID: request.Body.Pack,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teampacks.ErrNotFound) {
			return AttachTeamToPack404JSONResponse{
				Message: ToPtr("Failed to find team or pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teampacks.ErrAlreadyAssigned) {
			return AttachTeamToPack412JSONResponse{
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

			return AttachTeamToPack422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team pack"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachTeamToPack500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach team to pack"),
		}, nil
	}

	return AttachTeamToPack200JSONResponse{
		Message: ToPtr("Successfully attached team to pack"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitTeamPack implements the v1.ServerInterface.
func (a *API) PermitTeamPack(ctx context.Context, request PermitTeamPackRequestObject) (PermitTeamPackResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitTeamPack403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teampacks.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.TeamPackParams{
			TeamID: request.TeamId,
			PackID: request.Body.Pack,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teampacks.ErrNotFound) {
			return PermitTeamPack404JSONResponse{
				Message: ToPtr("Failed to find team or pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teampacks.ErrNotAssigned) {
			return PermitTeamPack412JSONResponse{
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

			return PermitTeamPack422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team pack"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitTeamPack500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update team pack perms"),
		}, nil
	}

	return PermitTeamPack200JSONResponse{
		Message: ToPtr("Successfully updated team pack perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteTeamFromPack implements the v1.ServerInterface.
func (a *API) DeleteTeamFromPack(ctx context.Context, request DeleteTeamFromPackRequestObject) (DeleteTeamFromPackResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteTeamFromPack403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teampacks.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.TeamPackParams{
			TeamID: request.TeamId,
			PackID: request.Body.Pack,
		},
	); err != nil {
		if errors.Is(err, teampacks.ErrNotFound) {
			return DeleteTeamFromPack404JSONResponse{
				Message: ToPtr("Failed to find team or pack"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, teampacks.ErrNotAssigned) {
			return DeleteTeamFromPack412JSONResponse{
				Message: ToPtr("Pack is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteTeamFromPack500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop team from pack"),
		}, nil
	}

	return DeleteTeamFromPack200JSONResponse{
		Message: ToPtr("Successfully dropped team from pack"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListTeamMods implements the v1.ServerInterface.
func (a *API) ListTeamMods(ctx context.Context, request ListTeamModsRequestObject) (ListTeamModsResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListTeamMods403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.teams.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.TeamId,
	)

	if err != nil {
		if errors.Is(err, teams.ErrNotFound) {
			return ListTeamMods404JSONResponse{
				Message: ToPtr("Failed to find team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListTeamMods500JSONResponse{
			Message: ToPtr("Failed to load team"),
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
			TeamID: record.ID,
		},
	)

	if err != nil {
		return ListTeamMods500JSONResponse{
			Message: ToPtr("Failed to load mods"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]TeamMod, len(records))
	for id, record := range records {
		payload[id] = a.convertTeamMod(record)
	}

	return ListTeamMods200JSONResponse{
		Total: ToPtr(count),
		Team:  ToPtr(a.convertTeam(record)),
		Mods:  ToPtr(payload),
	}, nil
}

// AttachTeamToMod implements the v1.ServerInterface.
func (a *API) AttachTeamToMod(ctx context.Context, request AttachTeamToModRequestObject) (AttachTeamToModResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachTeamToMod403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teammods.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.TeamModParams{
			TeamID: request.TeamId,
			ModID:  request.Body.Mod,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teammods.ErrNotFound) {
			return AttachTeamToMod404JSONResponse{
				Message: ToPtr("Failed to find team or mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teammods.ErrAlreadyAssigned) {
			return AttachTeamToMod412JSONResponse{
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

			return AttachTeamToMod422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team mod"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachTeamToMod500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach team to mod"),
		}, nil
	}

	return AttachTeamToMod200JSONResponse{
		Message: ToPtr("Successfully attached team to mod"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitTeamMod implements the v1.ServerInterface.
func (a *API) PermitTeamMod(ctx context.Context, request PermitTeamModRequestObject) (PermitTeamModResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitTeamMod403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teammods.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.TeamModParams{
			TeamID: request.TeamId,
			ModID:  request.Body.Mod,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teammods.ErrNotFound) {
			return PermitTeamMod404JSONResponse{
				Message: ToPtr("Failed to find team or mod"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teammods.ErrNotAssigned) {
			return PermitTeamMod412JSONResponse{
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

			return PermitTeamMod422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate team mod"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitTeamMod500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update team mod perms"),
		}, nil
	}

	return PermitTeamMod200JSONResponse{
		Message: ToPtr("Successfully updated team mod perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeleteTeamFromMod implements the v1.ServerInterface.
func (a *API) DeleteTeamFromMod(ctx context.Context, request DeleteTeamFromModRequestObject) (DeleteTeamFromModResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeleteTeamFromMod403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teammods.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.TeamModParams{
			TeamID: request.TeamId,
			ModID:  request.Body.Mod,
		},
	); err != nil {
		if errors.Is(err, teammods.ErrNotFound) {
			return DeleteTeamFromMod404JSONResponse{
				Message: ToPtr("Failed to find team or mod"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, teammods.ErrNotAssigned) {
			return DeleteTeamFromMod412JSONResponse{
				Message: ToPtr("Mod is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeleteTeamFromMod500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop team from mod"),
		}, nil
	}

	return DeleteTeamFromMod200JSONResponse{
		Message: ToPtr("Successfully dropped team from mod"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertTeam(record *model.Team) Team {
	result := Team{
		Id:        ToPtr(record.ID),
		Slug:      ToPtr(record.Slug),
		Name:      ToPtr(record.Name),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertTeamUser(record *model.UserTeam) UserTeam {
	result := UserTeam{
		UserId:    record.UserID,
		User:      ToPtr(a.convertUser(record.User)),
		TeamId:    record.TeamID,
		Perm:      ToPtr(UserTeamPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertTeamPack(record *model.TeamPack) TeamPack {
	result := TeamPack{
		TeamId:    record.TeamID,
		Team:      ToPtr(a.convertTeam(record.Team)),
		PackId:    record.PackID,
		Perm:      ToPtr(TeamPackPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertTeamMod(record *model.TeamMod) TeamMod {
	result := TeamMod{
		TeamId:    record.TeamID,
		Team:      ToPtr(a.convertTeam(record.Team)),
		ModId:     record.ModID,
		Perm:      ToPtr(TeamModPerm(record.Perm)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
