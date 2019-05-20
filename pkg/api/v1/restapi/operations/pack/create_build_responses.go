// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// CreateBuildOKCode is the HTTP code returned for type CreateBuildOK
const CreateBuildOKCode int = 200

/*CreateBuildOK The created build data

swagger:response createBuildOK
*/
type CreateBuildOK struct {

	/*
	  In: Body
	*/
	Payload *models.Build `json:"body,omitempty"`
}

// NewCreateBuildOK creates CreateBuildOK with default headers values
func NewCreateBuildOK() *CreateBuildOK {

	return &CreateBuildOK{}
}

// WithPayload adds the payload to the create build o k response
func (o *CreateBuildOK) WithPayload(payload *models.Build) *CreateBuildOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create build o k response
func (o *CreateBuildOK) SetPayload(payload *models.Build) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBuildOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBuildForbiddenCode is the HTTP code returned for type CreateBuildForbidden
const CreateBuildForbiddenCode int = 403

/*CreateBuildForbidden User is not authorized

swagger:response createBuildForbidden
*/
type CreateBuildForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreateBuildForbidden creates CreateBuildForbidden with default headers values
func NewCreateBuildForbidden() *CreateBuildForbidden {

	return &CreateBuildForbidden{}
}

// WithPayload adds the payload to the create build forbidden response
func (o *CreateBuildForbidden) WithPayload(payload *models.GeneralError) *CreateBuildForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create build forbidden response
func (o *CreateBuildForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBuildForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBuildPreconditionFailedCode is the HTTP code returned for type CreateBuildPreconditionFailed
const CreateBuildPreconditionFailedCode int = 412

/*CreateBuildPreconditionFailed Failed to parse request body

swagger:response createBuildPreconditionFailed
*/
type CreateBuildPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreateBuildPreconditionFailed creates CreateBuildPreconditionFailed with default headers values
func NewCreateBuildPreconditionFailed() *CreateBuildPreconditionFailed {

	return &CreateBuildPreconditionFailed{}
}

// WithPayload adds the payload to the create build precondition failed response
func (o *CreateBuildPreconditionFailed) WithPayload(payload *models.GeneralError) *CreateBuildPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create build precondition failed response
func (o *CreateBuildPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBuildPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBuildUnprocessableEntityCode is the HTTP code returned for type CreateBuildUnprocessableEntity
const CreateBuildUnprocessableEntityCode int = 422

/*CreateBuildUnprocessableEntity Failed to validate request

swagger:response createBuildUnprocessableEntity
*/
type CreateBuildUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewCreateBuildUnprocessableEntity creates CreateBuildUnprocessableEntity with default headers values
func NewCreateBuildUnprocessableEntity() *CreateBuildUnprocessableEntity {

	return &CreateBuildUnprocessableEntity{}
}

// WithPayload adds the payload to the create build unprocessable entity response
func (o *CreateBuildUnprocessableEntity) WithPayload(payload *models.ValidationError) *CreateBuildUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create build unprocessable entity response
func (o *CreateBuildUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBuildUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateBuildDefault Some error unrelated to the handler

swagger:response createBuildDefault
*/
type CreateBuildDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewCreateBuildDefault creates CreateBuildDefault with default headers values
func NewCreateBuildDefault(code int) *CreateBuildDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateBuildDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create build default response
func (o *CreateBuildDefault) WithStatusCode(code int) *CreateBuildDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create build default response
func (o *CreateBuildDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create build default response
func (o *CreateBuildDefault) WithPayload(payload *models.GeneralError) *CreateBuildDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create build default response
func (o *CreateBuildDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBuildDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
