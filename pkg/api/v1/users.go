package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/user"
	"github.com/kleister/kleister-api/pkg/model"
	"github.com/kleister/kleister-api/pkg/service/mods"
	"github.com/kleister/kleister-api/pkg/service/packs"
	"github.com/kleister/kleister-api/pkg/service/teams"
	"github.com/kleister/kleister-api/pkg/service/users"
	"github.com/kleister/kleister-api/pkg/validate"
)

// ListUsersHandler implements the handler for the ListUsers operation.
func ListUsersHandler(usersService users.Service) user.ListUsersHandlerFunc {
	return func(params user.ListUsersParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewListUsersForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := usersService.List(params.HTTPRequest.Context())

		if err != nil {
			return user.NewListUsersDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.User, len(records))
		for id, record := range records {
			payload[id] = convertUser(record)
		}

		return user.NewListUsersOK().WithPayload(payload)
	}
}

// ShowUserHandler implements the handler for the ShowUser operation.
func ShowUserHandler(usersService users.Service) user.ShowUserHandlerFunc {
	return func(params user.ShowUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewShowUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewShowUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewShowUserDefault(http.StatusInternalServerError)
		}

		return user.NewShowUserOK().WithPayload(convertUser(record))
	}
}

// CreateUserHandler implements the handler for the CreateUser operation.
func CreateUserHandler(usersService users.Service) user.CreateUserHandlerFunc {
	return func(params user.CreateUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewCreateUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record := &model.User{}

		if params.User.Slug != nil {
			record.Slug = *params.User.Slug
		}

		if params.User.Username != nil {
			record.Username = *params.User.Username
		}

		if params.User.Password != nil {
			record.Password = (*params.User.Password).String()
		}

		if params.User.Email != nil {
			record.Email = *params.User.Email
		}

		if params.User.Active != nil {
			record.Active = *params.User.Active
		}

		if params.User.Admin != nil {
			record.Admin = *params.User.Admin
		}

		created, err := usersService.Create(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewCreateUserUnprocessableEntity().WithPayload(payload)
			}

			return user.NewCreateUserDefault(http.StatusInternalServerError)
		}

		return user.NewCreateUserOK().WithPayload(convertUser(created))
	}
}

// UpdateUserHandler implements the handler for the UpdateUser operation.
func UpdateUserHandler(usersService users.Service) user.UpdateUserHandlerFunc {
	return func(params user.UpdateUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewUpdateUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewUpdateUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewUpdateUserDefault(http.StatusInternalServerError)
		}

		if params.User.Slug != nil {
			record.Slug = *params.User.Slug
		}

		if params.User.Username != nil {
			record.Username = *params.User.Username
		}

		if params.User.Password != nil {
			record.Password = (*params.User.Password).String()
		}

		if params.User.Email != nil {
			record.Email = *params.User.Email

		}

		if params.User.Active != nil {
			record.Active = *params.User.Active
		}

		if params.User.Admin != nil {
			record.Admin = *params.User.Admin
		}

		updated, err := usersService.Update(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewUpdateUserUnprocessableEntity().WithPayload(payload)
			}

			return user.NewUpdateUserDefault(http.StatusInternalServerError)
		}

		return user.NewUpdateUserOK().WithPayload(convertUser(updated))
	}
}

