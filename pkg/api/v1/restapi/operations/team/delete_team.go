// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteTeamHandlerFunc turns a function with the right signature into a delete team handler
type DeleteTeamHandlerFunc func(DeleteTeamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTeamHandlerFunc) Handle(params DeleteTeamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteTeamHandler interface for that can handle valid delete team params
type DeleteTeamHandler interface {
	Handle(DeleteTeamParams, *models.User) middleware.Responder
}

// NewDeleteTeam creates a new http.Handler for the delete team operation
func NewDeleteTeam(ctx *middleware.Context, handler DeleteTeamHandler) *DeleteTeam {
	return &DeleteTeam{Context: ctx, Handler: handler}
}

/*
	DeleteTeam swagger:route DELETE /teams/{team_id} team deleteTeam

Delete a specific team
*/
type DeleteTeam struct {
	Context *middleware.Context
	Handler DeleteTeamHandler
}

func (o *DeleteTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTeamParams()
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
