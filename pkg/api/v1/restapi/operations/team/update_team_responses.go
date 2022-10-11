// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UpdateTeamOKCode is the HTTP code returned for type UpdateTeamOK
const UpdateTeamOKCode int = 200

/*
UpdateTeamOK The updated team details

swagger:response updateTeamOK
*/
type UpdateTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.Team `json:"body,omitempty"`
}

// NewUpdateTeamOK creates UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {

	return &UpdateTeamOK{}
}

// WithPayload adds the payload to the update team o k response
func (o *UpdateTeamOK) WithPayload(payload *models.Team) *UpdateTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update team o k response
func (o *UpdateTeamOK) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTeamForbiddenCode is the HTTP code returned for type UpdateTeamForbidden
const UpdateTeamForbiddenCode int = 403

/*
UpdateTeamForbidden User is not authorized

swagger:response updateTeamForbidden
*/
type UpdateTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateTeamForbidden creates UpdateTeamForbidden with default headers values
func NewUpdateTeamForbidden() *UpdateTeamForbidden {

	return &UpdateTeamForbidden{}
}

// WithPayload adds the payload to the update team forbidden response
func (o *UpdateTeamForbidden) WithPayload(payload *models.GeneralError) *UpdateTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update team forbidden response
func (o *UpdateTeamForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTeamNotFoundCode is the HTTP code returned for type UpdateTeamNotFound
const UpdateTeamNotFoundCode int = 404

/*
UpdateTeamNotFound Team not found

swagger:response updateTeamNotFound
*/
type UpdateTeamNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateTeamNotFound creates UpdateTeamNotFound with default headers values
func NewUpdateTeamNotFound() *UpdateTeamNotFound {

	return &UpdateTeamNotFound{}
}

// WithPayload adds the payload to the update team not found response
func (o *UpdateTeamNotFound) WithPayload(payload *models.GeneralError) *UpdateTeamNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update team not found response
func (o *UpdateTeamNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeamNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTeamUnprocessableEntityCode is the HTTP code returned for type UpdateTeamUnprocessableEntity
const UpdateTeamUnprocessableEntityCode int = 422

/*
UpdateTeamUnprocessableEntity Failed to validate request

swagger:response updateTeamUnprocessableEntity
*/
type UpdateTeamUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewUpdateTeamUnprocessableEntity creates UpdateTeamUnprocessableEntity with default headers values
func NewUpdateTeamUnprocessableEntity() *UpdateTeamUnprocessableEntity {

	return &UpdateTeamUnprocessableEntity{}
}

// WithPayload adds the payload to the update team unprocessable entity response
func (o *UpdateTeamUnprocessableEntity) WithPayload(payload *models.ValidationError) *UpdateTeamUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update team unprocessable entity response
func (o *UpdateTeamUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeamUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
UpdateTeamDefault Some error unrelated to the handler

swagger:response updateTeamDefault
*/
type UpdateTeamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateTeamDefault creates UpdateTeamDefault with default headers values
func NewUpdateTeamDefault(code int) *UpdateTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update team default response
func (o *UpdateTeamDefault) WithStatusCode(code int) *UpdateTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update team default response
func (o *UpdateTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update team default response
func (o *UpdateTeamDefault) WithPayload(payload *models.GeneralError) *UpdateTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update team default response
func (o *UpdateTeamDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
