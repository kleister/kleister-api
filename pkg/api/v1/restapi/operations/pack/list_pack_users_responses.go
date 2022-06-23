// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListPackUsersOKCode is the HTTP code returned for type ListPackUsersOK
const ListPackUsersOKCode int = 200

/*ListPackUsersOK A collection of pack users

swagger:response listPackUsersOK
*/
type ListPackUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserPack `json:"body,omitempty"`
}

// NewListPackUsersOK creates ListPackUsersOK with default headers values
func NewListPackUsersOK() *ListPackUsersOK {

	return &ListPackUsersOK{}
}

// WithPayload adds the payload to the list pack users o k response
func (o *ListPackUsersOK) WithPayload(payload []*models.UserPack) *ListPackUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list pack users o k response
func (o *ListPackUsersOK) SetPayload(payload []*models.UserPack) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPackUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UserPack, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListPackUsersForbiddenCode is the HTTP code returned for type ListPackUsersForbidden
const ListPackUsersForbiddenCode int = 403

/*ListPackUsersForbidden User is not authorized

swagger:response listPackUsersForbidden
*/
type ListPackUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListPackUsersForbidden creates ListPackUsersForbidden with default headers values
func NewListPackUsersForbidden() *ListPackUsersForbidden {

	return &ListPackUsersForbidden{}
}

// WithPayload adds the payload to the list pack users forbidden response
func (o *ListPackUsersForbidden) WithPayload(payload *models.GeneralError) *ListPackUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list pack users forbidden response
func (o *ListPackUsersForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPackUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPackUsersNotFoundCode is the HTTP code returned for type ListPackUsersNotFound
const ListPackUsersNotFoundCode int = 404

/*ListPackUsersNotFound Pack not found

swagger:response listPackUsersNotFound
*/
type ListPackUsersNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListPackUsersNotFound creates ListPackUsersNotFound with default headers values
func NewListPackUsersNotFound() *ListPackUsersNotFound {

	return &ListPackUsersNotFound{}
}

// WithPayload adds the payload to the list pack users not found response
func (o *ListPackUsersNotFound) WithPayload(payload *models.GeneralError) *ListPackUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list pack users not found response
func (o *ListPackUsersNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPackUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListPackUsersDefault Some error unrelated to the handler

swagger:response listPackUsersDefault
*/
type ListPackUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListPackUsersDefault creates ListPackUsersDefault with default headers values
func NewListPackUsersDefault(code int) *ListPackUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &ListPackUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list pack users default response
func (o *ListPackUsersDefault) WithStatusCode(code int) *ListPackUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list pack users default response
func (o *ListPackUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list pack users default response
func (o *ListPackUsersDefault) WithPayload(payload *models.GeneralError) *ListPackUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list pack users default response
func (o *ListPackUsersDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPackUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}