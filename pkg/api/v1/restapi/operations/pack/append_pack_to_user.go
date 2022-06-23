// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendPackToUserHandlerFunc turns a function with the right signature into a append pack to user handler
type AppendPackToUserHandlerFunc func(AppendPackToUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AppendPackToUserHandlerFunc) Handle(params AppendPackToUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AppendPackToUserHandler interface for that can handle valid append pack to user params
type AppendPackToUserHandler interface {
	Handle(AppendPackToUserParams, *models.User) middleware.Responder
}

// NewAppendPackToUser creates a new http.Handler for the append pack to user operation
func NewAppendPackToUser(ctx *middleware.Context, handler AppendPackToUserHandler) *AppendPackToUser {
	return &AppendPackToUser{Context: ctx, Handler: handler}
}

/* AppendPackToUser swagger:route POST /packs/{pack_id}/users pack appendPackToUser

Assign a user to pack

*/
type AppendPackToUser struct {
	Context *middleware.Context
	Handler AppendPackToUserHandler
}

func (o *AppendPackToUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAppendPackToUserParams()
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