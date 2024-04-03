// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// StackStopReader is a Reader for the StackStop structure.
type StackStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackStopForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackStopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackStopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /stacks/{id}/stop] StackStop", response, response.Code())
	}
}

// NewStackStopOK creates a StackStopOK with default headers values
func NewStackStopOK() *StackStopOK {
	return &StackStopOK{}
}

/*
StackStopOK describes a response with status code 200, with default header values.

Success
*/
type StackStopOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack stop o k response has a 2xx status code
func (o *StackStopOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack stop o k response has a 3xx status code
func (o *StackStopOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack stop o k response has a 4xx status code
func (o *StackStopOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack stop o k response has a 5xx status code
func (o *StackStopOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack stop o k response a status code equal to that given
func (o *StackStopOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stack stop o k response
func (o *StackStopOK) Code() int {
	return 200
}

func (o *StackStopOK) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopOK  %+v", 200, o.Payload)
}

func (o *StackStopOK) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopOK  %+v", 200, o.Payload)
}

func (o *StackStopOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackStopBadRequest creates a StackStopBadRequest with default headers values
func NewStackStopBadRequest() *StackStopBadRequest {
	return &StackStopBadRequest{}
}

/*
StackStopBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackStopBadRequest struct {
}

// IsSuccess returns true when this stack stop bad request response has a 2xx status code
func (o *StackStopBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack stop bad request response has a 3xx status code
func (o *StackStopBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack stop bad request response has a 4xx status code
func (o *StackStopBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack stop bad request response has a 5xx status code
func (o *StackStopBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack stop bad request response a status code equal to that given
func (o *StackStopBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stack stop bad request response
func (o *StackStopBadRequest) Code() int {
	return 400
}

func (o *StackStopBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopBadRequest ", 400)
}

func (o *StackStopBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopBadRequest ", 400)
}

func (o *StackStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackStopForbidden creates a StackStopForbidden with default headers values
func NewStackStopForbidden() *StackStopForbidden {
	return &StackStopForbidden{}
}

/*
StackStopForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackStopForbidden struct {
}

// IsSuccess returns true when this stack stop forbidden response has a 2xx status code
func (o *StackStopForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack stop forbidden response has a 3xx status code
func (o *StackStopForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack stop forbidden response has a 4xx status code
func (o *StackStopForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack stop forbidden response has a 5xx status code
func (o *StackStopForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stack stop forbidden response a status code equal to that given
func (o *StackStopForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stack stop forbidden response
func (o *StackStopForbidden) Code() int {
	return 403
}

func (o *StackStopForbidden) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopForbidden ", 403)
}

func (o *StackStopForbidden) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopForbidden ", 403)
}

func (o *StackStopForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackStopNotFound creates a StackStopNotFound with default headers values
func NewStackStopNotFound() *StackStopNotFound {
	return &StackStopNotFound{}
}

/*
StackStopNotFound describes a response with status code 404, with default header values.

Not found
*/
type StackStopNotFound struct {
}

// IsSuccess returns true when this stack stop not found response has a 2xx status code
func (o *StackStopNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack stop not found response has a 3xx status code
func (o *StackStopNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack stop not found response has a 4xx status code
func (o *StackStopNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack stop not found response has a 5xx status code
func (o *StackStopNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stack stop not found response a status code equal to that given
func (o *StackStopNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stack stop not found response
func (o *StackStopNotFound) Code() int {
	return 404
}

func (o *StackStopNotFound) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopNotFound ", 404)
}

func (o *StackStopNotFound) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopNotFound ", 404)
}

func (o *StackStopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackStopInternalServerError creates a StackStopInternalServerError with default headers values
func NewStackStopInternalServerError() *StackStopInternalServerError {
	return &StackStopInternalServerError{}
}

/*
StackStopInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackStopInternalServerError struct {
}

// IsSuccess returns true when this stack stop internal server error response has a 2xx status code
func (o *StackStopInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack stop internal server error response has a 3xx status code
func (o *StackStopInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack stop internal server error response has a 4xx status code
func (o *StackStopInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack stop internal server error response has a 5xx status code
func (o *StackStopInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack stop internal server error response a status code equal to that given
func (o *StackStopInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stack stop internal server error response
func (o *StackStopInternalServerError) Code() int {
	return 500
}

func (o *StackStopInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopInternalServerError ", 500)
}

func (o *StackStopInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/stop][%d] stackStopInternalServerError ", 500)
}

func (o *StackStopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
