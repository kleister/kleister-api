// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteBuildFromVersionHandlerFunc turns a function with the right signature into a delete build from version handler
type DeleteBuildFromVersionHandlerFunc func(DeleteBuildFromVersionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBuildFromVersionHandlerFunc) Handle(params DeleteBuildFromVersionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteBuildFromVersionHandler interface for that can handle valid delete build from version params
type DeleteBuildFromVersionHandler interface {
	Handle(DeleteBuildFromVersionParams, *models.User) middleware.Responder
}

// NewDeleteBuildFromVersion creates a new http.Handler for the delete build from version operation
func NewDeleteBuildFromVersion(ctx *middleware.Context, handler DeleteBuildFromVersionHandler) *DeleteBuildFromVersion {
	return &DeleteBuildFromVersion{Context: ctx, Handler: handler}
}

/*
	DeleteBuildFromVersion swagger:route DELETE /packs/{pack_id}/builds/{build_id}/versions pack deleteBuildFromVersion

Unlink a version from a build
*/
type DeleteBuildFromVersion struct {
	Context *middleware.Context
	Handler DeleteBuildFromVersionHandler
}

func (o *DeleteBuildFromVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteBuildFromVersionParams()
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
