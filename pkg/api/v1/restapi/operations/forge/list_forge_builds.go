// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListForgeBuildsHandlerFunc turns a function with the right signature into a list forge builds handler
type ListForgeBuildsHandlerFunc func(ListForgeBuildsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListForgeBuildsHandlerFunc) Handle(params ListForgeBuildsParams) middleware.Responder {
	return fn(params)
}

// ListForgeBuildsHandler interface for that can handle valid list forge builds params
type ListForgeBuildsHandler interface {
	Handle(ListForgeBuildsParams) middleware.Responder
}

// NewListForgeBuilds creates a new http.Handler for the list forge builds operation
func NewListForgeBuilds(ctx *middleware.Context, handler ListForgeBuildsHandler) *ListForgeBuilds {
	return &ListForgeBuilds{Context: ctx, Handler: handler}
}

/*ListForgeBuilds swagger:route GET /forge/{forgeID}/builds forge listForgeBuilds

Fetch the builds assigned to a Forge version

*/
type ListForgeBuilds struct {
	Context *middleware.Context
	Handler ListForgeBuildsHandler
}

func (o *ListForgeBuilds) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListForgeBuildsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