// DeleteUserHandler implements the handler for the DeleteUser operation.
func DeleteUserHandler(usersService users.Service) user.DeleteUserHandlerFunc {
	return func(params user.DeleteUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewDeleteUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewDeleteUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserDefault(http.StatusInternalServerError)
		}

		if err := usersService.Delete(params.HTTPRequest.Context(), record.ID); err != nil {
			message := "failed to delete user"

			return user.NewDeleteUserBadRequest().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		message := "successfully deleted user"
		return user.NewDeleteUserOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// ListUserTeamsHandler implements the handler for the ListUserTeams operation.
func ListUserTeamsHandler(usersService users.Service) user.ListUserTeamsHandlerFunc {
	return func(params user.ListUserTeamsParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewListUserTeamsForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := usersService.ListTeams(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			// TODO: add handler if user not found
			return user.NewListUserTeamsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.TeamUser, len(records))
		for id, record := range records {
			payload[id] = convertTeamUser(record)
		}

		return user.NewListUserTeamsOK().WithPayload(payload)
	}
}

// AppendUserToTeamHandler implements the handler for the AppendUserToTeam operation.
func AppendUserToTeamHandler(usersService users.Service, teamsService teams.Service) user.AppendUserToTeamHandlerFunc {
	return func(params user.AppendUserToTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewAppendUserToTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		u, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewAppendUserToTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewAppendUserToTeamDefault(http.StatusInternalServerError)
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), *params.UserTeam.Team)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return user.NewAppendUserToTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewAppendUserToTeamDefault(http.StatusInternalServerError)
		}

		if err := usersService.AppendTeam(params.HTTPRequest.Context(), u.ID, t.ID, *params.UserTeam.Perm); err != nil {
			if err == users.ErrAlreadyAssigned {
				message := "team is already assigned"

				return user.NewAppendUserToTeamPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user team"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewAppendUserToTeamUnprocessableEntity().WithPayload(payload)
			}

			return user.NewAppendUserToTeamDefault(http.StatusInternalServerError)
		}

		message := "successfully assigned user to team"
		return user.NewAppendUserToTeamOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// PermitUserTeamHandler implements the handler for the PermitUserTeam operation.
func PermitUserTeamHandler(usersService users.Service, teamsService teams.Service) user.PermitUserTeamHandlerFunc {
	return func(params user.PermitUserTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewPermitUserTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		u, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewPermitUserTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewPermitUserTeamDefault(http.StatusInternalServerError)
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), *params.UserTeam.Team)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return user.NewPermitUserTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewPermitUserTeamDefault(http.StatusInternalServerError)
		}

		if err := usersService.PermitTeam(params.HTTPRequest.Context(), u.ID, t.ID, *params.UserTeam.Perm); err != nil {
			if err == users.ErrNotAssigned {
				message := "team is not assigned"

				return user.NewPermitUserTeamPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user team"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewPermitUserTeamUnprocessableEntity().WithPayload(payload)
			}

			return user.NewPermitUserTeamDefault(http.StatusInternalServerError)
		}

		message := "successfully updated team perms"
		return user.NewPermitUserTeamOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// DeleteUserFromTeamHandler implements the handler for the DeleteUserFromTeam operation.
