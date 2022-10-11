// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListModUsersOKCode is the HTTP code returned for type ListModUsersOK
const ListModUsersOKCode int = 200

/*
ListModUsersOK A collection of mod users

swagger:response listModUsersOK
*/
type ListModUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserMod `json:"body,omitempty"`
}

// NewListModUsersOK creates ListModUsersOK with default headers values
func NewListModUsersOK() *ListModUsersOK {

	return &ListModUsersOK{}
}

// WithPayload adds the payload to the list mod users o k response
func (o *ListModUsersOK) WithPayload(payload []*models.UserMod) *ListModUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list mod users o k response
func (o *ListModUsersOK) SetPayload(payload []*models.UserMod) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListModUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UserMod, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListModUsersForbiddenCode is the HTTP code returned for type ListModUsersForbidden
const ListModUsersForbiddenCode int = 403

/*
ListModUsersForbidden User is not authorized

swagger:response listModUsersForbidden
*/
type ListModUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListModUsersForbidden creates ListModUsersForbidden with default headers values
func NewListModUsersForbidden() *ListModUsersForbidden {

	return &ListModUsersForbidden{}
}

// WithPayload adds the payload to the list mod users forbidden response
func (o *ListModUsersForbidden) WithPayload(payload *models.GeneralError) *ListModUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list mod users forbidden response
func (o *ListModUsersForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListModUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListModUsersNotFoundCode is the HTTP code returned for type ListModUsersNotFound
const ListModUsersNotFoundCode int = 404

/*
ListModUsersNotFound Mod not found

swagger:response listModUsersNotFound
*/
type ListModUsersNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListModUsersNotFound creates ListModUsersNotFound with default headers values
func NewListModUsersNotFound() *ListModUsersNotFound {

	return &ListModUsersNotFound{}
}

// WithPayload adds the payload to the list mod users not found response
func (o *ListModUsersNotFound) WithPayload(payload *models.GeneralError) *ListModUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list mod users not found response
func (o *ListModUsersNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListModUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListModUsersDefault Some error unrelated to the handler

swagger:response listModUsersDefault
*/
type ListModUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListModUsersDefault creates ListModUsersDefault with default headers values
func NewListModUsersDefault(code int) *ListModUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &ListModUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list mod users default response
func (o *ListModUsersDefault) WithStatusCode(code int) *ListModUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list mod users default response
func (o *ListModUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list mod users default response
func (o *ListModUsersDefault) WithPayload(payload *models.GeneralError) *ListModUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list mod users default response
func (o *ListModUsersDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListModUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
