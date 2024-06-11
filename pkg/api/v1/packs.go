package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/packs"
	teampacks "github.com/kleister/kleister-api/pkg/service/team_packs"
	userpacks "github.com/kleister/kleister-api/pkg/service/user_packs"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListPacks implements the v1.ServerInterface.
func (a *API) ListPacks(ctx context.Context, request ListPacksRequestObject) (ListPacksResponseObject, error) {
	records, count, err := a.packs.WithPrincipal(
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
		return ListPacks500JSONResponse{
			Message: ToPtr("Failed to load packs"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]Pack, len(records))
	for id, record := range records {
		payload[id] = a.convertPack(record)
	}

	return ListPacks200JSONResponse{
		Total: ToPtr(count),
		Packs: ToPtr(payload),
	}, nil
}

// ShowPack implements the v1.ServerInterface.
func (a *API) ShowPack(ctx context.Context, request ShowPackRequestObject) (ShowPackResponseObject, error) {
	record, err := a.packs.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.PackId,
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return ShowPack404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ShowPack500JSONResponse{
			Message: ToPtr("Failed to load pack"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return ShowPack200JSONResponse(
		a.convertPack(record),
	), nil
}

// CreatePack implements the v1.ServerInterface.
func (a *API) CreatePack(ctx context.Context, request CreatePackRequestObject) (CreatePackResponseObject, error) {
	record := &model.Pack{}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Website != nil {
		record.Website = FromPtr(request.Body.Website)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	// TODO: Back
	// TODO: Icon
	// TODO: Logo

	if err := a.packs.WithPrincipal(
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

			return CreatePack422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate pack"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return CreatePack500JSONResponse{
			Message: ToPtr("Failed to create pack"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return CreatePack200JSONResponse(
		a.convertPack(record),
	), nil
}

// UpdatePack implements the v1.ServerInterface.
func (a *API) UpdatePack(ctx context.Context, request UpdatePackRequestObject) (UpdatePackResponseObject, error) {
	record, err := a.packs.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.PackId,
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return UpdatePack404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return UpdatePack500JSONResponse{
			Message: ToPtr("Failed to load pack"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if request.Body.Slug != nil {
		record.Slug = FromPtr(request.Body.Slug)
	}

	if request.Body.Name != nil {
		record.Name = FromPtr(request.Body.Name)
	}

	if request.Body.Website != nil {
		record.Website = FromPtr(request.Body.Website)
	}

	if request.Body.Public != nil {
		record.Public = FromPtr(request.Body.Public)
	}

	// TODO: Back
	// TODO: Icon
	// TODO: Logo

	if err := a.packs.WithPrincipal(
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

			return UpdatePack422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate pack"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdatePack500JSONResponse{
			Message: ToPtr("Failed to update pack"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdatePack200JSONResponse(
		a.convertPack(record),
	), nil
}

// DeletePack implements the v1.ServerInterface.
func (a *API) DeletePack(ctx context.Context, request DeletePackRequestObject) (DeletePackResponseObject, error) {
	record, err := a.packs.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.PackId,
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return DeletePack404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return DeletePack500JSONResponse{
			Message: ToPtr("Failed to load pack"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	if err := a.packs.WithPrincipal(
		current.GetUser(ctx),
	).Delete(
		ctx,
		record.ID,
	); err != nil {
		return DeletePack400JSONResponse{
			Status:  ToPtr(http.StatusBadRequest),
			Message: ToPtr("Failed to delete pack"),
		}, nil
	}

	return DeletePack200JSONResponse{
		Status:  ToPtr(http.StatusOK),
		Message: ToPtr("Successfully deleted pack"),
	}, nil
}

// ListPackTeams implements the v1.ServerInterface.
func (a *API) ListPackTeams(ctx context.Context, request ListPackTeamsRequestObject) (ListPackTeamsResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListPackTeams403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.packs.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.PackId,
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return ListPackTeams404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListPackTeams500JSONResponse{
			Message: ToPtr("Failed to load pack"),
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
			PackID: record.ID,
		},
	)

	if err != nil {
		return ListPackTeams500JSONResponse{
			Message: ToPtr("Failed to load teams"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]TeamPack, len(records))
	for id, record := range records {
		payload[id] = a.convertPackTeam(record)
	}

	return ListPackTeams200JSONResponse{
		Total: ToPtr(count),
		Pack:  ToPtr(a.convertPack(record)),
		Teams: ToPtr(payload),
	}, nil
}

// AttachPackToTeam implements the v1.ServerInterface.
func (a *API) AttachPackToTeam(ctx context.Context, request AttachPackToTeamRequestObject) (AttachPackToTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachPackToTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teampacks.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.TeamPackParams{
			PackID: request.PackId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teampacks.ErrNotFound) {
			return AttachPackToTeam404JSONResponse{
				Message: ToPtr("Failed to find pack or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teampacks.ErrAlreadyAssigned) {
			return AttachPackToTeam412JSONResponse{
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

			return AttachPackToTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate pack team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachPackToTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach pack to team"),
		}, nil
	}

	return AttachPackToTeam200JSONResponse{
		Message: ToPtr("Successfully attached pack to team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitPackTeam implements the v1.ServerInterface.
func (a *API) PermitPackTeam(ctx context.Context, request PermitPackTeamRequestObject) (PermitPackTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitPackTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teampacks.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.TeamPackParams{
			PackID: request.PackId,
			TeamID: request.Body.Team,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, teampacks.ErrNotFound) {
			return PermitPackTeam404JSONResponse{
				Message: ToPtr("Failed to find pack or team"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, teampacks.ErrNotAssigned) {
			return PermitPackTeam412JSONResponse{
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

			return PermitPackTeam422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate pack team"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitPackTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update pack team perms"),
		}, nil
	}

	return PermitPackTeam200JSONResponse{
		Message: ToPtr("Successfully updated pack team perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeletePackFromTeam implements the v1.ServerInterface.
func (a *API) DeletePackFromTeam(ctx context.Context, request DeletePackFromTeamRequestObject) (DeletePackFromTeamResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeletePackFromTeam403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.teampacks.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.TeamPackParams{
			PackID: request.PackId,
			TeamID: request.Body.Team,
		},
	); err != nil {
		if errors.Is(err, teampacks.ErrNotFound) {
			return DeletePackFromTeam404JSONResponse{
				Message: ToPtr("Failed to find pack or team"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, teampacks.ErrNotAssigned) {
			return DeletePackFromTeam412JSONResponse{
				Message: ToPtr("Team is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeletePackFromTeam500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop pack from team"),
		}, nil
	}

	return DeletePackFromTeam200JSONResponse{
		Message: ToPtr("Successfully dropped pack from team"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// ListPackUsers implements the v1.ServerInterface.
func (a *API) ListPackUsers(ctx context.Context, request ListPackUsersRequestObject) (ListPackUsersResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return ListPackUsers403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	record, err := a.packs.WithPrincipal(
		current.GetUser(ctx),
	).Show(
		ctx,
		request.PackId,
	)

	if err != nil {
		if errors.Is(err, packs.ErrNotFound) {
			return ListPackUsers404JSONResponse{
				Message: ToPtr("Failed to find pack"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		return ListPackUsers500JSONResponse{
			Message: ToPtr("Failed to load pack"),
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
			PackID: record.ID,
		},
	)

	if err != nil {
		return ListPackUsers500JSONResponse{
			Message: ToPtr("Failed to load users"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	payload := make([]UserPack, len(records))
	for id, record := range records {
		payload[id] = a.convertPackUser(record)
	}

	return ListPackUsers200JSONResponse{
		Total: ToPtr(count),
		Pack:  ToPtr(a.convertPack(record)),
		Users: ToPtr(payload),
	}, nil
}

// AttachPackToUser implements the v1.ServerInterface.
func (a *API) AttachPackToUser(ctx context.Context, request AttachPackToUserRequestObject) (AttachPackToUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return AttachPackToUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userpacks.WithPrincipal(
		current.GetUser(ctx),
	).Attach(
		ctx,
		model.UserPackParams{
			PackID: request.PackId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userpacks.ErrNotFound) {
			return AttachPackToUser404JSONResponse{
				Message: ToPtr("Failed to find pack or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userpacks.ErrAlreadyAssigned) {
			return AttachPackToUser412JSONResponse{
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

			return AttachPackToUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate pack user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return AttachPackToUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to attach pack to user"),
		}, nil
	}

	return AttachPackToUser200JSONResponse{
		Message: ToPtr("Successfully attached pack to user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// PermitPackUser implements the v1.ServerInterface.
func (a *API) PermitPackUser(ctx context.Context, request PermitPackUserRequestObject) (PermitPackUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return PermitPackUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userpacks.WithPrincipal(
		current.GetUser(ctx),
	).Permit(
		ctx,
		model.UserPackParams{
			PackID: request.PackId,
			UserID: request.Body.User,
			Perm:   string(FromPtr(request.Body.Perm)),
		},
	); err != nil {
		if errors.Is(err, userpacks.ErrNotFound) {
			return PermitPackUser404JSONResponse{
				Message: ToPtr("Failed to find pack or user"),
				Status:  ToPtr(http.StatusNotFound),
			}, nil
		}

		if errors.Is(err, userpacks.ErrNotAssigned) {
			return PermitPackUser412JSONResponse{
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

			return PermitPackUser422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate pack user"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return PermitPackUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to update pack user perms"),
		}, nil
	}

	return PermitPackUser200JSONResponse{
		Message: ToPtr("Successfully updated pack user perms"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

// DeletePackFromUser implements the v1.ServerInterface.
func (a *API) DeletePackFromUser(ctx context.Context, request DeletePackFromUserRequestObject) (DeletePackFromUserResponseObject, error) {
	if principal := current.GetUser(ctx); principal == nil || !principal.Admin {
		return DeletePackFromUser403JSONResponse{
			Message: ToPtr("Only admins can access this resource"),
			Status:  ToPtr(http.StatusForbidden),
		}, nil
	}

	if err := a.userpacks.WithPrincipal(
		current.GetUser(ctx),
	).Drop(
		ctx,
		model.UserPackParams{
			PackID: request.PackId,
			UserID: request.Body.User,
		},
	); err != nil {
		if errors.Is(err, userpacks.ErrNotFound) {
			return DeletePackFromUser404JSONResponse{
				Message: ToPtr("Failed to find pack or user"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		if errors.Is(err, userpacks.ErrNotAssigned) {
			return DeletePackFromUser412JSONResponse{
				Message: ToPtr("User is not attached"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		return DeletePackFromUser500JSONResponse{
			Status:  ToPtr(http.StatusUnprocessableEntity),
			Message: ToPtr("Failed to drop pack from user"),
		}, nil
	}

	return DeletePackFromUser200JSONResponse{
		Message: ToPtr("Successfully dropped pack from user"),
		Status:  ToPtr(http.StatusOK),
	}, nil
}

func (a *API) convertPack(record *model.Pack) Pack {
	result := Pack{
		Id:        ToPtr(record.ID),
		Slug:      ToPtr(record.Slug),
		Name:      ToPtr(record.Name),
		Website:   ToPtr(record.Website),
		Public:    ToPtr(record.Public),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	// TODO: Back
	// TODO: Icon
	// TODO: Logo

	return result
}

func (a *API) convertPackTeam(record *model.TeamPack) TeamPack {
	result := TeamPack{
		PackId:    record.PackID,
		TeamId:    record.TeamID,
		Team:      ToPtr(a.convertTeam(record.Team)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}

func (a *API) convertPackUser(record *model.UserPack) UserPack {
	result := UserPack{
		PackId:    record.PackID,
		UserId:    record.UserID,
		User:      ToPtr(a.convertUser(record.User)),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	return result
}
