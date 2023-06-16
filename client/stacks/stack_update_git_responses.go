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

// StackUpdateGitReader is a Reader for the StackUpdateGit structure.
type StackUpdateGitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackUpdateGitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackUpdateGitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackUpdateGitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackUpdateGitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackUpdateGitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackUpdateGitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackUpdateGitOK creates a StackUpdateGitOK with default headers values
func NewStackUpdateGitOK() *StackUpdateGitOK {
	return &StackUpdateGitOK{}
}

/*
StackUpdateGitOK describes a response with status code 200, with default header values.

Success
*/
type StackUpdateGitOK struct {
	Payload *models.PortainerStack
}

// IsSuccess returns true when this stack update git o k response has a 2xx status code
func (o *StackUpdateGitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack update git o k response has a 3xx status code
func (o *StackUpdateGitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack update git o k response has a 4xx status code
func (o *StackUpdateGitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack update git o k response has a 5xx status code
func (o *StackUpdateGitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack update git o k response a status code equal to that given
func (o *StackUpdateGitOK) IsCode(code int) bool {
	return code == 200
}

func (o *StackUpdateGitOK) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitOK  %+v", 200, o.Payload)
}

func (o *StackUpdateGitOK) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitOK  %+v", 200, o.Payload)
}

func (o *StackUpdateGitOK) GetPayload() *models.PortainerStack {
	return o.Payload
}

func (o *StackUpdateGitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackUpdateGitBadRequest creates a StackUpdateGitBadRequest with default headers values
func NewStackUpdateGitBadRequest() *StackUpdateGitBadRequest {
	return &StackUpdateGitBadRequest{}
}

/*
StackUpdateGitBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackUpdateGitBadRequest struct {
}

// IsSuccess returns true when this stack update git bad request response has a 2xx status code
func (o *StackUpdateGitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack update git bad request response has a 3xx status code
func (o *StackUpdateGitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack update git bad request response has a 4xx status code
func (o *StackUpdateGitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack update git bad request response has a 5xx status code
func (o *StackUpdateGitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack update git bad request response a status code equal to that given
func (o *StackUpdateGitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StackUpdateGitBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitBadRequest ", 400)
}

func (o *StackUpdateGitBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitBadRequest ", 400)
}

func (o *StackUpdateGitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackUpdateGitForbidden creates a StackUpdateGitForbidden with default headers values
func NewStackUpdateGitForbidden() *StackUpdateGitForbidden {
	return &StackUpdateGitForbidden{}
}

/*
StackUpdateGitForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackUpdateGitForbidden struct {
}

// IsSuccess returns true when this stack update git forbidden response has a 2xx status code
func (o *StackUpdateGitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack update git forbidden response has a 3xx status code
func (o *StackUpdateGitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack update git forbidden response has a 4xx status code
func (o *StackUpdateGitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack update git forbidden response has a 5xx status code
func (o *StackUpdateGitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stack update git forbidden response a status code equal to that given
func (o *StackUpdateGitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StackUpdateGitForbidden) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitForbidden ", 403)
}

func (o *StackUpdateGitForbidden) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitForbidden ", 403)
}

func (o *StackUpdateGitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackUpdateGitNotFound creates a StackUpdateGitNotFound with default headers values
func NewStackUpdateGitNotFound() *StackUpdateGitNotFound {
	return &StackUpdateGitNotFound{}
}

/*
StackUpdateGitNotFound describes a response with status code 404, with default header values.

Not found
*/
type StackUpdateGitNotFound struct {
}

// IsSuccess returns true when this stack update git not found response has a 2xx status code
func (o *StackUpdateGitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack update git not found response has a 3xx status code
func (o *StackUpdateGitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack update git not found response has a 4xx status code
func (o *StackUpdateGitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack update git not found response has a 5xx status code
func (o *StackUpdateGitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stack update git not found response a status code equal to that given
func (o *StackUpdateGitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StackUpdateGitNotFound) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitNotFound ", 404)
}

func (o *StackUpdateGitNotFound) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitNotFound ", 404)
}

func (o *StackUpdateGitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackUpdateGitInternalServerError creates a StackUpdateGitInternalServerError with default headers values
func NewStackUpdateGitInternalServerError() *StackUpdateGitInternalServerError {
	return &StackUpdateGitInternalServerError{}
}

/*
StackUpdateGitInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackUpdateGitInternalServerError struct {
}

// IsSuccess returns true when this stack update git internal server error response has a 2xx status code
func (o *StackUpdateGitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack update git internal server error response has a 3xx status code
func (o *StackUpdateGitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack update git internal server error response has a 4xx status code
func (o *StackUpdateGitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack update git internal server error response has a 5xx status code
func (o *StackUpdateGitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack update git internal server error response a status code equal to that given
func (o *StackUpdateGitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StackUpdateGitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitInternalServerError ", 500)
}

func (o *StackUpdateGitInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/git][%d] stackUpdateGitInternalServerError ", 500)
}

func (o *StackUpdateGitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
