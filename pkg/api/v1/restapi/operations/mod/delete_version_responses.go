// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteVersionOKCode is the HTTP code returned for type DeleteVersionOK
const DeleteVersionOKCode int = 200

/*DeleteVersionOK Plain success message

swagger:response deleteVersionOK
*/
type DeleteVersionOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteVersionOK creates DeleteVersionOK with default headers values
func NewDeleteVersionOK() *DeleteVersionOK {

	return &DeleteVersionOK{}
}

// WithPayload adds the payload to the delete version o k response
func (o *DeleteVersionOK) WithPayload(payload *models.GeneralError) *DeleteVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete version o k response
func (o *DeleteVersionOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVersionBadRequestCode is the HTTP code returned for type DeleteVersionBadRequest
const DeleteVersionBadRequestCode int = 400

/*DeleteVersionBadRequest Failed to delete the version

swagger:response deleteVersionBadRequest
*/
type DeleteVersionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteVersionBadRequest creates DeleteVersionBadRequest with default headers values
func NewDeleteVersionBadRequest() *DeleteVersionBadRequest {

	return &DeleteVersionBadRequest{}
}

// WithPayload adds the payload to the delete version bad request response
func (o *DeleteVersionBadRequest) WithPayload(payload *models.GeneralError) *DeleteVersionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete version bad request response
func (o *DeleteVersionBadRequest) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVersionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVersionForbiddenCode is the HTTP code returned for type DeleteVersionForbidden
const DeleteVersionForbiddenCode int = 403

/*DeleteVersionForbidden User is not authorized

swagger:response deleteVersionForbidden
*/
type DeleteVersionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteVersionForbidden creates DeleteVersionForbidden with default headers values
func NewDeleteVersionForbidden() *DeleteVersionForbidden {

	return &DeleteVersionForbidden{}
}

// WithPayload adds the payload to the delete version forbidden response
func (o *DeleteVersionForbidden) WithPayload(payload *models.GeneralError) *DeleteVersionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete version forbidden response
func (o *DeleteVersionForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVersionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVersionNotFoundCode is the HTTP code returned for type DeleteVersionNotFound
const DeleteVersionNotFoundCode int = 404

/*DeleteVersionNotFound Version or mod not found

swagger:response deleteVersionNotFound
*/
type DeleteVersionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteVersionNotFound creates DeleteVersionNotFound with default headers values
func NewDeleteVersionNotFound() *DeleteVersionNotFound {

	return &DeleteVersionNotFound{}
}

// WithPayload adds the payload to the delete version not found response
func (o *DeleteVersionNotFound) WithPayload(payload *models.GeneralError) *DeleteVersionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete version not found response
func (o *DeleteVersionNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVersionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteVersionDefault Some error unrelated to the handler

swagger:response deleteVersionDefault
*/
type DeleteVersionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteVersionDefault creates DeleteVersionDefault with default headers values
func NewDeleteVersionDefault(code int) *DeleteVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete version default response
func (o *DeleteVersionDefault) WithStatusCode(code int) *DeleteVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete version default response
func (o *DeleteVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete version default response
func (o *DeleteVersionDefault) WithPayload(payload *models.GeneralError) *DeleteVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete version default response
func (o *DeleteVersionDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}