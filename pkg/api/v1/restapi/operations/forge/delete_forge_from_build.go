// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteForgeFromBuildHandlerFunc turns a function with the right signature into a delete forge from build handler
type DeleteForgeFromBuildHandlerFunc func(DeleteForgeFromBuildParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteForgeFromBuildHandlerFunc) Handle(params DeleteForgeFromBuildParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteForgeFromBuildHandler interface for that can handle valid delete forge from build params
type DeleteForgeFromBuildHandler interface {
	Handle(DeleteForgeFromBuildParams, *models.User) middleware.Responder
}

// NewDeleteForgeFromBuild creates a new http.Handler for the delete forge from build operation
func NewDeleteForgeFromBuild(ctx *middleware.Context, handler DeleteForgeFromBuildHandler) *DeleteForgeFromBuild {
	return &DeleteForgeFromBuild{Context: ctx, Handler: handler}
}

/* DeleteForgeFromBuild swagger:route DELETE /forge/{forge_id}/builds forge deleteForgeFromBuild

Unlink a build from a Forge version

*/
type DeleteForgeFromBuild struct {
	Context *middleware.Context
	Handler DeleteForgeFromBuildHandler
}

func (o *DeleteForgeFromBuild) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteForgeFromBuildParams()
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