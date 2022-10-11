// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendModToTeamOKCode is the HTTP code returned for type AppendModToTeamOK
const AppendModToTeamOKCode int = 200

/*
AppendModToTeamOK Plain success message

swagger:response appendModToTeamOK
*/
type AppendModToTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendModToTeamOK creates AppendModToTeamOK with default headers values
func NewAppendModToTeamOK() *AppendModToTeamOK {

	return &AppendModToTeamOK{}
}

// WithPayload adds the payload to the append mod to team o k response
func (o *AppendModToTeamOK) WithPayload(payload *models.GeneralError) *AppendModToTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append mod to team o k response
func (o *AppendModToTeamOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendModToTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendModToTeamForbiddenCode is the HTTP code returned for type AppendModToTeamForbidden
const AppendModToTeamForbiddenCode int = 403

/*
AppendModToTeamForbidden User is not authorized

swagger:response appendModToTeamForbidden
*/
type AppendModToTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendModToTeamForbidden creates AppendModToTeamForbidden with default headers values
func NewAppendModToTeamForbidden() *AppendModToTeamForbidden {

	return &AppendModToTeamForbidden{}
}

// WithPayload adds the payload to the append mod to team forbidden response
func (o *AppendModToTeamForbidden) WithPayload(payload *models.GeneralError) *AppendModToTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append mod to team forbidden response
func (o *AppendModToTeamForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendModToTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendModToTeamNotFoundCode is the HTTP code returned for type AppendModToTeamNotFound
const AppendModToTeamNotFoundCode int = 404

/*
AppendModToTeamNotFound Mod or team not found

swagger:response appendModToTeamNotFound
*/
type AppendModToTeamNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendModToTeamNotFound creates AppendModToTeamNotFound with default headers values
func NewAppendModToTeamNotFound() *AppendModToTeamNotFound {

	return &AppendModToTeamNotFound{}
}

// WithPayload adds the payload to the append mod to team not found response
func (o *AppendModToTeamNotFound) WithPayload(payload *models.GeneralError) *AppendModToTeamNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append mod to team not found response
func (o *AppendModToTeamNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendModToTeamNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendModToTeamPreconditionFailedCode is the HTTP code returned for type AppendModToTeamPreconditionFailed
const AppendModToTeamPreconditionFailedCode int = 412

/*
AppendModToTeamPreconditionFailed Team is already assigned

swagger:response appendModToTeamPreconditionFailed
*/
type AppendModToTeamPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendModToTeamPreconditionFailed creates AppendModToTeamPreconditionFailed with default headers values
func NewAppendModToTeamPreconditionFailed() *AppendModToTeamPreconditionFailed {

	return &AppendModToTeamPreconditionFailed{}
}

// WithPayload adds the payload to the append mod to team precondition failed response
func (o *AppendModToTeamPreconditionFailed) WithPayload(payload *models.GeneralError) *AppendModToTeamPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append mod to team precondition failed response
func (o *AppendModToTeamPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendModToTeamPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendModToTeamUnprocessableEntityCode is the HTTP code returned for type AppendModToTeamUnprocessableEntity
const AppendModToTeamUnprocessableEntityCode int = 422

/*
AppendModToTeamUnprocessableEntity Failed to validate request

swagger:response appendModToTeamUnprocessableEntity
*/
type AppendModToTeamUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewAppendModToTeamUnprocessableEntity creates AppendModToTeamUnprocessableEntity with default headers values
func NewAppendModToTeamUnprocessableEntity() *AppendModToTeamUnprocessableEntity {

	return &AppendModToTeamUnprocessableEntity{}
}

// WithPayload adds the payload to the append mod to team unprocessable entity response
func (o *AppendModToTeamUnprocessableEntity) WithPayload(payload *models.ValidationError) *AppendModToTeamUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append mod to team unprocessable entity response
func (o *AppendModToTeamUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendModToTeamUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
AppendModToTeamDefault Some error unrelated to the handler

swagger:response appendModToTeamDefault
*/
type AppendModToTeamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendModToTeamDefault creates AppendModToTeamDefault with default headers values
func NewAppendModToTeamDefault(code int) *AppendModToTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &AppendModToTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the append mod to team default response
func (o *AppendModToTeamDefault) WithStatusCode(code int) *AppendModToTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the append mod to team default response
func (o *AppendModToTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the append mod to team default response
func (o *AppendModToTeamDefault) WithPayload(payload *models.GeneralError) *AppendModToTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append mod to team default response
func (o *AppendModToTeamDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendModToTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
