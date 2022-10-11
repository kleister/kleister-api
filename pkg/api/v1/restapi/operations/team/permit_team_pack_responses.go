// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// PermitTeamPackOKCode is the HTTP code returned for type PermitTeamPackOK
const PermitTeamPackOKCode int = 200

/*
PermitTeamPackOK Plain success message

swagger:response permitTeamPackOK
*/
type PermitTeamPackOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamPackOK creates PermitTeamPackOK with default headers values
func NewPermitTeamPackOK() *PermitTeamPackOK {

	return &PermitTeamPackOK{}
}

// WithPayload adds the payload to the permit team pack o k response
func (o *PermitTeamPackOK) WithPayload(payload *models.GeneralError) *PermitTeamPackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team pack o k response
func (o *PermitTeamPackOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamPackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamPackForbiddenCode is the HTTP code returned for type PermitTeamPackForbidden
const PermitTeamPackForbiddenCode int = 403

/*
PermitTeamPackForbidden User is not authorized

swagger:response permitTeamPackForbidden
*/
type PermitTeamPackForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamPackForbidden creates PermitTeamPackForbidden with default headers values
func NewPermitTeamPackForbidden() *PermitTeamPackForbidden {

	return &PermitTeamPackForbidden{}
}

// WithPayload adds the payload to the permit team pack forbidden response
func (o *PermitTeamPackForbidden) WithPayload(payload *models.GeneralError) *PermitTeamPackForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team pack forbidden response
func (o *PermitTeamPackForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamPackForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamPackNotFoundCode is the HTTP code returned for type PermitTeamPackNotFound
const PermitTeamPackNotFoundCode int = 404

/*
PermitTeamPackNotFound Team or pack not found

swagger:response permitTeamPackNotFound
*/
type PermitTeamPackNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamPackNotFound creates PermitTeamPackNotFound with default headers values
func NewPermitTeamPackNotFound() *PermitTeamPackNotFound {

	return &PermitTeamPackNotFound{}
}

// WithPayload adds the payload to the permit team pack not found response
func (o *PermitTeamPackNotFound) WithPayload(payload *models.GeneralError) *PermitTeamPackNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team pack not found response
func (o *PermitTeamPackNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamPackNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamPackPreconditionFailedCode is the HTTP code returned for type PermitTeamPackPreconditionFailed
const PermitTeamPackPreconditionFailedCode int = 412

/*
PermitTeamPackPreconditionFailed Pack is not assigned

swagger:response permitTeamPackPreconditionFailed
*/
type PermitTeamPackPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamPackPreconditionFailed creates PermitTeamPackPreconditionFailed with default headers values
func NewPermitTeamPackPreconditionFailed() *PermitTeamPackPreconditionFailed {

	return &PermitTeamPackPreconditionFailed{}
}

// WithPayload adds the payload to the permit team pack precondition failed response
func (o *PermitTeamPackPreconditionFailed) WithPayload(payload *models.GeneralError) *PermitTeamPackPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team pack precondition failed response
func (o *PermitTeamPackPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamPackPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitTeamPackUnprocessableEntityCode is the HTTP code returned for type PermitTeamPackUnprocessableEntity
const PermitTeamPackUnprocessableEntityCode int = 422

/*
PermitTeamPackUnprocessableEntity Failed to validate request

swagger:response permitTeamPackUnprocessableEntity
*/
type PermitTeamPackUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewPermitTeamPackUnprocessableEntity creates PermitTeamPackUnprocessableEntity with default headers values
func NewPermitTeamPackUnprocessableEntity() *PermitTeamPackUnprocessableEntity {

	return &PermitTeamPackUnprocessableEntity{}
}

// WithPayload adds the payload to the permit team pack unprocessable entity response
func (o *PermitTeamPackUnprocessableEntity) WithPayload(payload *models.ValidationError) *PermitTeamPackUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team pack unprocessable entity response
func (o *PermitTeamPackUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamPackUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PermitTeamPackDefault Some error unrelated to the handler

swagger:response permitTeamPackDefault
*/
type PermitTeamPackDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitTeamPackDefault creates PermitTeamPackDefault with default headers values
func NewPermitTeamPackDefault(code int) *PermitTeamPackDefault {
	if code <= 0 {
		code = 500
	}

	return &PermitTeamPackDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the permit team pack default response
func (o *PermitTeamPackDefault) WithStatusCode(code int) *PermitTeamPackDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the permit team pack default response
func (o *PermitTeamPackDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the permit team pack default response
func (o *PermitTeamPackDefault) WithPayload(payload *models.GeneralError) *PermitTeamPackDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit team pack default response
func (o *PermitTeamPackDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitTeamPackDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
