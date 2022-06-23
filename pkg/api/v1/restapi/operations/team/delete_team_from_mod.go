// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteTeamFromModHandlerFunc turns a function with the right signature into a delete team from mod handler
type DeleteTeamFromModHandlerFunc func(DeleteTeamFromModParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTeamFromModHandlerFunc) Handle(params DeleteTeamFromModParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteTeamFromModHandler interface for that can handle valid delete team from mod params
type DeleteTeamFromModHandler interface {
	Handle(DeleteTeamFromModParams, *models.User) middleware.Responder
}

// NewDeleteTeamFromMod creates a new http.Handler for the delete team from mod operation
func NewDeleteTeamFromMod(ctx *middleware.Context, handler DeleteTeamFromModHandler) *DeleteTeamFromMod {
	return &DeleteTeamFromMod{Context: ctx, Handler: handler}
}

/* DeleteTeamFromMod swagger:route DELETE /teams/{team_id}/mods team deleteTeamFromMod

Remove a mod from team

*/
type DeleteTeamFromMod struct {
	Context *middleware.Context
	Handler DeleteTeamFromModHandler
}

func (o *DeleteTeamFromMod) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTeamFromModParams()
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