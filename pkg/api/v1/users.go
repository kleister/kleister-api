package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kleister/kleister-api/pkg/api/v1/models"
	"github.com/kleister/kleister-api/pkg/api/v1/restapi/operations/user"
	"github.com/kleister/kleister-api/pkg/model"
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
