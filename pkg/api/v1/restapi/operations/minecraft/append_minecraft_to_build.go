// Code generated by go-swagger; DO NOT EDIT.

package minecraft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendMinecraftToBuildHandlerFunc turns a function with the right signature into a append minecraft to build handler
type AppendMinecraftToBuildHandlerFunc func(AppendMinecraftToBuildParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AppendMinecraftToBuildHandlerFunc) Handle(params AppendMinecraftToBuildParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AppendMinecraftToBuildHandler interface for that can handle valid append minecraft to build params
type AppendMinecraftToBuildHandler interface {
	Handle(AppendMinecraftToBuildParams, *models.User) middleware.Responder
}

// NewAppendMinecraftToBuild creates a new http.Handler for the append minecraft to build operation
func NewAppendMinecraftToBuild(ctx *middleware.Context, handler AppendMinecraftToBuildHandler) *AppendMinecraftToBuild {
	return &AppendMinecraftToBuild{Context: ctx, Handler: handler}
}

/*
	AppendMinecraftToBuild swagger:route POST /minecraft/{minecraft_id}/builds minecraft appendMinecraftToBuild

Assign a build to a Minecraft version
*/
type AppendMinecraftToBuild struct {
	Context *middleware.Context
	Handler AppendMinecraftToBuildHandler
}

func (o *AppendMinecraftToBuild) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAppendMinecraftToBuildParams()
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
