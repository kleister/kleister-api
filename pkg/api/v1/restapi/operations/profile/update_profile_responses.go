// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UpdateProfileOKCode is the HTTP code returned for type UpdateProfileOK
const UpdateProfileOKCode int = 200

/*UpdateProfileOK The updated profile data

swagger:response updateProfileOK
*/
type UpdateProfileOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewUpdateProfileOK creates UpdateProfileOK with default headers values
func NewUpdateProfileOK() *UpdateProfileOK {

	return &UpdateProfileOK{}
}

// WithPayload adds the payload to the update profile o k response
func (o *UpdateProfileOK) WithPayload(payload *models.Profile) *UpdateProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update profile o k response
func (o *UpdateProfileOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateProfileForbiddenCode is the HTTP code returned for type UpdateProfileForbidden
const UpdateProfileForbiddenCode int = 403

/*UpdateProfileForbidden User is not authorized

swagger:response updateProfileForbidden
*/
type UpdateProfileForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateProfileForbidden creates UpdateProfileForbidden with default headers values
func NewUpdateProfileForbidden() *UpdateProfileForbidden {

	return &UpdateProfileForbidden{}
}

// WithPayload adds the payload to the update profile forbidden response
func (o *UpdateProfileForbidden) WithPayload(payload *models.GeneralError) *UpdateProfileForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update profile forbidden response
func (o *UpdateProfileForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProfileForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateProfilePreconditionFailedCode is the HTTP code returned for type UpdateProfilePreconditionFailed
const UpdateProfilePreconditionFailedCode int = 412

/*UpdateProfilePreconditionFailed Failed to parse request body

swagger:response updateProfilePreconditionFailed
*/
type UpdateProfilePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateProfilePreconditionFailed creates UpdateProfilePreconditionFailed with default headers values
func NewUpdateProfilePreconditionFailed() *UpdateProfilePreconditionFailed {

	return &UpdateProfilePreconditionFailed{}
}

// WithPayload adds the payload to the update profile precondition failed response
func (o *UpdateProfilePreconditionFailed) WithPayload(payload *models.GeneralError) *UpdateProfilePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update profile precondition failed response
func (o *UpdateProfilePreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProfilePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateProfileUnprocessableEntityCode is the HTTP code returned for type UpdateProfileUnprocessableEntity
const UpdateProfileUnprocessableEntityCode int = 422

/*UpdateProfileUnprocessableEntity Failed to validate request

swagger:response updateProfileUnprocessableEntity
*/
type UpdateProfileUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewUpdateProfileUnprocessableEntity creates UpdateProfileUnprocessableEntity with default headers values
func NewUpdateProfileUnprocessableEntity() *UpdateProfileUnprocessableEntity {

	return &UpdateProfileUnprocessableEntity{}
}

// WithPayload adds the payload to the update profile unprocessable entity response
func (o *UpdateProfileUnprocessableEntity) WithPayload(payload *models.ValidationError) *UpdateProfileUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update profile unprocessable entity response
func (o *UpdateProfileUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProfileUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateProfileDefault Some error unrelated to the handler

swagger:response updateProfileDefault
*/
type UpdateProfileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateProfileDefault creates UpdateProfileDefault with default headers values
func NewUpdateProfileDefault(code int) *UpdateProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update profile default response
func (o *UpdateProfileDefault) WithStatusCode(code int) *UpdateProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update profile default response
func (o *UpdateProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update profile default response
func (o *UpdateProfileDefault) WithPayload(payload *models.GeneralError) *UpdateProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update profile default response
func (o *UpdateProfileDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
