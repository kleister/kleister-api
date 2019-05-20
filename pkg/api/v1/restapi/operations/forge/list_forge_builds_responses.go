// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListForgeBuildsOKCode is the HTTP code returned for type ListForgeBuildsOK
const ListForgeBuildsOKCode int = 200

/*ListForgeBuildsOK A collection of assigned builds

swagger:response listForgeBuildsOK
*/
type ListForgeBuildsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Build `json:"body,omitempty"`
}

// NewListForgeBuildsOK creates ListForgeBuildsOK with default headers values
func NewListForgeBuildsOK() *ListForgeBuildsOK {

	return &ListForgeBuildsOK{}
}

// WithPayload adds the payload to the list forge builds o k response
func (o *ListForgeBuildsOK) WithPayload(payload []*models.Build) *ListForgeBuildsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list forge builds o k response
func (o *ListForgeBuildsOK) SetPayload(payload []*models.Build) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListForgeBuildsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Build, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListForgeBuildsForbiddenCode is the HTTP code returned for type ListForgeBuildsForbidden
const ListForgeBuildsForbiddenCode int = 403

/*ListForgeBuildsForbidden User is not authorized

swagger:response listForgeBuildsForbidden
*/
type ListForgeBuildsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListForgeBuildsForbidden creates ListForgeBuildsForbidden with default headers values
func NewListForgeBuildsForbidden() *ListForgeBuildsForbidden {

	return &ListForgeBuildsForbidden{}
}

// WithPayload adds the payload to the list forge builds forbidden response
func (o *ListForgeBuildsForbidden) WithPayload(payload *models.GeneralError) *ListForgeBuildsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list forge builds forbidden response
func (o *ListForgeBuildsForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListForgeBuildsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListForgeBuildsDefault Some error unrelated to the handler

swagger:response listForgeBuildsDefault
*/
type ListForgeBuildsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListForgeBuildsDefault creates ListForgeBuildsDefault with default headers values
func NewListForgeBuildsDefault(code int) *ListForgeBuildsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListForgeBuildsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list forge builds default response
func (o *ListForgeBuildsDefault) WithStatusCode(code int) *ListForgeBuildsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list forge builds default response
func (o *ListForgeBuildsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list forge builds default response
func (o *ListForgeBuildsDefault) WithPayload(payload *models.GeneralError) *ListForgeBuildsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list forge builds default response
func (o *ListForgeBuildsDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListForgeBuildsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
