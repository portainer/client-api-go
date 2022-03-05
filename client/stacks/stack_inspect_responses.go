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

// StackInspectReader is a Reader for the StackInspect structure.
type StackInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackInspectOK creates a StackInspectOK with default headers values
func NewStackInspectOK() *StackInspectOK {
	return &StackInspectOK{}
}

/* StackInspectOK describes a response with status code 200, with default header values.

Success
*/
type StackInspectOK struct {
	Payload *models.PortainerStack
}

func (o *StackInspectOK) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectOK  %+v", 200, o.Payload)
}
func (o *StackInspectOK) GetPayload() *models.PortainerStack {
	return o.Payload
}

func (o *StackInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackInspectBadRequest creates a StackInspectBadRequest with default headers values
func NewStackInspectBadRequest() *StackInspectBadRequest {
	return &StackInspectBadRequest{}
}

/* StackInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackInspectBadRequest struct {
}

func (o *StackInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectBadRequest ", 400)
}

func (o *StackInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackInspectForbidden creates a StackInspectForbidden with default headers values
func NewStackInspectForbidden() *StackInspectForbidden {
	return &StackInspectForbidden{}
}

/* StackInspectForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackInspectForbidden struct {
}

func (o *StackInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectForbidden ", 403)
}

func (o *StackInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackInspectNotFound creates a StackInspectNotFound with default headers values
func NewStackInspectNotFound() *StackInspectNotFound {
	return &StackInspectNotFound{}
}

/* StackInspectNotFound describes a response with status code 404, with default header values.

Stack not found
*/
type StackInspectNotFound struct {
}

func (o *StackInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectNotFound ", 404)
}

func (o *StackInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackInspectInternalServerError creates a StackInspectInternalServerError with default headers values
func NewStackInspectInternalServerError() *StackInspectInternalServerError {
	return &StackInspectInternalServerError{}
}

/* StackInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackInspectInternalServerError struct {
}

func (o *StackInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectInternalServerError ", 500)
}

func (o *StackInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
