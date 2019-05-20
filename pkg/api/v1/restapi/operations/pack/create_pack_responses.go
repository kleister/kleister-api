// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// CreatePackOKCode is the HTTP code returned for type CreatePackOK
const CreatePackOKCode int = 200

/*CreatePackOK The created pack data

swagger:response createPackOK
*/
type CreatePackOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pack `json:"body,omitempty"`
}

// NewCreatePackOK creates CreatePackOK with default headers values
func NewCreatePackOK() *CreatePackOK {

	return &CreatePackOK{}
}

// WithPayload adds the payload to the create pack o k response
func (o *CreatePackOK) WithPayload(payload *models.Pack) *CreatePackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pack o k response
func (o *CreatePackOK) SetPayload(payload *models.Pack) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePackForbiddenCode is the HTTP code returned for type CreatePackForbidden
const CreatePackForbiddenCode int = 403

/*CreatePackForbidden User is not authorized

swagger:response createPackForbidden
*/
type CreatePackForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreatePackForbidden creates CreatePackForbidden with default headers values
func NewCreatePackForbidden() *CreatePackForbidden {

	return &CreatePackForbidden{}
}

// WithPayload adds the payload to the create pack forbidden response
func (o *CreatePackForbidden) WithPayload(payload *models.GeneralError) *CreatePackForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pack forbidden response
func (o *CreatePackForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePackForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePackPreconditionFailedCode is the HTTP code returned for type CreatePackPreconditionFailed
const CreatePackPreconditionFailedCode int = 412

/*CreatePackPreconditionFailed Failed to parse request body

swagger:response createPackPreconditionFailed
*/
type CreatePackPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreatePackPreconditionFailed creates CreatePackPreconditionFailed with default headers values
func NewCreatePackPreconditionFailed() *CreatePackPreconditionFailed {

	return &CreatePackPreconditionFailed{}
}

// WithPayload adds the payload to the create pack precondition failed response
func (o *CreatePackPreconditionFailed) WithPayload(payload *models.GeneralError) *CreatePackPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pack precondition failed response
func (o *CreatePackPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePackPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePackUnprocessableEntityCode is the HTTP code returned for type CreatePackUnprocessableEntity
const CreatePackUnprocessableEntityCode int = 422

/*CreatePackUnprocessableEntity Failed to validate request

swagger:response createPackUnprocessableEntity
*/
type CreatePackUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewCreatePackUnprocessableEntity creates CreatePackUnprocessableEntity with default headers values
func NewCreatePackUnprocessableEntity() *CreatePackUnprocessableEntity {

	return &CreatePackUnprocessableEntity{}
}

// WithPayload adds the payload to the create pack unprocessable entity response
func (o *CreatePackUnprocessableEntity) WithPayload(payload *models.ValidationError) *CreatePackUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pack unprocessable entity response
func (o *CreatePackUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePackUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreatePackDefault Some error unrelated to the handler

swagger:response createPackDefault
*/
type CreatePackDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreatePackDefault creates CreatePackDefault with default headers values
func NewCreatePackDefault(code int) *CreatePackDefault {
	if code <= 0 {
		code = 500
	}

	return &CreatePackDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create pack default response
func (o *CreatePackDefault) WithStatusCode(code int) *CreatePackDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create pack default response
func (o *CreatePackDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create pack default response
func (o *CreatePackDefault) WithPayload(payload *models.GeneralError) *CreatePackDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create pack default response
func (o *CreatePackDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePackDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
