// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListPackUsersParams creates a new ListPackUsersParams object
//
// There are no default values defined in the spec.
func NewListPackUsersParams() ListPackUsersParams {

	return ListPackUsersParams{}
}

// ListPackUsersParams contains all the bound params for the list pack users operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListPackUsers
type ListPackUsersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A pack UUID or slug
	  Required: true
	  In: path
	*/
	PackID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListPackUsersParams() beforehand.
func (o *ListPackUsersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPackID, rhkPackID, _ := route.Params.GetOK("pack_id")
	if err := o.bindPackID(rPackID, rhkPackID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPackID binds and validates parameter PackID from path.
func (o *ListPackUsersParams) bindPackID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.PackID = raw

	return nil
}
