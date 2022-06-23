// Code generated by go-swagger; DO NOT EDIT.

package minecraft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// SearchMinecraftsHandlerFunc turns a function with the right signature into a search minecrafts handler
type SearchMinecraftsHandlerFunc func(SearchMinecraftsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchMinecraftsHandlerFunc) Handle(params SearchMinecraftsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// SearchMinecraftsHandler interface for that can handle valid search minecrafts params
type SearchMinecraftsHandler interface {
	Handle(SearchMinecraftsParams, *models.User) middleware.Responder
}

// NewSearchMinecrafts creates a new http.Handler for the search minecrafts operation
func NewSearchMinecrafts(ctx *middleware.Context, handler SearchMinecraftsHandler) *SearchMinecrafts {
	return &SearchMinecrafts{Context: ctx, Handler: handler}
}

/* SearchMinecrafts swagger:route GET /minecraft/{minecraft_id} minecraft searchMinecrafts

Search for available Minecraft versions

*/
type SearchMinecrafts struct {
	Context *middleware.Context
	Handler SearchMinecraftsHandler
}

func (o *SearchMinecrafts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSearchMinecraftsParams()
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