// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeletePackFromUserHandlerFunc turns a function with the right signature into a delete pack from user handler
type DeletePackFromUserHandlerFunc func(DeletePackFromUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePackFromUserHandlerFunc) Handle(params DeletePackFromUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeletePackFromUserHandler interface for that can handle valid delete pack from user params
type DeletePackFromUserHandler interface {
	Handle(DeletePackFromUserParams, *models.User) middleware.Responder
}

// NewDeletePackFromUser creates a new http.Handler for the delete pack from user operation
func NewDeletePackFromUser(ctx *middleware.Context, handler DeletePackFromUserHandler) *DeletePackFromUser {
	return &DeletePackFromUser{Context: ctx, Handler: handler}
}

/* DeletePackFromUser swagger:route DELETE /packs/{pack_id}/users pack deletePackFromUser

Remove a user from pack

*/
type DeletePackFromUser struct {
	Context *middleware.Context
	Handler DeletePackFromUserHandler
}

func (o *DeletePackFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePackFromUserParams()
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