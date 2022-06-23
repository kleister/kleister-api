// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListTeamPacksOKCode is the HTTP code returned for type ListTeamPacksOK
const ListTeamPacksOKCode int = 200

/*ListTeamPacksOK A collection of team packs

swagger:response listTeamPacksOK
*/
type ListTeamPacksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TeamPack `json:"body,omitempty"`
}

// NewListTeamPacksOK creates ListTeamPacksOK with default headers values
func NewListTeamPacksOK() *ListTeamPacksOK {

	return &ListTeamPacksOK{}
}

// WithPayload adds the payload to the list team packs o k response
func (o *ListTeamPacksOK) WithPayload(payload []*models.TeamPack) *ListTeamPacksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team packs o k response
func (o *ListTeamPacksOK) SetPayload(payload []*models.TeamPack) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamPacksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.TeamPack, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListTeamPacksForbiddenCode is the HTTP code returned for type ListTeamPacksForbidden
const ListTeamPacksForbiddenCode int = 403

/*ListTeamPacksForbidden User is not authorized

swagger:response listTeamPacksForbidden
*/
type ListTeamPacksForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamPacksForbidden creates ListTeamPacksForbidden with default headers values
func NewListTeamPacksForbidden() *ListTeamPacksForbidden {

	return &ListTeamPacksForbidden{}
}

// WithPayload adds the payload to the list team packs forbidden response
func (o *ListTeamPacksForbidden) WithPayload(payload *models.GeneralError) *ListTeamPacksForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team packs forbidden response
func (o *ListTeamPacksForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamPacksForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListTeamPacksNotFoundCode is the HTTP code returned for type ListTeamPacksNotFound
const ListTeamPacksNotFoundCode int = 404

/*ListTeamPacksNotFound Team not found

swagger:response listTeamPacksNotFound
*/
type ListTeamPacksNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamPacksNotFound creates ListTeamPacksNotFound with default headers values
func NewListTeamPacksNotFound() *ListTeamPacksNotFound {

	return &ListTeamPacksNotFound{}
}

// WithPayload adds the payload to the list team packs not found response
func (o *ListTeamPacksNotFound) WithPayload(payload *models.GeneralError) *ListTeamPacksNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team packs not found response
func (o *ListTeamPacksNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamPacksNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListTeamPacksDefault Some error unrelated to the handler

swagger:response listTeamPacksDefault
*/
type ListTeamPacksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListTeamPacksDefault creates ListTeamPacksDefault with default headers values
func NewListTeamPacksDefault(code int) *ListTeamPacksDefault {
	if code <= 0 {
		code = 500
	}

	return &ListTeamPacksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list team packs default response
func (o *ListTeamPacksDefault) WithStatusCode(code int) *ListTeamPacksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list team packs default response
func (o *ListTeamPacksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list team packs default response
func (o *ListTeamPacksDefault) WithPayload(payload *models.GeneralError) *ListTeamPacksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list team packs default response
func (o *ListTeamPacksDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTeamPacksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}