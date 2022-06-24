package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/team"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListTeamsHandler implements the handler for the ListTeams operation.
func ListTeamsHandler(teamsService teams.Service) team.ListTeamsHandlerFunc {
	return func(params team.ListTeamsParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewListTeamsForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := teamsService.List(params.HTTPRequest.Context())

		if err != nil {
			return team.NewListTeamsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Team, len(records))
		for id, record := range records {
			payload[id] = convertTeam(record)
		}

		return team.NewListTeamsOK().WithPayload(payload)
	}
}

// ShowTeamHandler implements the handler for the ShowTeam operation.
func ShowTeamHandler(teamsService teams.Service) team.ShowTeamHandlerFunc {
	return func(params team.ShowTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewShowTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewShowTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewShowTeamDefault(http.StatusInternalServerError)
		}

		return team.NewShowTeamOK().WithPayload(convertTeam(record))
	}
}

// CreateTeamHandler implements the handler for the CreateTeam operation.
func CreateTeamHandler(teamsService teams.Service) team.CreateTeamHandlerFunc {
	return func(params team.CreateTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewCreateTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record := &model.Team{}

		if params.Team.Slug != nil {
			record.Slug = *params.Team.Slug
		}

		if params.Team.Name != nil {
			record.Name = *params.Team.Name
		}

		created, err := teamsService.Create(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewCreateTeamUnprocessableEntity().WithPayload(payload)
			}

			return team.NewCreateTeamDefault(http.StatusInternalServerError)
		}

		return team.NewCreateTeamOK().WithPayload(convertTeam(created))
	}
}

// UpdateTeamHandler implements the handler for the UpdateTeam operation.
func UpdateTeamHandler(teamsService teams.Service) team.UpdateTeamHandlerFunc {
	return func(params team.UpdateTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewUpdateTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewUpdateTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewUpdateTeamDefault(http.StatusInternalServerError)
		}

		if params.Team.Slug != nil {
			record.Slug = *params.Team.Slug
		}

		if params.Team.Name != nil {
			record.Name = *params.Team.Name
		}

		updated, err := teamsService.Update(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewUpdateTeamUnprocessableEntity().WithPayload(payload)
			}

			return team.NewUpdateTeamDefault(http.StatusInternalServerError)
		}

		return team.NewUpdateTeamOK().WithPayload(convertTeam(updated))
	}
}

