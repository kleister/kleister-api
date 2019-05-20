// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteUserOKCode is the HTTP code returned for type DeleteUserOK
const DeleteUserOKCode int = 200

/*DeleteUserOK Plain success message

swagger:response deleteUserOK
*/
type DeleteUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserOK creates DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {

	return &DeleteUserOK{}
}

// WithPayload adds the payload to the delete user o k response
func (o *DeleteUserOK) WithPayload(payload *models.GeneralError) *DeleteUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user o k response
func (o *DeleteUserOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserBadRequestCode is the HTTP code returned for type DeleteUserBadRequest
const DeleteUserBadRequestCode int = 400

/*DeleteUserBadRequest Failed to delete the user

swagger:response deleteUserBadRequest
*/
type DeleteUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserBadRequest creates DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {

	return &DeleteUserBadRequest{}
}

// WithPayload adds the payload to the delete user bad request response
func (o *DeleteUserBadRequest) WithPayload(payload *models.GeneralError) *DeleteUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user bad request response
func (o *DeleteUserBadRequest) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserForbiddenCode is the HTTP code returned for type DeleteUserForbidden
const DeleteUserForbiddenCode int = 403

/*DeleteUserForbidden User is not authorized

swagger:response deleteUserForbidden
*/
type DeleteUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserForbidden creates DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {

	return &DeleteUserForbidden{}
}

// WithPayload adds the payload to the delete user forbidden response
func (o *DeleteUserForbidden) WithPayload(payload *models.GeneralError) *DeleteUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user forbidden response
func (o *DeleteUserForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteUserDefault Some error unrelated to the handler

swagger:response deleteUserDefault
*/
type DeleteUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserDefault creates DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete user default response
func (o *DeleteUserDefault) WithStatusCode(code int) *DeleteUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete user default response
func (o *DeleteUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete user default response
func (o *DeleteUserDefault) WithPayload(payload *models.GeneralError) *DeleteUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user default response
func (o *DeleteUserDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
