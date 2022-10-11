// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListVersionBuildsHandlerFunc turns a function with the right signature into a list version builds handler
type ListVersionBuildsHandlerFunc func(ListVersionBuildsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ListVersionBuildsHandlerFunc) Handle(params ListVersionBuildsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ListVersionBuildsHandler interface for that can handle valid list version builds params
type ListVersionBuildsHandler interface {
	Handle(ListVersionBuildsParams, *models.User) middleware.Responder
}

// NewListVersionBuilds creates a new http.Handler for the list version builds operation
func NewListVersionBuilds(ctx *middleware.Context, handler ListVersionBuildsHandler) *ListVersionBuilds {
	return &ListVersionBuilds{Context: ctx, Handler: handler}
}

/*
	ListVersionBuilds swagger:route GET /mods/{mod_id}/versions/{version_id}/builds mod listVersionBuilds

Fetch all builds assigned to version
*/
type ListVersionBuilds struct {
	Context *middleware.Context
	Handler ListVersionBuildsHandler
}

func (o *ListVersionBuilds) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListVersionBuildsParams()
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
