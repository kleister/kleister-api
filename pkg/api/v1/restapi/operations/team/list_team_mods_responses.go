// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListTeamModsOKCode is the HTTP code returned for type ListTeamModsOK
const ListTeamModsOKCode int = 200

/*
ListTeamModsOK A collection of team mods

swagger:response listTeamModsOK
*/
type ListTeamModsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TeamMod `json:"body,omitempty"`
}

// NewListTeamModsOK creates ListTeamModsOK with default headers values
func NewListTeamModsOK() *ListTeamModsOK {

	return &ListTeamModsOK{}
}

// WithPayload adds the payload to the list team mods o k response
func (o *ListTeamModsOK) WithPayload(payload []*models.TeamMod) *ListTeamModsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team mods o k response
func (o *ListTeamModsOK) SetPayload(payload []*models.TeamMod) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamModsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.TeamMod, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListTeamModsForbiddenCode is the HTTP code returned for type ListTeamModsForbidden
const ListTeamModsForbiddenCode int = 403

/*
ListTeamModsForbidden User is not authorized

swagger:response listTeamModsForbidden
*/
type ListTeamModsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamModsForbidden creates ListTeamModsForbidden with default headers values
func NewListTeamModsForbidden() *ListTeamModsForbidden {

	return &ListTeamModsForbidden{}
}

// WithPayload adds the payload to the list team mods forbidden response
func (o *ListTeamModsForbidden) WithPayload(payload *models.GeneralError) *ListTeamModsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team mods forbidden response
func (o *ListTeamModsForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamModsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListTeamModsNotFoundCode is the HTTP code returned for type ListTeamModsNotFound
const ListTeamModsNotFoundCode int = 404

/*
ListTeamModsNotFound Team not found

swagger:response listTeamModsNotFound
*/
type ListTeamModsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamModsNotFound creates ListTeamModsNotFound with default headers values
func NewListTeamModsNotFound() *ListTeamModsNotFound {

	return &ListTeamModsNotFound{}
}

// WithPayload adds the payload to the list team mods not found response
func (o *ListTeamModsNotFound) WithPayload(payload *models.GeneralError) *ListTeamModsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team mods not found response
func (o *ListTeamModsNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamModsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListTeamModsDefault Some error unrelated to the handler

swagger:response listTeamModsDefault
*/
type ListTeamModsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamModsDefault creates ListTeamModsDefault with default headers values
func NewListTeamModsDefault(code int) *ListTeamModsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListTeamModsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list team mods default response
func (o *ListTeamModsDefault) WithStatusCode(code int) *ListTeamModsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list team mods default response
func (o *ListTeamModsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list team mods default response
func (o *ListTeamModsDefault) WithPayload(payload *models.GeneralError) *ListTeamModsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team mods default response
func (o *ListTeamModsDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamModsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
