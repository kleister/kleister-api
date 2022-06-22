// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendUserToModOKCode is the HTTP code returned for type AppendUserToModOK
const AppendUserToModOKCode int = 200

/*AppendUserToModOK Plain success message

swagger:response appendUserToModOK
*/
type AppendUserToModOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToModOK creates AppendUserToModOK with default headers values
func NewAppendUserToModOK() *AppendUserToModOK {

	return &AppendUserToModOK{}
}

// WithPayload adds the payload to the append user to mod o k response
func (o *AppendUserToModOK) WithPayload(payload *models.GeneralError) *AppendUserToModOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to mod o k response
func (o *AppendUserToModOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToModOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendUserToModForbiddenCode is the HTTP code returned for type AppendUserToModForbidden
const AppendUserToModForbiddenCode int = 403

/*AppendUserToModForbidden User is not authorized

swagger:response appendUserToModForbidden
*/
type AppendUserToModForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToModForbidden creates AppendUserToModForbidden with default headers values
func NewAppendUserToModForbidden() *AppendUserToModForbidden {

	return &AppendUserToModForbidden{}
}

// WithPayload adds the payload to the append user to mod forbidden response
func (o *AppendUserToModForbidden) WithPayload(payload *models.GeneralError) *AppendUserToModForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to mod forbidden response
func (o *AppendUserToModForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToModForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendUserToModNotFoundCode is the HTTP code returned for type AppendUserToModNotFound
const AppendUserToModNotFoundCode int = 404

/*AppendUserToModNotFound User or mod not found

swagger:response appendUserToModNotFound
*/
type AppendUserToModNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToModNotFound creates AppendUserToModNotFound with default headers values
func NewAppendUserToModNotFound() *AppendUserToModNotFound {

	return &AppendUserToModNotFound{}
}

// WithPayload adds the payload to the append user to mod not found response
func (o *AppendUserToModNotFound) WithPayload(payload *models.GeneralError) *AppendUserToModNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to mod not found response
func (o *AppendUserToModNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToModNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendUserToModPreconditionFailedCode is the HTTP code returned for type AppendUserToModPreconditionFailed
const AppendUserToModPreconditionFailedCode int = 412

/*AppendUserToModPreconditionFailed Mod is already assigned

swagger:response appendUserToModPreconditionFailed
*/
type AppendUserToModPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToModPreconditionFailed creates AppendUserToModPreconditionFailed with default headers values
func NewAppendUserToModPreconditionFailed() *AppendUserToModPreconditionFailed {

	return &AppendUserToModPreconditionFailed{}
}

// WithPayload adds the payload to the append user to mod precondition failed response
func (o *AppendUserToModPreconditionFailed) WithPayload(payload *models.GeneralError) *AppendUserToModPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to mod precondition failed response
func (o *AppendUserToModPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToModPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendUserToModUnprocessableEntityCode is the HTTP code returned for type AppendUserToModUnprocessableEntity
const AppendUserToModUnprocessableEntityCode int = 422

/*AppendUserToModUnprocessableEntity Failed to validate request

swagger:response appendUserToModUnprocessableEntity
*/
type AppendUserToModUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewAppendUserToModUnprocessableEntity creates AppendUserToModUnprocessableEntity with default headers values
func NewAppendUserToModUnprocessableEntity() *AppendUserToModUnprocessableEntity {

	return &AppendUserToModUnprocessableEntity{}
}

// WithPayload adds the payload to the append user to mod unprocessable entity response
func (o *AppendUserToModUnprocessableEntity) WithPayload(payload *models.ValidationError) *AppendUserToModUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to mod unprocessable entity response
func (o *AppendUserToModUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToModUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AppendUserToModDefault Some error unrelated to the handler

swagger:response appendUserToModDefault
*/
type AppendUserToModDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToModDefault creates AppendUserToModDefault with default headers values
func NewAppendUserToModDefault(code int) *AppendUserToModDefault {
	if code <= 0 {
		code = 500
	}

	return &AppendUserToModDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the append user to mod default response
func (o *AppendUserToModDefault) WithStatusCode(code int) *AppendUserToModDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the append user to mod default response
func (o *AppendUserToModDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the append user to mod default response
func (o *AppendUserToModDefault) WithPayload(payload *models.GeneralError) *AppendUserToModDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to mod default response
func (o *AppendUserToModDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToModDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
