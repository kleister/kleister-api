// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListTeamModsHandlerFunc turns a function with the right signature into a list team mods handler
type ListTeamModsHandlerFunc func(ListTeamModsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ListTeamModsHandlerFunc) Handle(params ListTeamModsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ListTeamModsHandler interface for that can handle valid list team mods params
type ListTeamModsHandler interface {
	Handle(ListTeamModsParams, *models.User) middleware.Responder
}

// NewListTeamMods creates a new http.Handler for the list team mods operation
func NewListTeamMods(ctx *middleware.Context, handler ListTeamModsHandler) *ListTeamMods {
	return &ListTeamMods{Context: ctx, Handler: handler}
}

/* ListTeamMods swagger:route GET /teams/{team_id}/mods team listTeamMods

Fetch all mods assigned to team

*/
type ListTeamMods struct {
	Context *middleware.Context
	Handler ListTeamModsHandler
}

func (o *ListTeamMods) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListTeamModsParams()
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
