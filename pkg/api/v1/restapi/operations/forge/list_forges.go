// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListForgesHandlerFunc turns a function with the right signature into a list forges handler
type ListForgesHandlerFunc func(ListForgesParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ListForgesHandlerFunc) Handle(params ListForgesParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ListForgesHandler interface for that can handle valid list forges params
type ListForgesHandler interface {
	Handle(ListForgesParams, *models.User) middleware.Responder
}

// NewListForges creates a new http.Handler for the list forges operation
func NewListForges(ctx *middleware.Context, handler ListForgesHandler) *ListForges {
	return &ListForges{Context: ctx, Handler: handler}
}

/*
	ListForges swagger:route GET /forge forge listForges

Fetch the available Forge versions
*/
type ListForges struct {
	Context *middleware.Context
	Handler ListForgesHandler
}

func (o *ListForges) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListForgesParams()
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
