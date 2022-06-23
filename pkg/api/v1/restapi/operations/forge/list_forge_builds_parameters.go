// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListForgeBuildsParams creates a new ListForgeBuildsParams object
//
// There are no default values defined in the spec.
func NewListForgeBuildsParams() ListForgeBuildsParams {

	return ListForgeBuildsParams{}
}

// ListForgeBuildsParams contains all the bound params for the list forge builds operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListForgeBuilds
type ListForgeBuildsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A forge UUID or slug
	  Required: true
	  In: path
	*/
	ForgeID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListForgeBuildsParams() beforehand.
func (o *ListForgeBuildsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rForgeID, rhkForgeID, _ := route.Params.GetOK("forge_id")
	if err := o.bindForgeID(rForgeID, rhkForgeID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindForgeID binds and validates parameter ForgeID from path.
func (o *ListForgeBuildsParams) bindForgeID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ForgeID = raw

	return nil
}