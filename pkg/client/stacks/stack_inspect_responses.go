// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/pkg/models"
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
		return nil, runtime.NewAPIError("[GET /stacks/{id}] StackInspect", response, response.Code())
	}
}

// NewStackInspectOK creates a StackInspectOK with default headers values
func NewStackInspectOK() *StackInspectOK {
	return &StackInspectOK{}
}

/*
StackInspectOK describes a response with status code 200, with default header values.

Success
*/
type StackInspectOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack inspect o k response has a 2xx status code
func (o *StackInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack inspect o k response has a 3xx status code
func (o *StackInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack inspect o k response has a 4xx status code
func (o *StackInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack inspect o k response has a 5xx status code
func (o *StackInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack inspect o k response a status code equal to that given
func (o *StackInspectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stack inspect o k response
func (o *StackInspectOK) Code() int {
	return 200
}

func (o *StackInspectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectOK %s", 200, payload)
}

func (o *StackInspectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectOK %s", 200, payload)
}

func (o *StackInspectOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

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

/*
StackInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackInspectBadRequest struct {
}

// IsSuccess returns true when this stack inspect bad request response has a 2xx status code
func (o *StackInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack inspect bad request response has a 3xx status code
func (o *StackInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack inspect bad request response has a 4xx status code
func (o *StackInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack inspect bad request response has a 5xx status code
func (o *StackInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack inspect bad request response a status code equal to that given
func (o *StackInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stack inspect bad request response
func (o *StackInspectBadRequest) Code() int {
	return 400
}

func (o *StackInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectBadRequest", 400)
}

func (o *StackInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectBadRequest", 400)
}

func (o *StackInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackInspectForbidden creates a StackInspectForbidden with default headers values
func NewStackInspectForbidden() *StackInspectForbidden {
	return &StackInspectForbidden{}
}

/*
StackInspectForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackInspectForbidden struct {
}

// IsSuccess returns true when this stack inspect forbidden response has a 2xx status code
func (o *StackInspectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack inspect forbidden response has a 3xx status code
func (o *StackInspectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack inspect forbidden response has a 4xx status code
func (o *StackInspectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack inspect forbidden response has a 5xx status code
func (o *StackInspectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stack inspect forbidden response a status code equal to that given
func (o *StackInspectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stack inspect forbidden response
func (o *StackInspectForbidden) Code() int {
	return 403
}

func (o *StackInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectForbidden", 403)
}

func (o *StackInspectForbidden) String() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectForbidden", 403)
}

func (o *StackInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackInspectNotFound creates a StackInspectNotFound with default headers values
func NewStackInspectNotFound() *StackInspectNotFound {
	return &StackInspectNotFound{}
}

/*
StackInspectNotFound describes a response with status code 404, with default header values.

Stack not found
*/
type StackInspectNotFound struct {
}

// IsSuccess returns true when this stack inspect not found response has a 2xx status code
func (o *StackInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack inspect not found response has a 3xx status code
func (o *StackInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack inspect not found response has a 4xx status code
func (o *StackInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack inspect not found response has a 5xx status code
func (o *StackInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stack inspect not found response a status code equal to that given
func (o *StackInspectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stack inspect not found response
func (o *StackInspectNotFound) Code() int {
	return 404
}

func (o *StackInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectNotFound", 404)
}

func (o *StackInspectNotFound) String() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectNotFound", 404)
}

func (o *StackInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackInspectInternalServerError creates a StackInspectInternalServerError with default headers values
func NewStackInspectInternalServerError() *StackInspectInternalServerError {
	return &StackInspectInternalServerError{}
}

/*
StackInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackInspectInternalServerError struct {
}

// IsSuccess returns true when this stack inspect internal server error response has a 2xx status code
func (o *StackInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack inspect internal server error response has a 3xx status code
func (o *StackInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack inspect internal server error response has a 4xx status code
func (o *StackInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack inspect internal server error response has a 5xx status code
func (o *StackInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack inspect internal server error response a status code equal to that given
func (o *StackInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stack inspect internal server error response
func (o *StackInspectInternalServerError) Code() int {
	return 500
}

func (o *StackInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectInternalServerError", 500)
}

func (o *StackInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /stacks/{id}][%d] stackInspectInternalServerError", 500)
}

func (o *StackInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
