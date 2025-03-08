package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/kleister/kleister-api/pkg/middleware/current"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/token"
	"github.com/kleister/kleister-api/pkg/validate"
	"github.com/rs/zerolog/log"
)

// TokenProfile implements the v1.ServerInterface.
func (a *API) TokenProfile(w http.ResponseWriter, r *http.Request) {
	principal := current.GetUser(
		r.Context(),
	)

	result, err := token.Authed(
		a.config.Token.Secret,
		10*365*24*time.Hour,
		principal.ID,
		principal.Username,
		principal.Email,
		principal.Fullname,
		principal.Admin,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("action", "TokenProfile").
			Str("username", principal.Username).
			Str("uid", principal.ID).
			Msg("Failed to generate a token")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to generate a token"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, TokenResponse(
		a.convertAuthToken(result),
	))
}

// ShowProfile implements the v1.ServerInterface.
func (a *API) ShowProfile(w http.ResponseWriter, r *http.Request) {
	record := current.GetUser(
		r.Context(),
	)

	render.JSON(w, r, ProfileResponse(
		a.convertProfile(
			record,
		),
	))
}

// UpdateProfile implements the v1.ServerInterface.
func (a *API) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	body := &UpdateProfileBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Error().
			Err(err).
			Str("action", "UpdateProfile").
			Msg("Failed to decode request body")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to decode request"),
			Status:  ToPtr(http.StatusBadRequest),
		})

		return
	}

	record := current.GetUser(
		r.Context(),
	)

	if body.Username != nil {
		record.Username = FromPtr(body.Username)
	}

	if body.Password != nil {
		record.Password = FromPtr(body.Password)
	}

	if body.Email != nil {
		record.Email = FromPtr(body.Email)
	}

	if body.Fullname != nil {
		record.Fullname = FromPtr(body.Fullname)
	}

	if err := a.storage.Users.Update(
		r.Context(),
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
				Message: ToPtr("Failed to validate profile"),
				Status:  ToPtr(http.StatusUnprocessableEntity),
				Errors:  ToPtr(errors),
			})

			return
		}

		log.Error().
			Err(err).
			Str("user", record.ID).
			Str("action", "UpdateProfile").
			Msg("Failed to update profile")

		a.RenderNotify(w, r, Notification{
			Message: ToPtr("Failed to update profile"),
			Status:  ToPtr(http.StatusInternalServerError),
		})

		return
	}

	render.JSON(w, r, ProfileResponse(
		a.convertProfile(
			record,
		),
	))
}

func (a *API) convertProfile(record *model.User) Profile {
	result := Profile{
		ID:        ToPtr(record.ID),
		Username:  ToPtr(record.Username),
		Email:     ToPtr(record.Email),
		Fullname:  ToPtr(record.Fullname),
		Profile:   ToPtr(gravatarFor(record.Email)),
		Active:    ToPtr(record.Active),
		Admin:     ToPtr(record.Admin),
		CreatedAt: ToPtr(record.CreatedAt),
		UpdatedAt: ToPtr(record.UpdatedAt),
	}

	if len(record.Auths) > 0 {
		auths := make([]UserAuth, 0)

		for _, auth := range record.Auths {
			auths = append(
				auths,
				a.convertUserAuth(auth),
			)
		}

		result.Auths = ToPtr(auths)
	}

	if len(record.Groups) > 0 {
		groups := make([]UserGroup, 0)

		for _, group := range record.Groups {
			groups = append(
				groups,
				a.convertUserGroup(group),
			)
		}

		result.Groups = ToPtr(groups)
	}

	if len(record.Mods) > 0 {
		mods := make([]UserMod, 0)

		for _, mod := range record.Mods {
			mods = append(
				mods,
				a.convertUserMod(mod),
			)
		}

		result.Mods = ToPtr(mods)
	}

	if len(record.Packs) > 0 {
		packs := make([]UserPack, 0)

		for _, pack := range record.Packs {
			packs = append(
				packs,
				a.convertUserPack(pack),
			)
		}

		result.Packs = ToPtr(packs)
	}

	return result
}
