// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListUserPacksOKCode is the HTTP code returned for type ListUserPacksOK
const ListUserPacksOKCode int = 200

/*ListUserPacksOK A collection of team packs

swagger:response listUserPacksOK
*/
type ListUserPacksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserPack `json:"body,omitempty"`
}

// NewListUserPacksOK creates ListUserPacksOK with default headers values
func NewListUserPacksOK() *ListUserPacksOK {

	return &ListUserPacksOK{}
}

// WithPayload adds the payload to the list user packs o k response
func (o *ListUserPacksOK) WithPayload(payload []*models.UserPack) *ListUserPacksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user packs o k response
func (o *ListUserPacksOK) SetPayload(payload []*models.UserPack) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserPacksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListUserPacksForbiddenCode is the HTTP code returned for type ListUserPacksForbidden
const ListUserPacksForbiddenCode int = 403

/*ListUserPacksForbidden User is not authorized

swagger:response listUserPacksForbidden
*/
type ListUserPacksForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListUserPacksForbidden creates ListUserPacksForbidden with default headers values
func NewListUserPacksForbidden() *ListUserPacksForbidden {

	return &ListUserPacksForbidden{}
}

// WithPayload adds the payload to the list user packs forbidden response
func (o *ListUserPacksForbidden) WithPayload(payload *models.GeneralError) *ListUserPacksForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user packs forbidden response
func (o *ListUserPacksForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserPacksForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUserPacksNotFoundCode is the HTTP code returned for type ListUserPacksNotFound
const ListUserPacksNotFoundCode int = 404

/*ListUserPacksNotFound User not found

swagger:response listUserPacksNotFound
*/
type ListUserPacksNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListUserPacksNotFound creates ListUserPacksNotFound with default headers values
func NewListUserPacksNotFound() *ListUserPacksNotFound {

	return &ListUserPacksNotFound{}
}

// WithPayload adds the payload to the list user packs not found response
func (o *ListUserPacksNotFound) WithPayload(payload *models.GeneralError) *ListUserPacksNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user packs not found response
func (o *ListUserPacksNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserPacksNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListUserPacksDefault Some error unrelated to the handler

swagger:response listUserPacksDefault
*/
type ListUserPacksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListUserPacksDefault creates ListUserPacksDefault with default headers values
func NewListUserPacksDefault(code int) *ListUserPacksDefault {
	if code <= 0 {
		code = 500
	}

	return &ListUserPacksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list user packs default response
func (o *ListUserPacksDefault) WithStatusCode(code int) *ListUserPacksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list user packs default response
func (o *ListUserPacksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list user packs default response
func (o *ListUserPacksDefault) WithPayload(payload *models.GeneralError) *ListUserPacksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user packs default response
func (o *ListUserPacksDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserPacksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}