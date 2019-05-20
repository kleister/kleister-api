// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteTeamOKCode is the HTTP code returned for type DeleteTeamOK
const DeleteTeamOKCode int = 200

/*DeleteTeamOK Plain success message

swagger:response deleteTeamOK
*/
type DeleteTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteTeamOK creates DeleteTeamOK with default headers values
func NewDeleteTeamOK() *DeleteTeamOK {

	return &DeleteTeamOK{}
}

// WithPayload adds the payload to the delete team o k response
func (o *DeleteTeamOK) WithPayload(payload *models.GeneralError) *DeleteTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete team o k response
func (o *DeleteTeamOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTeamBadRequestCode is the HTTP code returned for type DeleteTeamBadRequest
const DeleteTeamBadRequestCode int = 400

/*DeleteTeamBadRequest Failed to delete the team

swagger:response deleteTeamBadRequest
*/
type DeleteTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteTeamBadRequest creates DeleteTeamBadRequest with default headers values
func NewDeleteTeamBadRequest() *DeleteTeamBadRequest {

	return &DeleteTeamBadRequest{}
}

// WithPayload adds the payload to the delete team bad request response
func (o *DeleteTeamBadRequest) WithPayload(payload *models.GeneralError) *DeleteTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete team bad request response
func (o *DeleteTeamBadRequest) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTeamForbiddenCode is the HTTP code returned for type DeleteTeamForbidden
const DeleteTeamForbiddenCode int = 403

/*DeleteTeamForbidden User is not authorized

swagger:response deleteTeamForbidden
*/
type DeleteTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteTeamForbidden creates DeleteTeamForbidden with default headers values
func NewDeleteTeamForbidden() *DeleteTeamForbidden {

	return &DeleteTeamForbidden{}
}

// WithPayload adds the payload to the delete team forbidden response
func (o *DeleteTeamForbidden) WithPayload(payload *models.GeneralError) *DeleteTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete team forbidden response
func (o *DeleteTeamForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteTeamDefault Some error unrelated to the handler

swagger:response deleteTeamDefault
*/
type DeleteTeamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteTeamDefault creates DeleteTeamDefault with default headers values
func NewDeleteTeamDefault(code int) *DeleteTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete team default response
func (o *DeleteTeamDefault) WithStatusCode(code int) *DeleteTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete team default response
func (o *DeleteTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete team default response
func (o *DeleteTeamDefault) WithPayload(payload *models.GeneralError) *DeleteTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete team default response
func (o *DeleteTeamDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
