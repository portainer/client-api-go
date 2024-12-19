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

	"github.com/portainer/client-api-go/v2/models"
)

// StackFileInspectReader is a Reader for the StackFileInspect structure.
type StackFileInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackFileInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackFileInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackFileInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackFileInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackFileInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackFileInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /stacks/{id}/file] StackFileInspect", response, response.Code())
	}
}

// NewStackFileInspectOK creates a StackFileInspectOK with default headers values
func NewStackFileInspectOK() *StackFileInspectOK {
	return &StackFileInspectOK{}
}

/*
StackFileInspectOK describes a response with status code 200, with default header values.

Success
*/
type StackFileInspectOK struct {
	Payload *models.StacksStackFileResponse
}

// IsSuccess returns true when this stack file inspect o k response has a 2xx status code
func (o *StackFileInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack file inspect o k response has a 3xx status code
func (o *StackFileInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack file inspect o k response has a 4xx status code
func (o *StackFileInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack file inspect o k response has a 5xx status code
func (o *StackFileInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack file inspect o k response a status code equal to that given
func (o *StackFileInspectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stack file inspect o k response
func (o *StackFileInspectOK) Code() int {
	return 200
}

func (o *StackFileInspectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectOK %s", 200, payload)
}

func (o *StackFileInspectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectOK %s", 200, payload)
}

func (o *StackFileInspectOK) GetPayload() *models.StacksStackFileResponse {
	return o.Payload
}

func (o *StackFileInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StacksStackFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackFileInspectBadRequest creates a StackFileInspectBadRequest with default headers values
func NewStackFileInspectBadRequest() *StackFileInspectBadRequest {
	return &StackFileInspectBadRequest{}
}

/*
StackFileInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackFileInspectBadRequest struct {
}

// IsSuccess returns true when this stack file inspect bad request response has a 2xx status code
func (o *StackFileInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack file inspect bad request response has a 3xx status code
func (o *StackFileInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack file inspect bad request response has a 4xx status code
func (o *StackFileInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack file inspect bad request response has a 5xx status code
func (o *StackFileInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack file inspect bad request response a status code equal to that given
func (o *StackFileInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stack file inspect bad request response
func (o *StackFileInspectBadRequest) Code() int {
	return 400
}

func (o *StackFileInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectBadRequest", 400)
}

func (o *StackFileInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectBadRequest", 400)
}

func (o *StackFileInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackFileInspectForbidden creates a StackFileInspectForbidden with default headers values
func NewStackFileInspectForbidden() *StackFileInspectForbidden {
	return &StackFileInspectForbidden{}
}

/*
StackFileInspectForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackFileInspectForbidden struct {
}

// IsSuccess returns true when this stack file inspect forbidden response has a 2xx status code
func (o *StackFileInspectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack file inspect forbidden response has a 3xx status code
func (o *StackFileInspectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack file inspect forbidden response has a 4xx status code
func (o *StackFileInspectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack file inspect forbidden response has a 5xx status code
func (o *StackFileInspectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stack file inspect forbidden response a status code equal to that given
func (o *StackFileInspectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stack file inspect forbidden response
func (o *StackFileInspectForbidden) Code() int {
	return 403
}

func (o *StackFileInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectForbidden", 403)
}

func (o *StackFileInspectForbidden) String() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectForbidden", 403)
}

func (o *StackFileInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackFileInspectNotFound creates a StackFileInspectNotFound with default headers values
func NewStackFileInspectNotFound() *StackFileInspectNotFound {
	return &StackFileInspectNotFound{}
}

/*
StackFileInspectNotFound describes a response with status code 404, with default header values.

Stack not found
*/
type StackFileInspectNotFound struct {
}

// IsSuccess returns true when this stack file inspect not found response has a 2xx status code
func (o *StackFileInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack file inspect not found response has a 3xx status code
func (o *StackFileInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack file inspect not found response has a 4xx status code
func (o *StackFileInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack file inspect not found response has a 5xx status code
func (o *StackFileInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stack file inspect not found response a status code equal to that given
func (o *StackFileInspectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stack file inspect not found response
func (o *StackFileInspectNotFound) Code() int {
	return 404
}

func (o *StackFileInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectNotFound", 404)
}

func (o *StackFileInspectNotFound) String() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectNotFound", 404)
}

func (o *StackFileInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackFileInspectInternalServerError creates a StackFileInspectInternalServerError with default headers values
func NewStackFileInspectInternalServerError() *StackFileInspectInternalServerError {
	return &StackFileInspectInternalServerError{}
}

/*
StackFileInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackFileInspectInternalServerError struct {
}

// IsSuccess returns true when this stack file inspect internal server error response has a 2xx status code
func (o *StackFileInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack file inspect internal server error response has a 3xx status code
func (o *StackFileInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack file inspect internal server error response has a 4xx status code
func (o *StackFileInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack file inspect internal server error response has a 5xx status code
func (o *StackFileInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack file inspect internal server error response a status code equal to that given
func (o *StackFileInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stack file inspect internal server error response
func (o *StackFileInspectInternalServerError) Code() int {
	return 500
}

func (o *StackFileInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectInternalServerError", 500)
}

func (o *StackFileInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /stacks/{id}/file][%d] stackFileInspectInternalServerError", 500)
}

func (o *StackFileInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
