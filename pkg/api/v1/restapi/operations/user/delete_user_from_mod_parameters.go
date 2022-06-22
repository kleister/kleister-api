// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewDeleteUserFromModParams creates a new DeleteUserFromModParams object
//
// There are no default values defined in the spec.
func NewDeleteUserFromModParams() DeleteUserFromModParams {

	return DeleteUserFromModParams{}
}

// DeleteUserFromModParams contains all the bound params for the delete user from mod operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteUserFromMod
type DeleteUserFromModParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A user UUID or slug
	  Required: true
	  In: path
	*/
	UserID string
	/*The user mod data to delete
	  Required: true
	  In: body
	*/
	UserMod *models.UserModParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteUserFromModParams() beforehand.
func (o *DeleteUserFromModParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUserID, rhkUserID, _ := route.Params.GetOK("user_id")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserModParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("userMod", "body", ""))
			} else {
				res = append(res, errors.NewParseError("userMod", "body", "", err))
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
				o.UserMod = &body
			}
		}
	} else {
		res = append(res, errors.Required("userMod", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *DeleteUserFromModParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UserID = raw

	return nil
}
