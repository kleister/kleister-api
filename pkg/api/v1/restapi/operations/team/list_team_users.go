// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListTeamUsersHandlerFunc turns a function with the right signature into a list team users handler
type ListTeamUsersHandlerFunc func(ListTeamUsersParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ListTeamUsersHandlerFunc) Handle(params ListTeamUsersParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ListTeamUsersHandler interface for that can handle valid list team users params
type ListTeamUsersHandler interface {
	Handle(ListTeamUsersParams, *models.User) middleware.Responder
}

// NewListTeamUsers creates a new http.Handler for the list team users operation
func NewListTeamUsers(ctx *middleware.Context, handler ListTeamUsersHandler) *ListTeamUsers {
	return &ListTeamUsers{Context: ctx, Handler: handler}
}

/* ListTeamUsers swagger:route GET /teams/{team_id}/users team listTeamUsers

Fetch all users assigned to team

*/
type ListTeamUsers struct {
	Context *middleware.Context
	Handler ListTeamUsersHandler
}

func (o *ListTeamUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListTeamUsersParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}