// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteUserFromPackOKCode is the HTTP code returned for type DeleteUserFromPackOK
const DeleteUserFromPackOKCode int = 200

/*DeleteUserFromPackOK Plain success message

swagger:response deleteUserFromPackOK
*/
type DeleteUserFromPackOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromPackOK creates DeleteUserFromPackOK with default headers values
func NewDeleteUserFromPackOK() *DeleteUserFromPackOK {

	return &DeleteUserFromPackOK{}
}

// WithPayload adds the payload to the delete user from pack o k response
func (o *DeleteUserFromPackOK) WithPayload(payload *models.GeneralError) *DeleteUserFromPackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from pack o k response
func (o *DeleteUserFromPackOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromPackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserFromPackForbiddenCode is the HTTP code returned for type DeleteUserFromPackForbidden
const DeleteUserFromPackForbiddenCode int = 403

/*DeleteUserFromPackForbidden User is not authorized

swagger:response deleteUserFromPackForbidden
*/
type DeleteUserFromPackForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromPackForbidden creates DeleteUserFromPackForbidden with default headers values
func NewDeleteUserFromPackForbidden() *DeleteUserFromPackForbidden {

	return &DeleteUserFromPackForbidden{}
}

// WithPayload adds the payload to the delete user from pack forbidden response
func (o *DeleteUserFromPackForbidden) WithPayload(payload *models.GeneralError) *DeleteUserFromPackForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from pack forbidden response
func (o *DeleteUserFromPackForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromPackForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserFromPackNotFoundCode is the HTTP code returned for type DeleteUserFromPackNotFound
const DeleteUserFromPackNotFoundCode int = 404

/*DeleteUserFromPackNotFound User or pack not found

swagger:response deleteUserFromPackNotFound
*/
type DeleteUserFromPackNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromPackNotFound creates DeleteUserFromPackNotFound with default headers values
func NewDeleteUserFromPackNotFound() *DeleteUserFromPackNotFound {

	return &DeleteUserFromPackNotFound{}
}

// WithPayload adds the payload to the delete user from pack not found response
func (o *DeleteUserFromPackNotFound) WithPayload(payload *models.GeneralError) *DeleteUserFromPackNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from pack not found response
func (o *DeleteUserFromPackNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromPackNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserFromPackPreconditionFailedCode is the HTTP code returned for type DeleteUserFromPackPreconditionFailed
const DeleteUserFromPackPreconditionFailedCode int = 412

/*DeleteUserFromPackPreconditionFailed Pack is not assigned

swagger:response deleteUserFromPackPreconditionFailed
*/
type DeleteUserFromPackPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromPackPreconditionFailed creates DeleteUserFromPackPreconditionFailed with default headers values
func NewDeleteUserFromPackPreconditionFailed() *DeleteUserFromPackPreconditionFailed {

	return &DeleteUserFromPackPreconditionFailed{}
}

// WithPayload adds the payload to the delete user from pack precondition failed response
func (o *DeleteUserFromPackPreconditionFailed) WithPayload(payload *models.GeneralError) *DeleteUserFromPackPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from pack precondition failed response
func (o *DeleteUserFromPackPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromPackPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteUserFromPackDefault Some error unrelated to the handler

swagger:response deleteUserFromPackDefault
*/
type DeleteUserFromPackDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromPackDefault creates DeleteUserFromPackDefault with default headers values
func NewDeleteUserFromPackDefault(code int) *DeleteUserFromPackDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUserFromPackDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete user from pack default response
func (o *DeleteUserFromPackDefault) WithStatusCode(code int) *DeleteUserFromPackDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete user from pack default response
func (o *DeleteUserFromPackDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete user from pack default response
func (o *DeleteUserFromPackDefault) WithPayload(payload *models.GeneralError) *DeleteUserFromPackDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from pack default response
func (o *DeleteUserFromPackDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromPackDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}