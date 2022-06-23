// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// PermitModUserHandlerFunc turns a function with the right signature into a permit mod user handler
type PermitModUserHandlerFunc func(PermitModUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PermitModUserHandlerFunc) Handle(params PermitModUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PermitModUserHandler interface for that can handle valid permit mod user params
type PermitModUserHandler interface {
	Handle(PermitModUserParams, *models.User) middleware.Responder
}

// NewPermitModUser creates a new http.Handler for the permit mod user operation
func NewPermitModUser(ctx *middleware.Context, handler PermitModUserHandler) *PermitModUser {
	return &PermitModUser{Context: ctx, Handler: handler}
}

/* PermitModUser swagger:route PUT /mods/{mod_id}/users mod permitModUser

Update user perms for mod

*/
type PermitModUser struct {
	Context *middleware.Context
	Handler PermitModUserHandler
}

func (o *PermitModUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPermitModUserParams()
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