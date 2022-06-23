// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// CreateBuildHandlerFunc turns a function with the right signature into a create build handler
type CreateBuildHandlerFunc func(CreateBuildParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBuildHandlerFunc) Handle(params CreateBuildParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// CreateBuildHandler interface for that can handle valid create build params
type CreateBuildHandler interface {
	Handle(CreateBuildParams, *models.User) middleware.Responder
}

// NewCreateBuild creates a new http.Handler for the create build operation
func NewCreateBuild(ctx *middleware.Context, handler CreateBuildHandler) *CreateBuild {
	return &CreateBuild{Context: ctx, Handler: handler}
}

/* CreateBuild swagger:route POST /packs/{pack_id}/builds pack createBuild

Create a new build for a pack

*/
type CreateBuild struct {
	Context *middleware.Context
	Handler CreateBuildHandler
}

func (o *CreateBuild) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateBuildParams()
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