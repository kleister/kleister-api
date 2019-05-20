// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PermitPackUserHandlerFunc turns a function with the right signature into a permit pack user handler
type PermitPackUserHandlerFunc func(PermitPackUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PermitPackUserHandlerFunc) Handle(params PermitPackUserParams) middleware.Responder {
	return fn(params)
}

// PermitPackUserHandler interface for that can handle valid permit pack user params
type PermitPackUserHandler interface {
	Handle(PermitPackUserParams) middleware.Responder
}

// NewPermitPackUser creates a new http.Handler for the permit pack user operation
func NewPermitPackUser(ctx *middleware.Context, handler PermitPackUserHandler) *PermitPackUser {
	return &PermitPackUser{Context: ctx, Handler: handler}
}

/*PermitPackUser swagger:route PUT /packs/{packID}/users pack permitPackUser

Update user perms for pack

*/
type PermitPackUser struct {
	Context *middleware.Context
	Handler PermitPackUserHandler
}

func (o *PermitPackUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPermitPackUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
