// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeletePackFromTeamHandlerFunc turns a function with the right signature into a delete pack from team handler
type DeletePackFromTeamHandlerFunc func(DeletePackFromTeamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePackFromTeamHandlerFunc) Handle(params DeletePackFromTeamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeletePackFromTeamHandler interface for that can handle valid delete pack from team params
type DeletePackFromTeamHandler interface {
	Handle(DeletePackFromTeamParams, *models.User) middleware.Responder
}

// NewDeletePackFromTeam creates a new http.Handler for the delete pack from team operation
func NewDeletePackFromTeam(ctx *middleware.Context, handler DeletePackFromTeamHandler) *DeletePackFromTeam {
	return &DeletePackFromTeam{Context: ctx, Handler: handler}
}

/*
	DeletePackFromTeam swagger:route DELETE /packs/{pack_id}/teams pack deletePackFromTeam

Remove a team from pack
*/
type DeletePackFromTeam struct {
	Context *middleware.Context
	Handler DeletePackFromTeamHandler
}

func (o *DeletePackFromTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePackFromTeamParams()
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
