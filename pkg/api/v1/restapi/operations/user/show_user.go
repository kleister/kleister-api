// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ShowUserHandlerFunc turns a function with the right signature into a show user handler
type ShowUserHandlerFunc func(ShowUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ShowUserHandlerFunc) Handle(params ShowUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ShowUserHandler interface for that can handle valid show user params
type ShowUserHandler interface {
	Handle(ShowUserParams, *models.User) middleware.Responder
}

// NewShowUser creates a new http.Handler for the show user operation
func NewShowUser(ctx *middleware.Context, handler ShowUserHandler) *ShowUser {
	return &ShowUser{Context: ctx, Handler: handler}
}

/*
	ShowUser swagger:route GET /users/{user_id} user showUser

Fetch a specific user
*/
type ShowUser struct {
	Context *middleware.Context
	Handler ShowUserHandler
}

func (o *ShowUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewShowUserParams()
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
