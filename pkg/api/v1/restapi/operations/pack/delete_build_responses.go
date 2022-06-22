// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteBuildOKCode is the HTTP code returned for type DeleteBuildOK
const DeleteBuildOKCode int = 200

/*DeleteBuildOK Plain success message

swagger:response deleteBuildOK
*/
type DeleteBuildOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteBuildOK creates DeleteBuildOK with default headers values
func NewDeleteBuildOK() *DeleteBuildOK {

	return &DeleteBuildOK{}
}

// WithPayload adds the payload to the delete build o k response
func (o *DeleteBuildOK) WithPayload(payload *models.GeneralError) *DeleteBuildOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete build o k response
func (o *DeleteBuildOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBuildOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBuildBadRequestCode is the HTTP code returned for type DeleteBuildBadRequest
const DeleteBuildBadRequestCode int = 400

/*DeleteBuildBadRequest Failed to delete the build

swagger:response deleteBuildBadRequest
*/
type DeleteBuildBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteBuildBadRequest creates DeleteBuildBadRequest with default headers values
func NewDeleteBuildBadRequest() *DeleteBuildBadRequest {

	return &DeleteBuildBadRequest{}
}

// WithPayload adds the payload to the delete build bad request response
func (o *DeleteBuildBadRequest) WithPayload(payload *models.GeneralError) *DeleteBuildBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete build bad request response
func (o *DeleteBuildBadRequest) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBuildBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBuildForbiddenCode is the HTTP code returned for type DeleteBuildForbidden
const DeleteBuildForbiddenCode int = 403

/*DeleteBuildForbidden User is not authorized

swagger:response deleteBuildForbidden
*/
type DeleteBuildForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteBuildForbidden creates DeleteBuildForbidden with default headers values
func NewDeleteBuildForbidden() *DeleteBuildForbidden {

	return &DeleteBuildForbidden{}
}

// WithPayload adds the payload to the delete build forbidden response
func (o *DeleteBuildForbidden) WithPayload(payload *models.GeneralError) *DeleteBuildForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete build forbidden response
func (o *DeleteBuildForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBuildForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBuildNotFoundCode is the HTTP code returned for type DeleteBuildNotFound
const DeleteBuildNotFoundCode int = 404

/*DeleteBuildNotFound Build or pack not found

swagger:response deleteBuildNotFound
*/
type DeleteBuildNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteBuildNotFound creates DeleteBuildNotFound with default headers values
func NewDeleteBuildNotFound() *DeleteBuildNotFound {

	return &DeleteBuildNotFound{}
}

// WithPayload adds the payload to the delete build not found response
func (o *DeleteBuildNotFound) WithPayload(payload *models.GeneralError) *DeleteBuildNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete build not found response
func (o *DeleteBuildNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBuildNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteBuildDefault Some error unrelated to the handler

swagger:response deleteBuildDefault
*/
type DeleteBuildDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteBuildDefault creates DeleteBuildDefault with default headers values
func NewDeleteBuildDefault(code int) *DeleteBuildDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteBuildDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete build default response
func (o *DeleteBuildDefault) WithStatusCode(code int) *DeleteBuildDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete build default response
func (o *DeleteBuildDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete build default response
func (o *DeleteBuildDefault) WithPayload(payload *models.GeneralError) *DeleteBuildDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete build default response
func (o *DeleteBuildDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBuildDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
