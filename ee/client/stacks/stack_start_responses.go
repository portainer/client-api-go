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

// StackStartReader is a Reader for the StackStart structure.
type StackStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackStartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackStartForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackStartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackStartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackStartOK creates a StackStartOK with default headers values
func NewStackStartOK() *StackStartOK {
	return &StackStartOK{}
}

/*
StackStartOK describes a response with status code 200, with default header values.

Success
*/
type StackStartOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack start o k response has a 2xx status code
func (o *StackStartOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack start o k response has a 3xx status code
func (o *StackStartOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack start o k response has a 4xx status code
func (o *StackStartOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack start o k response has a 5xx status code
func (o *StackStartOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack start o k response a status code equal to that given
func (o *StackStartOK) IsCode(code int) bool {
	return code == 200
}

func (o *StackStartOK) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartOK  %+v", 200, o.Payload)
}

func (o *StackStartOK) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartOK  %+v", 200, o.Payload)
}

func (o *StackStartOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackStartBadRequest creates a StackStartBadRequest with default headers values
func NewStackStartBadRequest() *StackStartBadRequest {
	return &StackStartBadRequest{}
}

/*
StackStartBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackStartBadRequest struct {
}

// IsSuccess returns true when this stack start bad request response has a 2xx status code
func (o *StackStartBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack start bad request response has a 3xx status code
func (o *StackStartBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack start bad request response has a 4xx status code
func (o *StackStartBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack start bad request response has a 5xx status code
func (o *StackStartBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack start bad request response a status code equal to that given
func (o *StackStartBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StackStartBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartBadRequest ", 400)
}

func (o *StackStartBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartBadRequest ", 400)
}

func (o *StackStartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackStartForbidden creates a StackStartForbidden with default headers values
func NewStackStartForbidden() *StackStartForbidden {
	return &StackStartForbidden{}
}

/*
StackStartForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackStartForbidden struct {
}

// IsSuccess returns true when this stack start forbidden response has a 2xx status code
func (o *StackStartForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack start forbidden response has a 3xx status code
func (o *StackStartForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack start forbidden response has a 4xx status code
func (o *StackStartForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack start forbidden response has a 5xx status code
func (o *StackStartForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stack start forbidden response a status code equal to that given
func (o *StackStartForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StackStartForbidden) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartForbidden ", 403)
}

func (o *StackStartForbidden) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartForbidden ", 403)
}

func (o *StackStartForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackStartNotFound creates a StackStartNotFound with default headers values
func NewStackStartNotFound() *StackStartNotFound {
	return &StackStartNotFound{}
}

/*
StackStartNotFound describes a response with status code 404, with default header values.

Not found
*/
type StackStartNotFound struct {
}

// IsSuccess returns true when this stack start not found response has a 2xx status code
func (o *StackStartNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack start not found response has a 3xx status code
func (o *StackStartNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack start not found response has a 4xx status code
func (o *StackStartNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack start not found response has a 5xx status code
func (o *StackStartNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stack start not found response a status code equal to that given
func (o *StackStartNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StackStartNotFound) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartNotFound ", 404)
}

func (o *StackStartNotFound) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartNotFound ", 404)
}

func (o *StackStartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackStartInternalServerError creates a StackStartInternalServerError with default headers values
func NewStackStartInternalServerError() *StackStartInternalServerError {
	return &StackStartInternalServerError{}
}

/*
StackStartInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackStartInternalServerError struct {
}

// IsSuccess returns true when this stack start internal server error response has a 2xx status code
func (o *StackStartInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack start internal server error response has a 3xx status code
func (o *StackStartInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack start internal server error response has a 4xx status code
func (o *StackStartInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack start internal server error response has a 5xx status code
func (o *StackStartInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack start internal server error response a status code equal to that given
func (o *StackStartInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StackStartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartInternalServerError ", 500)
}

func (o *StackStartInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/start][%d] stackStartInternalServerError ", 500)
}

func (o *StackStartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
