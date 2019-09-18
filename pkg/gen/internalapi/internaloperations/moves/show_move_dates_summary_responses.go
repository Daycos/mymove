// Code generated by go-swagger; DO NOT EDIT.

package moves

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowMoveDatesSummaryOKCode is the HTTP code returned for type ShowMoveDatesSummaryOK
const ShowMoveDatesSummaryOKCode int = 200

/*ShowMoveDatesSummaryOK List of projected move-related dates

swagger:response showMoveDatesSummaryOK
*/
type ShowMoveDatesSummaryOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.MoveDatesSummary `json:"body,omitempty"`
}

// NewShowMoveDatesSummaryOK creates ShowMoveDatesSummaryOK with default headers values
func NewShowMoveDatesSummaryOK() *ShowMoveDatesSummaryOK {

	return &ShowMoveDatesSummaryOK{}
}

// WithPayload adds the payload to the show move dates summary o k response
func (o *ShowMoveDatesSummaryOK) WithPayload(payload *internalmessages.MoveDatesSummary) *ShowMoveDatesSummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show move dates summary o k response
func (o *ShowMoveDatesSummaryOK) SetPayload(payload *internalmessages.MoveDatesSummary) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowMoveDatesSummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowMoveDatesSummaryBadRequestCode is the HTTP code returned for type ShowMoveDatesSummaryBadRequest
const ShowMoveDatesSummaryBadRequestCode int = 400

/*ShowMoveDatesSummaryBadRequest invalid request

swagger:response showMoveDatesSummaryBadRequest
*/
type ShowMoveDatesSummaryBadRequest struct {
}

// NewShowMoveDatesSummaryBadRequest creates ShowMoveDatesSummaryBadRequest with default headers values
func NewShowMoveDatesSummaryBadRequest() *ShowMoveDatesSummaryBadRequest {

	return &ShowMoveDatesSummaryBadRequest{}
}

// WriteResponse to the client
func (o *ShowMoveDatesSummaryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowMoveDatesSummaryUnauthorizedCode is the HTTP code returned for type ShowMoveDatesSummaryUnauthorized
const ShowMoveDatesSummaryUnauthorizedCode int = 401

/*ShowMoveDatesSummaryUnauthorized request requires user authentication

swagger:response showMoveDatesSummaryUnauthorized
*/
type ShowMoveDatesSummaryUnauthorized struct {
}

// NewShowMoveDatesSummaryUnauthorized creates ShowMoveDatesSummaryUnauthorized with default headers values
func NewShowMoveDatesSummaryUnauthorized() *ShowMoveDatesSummaryUnauthorized {

	return &ShowMoveDatesSummaryUnauthorized{}
}

// WriteResponse to the client
func (o *ShowMoveDatesSummaryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ShowMoveDatesSummaryForbiddenCode is the HTTP code returned for type ShowMoveDatesSummaryForbidden
const ShowMoveDatesSummaryForbiddenCode int = 403

/*ShowMoveDatesSummaryForbidden user is not authorized

swagger:response showMoveDatesSummaryForbidden
*/
type ShowMoveDatesSummaryForbidden struct {
}

// NewShowMoveDatesSummaryForbidden creates ShowMoveDatesSummaryForbidden with default headers values
func NewShowMoveDatesSummaryForbidden() *ShowMoveDatesSummaryForbidden {

	return &ShowMoveDatesSummaryForbidden{}
}

// WriteResponse to the client
func (o *ShowMoveDatesSummaryForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowMoveDatesSummaryInternalServerErrorCode is the HTTP code returned for type ShowMoveDatesSummaryInternalServerError
const ShowMoveDatesSummaryInternalServerErrorCode int = 500

/*ShowMoveDatesSummaryInternalServerError internal server error

swagger:response showMoveDatesSummaryInternalServerError
*/
type ShowMoveDatesSummaryInternalServerError struct {
}

// NewShowMoveDatesSummaryInternalServerError creates ShowMoveDatesSummaryInternalServerError with default headers values
func NewShowMoveDatesSummaryInternalServerError() *ShowMoveDatesSummaryInternalServerError {

	return &ShowMoveDatesSummaryInternalServerError{}
}

// WriteResponse to the client
func (o *ShowMoveDatesSummaryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
