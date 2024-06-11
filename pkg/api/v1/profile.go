package v1

import (
	"context"
	"net/http"

	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/token"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
)

// TokenProfile implements the v1.ServerInterface.
func (a *API) TokenProfile(ctx context.Context, _ TokenProfileRequestObject) (TokenProfileResponseObject, error) {
	principal := current.GetUser(
		ctx,
	)

	result, err := token.New(
		principal.Username,
	).Unlimited(
		a.config.Session.Secret,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("username", principal.Username).
			Msg("Failed to generate a token")

		return TokenProfile500JSONResponse{
			Message: ToPtr("Failed to generate a token"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return TokenProfile200JSONResponse(
		a.convertAuthToken(result),
	), nil
}

// ShowProfile implements the v1.ServerInterface.
func (a *API) ShowProfile(ctx context.Context, _ ShowProfileRequestObject) (ShowProfileResponseObject, error) {
	record := current.GetUser(
		ctx,
	)

	return ShowProfile200JSONResponse(
		a.convertProfile(record),
	), nil
}

// UpdateProfile implements the v1.ServerInterface.
func (a *API) UpdateProfile(ctx context.Context, request UpdateProfileRequestObject) (UpdateProfileResponseObject, error) {
	record := current.GetUser(
		ctx,
	)

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

	if err := a.users.Update(
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

			return UpdateProfile422JSONResponse{
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Message: ToPtr("Failed to validate profile"),
				Errors:  ToPtr(errors),
			}, nil
		}

		return UpdateProfile500JSONResponse{
			Message: ToPtr("Failed to update profile"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	return UpdateProfile200JSONResponse(
		a.convertProfile(record),
	), nil
}

func (a *API) convertProfile(record *model.User) Profile {
	result := Profile{
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

	teams := make([]UserTeam, 0)

	for _, team := range record.Teams {
		teams = append(
			teams,
			a.convertUserTeam(team),
		)
	}

	result.Teams = ToPtr(teams)

	packs := make([]UserPack, 0)

	for _, pack := range record.Packs {
		packs = append(
			packs,
			a.convertUserPack(pack),
		)
	}

	result.Packs = ToPtr(packs)

	mods := make([]UserMod, 0)

	for _, mod := range record.Mods {
		mods = append(
			mods,
			a.convertUserMod(mod),
		)
	}

	result.Mods = ToPtr(mods)

	return result
}
