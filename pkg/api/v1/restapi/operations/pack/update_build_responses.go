// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UpdateBuildOKCode is the HTTP code returned for type UpdateBuildOK
const UpdateBuildOKCode int = 200

/*UpdateBuildOK The updated build details

swagger:response updateBuildOK
*/
type UpdateBuildOK struct {

	/*
	  In: Body
	*/
	Payload *models.Build `json:"body,omitempty"`
}

// NewUpdateBuildOK creates UpdateBuildOK with default headers values
func NewUpdateBuildOK() *UpdateBuildOK {

	return &UpdateBuildOK{}
}

// WithPayload adds the payload to the update build o k response
func (o *UpdateBuildOK) WithPayload(payload *models.Build) *UpdateBuildOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update build o k response
func (o *UpdateBuildOK) SetPayload(payload *models.Build) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBuildOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBuildForbiddenCode is the HTTP code returned for type UpdateBuildForbidden
const UpdateBuildForbiddenCode int = 403

/*UpdateBuildForbidden User is not authorized

swagger:response updateBuildForbidden
*/
type UpdateBuildForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateBuildForbidden creates UpdateBuildForbidden with default headers values
func NewUpdateBuildForbidden() *UpdateBuildForbidden {

	return &UpdateBuildForbidden{}
}

// WithPayload adds the payload to the update build forbidden response
func (o *UpdateBuildForbidden) WithPayload(payload *models.GeneralError) *UpdateBuildForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update build forbidden response
func (o *UpdateBuildForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBuildForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBuildNotFoundCode is the HTTP code returned for type UpdateBuildNotFound
const UpdateBuildNotFoundCode int = 404

/*UpdateBuildNotFound Build or pack not found

swagger:response updateBuildNotFound
*/
type UpdateBuildNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateBuildNotFound creates UpdateBuildNotFound with default headers values
func NewUpdateBuildNotFound() *UpdateBuildNotFound {

	return &UpdateBuildNotFound{}
}

// WithPayload adds the payload to the update build not found response
func (o *UpdateBuildNotFound) WithPayload(payload *models.GeneralError) *UpdateBuildNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update build not found response
func (o *UpdateBuildNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBuildNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBuildUnprocessableEntityCode is the HTTP code returned for type UpdateBuildUnprocessableEntity
const UpdateBuildUnprocessableEntityCode int = 422

/*UpdateBuildUnprocessableEntity Failed to validate request

swagger:response updateBuildUnprocessableEntity
*/
type UpdateBuildUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewUpdateBuildUnprocessableEntity creates UpdateBuildUnprocessableEntity with default headers values
func NewUpdateBuildUnprocessableEntity() *UpdateBuildUnprocessableEntity {

	return &UpdateBuildUnprocessableEntity{}
}

// WithPayload adds the payload to the update build unprocessable entity response
func (o *UpdateBuildUnprocessableEntity) WithPayload(payload *models.ValidationError) *UpdateBuildUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update build unprocessable entity response
func (o *UpdateBuildUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBuildUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateBuildDefault Some error unrelated to the handler

swagger:response updateBuildDefault
*/
type UpdateBuildDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateBuildDefault creates UpdateBuildDefault with default headers values
func NewUpdateBuildDefault(code int) *UpdateBuildDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateBuildDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update build default response
func (o *UpdateBuildDefault) WithStatusCode(code int) *UpdateBuildDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update build default response
func (o *UpdateBuildDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update build default response
func (o *UpdateBuildDefault) WithPayload(payload *models.GeneralError) *UpdateBuildDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update build default response
func (o *UpdateBuildDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBuildDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}