// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListTeamUsersOKCode is the HTTP code returned for type ListTeamUsersOK
const ListTeamUsersOKCode int = 200

/*ListTeamUsersOK A collection of team users

swagger:response listTeamUsersOK
*/
type ListTeamUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TeamUser `json:"body,omitempty"`
}

// NewListTeamUsersOK creates ListTeamUsersOK with default headers values
func NewListTeamUsersOK() *ListTeamUsersOK {

	return &ListTeamUsersOK{}
}

// WithPayload adds the payload to the list team users o k response
func (o *ListTeamUsersOK) WithPayload(payload []*models.TeamUser) *ListTeamUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team users o k response
func (o *ListTeamUsersOK) SetPayload(payload []*models.TeamUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.TeamUser, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListTeamUsersForbiddenCode is the HTTP code returned for type ListTeamUsersForbidden
const ListTeamUsersForbiddenCode int = 403

/*ListTeamUsersForbidden User is not authorized

swagger:response listTeamUsersForbidden
*/
type ListTeamUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamUsersForbidden creates ListTeamUsersForbidden with default headers values
func NewListTeamUsersForbidden() *ListTeamUsersForbidden {

	return &ListTeamUsersForbidden{}
}

// WithPayload adds the payload to the list team users forbidden response
func (o *ListTeamUsersForbidden) WithPayload(payload *models.GeneralError) *ListTeamUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team users forbidden response
func (o *ListTeamUsersForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListTeamUsersNotFoundCode is the HTTP code returned for type ListTeamUsersNotFound
const ListTeamUsersNotFoundCode int = 404

/*ListTeamUsersNotFound Team not found

swagger:response listTeamUsersNotFound
*/
type ListTeamUsersNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamUsersNotFound creates ListTeamUsersNotFound with default headers values
func NewListTeamUsersNotFound() *ListTeamUsersNotFound {

	return &ListTeamUsersNotFound{}
}

// WithPayload adds the payload to the list team users not found response
func (o *ListTeamUsersNotFound) WithPayload(payload *models.GeneralError) *ListTeamUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team users not found response
func (o *ListTeamUsersNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListTeamUsersDefault Some error unrelated to the handler

swagger:response listTeamUsersDefault
*/
type ListTeamUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamUsersDefault creates ListTeamUsersDefault with default headers values
func NewListTeamUsersDefault(code int) *ListTeamUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &ListTeamUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list team users default response
func (o *ListTeamUsersDefault) WithStatusCode(code int) *ListTeamUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list team users default response
func (o *ListTeamUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list team users default response
func (o *ListTeamUsersDefault) WithPayload(payload *models.GeneralError) *ListTeamUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team users default response
func (o *ListTeamUsersDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}