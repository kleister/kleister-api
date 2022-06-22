// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendVersionToBuildOKCode is the HTTP code returned for type AppendVersionToBuildOK
const AppendVersionToBuildOKCode int = 200

/*AppendVersionToBuildOK Plain success message

swagger:response appendVersionToBuildOK
*/
type AppendVersionToBuildOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendVersionToBuildOK creates AppendVersionToBuildOK with default headers values
func NewAppendVersionToBuildOK() *AppendVersionToBuildOK {

	return &AppendVersionToBuildOK{}
}

// WithPayload adds the payload to the append version to build o k response
func (o *AppendVersionToBuildOK) WithPayload(payload *models.GeneralError) *AppendVersionToBuildOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append version to build o k response
func (o *AppendVersionToBuildOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendVersionToBuildOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendVersionToBuildForbiddenCode is the HTTP code returned for type AppendVersionToBuildForbidden
const AppendVersionToBuildForbiddenCode int = 403

/*AppendVersionToBuildForbidden User is not authorized

swagger:response appendVersionToBuildForbidden
*/
type AppendVersionToBuildForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendVersionToBuildForbidden creates AppendVersionToBuildForbidden with default headers values
func NewAppendVersionToBuildForbidden() *AppendVersionToBuildForbidden {

	return &AppendVersionToBuildForbidden{}
}

// WithPayload adds the payload to the append version to build forbidden response
func (o *AppendVersionToBuildForbidden) WithPayload(payload *models.GeneralError) *AppendVersionToBuildForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append version to build forbidden response
func (o *AppendVersionToBuildForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendVersionToBuildForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendVersionToBuildNotFoundCode is the HTTP code returned for type AppendVersionToBuildNotFound
const AppendVersionToBuildNotFoundCode int = 404

/*AppendVersionToBuildNotFound Build, version or mod not found

swagger:response appendVersionToBuildNotFound
*/
type AppendVersionToBuildNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendVersionToBuildNotFound creates AppendVersionToBuildNotFound with default headers values
func NewAppendVersionToBuildNotFound() *AppendVersionToBuildNotFound {

	return &AppendVersionToBuildNotFound{}
}

// WithPayload adds the payload to the append version to build not found response
func (o *AppendVersionToBuildNotFound) WithPayload(payload *models.GeneralError) *AppendVersionToBuildNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append version to build not found response
func (o *AppendVersionToBuildNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendVersionToBuildNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendVersionToBuildPreconditionFailedCode is the HTTP code returned for type AppendVersionToBuildPreconditionFailed
const AppendVersionToBuildPreconditionFailedCode int = 412

/*AppendVersionToBuildPreconditionFailed Build is already assigned

swagger:response appendVersionToBuildPreconditionFailed
*/
type AppendVersionToBuildPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendVersionToBuildPreconditionFailed creates AppendVersionToBuildPreconditionFailed with default headers values
func NewAppendVersionToBuildPreconditionFailed() *AppendVersionToBuildPreconditionFailed {

	return &AppendVersionToBuildPreconditionFailed{}
}

// WithPayload adds the payload to the append version to build precondition failed response
func (o *AppendVersionToBuildPreconditionFailed) WithPayload(payload *models.GeneralError) *AppendVersionToBuildPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append version to build precondition failed response
func (o *AppendVersionToBuildPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendVersionToBuildPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendVersionToBuildUnprocessableEntityCode is the HTTP code returned for type AppendVersionToBuildUnprocessableEntity
const AppendVersionToBuildUnprocessableEntityCode int = 422

/*AppendVersionToBuildUnprocessableEntity Failed to validate request

swagger:response appendVersionToBuildUnprocessableEntity
*/
type AppendVersionToBuildUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewAppendVersionToBuildUnprocessableEntity creates AppendVersionToBuildUnprocessableEntity with default headers values
func NewAppendVersionToBuildUnprocessableEntity() *AppendVersionToBuildUnprocessableEntity {

	return &AppendVersionToBuildUnprocessableEntity{}
}

// WithPayload adds the payload to the append version to build unprocessable entity response
func (o *AppendVersionToBuildUnprocessableEntity) WithPayload(payload *models.ValidationError) *AppendVersionToBuildUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append version to build unprocessable entity response
func (o *AppendVersionToBuildUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendVersionToBuildUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AppendVersionToBuildDefault Some error unrelated to the handler

swagger:response appendVersionToBuildDefault
*/
type AppendVersionToBuildDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendVersionToBuildDefault creates AppendVersionToBuildDefault with default headers values
func NewAppendVersionToBuildDefault(code int) *AppendVersionToBuildDefault {
	if code <= 0 {
		code = 500
	}

	return &AppendVersionToBuildDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the append version to build default response
func (o *AppendVersionToBuildDefault) WithStatusCode(code int) *AppendVersionToBuildDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the append version to build default response
func (o *AppendVersionToBuildDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the append version to build default response
func (o *AppendVersionToBuildDefault) WithPayload(payload *models.GeneralError) *AppendVersionToBuildDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append version to build default response
func (o *AppendVersionToBuildDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendVersionToBuildDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
