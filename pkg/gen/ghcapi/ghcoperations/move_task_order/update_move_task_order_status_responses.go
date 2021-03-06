// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdateMoveTaskOrderStatusOKCode is the HTTP code returned for type UpdateMoveTaskOrderStatusOK
const UpdateMoveTaskOrderStatusOKCode int = 200

/*UpdateMoveTaskOrderStatusOK Successfully updated move task order status

swagger:response updateMoveTaskOrderStatusOK
*/
type UpdateMoveTaskOrderStatusOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderStatusOK creates UpdateMoveTaskOrderStatusOK with default headers values
func NewUpdateMoveTaskOrderStatusOK() *UpdateMoveTaskOrderStatusOK {

	return &UpdateMoveTaskOrderStatusOK{}
}

// WithPayload adds the payload to the update move task order status o k response
func (o *UpdateMoveTaskOrderStatusOK) WithPayload(payload *ghcmessages.MoveTaskOrder) *UpdateMoveTaskOrderStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order status o k response
func (o *UpdateMoveTaskOrderStatusOK) SetPayload(payload *ghcmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMoveTaskOrderStatusBadRequestCode is the HTTP code returned for type UpdateMoveTaskOrderStatusBadRequest
const UpdateMoveTaskOrderStatusBadRequestCode int = 400

/*UpdateMoveTaskOrderStatusBadRequest The request payload is invalid

swagger:response updateMoveTaskOrderStatusBadRequest
*/
type UpdateMoveTaskOrderStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderStatusBadRequest creates UpdateMoveTaskOrderStatusBadRequest with default headers values
func NewUpdateMoveTaskOrderStatusBadRequest() *UpdateMoveTaskOrderStatusBadRequest {

	return &UpdateMoveTaskOrderStatusBadRequest{}
}

// WithPayload adds the payload to the update move task order status bad request response
func (o *UpdateMoveTaskOrderStatusBadRequest) WithPayload(payload interface{}) *UpdateMoveTaskOrderStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order status bad request response
func (o *UpdateMoveTaskOrderStatusBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderStatusUnauthorizedCode is the HTTP code returned for type UpdateMoveTaskOrderStatusUnauthorized
const UpdateMoveTaskOrderStatusUnauthorizedCode int = 401

/*UpdateMoveTaskOrderStatusUnauthorized The request was denied

swagger:response updateMoveTaskOrderStatusUnauthorized
*/
type UpdateMoveTaskOrderStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderStatusUnauthorized creates UpdateMoveTaskOrderStatusUnauthorized with default headers values
func NewUpdateMoveTaskOrderStatusUnauthorized() *UpdateMoveTaskOrderStatusUnauthorized {

	return &UpdateMoveTaskOrderStatusUnauthorized{}
}

// WithPayload adds the payload to the update move task order status unauthorized response
func (o *UpdateMoveTaskOrderStatusUnauthorized) WithPayload(payload interface{}) *UpdateMoveTaskOrderStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order status unauthorized response
func (o *UpdateMoveTaskOrderStatusUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderStatusForbiddenCode is the HTTP code returned for type UpdateMoveTaskOrderStatusForbidden
const UpdateMoveTaskOrderStatusForbiddenCode int = 403

/*UpdateMoveTaskOrderStatusForbidden The request was denied

swagger:response updateMoveTaskOrderStatusForbidden
*/
type UpdateMoveTaskOrderStatusForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderStatusForbidden creates UpdateMoveTaskOrderStatusForbidden with default headers values
func NewUpdateMoveTaskOrderStatusForbidden() *UpdateMoveTaskOrderStatusForbidden {

	return &UpdateMoveTaskOrderStatusForbidden{}
}

// WithPayload adds the payload to the update move task order status forbidden response
func (o *UpdateMoveTaskOrderStatusForbidden) WithPayload(payload interface{}) *UpdateMoveTaskOrderStatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order status forbidden response
func (o *UpdateMoveTaskOrderStatusForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderStatusNotFoundCode is the HTTP code returned for type UpdateMoveTaskOrderStatusNotFound
const UpdateMoveTaskOrderStatusNotFoundCode int = 404

/*UpdateMoveTaskOrderStatusNotFound The requested resource wasn't found

swagger:response updateMoveTaskOrderStatusNotFound
*/
type UpdateMoveTaskOrderStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderStatusNotFound creates UpdateMoveTaskOrderStatusNotFound with default headers values
func NewUpdateMoveTaskOrderStatusNotFound() *UpdateMoveTaskOrderStatusNotFound {

	return &UpdateMoveTaskOrderStatusNotFound{}
}

// WithPayload adds the payload to the update move task order status not found response
func (o *UpdateMoveTaskOrderStatusNotFound) WithPayload(payload interface{}) *UpdateMoveTaskOrderStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order status not found response
func (o *UpdateMoveTaskOrderStatusNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderStatusInternalServerErrorCode is the HTTP code returned for type UpdateMoveTaskOrderStatusInternalServerError
const UpdateMoveTaskOrderStatusInternalServerErrorCode int = 500

/*UpdateMoveTaskOrderStatusInternalServerError A server error occurred

swagger:response updateMoveTaskOrderStatusInternalServerError
*/
type UpdateMoveTaskOrderStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderStatusInternalServerError creates UpdateMoveTaskOrderStatusInternalServerError with default headers values
func NewUpdateMoveTaskOrderStatusInternalServerError() *UpdateMoveTaskOrderStatusInternalServerError {

	return &UpdateMoveTaskOrderStatusInternalServerError{}
}

// WithPayload adds the payload to the update move task order status internal server error response
func (o *UpdateMoveTaskOrderStatusInternalServerError) WithPayload(payload interface{}) *UpdateMoveTaskOrderStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order status internal server error response
func (o *UpdateMoveTaskOrderStatusInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
