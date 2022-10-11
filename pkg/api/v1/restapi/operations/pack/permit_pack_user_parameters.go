// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// NewPermitPackUserParams creates a new PermitPackUserParams object
//
// There are no default values defined in the spec.
func NewPermitPackUserParams() PermitPackUserParams {

	return PermitPackUserParams{}
}

// PermitPackUserParams contains all the bound params for the permit pack user operation
// typically these are obtained from a http.Request
//
// swagger:parameters PermitPackUser
type PermitPackUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A pack UUID or slug
	  Required: true
	  In: path
	*/
	PackID string
	/*The pack user data to update
	  Required: true
	  In: body
	*/
	PackUser *models.PackUserParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPermitPackUserParams() beforehand.
func (o *PermitPackUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPackID, rhkPackID, _ := route.Params.GetOK("pack_id")
	if err := o.bindPackID(rPackID, rhkPackID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PackUserParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("packUser", "body", ""))
			} else {
				res = append(res, errors.NewParseError("packUser", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.PackUser = &body
			}
		}
	} else {
		res = append(res, errors.Required("packUser", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPackID binds and validates parameter PackID from path.
func (o *PermitPackUserParams) bindPackID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.PackID = raw

	return nil
}
