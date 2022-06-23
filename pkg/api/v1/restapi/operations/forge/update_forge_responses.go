// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UpdateForgeOKCode is the HTTP code returned for type UpdateForgeOK
const UpdateForgeOKCode int = 200

/*UpdateForgeOK Plain success message

swagger:response updateForgeOK
*/
type UpdateForgeOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateForgeOK creates UpdateForgeOK with default headers values
func NewUpdateForgeOK() *UpdateForgeOK {

	return &UpdateForgeOK{}
}

// WithPayload adds the payload to the update forge o k response
func (o *UpdateForgeOK) WithPayload(payload *models.GeneralError) *UpdateForgeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update forge o k response
func (o *UpdateForgeOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateForgeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateForgeForbiddenCode is the HTTP code returned for type UpdateForgeForbidden
const UpdateForgeForbiddenCode int = 403

/*UpdateForgeForbidden User is not authorized

swagger:response updateForgeForbidden
*/
type UpdateForgeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateForgeForbidden creates UpdateForgeForbidden with default headers values
func NewUpdateForgeForbidden() *UpdateForgeForbidden {

	return &UpdateForgeForbidden{}
}

// WithPayload adds the payload to the update forge forbidden response
func (o *UpdateForgeForbidden) WithPayload(payload *models.GeneralError) *UpdateForgeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update forge forbidden response
func (o *UpdateForgeForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateForgeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateForgeServiceUnavailableCode is the HTTP code returned for type UpdateForgeServiceUnavailable
const UpdateForgeServiceUnavailableCode int = 503

/*UpdateForgeServiceUnavailable If remote source is not available

swagger:response updateForgeServiceUnavailable
*/
type UpdateForgeServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateForgeServiceUnavailable creates UpdateForgeServiceUnavailable with default headers values
func NewUpdateForgeServiceUnavailable() *UpdateForgeServiceUnavailable {

	return &UpdateForgeServiceUnavailable{}
}

// WithPayload adds the payload to the update forge service unavailable response
func (o *UpdateForgeServiceUnavailable) WithPayload(payload *models.GeneralError) *UpdateForgeServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update forge service unavailable response
func (o *UpdateForgeServiceUnavailable) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateForgeServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateForgeDefault Some error unrelated to the handler

swagger:response updateForgeDefault
*/
type UpdateForgeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUpdateForgeDefault creates UpdateForgeDefault with default headers values
func NewUpdateForgeDefault(code int) *UpdateForgeDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateForgeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update forge default response
func (o *UpdateForgeDefault) WithStatusCode(code int) *UpdateForgeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update forge default response
func (o *UpdateForgeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update forge default response
func (o *UpdateForgeDefault) WithPayload(payload *models.GeneralError) *UpdateForgeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update forge default response
func (o *UpdateForgeDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateForgeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}