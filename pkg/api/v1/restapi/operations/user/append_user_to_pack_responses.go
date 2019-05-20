// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendUserToPackOKCode is the HTTP code returned for type AppendUserToPackOK
const AppendUserToPackOKCode int = 200

/*AppendUserToPackOK Plain success message

swagger:response appendUserToPackOK
*/
type AppendUserToPackOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToPackOK creates AppendUserToPackOK with default headers values
func NewAppendUserToPackOK() *AppendUserToPackOK {

	return &AppendUserToPackOK{}
}

// WithPayload adds the payload to the append user to pack o k response
func (o *AppendUserToPackOK) WithPayload(payload *models.GeneralError) *AppendUserToPackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to pack o k response
func (o *AppendUserToPackOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToPackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendUserToPackForbiddenCode is the HTTP code returned for type AppendUserToPackForbidden
const AppendUserToPackForbiddenCode int = 403

/*AppendUserToPackForbidden User is not authorized

swagger:response appendUserToPackForbidden
*/
type AppendUserToPackForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToPackForbidden creates AppendUserToPackForbidden with default headers values
func NewAppendUserToPackForbidden() *AppendUserToPackForbidden {

	return &AppendUserToPackForbidden{}
}

// WithPayload adds the payload to the append user to pack forbidden response
func (o *AppendUserToPackForbidden) WithPayload(payload *models.GeneralError) *AppendUserToPackForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to pack forbidden response
func (o *AppendUserToPackForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToPackForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendUserToPackPreconditionFailedCode is the HTTP code returned for type AppendUserToPackPreconditionFailed
const AppendUserToPackPreconditionFailedCode int = 412

/*AppendUserToPackPreconditionFailed Failed to parse request body

swagger:response appendUserToPackPreconditionFailed
*/
type AppendUserToPackPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToPackPreconditionFailed creates AppendUserToPackPreconditionFailed with default headers values
func NewAppendUserToPackPreconditionFailed() *AppendUserToPackPreconditionFailed {

	return &AppendUserToPackPreconditionFailed{}
}

// WithPayload adds the payload to the append user to pack precondition failed response
func (o *AppendUserToPackPreconditionFailed) WithPayload(payload *models.GeneralError) *AppendUserToPackPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to pack precondition failed response
func (o *AppendUserToPackPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToPackPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppendUserToPackUnprocessableEntityCode is the HTTP code returned for type AppendUserToPackUnprocessableEntity
const AppendUserToPackUnprocessableEntityCode int = 422

/*AppendUserToPackUnprocessableEntity Pack is already assigned

swagger:response appendUserToPackUnprocessableEntity
*/
type AppendUserToPackUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToPackUnprocessableEntity creates AppendUserToPackUnprocessableEntity with default headers values
func NewAppendUserToPackUnprocessableEntity() *AppendUserToPackUnprocessableEntity {

	return &AppendUserToPackUnprocessableEntity{}
}

// WithPayload adds the payload to the append user to pack unprocessable entity response
func (o *AppendUserToPackUnprocessableEntity) WithPayload(payload *models.GeneralError) *AppendUserToPackUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to pack unprocessable entity response
func (o *AppendUserToPackUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToPackUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AppendUserToPackDefault Some error unrelated to the handler

swagger:response appendUserToPackDefault
*/
type AppendUserToPackDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewAppendUserToPackDefault creates AppendUserToPackDefault with default headers values
func NewAppendUserToPackDefault(code int) *AppendUserToPackDefault {
	if code <= 0 {
		code = 500
	}

	return &AppendUserToPackDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the append user to pack default response
func (o *AppendUserToPackDefault) WithStatusCode(code int) *AppendUserToPackDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the append user to pack default response
func (o *AppendUserToPackDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the append user to pack default response
func (o *AppendUserToPackDefault) WithPayload(payload *models.GeneralError) *AppendUserToPackDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the append user to pack default response
func (o *AppendUserToPackDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppendUserToPackDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
