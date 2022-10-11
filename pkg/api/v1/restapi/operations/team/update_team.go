// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UpdateTeamHandlerFunc turns a function with the right signature into a update team handler
type UpdateTeamHandlerFunc func(UpdateTeamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateTeamHandlerFunc) Handle(params UpdateTeamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// UpdateTeamHandler interface for that can handle valid update team params
type UpdateTeamHandler interface {
	Handle(UpdateTeamParams, *models.User) middleware.Responder
}

// NewUpdateTeam creates a new http.Handler for the update team operation
func NewUpdateTeam(ctx *middleware.Context, handler UpdateTeamHandler) *UpdateTeam {
	return &UpdateTeam{Context: ctx, Handler: handler}
}

/*
	UpdateTeam swagger:route PUT /teams/{team_id} team updateTeam

Update a specific team
*/
type UpdateTeam struct {
	Context *middleware.Context
	Handler UpdateTeamHandler
}

func (o *UpdateTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateTeamParams()
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
