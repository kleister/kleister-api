// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListUserModsOKCode is the HTTP code returned for type ListUserModsOK
const ListUserModsOKCode int = 200

/*ListUserModsOK A collection of user mods

swagger:response listUserModsOK
*/
type ListUserModsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserMod `json:"body,omitempty"`
}

// NewListUserModsOK creates ListUserModsOK with default headers values
func NewListUserModsOK() *ListUserModsOK {

	return &ListUserModsOK{}
}

// WithPayload adds the payload to the list user mods o k response
func (o *ListUserModsOK) WithPayload(payload []*models.UserMod) *ListUserModsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user mods o k response
func (o *ListUserModsOK) SetPayload(payload []*models.UserMod) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserModsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListUserModsForbiddenCode is the HTTP code returned for type ListUserModsForbidden
const ListUserModsForbiddenCode int = 403

/*ListUserModsForbidden User is not authorized

swagger:response listUserModsForbidden
*/
type ListUserModsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListUserModsForbidden creates ListUserModsForbidden with default headers values
func NewListUserModsForbidden() *ListUserModsForbidden {

	return &ListUserModsForbidden{}
}

// WithPayload adds the payload to the list user mods forbidden response
func (o *ListUserModsForbidden) WithPayload(payload *models.GeneralError) *ListUserModsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user mods forbidden response
func (o *ListUserModsForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserModsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUserModsNotFoundCode is the HTTP code returned for type ListUserModsNotFound
const ListUserModsNotFoundCode int = 404

/*ListUserModsNotFound User not found

swagger:response listUserModsNotFound
*/
type ListUserModsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListUserModsNotFound creates ListUserModsNotFound with default headers values
func NewListUserModsNotFound() *ListUserModsNotFound {

	return &ListUserModsNotFound{}
}

// WithPayload adds the payload to the list user mods not found response
func (o *ListUserModsNotFound) WithPayload(payload *models.GeneralError) *ListUserModsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user mods not found response
func (o *ListUserModsNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserModsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListUserModsDefault Some error unrelated to the handler

swagger:response listUserModsDefault
*/
type ListUserModsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListUserModsDefault creates ListUserModsDefault with default headers values
func NewListUserModsDefault(code int) *ListUserModsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListUserModsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list user mods default response
func (o *ListUserModsDefault) WithStatusCode(code int) *ListUserModsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list user mods default response
func (o *ListUserModsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list user mods default response
func (o *ListUserModsDefault) WithPayload(payload *models.GeneralError) *ListUserModsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user mods default response
func (o *ListUserModsDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserModsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}