func DeleteUserFromTeamHandler(usersService users.Service, teamsService teams.Service) user.DeleteUserFromTeamHandlerFunc {
	return func(params user.DeleteUserFromTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewDeleteUserFromTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		u, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewDeleteUserFromTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromTeamDefault(http.StatusInternalServerError)
		}

		t, err := teamsService.Show(params.HTTPRequest.Context(), *params.UserTeam.Team)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return user.NewDeleteUserFromTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromTeamDefault(http.StatusInternalServerError)
		}

		if err := usersService.DropTeam(params.HTTPRequest.Context(), u.ID, t.ID); err != nil {
			if err == users.ErrNotAssigned {
				message := "team is not assigned"

				return user.NewDeleteUserFromTeamPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromTeamDefault(http.StatusInternalServerError)
		}

		message := "successfully removed from team"
		return user.NewDeleteUserFromTeamOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// ListUserModsHandler implements the handler for the ListUserMods operation.
func ListUserModsHandler(usersService users.Service) user.ListUserModsHandlerFunc {
	return func(params user.ListUserModsParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewListUserModsForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := usersService.ListMods(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			// TODO: add handler if user not found
			return user.NewListUserModsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.UserMod, len(records))
		for id, record := range records {
			payload[id] = convertUserMod(record)
		}

		return user.NewListUserModsOK().WithPayload(payload)
	}
}

// AppendUserToModHandler implements the handler for the AppendUserToMod operation.
func AppendUserToModHandler(usersService users.Service, modsService mods.Service) user.AppendUserToModHandlerFunc {
	return func(params user.AppendUserToModParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewAppendUserToModForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewAppendUserToModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewAppendUserToModDefault(http.StatusInternalServerError)
		}

		u, err := modsService.Show(params.HTTPRequest.Context(), *params.UserMod.Mod)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return user.NewAppendUserToModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewAppendUserToModDefault(http.StatusInternalServerError)
		}

		if err := usersService.AppendMod(params.HTTPRequest.Context(), t.ID, u.ID, *params.UserMod.Perm); err != nil {
			if err == users.ErrAlreadyAssigned {
				message := "mod is already assigned"

				return user.NewAppendUserToModPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user mod"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewAppendUserToModUnprocessableEntity().WithPayload(payload)
			}

			return user.NewAppendUserToModDefault(http.StatusInternalServerError)
		}

		message := "successfully assigned user to mod"
		return user.NewAppendUserToModOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// PermitUserModHandler implements the handler for the PermitUserMod operation.
func PermitUserModHandler(usersService users.Service, modsService mods.Service) user.PermitUserModHandlerFunc {
	return func(params user.PermitUserModParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewPermitUserModForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewPermitUserModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewPermitUserModDefault(http.StatusInternalServerError)
		}

		u, err := modsService.Show(params.HTTPRequest.Context(), *params.UserMod.Mod)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return user.NewPermitUserModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewPermitUserModDefault(http.StatusInternalServerError)
		}

		if err := usersService.PermitMod(params.HTTPRequest.Context(), t.ID, u.ID, *params.UserMod.Perm); err != nil {
			if err == users.ErrNotAssigned {
				message := "mod is not assigned"

				return user.NewPermitUserModPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user mod"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewPermitUserModUnprocessableEntity().WithPayload(payload)
			}

			return user.NewPermitUserModDefault(http.StatusInternalServerError)
		}

		message := "successfully updated mod perms"
		return user.NewPermitUserModOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// DeleteUserFromModHandler implements the handler for the DeleteUserFromMod operation.
func DeleteUserFromModHandler(usersService users.Service, modsService mods.Service) user.DeleteUserFromModHandlerFunc {
	return func(params user.DeleteUserFromModParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewDeleteUserFromModForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewDeleteUserFromModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromModDefault(http.StatusInternalServerError)
		}

		u, err := modsService.Show(params.HTTPRequest.Context(), *params.UserMod.Mod)

		if err != nil {
			if err == mods.ErrNotFound {
				message := "mod not found"

				return user.NewDeleteUserFromModNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromModDefault(http.StatusInternalServerError)
		}

		if err := usersService.DropMod(params.HTTPRequest.Context(), t.ID, u.ID); err != nil {
			if err == users.ErrNotAssigned {
				message := "mod is not assigned"

				return user.NewDeleteUserFromModPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromModDefault(http.StatusInternalServerError)
		}

		message := "successfully removed from mod"
		return user.NewDeleteUserFromModOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// ListUserPacksHandler implements the handler for the ListUserPacks operation.
func ListUserPacksHandler(usersService users.Service) user.ListUserPacksHandlerFunc {
	return func(params user.ListUserPacksParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewListUserPacksForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := usersService.ListPacks(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			// TODO: add handler if user not found
			return user.NewListUserPacksDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.UserPack, len(records))
		for id, record := range records {
			payload[id] = convertUserPack(record)
		}

		return user.NewListUserPacksOK().WithPayload(payload)
	}
}

// AppendUserToPackHandler implements the handler for the AppendUserToPack operation.
func AppendUserToPackHandler(usersService users.Service, packsService packs.Service) user.AppendUserToPackHandlerFunc {
	return func(params user.AppendUserToPackParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewAppendUserToPackForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewAppendUserToPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewAppendUserToPackDefault(http.StatusInternalServerError)
		}

		u, err := packsService.Show(params.HTTPRequest.Context(), *params.UserPack.Pack)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return user.NewAppendUserToPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewAppendUserToPackDefault(http.StatusInternalServerError)
		}

		if err := usersService.AppendPack(params.HTTPRequest.Context(), t.ID, u.ID, *params.UserPack.Perm); err != nil {
			if err == users.ErrAlreadyAssigned {
				message := "pack is already assigned"

				return user.NewAppendUserToPackPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user pack"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewAppendUserToPackUnprocessableEntity().WithPayload(payload)
			}

			return user.NewAppendUserToPackDefault(http.StatusInternalServerError)
		}

		message := "successfully assigned user to pack"
		return user.NewAppendUserToPackOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// PermitUserPackHandler implements the handler for the PermitUserPack operation.
func PermitUserPackHandler(usersService users.Service, packsService packs.Service) user.PermitUserPackHandlerFunc {
	return func(params user.PermitUserPackParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewPermitUserPackForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewPermitUserPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewPermitUserPackDefault(http.StatusInternalServerError)
		}

		u, err := packsService.Show(params.HTTPRequest.Context(), *params.UserPack.Pack)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return user.NewPermitUserPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewPermitUserPackDefault(http.StatusInternalServerError)
		}

		if err := usersService.PermitPack(params.HTTPRequest.Context(), t.ID, u.ID, *params.UserPack.Perm); err != nil {
			if err == users.ErrNotAssigned {
				message := "pack is not assigned"

				return user.NewPermitUserPackPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(validate.Errors); ok {
				message := "failed to validate user pack"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return user.NewPermitUserPackUnprocessableEntity().WithPayload(payload)
			}

			return user.NewPermitUserPackDefault(http.StatusInternalServerError)
		}

		message := "successfully updated pack perms"
		return user.NewPermitUserPackOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// DeleteUserFromPackHandler implements the handler for the DeleteUserFromPack operation.
func DeleteUserFromPackHandler(usersService users.Service, packsService packs.Service) user.DeleteUserFromPackHandlerFunc {
	return func(params user.DeleteUserFromPackParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return user.NewDeleteUserFromPackForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := usersService.Show(params.HTTPRequest.Context(), params.UserID)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return user.NewDeleteUserFromPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromPackDefault(http.StatusInternalServerError)
		}

		u, err := packsService.Show(params.HTTPRequest.Context(), *params.UserPack.Pack)

		if err != nil {
			if err == packs.ErrNotFound {
				message := "pack not found"

				return user.NewDeleteUserFromPackNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromPackDefault(http.StatusInternalServerError)
		}

		if err := usersService.DropPack(params.HTTPRequest.Context(), t.ID, u.ID); err != nil {
			if err == users.ErrNotAssigned {
				message := "pack is not assigned"

				return user.NewDeleteUserFromPackPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return user.NewDeleteUserFromPackDefault(http.StatusInternalServerError)
		}

		message := "successfully removed from pack"
		return user.NewDeleteUserFromPackOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// convertUser is a simple helper to convert between different model formats.
func convertUser(record *model.User) *models.User {
	teams := make([]*models.TeamUser, 0)

	for _, team := range record.Teams {
		teams = append(teams, convertTeamUser(team))
	}

	mods := make([]*models.UserMod, 0)

	for _, mod := range record.Mods {
		mods = append(mods, convertUserMod(mod))
	}

	packs := make([]*models.UserPack, 0)

	for _, pack := range record.Packs {
		packs = append(packs, convertUserPack(pack))
	}

	return &models.User{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Email:     &record.Email,
		Username:  &record.Username,
		Password:  nil,
		Avatar:    &record.Avatar,
		Active:    &record.Active,
		Admin:     &record.Admin,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
		Teams:     teams,
		Mods:      mods,
		Packs:     packs,
	}
}

// convertUserPack is a simple helper to convert between different model formats.
func convertUserPack(record *model.UserPack) *models.UserPack {
	userID := strfmt.UUID(record.UserID)
	packID := strfmt.UUID(record.PackID)

	return &models.UserPack{
		UserID:    &userID,
		User:      convertUser(record.User),
		PackID:    &packID,
		Pack:      convertPack(record.Pack),
		Perm:      &record.Perm,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}

// convertUserMod is a simple helper to convert between different model formats.
func convertUserMod(record *model.UserMod) *models.UserMod {
	userID := strfmt.UUID(record.UserID)
	modID := strfmt.UUID(record.ModID)

	return &models.UserMod{
		UserID:    &userID,
		User:      convertUser(record.User),
		ModID:     &modID,
		Mod:       convertMod(record.Mod),
		Perm:      &record.Perm,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}
