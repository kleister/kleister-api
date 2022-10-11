// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ShowTeamOKCode is the HTTP code returned for type ShowTeamOK
const ShowTeamOKCode int = 200

/*
ShowTeamOK The fetched team details

swagger:response showTeamOK
*/
type ShowTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.Team `json:"body,omitempty"`
}

// NewShowTeamOK creates ShowTeamOK with default headers values
func NewShowTeamOK() *ShowTeamOK {

	return &ShowTeamOK{}
}

// WithPayload adds the payload to the show team o k response
func (o *ShowTeamOK) WithPayload(payload *models.Team) *ShowTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show team o k response
func (o *ShowTeamOK) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowTeamForbiddenCode is the HTTP code returned for type ShowTeamForbidden
const ShowTeamForbiddenCode int = 403

/*
ShowTeamForbidden User is not authorized

swagger:response showTeamForbidden
*/
type ShowTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewShowTeamForbidden creates ShowTeamForbidden with default headers values
func NewShowTeamForbidden() *ShowTeamForbidden {

	return &ShowTeamForbidden{}
}

// WithPayload adds the payload to the show team forbidden response
func (o *ShowTeamForbidden) WithPayload(payload *models.GeneralError) *ShowTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show team forbidden response
func (o *ShowTeamForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowTeamNotFoundCode is the HTTP code returned for type ShowTeamNotFound
const ShowTeamNotFoundCode int = 404

/*
ShowTeamNotFound Team not found

swagger:response showTeamNotFound
*/
type ShowTeamNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewShowTeamNotFound creates ShowTeamNotFound with default headers values
func NewShowTeamNotFound() *ShowTeamNotFound {

	return &ShowTeamNotFound{}
}

// WithPayload adds the payload to the show team not found response
func (o *ShowTeamNotFound) WithPayload(payload *models.GeneralError) *ShowTeamNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show team not found response
func (o *ShowTeamNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowTeamNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ShowTeamDefault Some error unrelated to the handler

swagger:response showTeamDefault
*/
type ShowTeamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewShowTeamDefault creates ShowTeamDefault with default headers values
func NewShowTeamDefault(code int) *ShowTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &ShowTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the show team default response
func (o *ShowTeamDefault) WithStatusCode(code int) *ShowTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the show team default response
func (o *ShowTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the show team default response
func (o *ShowTeamDefault) WithPayload(payload *models.GeneralError) *ShowTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show team default response
func (o *ShowTeamDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
