// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// NewCreateVersionParams creates a new CreateVersionParams object
//
// There are no default values defined in the spec.
func NewCreateVersionParams() CreateVersionParams {

	return CreateVersionParams{}
}

// CreateVersionParams contains all the bound params for the create version operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateVersion
type CreateVersionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A mod UUID or slug
	  Required: true
	  In: path
	*/
	ModID string
	/*The version data to create
	  Required: true
	  In: body
	*/
	Version *models.Version
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateVersionParams() beforehand.
func (o *CreateVersionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rModID, rhkModID, _ := route.Params.GetOK("mod_id")
	if err := o.bindModID(rModID, rhkModID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Version
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("version", "body", ""))
			} else {
				res = append(res, errors.NewParseError("version", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Version = &body
			}
		}
	} else {
		res = append(res, errors.Required("version", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindModID binds and validates parameter ModID from path.
func (o *CreateVersionParams) bindModID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ModID = raw

	return nil
}
