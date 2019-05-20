// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendPackToUserOKCode is the HTTP code returned for type AppendPackToUserOK
const AppendPackToUserOKCode int = 200

/*AppendPackToUserOK Plain success message

swagger:response appendPackToUserOK
*/
type AppendPackToUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendPackToUserOK creates AppendPackToUserOK with default headers values
func NewAppendPackToUserOK() *AppendPackToUserOK {

	return &AppendPackToUserOK{}
}

// WithPayload adds the payload to the append pack to user o k response
func (o *AppendPackToUserOK) WithPayload(payload *models.GeneralError) *AppendPackToUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append pack to user o k response
func (o *AppendPackToUserOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendPackToUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendPackToUserForbiddenCode is the HTTP code returned for type AppendPackToUserForbidden
const AppendPackToUserForbiddenCode int = 403

/*AppendPackToUserForbidden User is not authorized

swagger:response appendPackToUserForbidden
*/
type AppendPackToUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendPackToUserForbidden creates AppendPackToUserForbidden with default headers values
func NewAppendPackToUserForbidden() *AppendPackToUserForbidden {

	return &AppendPackToUserForbidden{}
}

// WithPayload adds the payload to the append pack to user forbidden response
func (o *AppendPackToUserForbidden) WithPayload(payload *models.GeneralError) *AppendPackToUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append pack to user forbidden response
func (o *AppendPackToUserForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendPackToUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendPackToUserPreconditionFailedCode is the HTTP code returned for type AppendPackToUserPreconditionFailed
const AppendPackToUserPreconditionFailedCode int = 412

/*AppendPackToUserPreconditionFailed Failed to parse request body

swagger:response appendPackToUserPreconditionFailed
*/
type AppendPackToUserPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendPackToUserPreconditionFailed creates AppendPackToUserPreconditionFailed with default headers values
func NewAppendPackToUserPreconditionFailed() *AppendPackToUserPreconditionFailed {

	return &AppendPackToUserPreconditionFailed{}
}

// WithPayload adds the payload to the append pack to user precondition failed response
func (o *AppendPackToUserPreconditionFailed) WithPayload(payload *models.GeneralError) *AppendPackToUserPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append pack to user precondition failed response
func (o *AppendPackToUserPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendPackToUserPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendPackToUserUnprocessableEntityCode is the HTTP code returned for type AppendPackToUserUnprocessableEntity
const AppendPackToUserUnprocessableEntityCode int = 422

/*AppendPackToUserUnprocessableEntity User is already assigned

swagger:response appendPackToUserUnprocessableEntity
*/
type AppendPackToUserUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendPackToUserUnprocessableEntity creates AppendPackToUserUnprocessableEntity with default headers values
func NewAppendPackToUserUnprocessableEntity() *AppendPackToUserUnprocessableEntity {

	return &AppendPackToUserUnprocessableEntity{}
}

// WithPayload adds the payload to the append pack to user unprocessable entity response
func (o *AppendPackToUserUnprocessableEntity) WithPayload(payload *models.GeneralError) *AppendPackToUserUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append pack to user unprocessable entity response
func (o *AppendPackToUserUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendPackToUserUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AppendPackToUserDefault Some error unrelated to the handler

swagger:response appendPackToUserDefault
*/
type AppendPackToUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendPackToUserDefault creates AppendPackToUserDefault with default headers values
func NewAppendPackToUserDefault(code int) *AppendPackToUserDefault {
	if code <= 0 {
		code = 500
	}

	return &AppendPackToUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the append pack to user default response
func (o *AppendPackToUserDefault) WithStatusCode(code int) *AppendPackToUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the append pack to user default response
func (o *AppendPackToUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the append pack to user default response
func (o *AppendPackToUserDefault) WithPayload(payload *models.GeneralError) *AppendPackToUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append pack to user default response
func (o *AppendPackToUserDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendPackToUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
