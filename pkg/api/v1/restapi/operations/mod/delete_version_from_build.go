// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteVersionFromBuildHandlerFunc turns a function with the right signature into a delete version from build handler
type DeleteVersionFromBuildHandlerFunc func(DeleteVersionFromBuildParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteVersionFromBuildHandlerFunc) Handle(params DeleteVersionFromBuildParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteVersionFromBuildHandler interface for that can handle valid delete version from build params
type DeleteVersionFromBuildHandler interface {
	Handle(DeleteVersionFromBuildParams, *models.User) middleware.Responder
}

// NewDeleteVersionFromBuild creates a new http.Handler for the delete version from build operation
func NewDeleteVersionFromBuild(ctx *middleware.Context, handler DeleteVersionFromBuildHandler) *DeleteVersionFromBuild {
	return &DeleteVersionFromBuild{Context: ctx, Handler: handler}
}

/*
	DeleteVersionFromBuild swagger:route DELETE /mods/{mod_id}/versions/{version_id}/builds mod deleteVersionFromBuild

Unlink a build from a version
*/
type DeleteVersionFromBuild struct {
	Context *middleware.Context
	Handler DeleteVersionFromBuildHandler
}

func (o *DeleteVersionFromBuild) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteVersionFromBuildParams()
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
