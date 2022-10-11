// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UpdatePackHandlerFunc turns a function with the right signature into a update pack handler
type UpdatePackHandlerFunc func(UpdatePackParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePackHandlerFunc) Handle(params UpdatePackParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// UpdatePackHandler interface for that can handle valid update pack params
type UpdatePackHandler interface {
	Handle(UpdatePackParams, *models.User) middleware.Responder
}

// NewUpdatePack creates a new http.Handler for the update pack operation
func NewUpdatePack(ctx *middleware.Context, handler UpdatePackHandler) *UpdatePack {
	return &UpdatePack{Context: ctx, Handler: handler}
}

/*
	UpdatePack swagger:route PUT /packs/{pack_id} pack updatePack

Update a specific pack
*/
type UpdatePack struct {
	Context *middleware.Context
	Handler UpdatePackHandler
}

func (o *UpdatePack) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdatePackParams()
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
