// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// PermitUserPackOKCode is the HTTP code returned for type PermitUserPackOK
const PermitUserPackOKCode int = 200

/*PermitUserPackOK Plain success message

swagger:response permitUserPackOK
*/
type PermitUserPackOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserPackOK creates PermitUserPackOK with default headers values
func NewPermitUserPackOK() *PermitUserPackOK {

	return &PermitUserPackOK{}
}

// WithPayload adds the payload to the permit user pack o k response
func (o *PermitUserPackOK) WithPayload(payload *models.GeneralError) *PermitUserPackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user pack o k response
func (o *PermitUserPackOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserPackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitUserPackForbiddenCode is the HTTP code returned for type PermitUserPackForbidden
const PermitUserPackForbiddenCode int = 403

/*PermitUserPackForbidden User is not authorized

swagger:response permitUserPackForbidden
*/
type PermitUserPackForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserPackForbidden creates PermitUserPackForbidden with default headers values
func NewPermitUserPackForbidden() *PermitUserPackForbidden {

	return &PermitUserPackForbidden{}
}

// WithPayload adds the payload to the permit user pack forbidden response
func (o *PermitUserPackForbidden) WithPayload(payload *models.GeneralError) *PermitUserPackForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user pack forbidden response
func (o *PermitUserPackForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserPackForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitUserPackPreconditionFailedCode is the HTTP code returned for type PermitUserPackPreconditionFailed
const PermitUserPackPreconditionFailedCode int = 412

/*PermitUserPackPreconditionFailed Failed to parse request body

swagger:response permitUserPackPreconditionFailed
*/
type PermitUserPackPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserPackPreconditionFailed creates PermitUserPackPreconditionFailed with default headers values
func NewPermitUserPackPreconditionFailed() *PermitUserPackPreconditionFailed {

	return &PermitUserPackPreconditionFailed{}
}

// WithPayload adds the payload to the permit user pack precondition failed response
func (o *PermitUserPackPreconditionFailed) WithPayload(payload *models.GeneralError) *PermitUserPackPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user pack precondition failed response
func (o *PermitUserPackPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserPackPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PermitUserPackUnprocessableEntityCode is the HTTP code returned for type PermitUserPackUnprocessableEntity
const PermitUserPackUnprocessableEntityCode int = 422

/*PermitUserPackUnprocessableEntity Pack is not assigned

swagger:response permitUserPackUnprocessableEntity
*/
type PermitUserPackUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserPackUnprocessableEntity creates PermitUserPackUnprocessableEntity with default headers values
func NewPermitUserPackUnprocessableEntity() *PermitUserPackUnprocessableEntity {

	return &PermitUserPackUnprocessableEntity{}
}

// WithPayload adds the payload to the permit user pack unprocessable entity response
func (o *PermitUserPackUnprocessableEntity) WithPayload(payload *models.GeneralError) *PermitUserPackUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user pack unprocessable entity response
func (o *PermitUserPackUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserPackUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PermitUserPackDefault Some error unrelated to the handler

swagger:response permitUserPackDefault
*/
type PermitUserPackDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPermitUserPackDefault creates PermitUserPackDefault with default headers values
func NewPermitUserPackDefault(code int) *PermitUserPackDefault {
	if code <= 0 {
		code = 500
	}

	return &PermitUserPackDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the permit user pack default response
func (o *PermitUserPackDefault) WithStatusCode(code int) *PermitUserPackDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the permit user pack default response
func (o *PermitUserPackDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the permit user pack default response
func (o *PermitUserPackDefault) WithPayload(payload *models.GeneralError) *PermitUserPackDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the permit user pack default response
func (o *PermitUserPackDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PermitUserPackDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
