// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListBuildVersionsHandlerFunc turns a function with the right signature into a list build versions handler
type ListBuildVersionsHandlerFunc func(ListBuildVersionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListBuildVersionsHandlerFunc) Handle(params ListBuildVersionsParams) middleware.Responder {
	return fn(params)
}

// ListBuildVersionsHandler interface for that can handle valid list build versions params
type ListBuildVersionsHandler interface {
	Handle(ListBuildVersionsParams) middleware.Responder
}

// NewListBuildVersions creates a new http.Handler for the list build versions operation
func NewListBuildVersions(ctx *middleware.Context, handler ListBuildVersionsHandler) *ListBuildVersions {
	return &ListBuildVersions{Context: ctx, Handler: handler}
}

/*ListBuildVersions swagger:route GET /packs/{packID}/builds/{buildID}/versions pack listBuildVersions

Fetch all versions assigned to build

*/
type ListBuildVersions struct {
	Context *middleware.Context
	Handler ListBuildVersionsHandler
}

func (o *ListBuildVersions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListBuildVersionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
