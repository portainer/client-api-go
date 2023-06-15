// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// StackCreateReader is a Reader for the StackCreate structure.
type StackCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackCreateOK creates a StackCreateOK with default headers values
func NewStackCreateOK() *StackCreateOK {
	return &StackCreateOK{}
}

/*
StackCreateOK describes a response with status code 200, with default header values.

OK
*/
type StackCreateOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack create o k response has a 2xx status code
func (o *StackCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack create o k response has a 3xx status code
func (o *StackCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create o k response has a 4xx status code
func (o *StackCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create o k response has a 5xx status code
func (o *StackCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create o k response a status code equal to that given
func (o *StackCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *StackCreateOK) Error() string {
	return fmt.Sprintf("[POST /stacks][%d] stackCreateOK  %+v", 200, o.Payload)
}

func (o *StackCreateOK) String() string {
	return fmt.Sprintf("[POST /stacks][%d] stackCreateOK  %+v", 200, o.Payload)
}

func (o *StackCreateOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCreateBadRequest creates a StackCreateBadRequest with default headers values
func NewStackCreateBadRequest() *StackCreateBadRequest {
	return &StackCreateBadRequest{}
}

/*
StackCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackCreateBadRequest struct {
}

// IsSuccess returns true when this stack create bad request response has a 2xx status code
func (o *StackCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create bad request response has a 3xx status code
func (o *StackCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create bad request response has a 4xx status code
func (o *StackCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack create bad request response has a 5xx status code
func (o *StackCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create bad request response a status code equal to that given
func (o *StackCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StackCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks][%d] stackCreateBadRequest ", 400)
}

func (o *StackCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks][%d] stackCreateBadRequest ", 400)
}

func (o *StackCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackCreateInternalServerError creates a StackCreateInternalServerError with default headers values
func NewStackCreateInternalServerError() *StackCreateInternalServerError {
	return &StackCreateInternalServerError{}
}

/*
StackCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackCreateInternalServerError struct {
}

// IsSuccess returns true when this stack create internal server error response has a 2xx status code
func (o *StackCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create internal server error response has a 3xx status code
func (o *StackCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create internal server error response has a 4xx status code
func (o *StackCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create internal server error response has a 5xx status code
func (o *StackCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack create internal server error response a status code equal to that given
func (o *StackCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StackCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks][%d] stackCreateInternalServerError ", 500)
}

func (o *StackCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks][%d] stackCreateInternalServerError ", 500)
}

func (o *StackCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
