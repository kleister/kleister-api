// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteBuildFromVersionHandlerFunc turns a function with the right signature into a delete build from version handler
type DeleteBuildFromVersionHandlerFunc func(DeleteBuildFromVersionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBuildFromVersionHandlerFunc) Handle(params DeleteBuildFromVersionParams) middleware.Responder {
	return fn(params)
}

// DeleteBuildFromVersionHandler interface for that can handle valid delete build from version params
type DeleteBuildFromVersionHandler interface {
	Handle(DeleteBuildFromVersionParams) middleware.Responder
}

// NewDeleteBuildFromVersion creates a new http.Handler for the delete build from version operation
func NewDeleteBuildFromVersion(ctx *middleware.Context, handler DeleteBuildFromVersionHandler) *DeleteBuildFromVersion {
	return &DeleteBuildFromVersion{Context: ctx, Handler: handler}
}

/*DeleteBuildFromVersion swagger:route DELETE /packs/{packID}/builds/{buildID}/versions pack deleteBuildFromVersion

Unlink a version from a build

*/
type DeleteBuildFromVersion struct {
	Context *middleware.Context
	Handler DeleteBuildFromVersionHandler
}

func (o *DeleteBuildFromVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBuildFromVersionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
