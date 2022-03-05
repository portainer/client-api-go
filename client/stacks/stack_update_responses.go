// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
)

// StackUpdateReader is a Reader for the StackUpdate structure.
type StackUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackUpdateOK creates a StackUpdateOK with default headers values
func NewStackUpdateOK() *StackUpdateOK {
	return &StackUpdateOK{}
}

/* StackUpdateOK describes a response with status code 200, with default header values.

Success
*/
type StackUpdateOK struct {
	Payload *models.PortainerStack
}

func (o *StackUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /stacks/{id}][%d] stackUpdateOK  %+v", 200, o.Payload)
}
func (o *StackUpdateOK) GetPayload() *models.PortainerStack {
	return o.Payload
}

func (o *StackUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackUpdateBadRequest creates a StackUpdateBadRequest with default headers values
func NewStackUpdateBadRequest() *StackUpdateBadRequest {
	return &StackUpdateBadRequest{}
}

/* StackUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackUpdateBadRequest struct {
}

func (o *StackUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /stacks/{id}][%d] stackUpdateBadRequest ", 400)
}

func (o *StackUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackUpdateForbidden creates a StackUpdateForbidden with default headers values
func NewStackUpdateForbidden() *StackUpdateForbidden {
	return &StackUpdateForbidden{}
}

/* StackUpdateForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackUpdateForbidden struct {
}

func (o *StackUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /stacks/{id}][%d] stackUpdateForbidden ", 403)
}

func (o *StackUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackUpdateNotFound creates a StackUpdateNotFound with default headers values
func NewStackUpdateNotFound() *StackUpdateNotFound {
	return &StackUpdateNotFound{}
}

/* StackUpdateNotFound describes a response with status code 404, with default header values.

Not found
*/
type StackUpdateNotFound struct {
}

func (o *StackUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /stacks/{id}][%d] stackUpdateNotFound ", 404)
}

func (o *StackUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackUpdateInternalServerError creates a StackUpdateInternalServerError with default headers values
func NewStackUpdateInternalServerError() *StackUpdateInternalServerError {
	return &StackUpdateInternalServerError{}
}

/* StackUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackUpdateInternalServerError struct {
}

func (o *StackUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /stacks/{id}][%d] stackUpdateInternalServerError ", 500)
}

func (o *StackUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