// DeleteTeamHandler implements the handler for the DeleteTeam operation.
func DeleteTeamHandler(teamsService teams.Service) team.DeleteTeamHandlerFunc {
	return func(params team.DeleteTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewDeleteTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewDeleteTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamDefault(http.StatusInternalServerError)
		}

		if err := teamsService.Delete(params.HTTPRequest.Context(), record.ID); err != nil {
			message := "failed to delete team"

			return team.NewDeleteTeamBadRequest().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		message := "successfully deleted team"
		return team.NewDeleteTeamOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// ListTeamUsersHandler implements the handler for the ListTeamUsers operation.
func ListTeamUsersHandler(teamsService teams.Service) team.ListTeamUsersHandlerFunc {
	return func(params team.ListTeamUsersParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewListTeamUsersForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := teamsService.ListUsers(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			// TODO: add handler if team not found
			return team.NewListTeamUsersDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.TeamUser, len(records))
		for id, record := range records {
			payload[id] = convertTeamUser(record)
		}

		return team.NewListTeamUsersOK().WithPayload(payload)
	}
}

// AppendTeamToUserHandler implements the handler for the AppendTeamToUser operation.
func AppendTeamToUserHandler(teamsService teams.Service, usersService users.Service) team.AppendTeamToUserHandlerFunc {
	return func(params team.AppendTeamToUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewAppendTeamToUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewAppendTeamToUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToUserDefault(http.StatusInternalServerError)
		}

		u, err := usersService.Show(params.HTTPRequest.Context(), *params.TeamUser.User)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return team.NewAppendTeamToUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToUserDefault(http.StatusInternalServerError)
		}

		if err := teamsService.AppendUser(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamUser.Perm); err != nil {
			if err == teams.ErrAlreadyAssigned {
				message := "user is already assigned"

				return team.NewAppendTeamToUserPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team user"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewAppendTeamToUserUnprocessableEntity().WithPayload(payload)
			}

			return team.NewAppendTeamToUserDefault(http.StatusInternalServerError)
		}

		message := "successfully assigned team to user"
		return team.NewAppendTeamToUserOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// PermitTeamUserHandler implements the handler for the PermitTeamUser operation.
func PermitTeamUserHandler(teamsService teams.Service, usersService users.Service) team.PermitTeamUserHandlerFunc {
	return func(params team.PermitTeamUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewPermitTeamUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewPermitTeamUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamUserDefault(http.StatusInternalServerError)
		}

		u, err := usersService.Show(params.HTTPRequest.Context(), *params.TeamUser.User)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return team.NewPermitTeamUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamUserDefault(http.StatusInternalServerError)
		}

		if err := teamsService.PermitUser(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamUser.Perm); err != nil {
			if err == teams.ErrNotAssigned {
				message := "user is not assigned"

				return team.NewPermitTeamUserPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team user"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewPermitTeamUserUnprocessableEntity().WithPayload(payload)
			}

			return team.NewPermitTeamUserDefault(http.StatusInternalServerError)
		}

		message := "successfully updated user perms"
		return team.NewPermitTeamUserOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// DeleteTeamFromUserHandler implements the handler for the DeleteTeamFromUser operation.
func DeleteTeamFromUserHandler(teamsService teams.Service, usersService users.Service) team.DeleteTeamFromUserHandlerFunc {
	return func(params team.DeleteTeamFromUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewDeleteTeamFromUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewDeleteTeamFromUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromUserDefault(http.StatusInternalServerError)
		}

		u, err := usersService.Show(params.HTTPRequest.Context(), *params.TeamUser.User)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return team.NewDeleteTeamFromUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromUserDefault(http.StatusInternalServerError)
		}

		if err := teamsService.DropUser(params.HTTPRequest.Context(), t.ID, u.ID); err != nil {
			if err == teams.ErrNotAssigned {
				message := "user is not assigned"

				return team.NewDeleteTeamFromUserPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromUserDefault(http.StatusInternalServerError)
		}

		message := "successfully removed from user"
		return team.NewDeleteTeamFromUserOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// ListTeamModsHandler implements the handler for the ListTeamMods operation.
func ListTeamModsHandler(teamsService teams.Service) team.ListTeamModsHandlerFunc {
	return func(params team.ListTeamModsParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewListTeamModsForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := teamsService.ListMods(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			// TODO: add handler if team not found
			return team.NewListTeamModsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.TeamMod, len(records))
		for id, record := range records {
			payload[id] = convertTeamMod(record)
		}

		return team.NewListTeamModsOK().WithPayload(payload)
	}
}

// AppendTeamToModHandler implements the handler for the AppendTeamToMod operation.
func AppendTeamToModHandler(teamsService teams.Service, modsService mods.Service) team.AppendTeamToModHandlerFunc {
	return func(params team.AppendTeamToModParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewAppendTeamToModForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewAppendTeamToModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToModDefault(http.StatusInternalServerError)
		}

		u, err := modsService.Show(params.HTTPRequest.Context(), *params.TeamMod.Mod)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return team.NewAppendTeamToModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToModDefault(http.StatusInternalServerError)
		}

		if err := teamsService.AppendMod(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamMod.Perm); err != nil {
			if err == teams.ErrAlreadyAssigned {
				message := "mod is already assigned"

				return team.NewAppendTeamToModPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team mod"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewAppendTeamToModUnprocessableEntity().WithPayload(payload)
			}

			return team.NewAppendTeamToModDefault(http.StatusInternalServerError)
		}

		message := "successfully assigned team to mod"
		return team.NewAppendTeamToModOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// PermitTeamModHandler implements the handler for the PermitTeamMod operation.
func PermitTeamModHandler(teamsService teams.Service, modsService mods.Service) team.PermitTeamModHandlerFunc {
	return func(params team.PermitTeamModParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewPermitTeamModForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewPermitTeamModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamModDefault(http.StatusInternalServerError)
		}

		u, err := modsService.Show(params.HTTPRequest.Context(), *params.TeamMod.Mod)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return team.NewPermitTeamModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamModDefault(http.StatusInternalServerError)
		}

		if err := teamsService.PermitMod(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamMod.Perm); err != nil {
			if err == teams.ErrNotAssigned {
				message := "mod is not assigned"

				return team.NewPermitTeamModPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team mod"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewPermitTeamModUnprocessableEntity().WithPayload(payload)
			}

			return team.NewPermitTeamModDefault(http.StatusInternalServerError)
		}

		message := "successfully updated mod perms"
		return team.NewPermitTeamModOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// DeleteTeamFromModHandler implements the handler for the DeleteTeamFromMod operation.
func DeleteTeamFromModHandler(teamsService teams.Service, modsService mods.Service) team.DeleteTeamFromModHandlerFunc {
	return func(params team.DeleteTeamFromModParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewDeleteTeamFromModForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewDeleteTeamFromModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromModDefault(http.StatusInternalServerError)
		}

		u, err := modsService.Show(params.HTTPRequest.Context(), *params.TeamMod.Mod)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return team.NewDeleteTeamFromModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromModDefault(http.StatusInternalServerError)
		}

		if err := teamsService.DropMod(params.HTTPRequest.Context(), t.ID, u.ID); err != nil {
			if err == teams.ErrNotAssigned {
				message := "mod is not assigned"

				return team.NewDeleteTeamFromModPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromModDefault(http.StatusInternalServerError)
		}

		message := "successfully removed from mod"
		return team.NewDeleteTeamFromModOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// ListTeamPacksHandler implements the handler for the ListTeamPacks operation.
func ListTeamPacksHandler(teamsService teams.Service) team.ListTeamPacksHandlerFunc {
	return func(params team.ListTeamPacksParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewListTeamPacksForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := teamsService.ListPacks(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			// TODO: add handler if team not found
			return team.NewListTeamPacksDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.TeamPack, len(records))
		for id, record := range records {
			payload[id] = convertTeamPack(record)
		}

		return team.NewListTeamPacksOK().WithPayload(payload)
	}
}

// AppendTeamToPackHandler implements the handler for the AppendTeamToPack operation.
func AppendTeamToPackHandler(teamsService teams.Service, packsService packs.Service) team.AppendTeamToPackHandlerFunc {
	return func(params team.AppendTeamToPackParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewAppendTeamToPackForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewAppendTeamToPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToPackDefault(http.StatusInternalServerError)
		}

		u, err := packsService.Show(params.HTTPRequest.Context(), *params.TeamPack.Pack)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return team.NewAppendTeamToPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToPackDefault(http.StatusInternalServerError)
		}

		if err := teamsService.AppendPack(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamPack.Perm); err != nil {
			if err == teams.ErrAlreadyAssigned {
				message := "pack is already assigned"

				return team.NewAppendTeamToPackPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team pack"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewAppendTeamToPackUnprocessableEntity().WithPayload(payload)
			}

			return team.NewAppendTeamToPackDefault(http.StatusInternalServerError)
		}

		message := "successfully assigned team to pack"
		return team.NewAppendTeamToPackOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// PermitTeamPackHandler implements the handler for the PermitTeamPack operation.
func PermitTeamPackHandler(teamsService teams.Service, packsService packs.Service) team.PermitTeamPackHandlerFunc {
	return func(params team.PermitTeamPackParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewPermitTeamPackForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewPermitTeamPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamPackDefault(http.StatusInternalServerError)
		}

		u, err := packsService.Show(params.HTTPRequest.Context(), *params.TeamPack.Pack)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return team.NewPermitTeamPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamPackDefault(http.StatusInternalServerError)
		}

		if err := teamsService.PermitPack(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamPack.Perm); err != nil {
			if err == teams.ErrNotAssigned {
				message := "pack is not assigned"

				return team.NewPermitTeamPackPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate team pack"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewPermitTeamPackUnprocessableEntity().WithPayload(payload)
			}

			return team.NewPermitTeamPackDefault(http.StatusInternalServerError)
		}

		message := "successfully updated pack perms"
		return team.NewPermitTeamPackOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// DeleteTeamFromPackHandler implements the handler for the DeleteTeamFromPack operation.
func DeleteTeamFromPackHandler(teamsService teams.Service, packsService packs.Service) team.DeleteTeamFromPackHandlerFunc {
	return func(params team.DeleteTeamFromPackParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewDeleteTeamFromPackForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewDeleteTeamFromPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromPackDefault(http.StatusInternalServerError)
		}

		u, err := packsService.Show(params.HTTPRequest.Context(), *params.TeamPack.Pack)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return team.NewDeleteTeamFromPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromPackDefault(http.StatusInternalServerError)
		}

		if err := teamsService.DropPack(params.HTTPRequest.Context(), t.ID, u.ID); err != nil {
			if err == teams.ErrNotAssigned {
				message := "pack is not assigned"

				return team.NewDeleteTeamFromPackPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromPackDefault(http.StatusInternalServerError)
		}

		message := "successfully removed from pack"
		return team.NewDeleteTeamFromPackOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// convertTeam is a simple helper to convert between different model formats.
func convertTeam(record *model.Team) *models.Team {
	users := make([]*models.TeamUser, 0)

	for _, user := range record.Users {
		users = append(users, convertTeamUser(user))
	}

	mods := make([]*models.TeamMod, 0)

	for _, mod := range record.Mods {
		mods = append(mods, convertTeamMod(mod))
	}

	packs := make([]*models.TeamPack, 0)

	for _, pack := range record.Packs {
		packs = append(packs, convertTeamPack(pack))
	}

	return &models.Team{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Name:      &record.Name,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
		Users:     users,
		Mods:      mods,
		Packs:     packs,
	}
}

// convertTeamUser is a simple helper to convert between different model formats.
func convertTeamUser(record *model.TeamUser) *models.TeamUser {
	userID := strfmt.UUID(record.UserID)
	teamID := strfmt.UUID(record.TeamID)

	return &models.TeamUser{
		TeamID:    &teamID,
		Team:      convertTeam(record.Team),
		UserID:    &userID,
		User:      convertUser(record.User),
		Perm:      &record.Perm,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}

// convertTeamPack is a simple helper to convert between different model formats.
func convertTeamPack(record *model.TeamPack) *models.TeamPack {
	teamID := strfmt.UUID(record.TeamID)
	packID := strfmt.UUID(record.PackID)

	return &models.TeamPack{
		TeamID:    &teamID,
		Team:      convertTeam(record.Team),
		PackID:    &packID,
		Pack:      convertPack(record.Pack),
		Perm:      &record.Perm,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}

// convertTeamMod is a simple helper to convert between different model formats.
func convertTeamMod(record *model.TeamMod) *models.TeamMod {
	teamID := strfmt.UUID(record.TeamID)
	modID := strfmt.UUID(record.ModID)

	return &models.TeamMod{
		TeamID:    &teamID,
		Team:      convertTeam(record.Team),
		ModID:     &modID,
		Mod:       convertMod(record.Mod),
		Perm:      &record.Perm,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}
