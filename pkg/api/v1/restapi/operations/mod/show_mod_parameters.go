// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewShowModParams creates a new ShowModParams object
//
// There are no default values defined in the spec.
func NewShowModParams() ShowModParams {

	return ShowModParams{}
}

// ShowModParams contains all the bound params for the show mod operation
// typically these are obtained from a http.Request
//
// swagger:parameters ShowMod
type ShowModParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A mod UUID or slug
	  Required: true
	  In: path
	*/
	ModID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShowModParams() beforehand.
func (o *ShowModParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rModID, rhkModID, _ := route.Params.GetOK("mod_id")
	if err := o.bindModID(rModID, rhkModID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindModID binds and validates parameter ModID from path.
func (o *ShowModParams) bindModID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ModID = raw

	return nil
}
