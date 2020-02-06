// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateUploadReader is a Reader for the CreateUpload structure.
type CreateUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUploadCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUploadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUploadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUploadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUploadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUploadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUploadCreated creates a CreateUploadCreated with default headers values
func NewCreateUploadCreated() *CreateUploadCreated {
	return &CreateUploadCreated{}
}

/*CreateUploadCreated handles this case with default header values.

Created upload
*/
type CreateUploadCreated struct {
	Payload *primemessages.Upload
}

func (o *CreateUploadCreated) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadCreated  %+v", 201, o.Payload)
}

func (o *CreateUploadCreated) GetPayload() *primemessages.Upload {
	return o.Payload
}

func (o *CreateUploadCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Upload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadBadRequest creates a CreateUploadBadRequest with default headers values
func NewCreateUploadBadRequest() *CreateUploadBadRequest {
	return &CreateUploadBadRequest{}
}

/*CreateUploadBadRequest handles this case with default header values.

Invalid request
*/
type CreateUploadBadRequest struct {
	Payload interface{}
}

func (o *CreateUploadBadRequest) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUploadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUploadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadUnauthorized creates a CreateUploadUnauthorized with default headers values
func NewCreateUploadUnauthorized() *CreateUploadUnauthorized {
	return &CreateUploadUnauthorized{}
}

/*CreateUploadUnauthorized handles this case with default header values.

The request was denied
*/
type CreateUploadUnauthorized struct {
	Payload interface{}
}

func (o *CreateUploadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUploadUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUploadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadForbidden creates a CreateUploadForbidden with default headers values
func NewCreateUploadForbidden() *CreateUploadForbidden {
	return &CreateUploadForbidden{}
}

/*CreateUploadForbidden handles this case with default header values.

The request was denied
*/
type CreateUploadForbidden struct {
	Payload interface{}
}

func (o *CreateUploadForbidden) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadForbidden  %+v", 403, o.Payload)
}

func (o *CreateUploadForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUploadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadNotFound creates a CreateUploadNotFound with default headers values
func NewCreateUploadNotFound() *CreateUploadNotFound {
	return &CreateUploadNotFound{}
}

/*CreateUploadNotFound handles this case with default header values.

The requested resource wasn't found
*/
type CreateUploadNotFound struct {
	Payload interface{}
}

func (o *CreateUploadNotFound) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadNotFound  %+v", 404, o.Payload)
}

func (o *CreateUploadNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUploadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadInternalServerError creates a CreateUploadInternalServerError with default headers values
func NewCreateUploadInternalServerError() *CreateUploadInternalServerError {
	return &CreateUploadInternalServerError{}
}

/*CreateUploadInternalServerError handles this case with default header values.

A server error occurred
*/
type CreateUploadInternalServerError struct {
	Payload interface{}
}

func (o *CreateUploadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUploadInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUploadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}