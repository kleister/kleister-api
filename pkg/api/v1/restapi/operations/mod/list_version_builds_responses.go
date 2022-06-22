// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListVersionBuildsOKCode is the HTTP code returned for type ListVersionBuildsOK
const ListVersionBuildsOKCode int = 200

/*ListVersionBuildsOK A collection of version builds

swagger:response listVersionBuildsOK
*/
type ListVersionBuildsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.BuildVersion `json:"body,omitempty"`
}

// NewListVersionBuildsOK creates ListVersionBuildsOK with default headers values
func NewListVersionBuildsOK() *ListVersionBuildsOK {

	return &ListVersionBuildsOK{}
}

// WithPayload adds the payload to the list version builds o k response
func (o *ListVersionBuildsOK) WithPayload(payload []*models.BuildVersion) *ListVersionBuildsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list version builds o k response
func (o *ListVersionBuildsOK) SetPayload(payload []*models.BuildVersion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionBuildsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.BuildVersion, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListVersionBuildsForbiddenCode is the HTTP code returned for type ListVersionBuildsForbidden
const ListVersionBuildsForbiddenCode int = 403

/*ListVersionBuildsForbidden User is not authorized

swagger:response listVersionBuildsForbidden
*/
type ListVersionBuildsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListVersionBuildsForbidden creates ListVersionBuildsForbidden with default headers values
func NewListVersionBuildsForbidden() *ListVersionBuildsForbidden {

	return &ListVersionBuildsForbidden{}
}

// WithPayload adds the payload to the list version builds forbidden response
func (o *ListVersionBuildsForbidden) WithPayload(payload *models.GeneralError) *ListVersionBuildsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list version builds forbidden response
func (o *ListVersionBuildsForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionBuildsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListVersionBuildsNotFoundCode is the HTTP code returned for type ListVersionBuildsNotFound
const ListVersionBuildsNotFoundCode int = 404

/*ListVersionBuildsNotFound Version or mod not found

swagger:response listVersionBuildsNotFound
*/
type ListVersionBuildsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListVersionBuildsNotFound creates ListVersionBuildsNotFound with default headers values
func NewListVersionBuildsNotFound() *ListVersionBuildsNotFound {

	return &ListVersionBuildsNotFound{}
}

// WithPayload adds the payload to the list version builds not found response
func (o *ListVersionBuildsNotFound) WithPayload(payload *models.GeneralError) *ListVersionBuildsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list version builds not found response
func (o *ListVersionBuildsNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionBuildsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListVersionBuildsDefault Some error unrelated to the handler

swagger:response listVersionBuildsDefault
*/
type ListVersionBuildsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListVersionBuildsDefault creates ListVersionBuildsDefault with default headers values
func NewListVersionBuildsDefault(code int) *ListVersionBuildsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListVersionBuildsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list version builds default response
func (o *ListVersionBuildsDefault) WithStatusCode(code int) *ListVersionBuildsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list version builds default response
func (o *ListVersionBuildsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list version builds default response
func (o *ListVersionBuildsDefault) WithPayload(payload *models.GeneralError) *ListVersionBuildsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list version builds default response
func (o *ListVersionBuildsDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionBuildsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
