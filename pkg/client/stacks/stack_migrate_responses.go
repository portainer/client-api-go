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

// StackMigrateReader is a Reader for the StackMigrate structure.
type StackMigrateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackMigrateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackMigrateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackMigrateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackMigrateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackMigrateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewStackMigrateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackMigrateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /stacks/{id}/migrate] StackMigrate", response, response.Code())
	}
}

// NewStackMigrateOK creates a StackMigrateOK with default headers values
func NewStackMigrateOK() *StackMigrateOK {
	return &StackMigrateOK{}
}

/*
StackMigrateOK describes a response with status code 200, with default header values.

Success
*/
type StackMigrateOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack migrate o k response has a 2xx status code
func (o *StackMigrateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack migrate o k response has a 3xx status code
func (o *StackMigrateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack migrate o k response has a 4xx status code
func (o *StackMigrateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack migrate o k response has a 5xx status code
func (o *StackMigrateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack migrate o k response a status code equal to that given
func (o *StackMigrateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stack migrate o k response
func (o *StackMigrateOK) Code() int {
	return 200
}

func (o *StackMigrateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateOK %s", 200, payload)
}

func (o *StackMigrateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateOK %s", 200, payload)
}

func (o *StackMigrateOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackMigrateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackMigrateBadRequest creates a StackMigrateBadRequest with default headers values
func NewStackMigrateBadRequest() *StackMigrateBadRequest {
	return &StackMigrateBadRequest{}
}

/*
StackMigrateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackMigrateBadRequest struct {
}

// IsSuccess returns true when this stack migrate bad request response has a 2xx status code
func (o *StackMigrateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack migrate bad request response has a 3xx status code
func (o *StackMigrateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack migrate bad request response has a 4xx status code
func (o *StackMigrateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack migrate bad request response has a 5xx status code
func (o *StackMigrateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack migrate bad request response a status code equal to that given
func (o *StackMigrateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stack migrate bad request response
func (o *StackMigrateBadRequest) Code() int {
	return 400
}

func (o *StackMigrateBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateBadRequest", 400)
}

func (o *StackMigrateBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateBadRequest", 400)
}

func (o *StackMigrateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackMigrateForbidden creates a StackMigrateForbidden with default headers values
func NewStackMigrateForbidden() *StackMigrateForbidden {
	return &StackMigrateForbidden{}
}

/*
StackMigrateForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackMigrateForbidden struct {
}

// IsSuccess returns true when this stack migrate forbidden response has a 2xx status code
func (o *StackMigrateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack migrate forbidden response has a 3xx status code
func (o *StackMigrateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack migrate forbidden response has a 4xx status code
func (o *StackMigrateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack migrate forbidden response has a 5xx status code
func (o *StackMigrateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stack migrate forbidden response a status code equal to that given
func (o *StackMigrateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stack migrate forbidden response
func (o *StackMigrateForbidden) Code() int {
	return 403
}

func (o *StackMigrateForbidden) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateForbidden", 403)
}

func (o *StackMigrateForbidden) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateForbidden", 403)
}

func (o *StackMigrateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackMigrateNotFound creates a StackMigrateNotFound with default headers values
func NewStackMigrateNotFound() *StackMigrateNotFound {
	return &StackMigrateNotFound{}
}

/*
StackMigrateNotFound describes a response with status code 404, with default header values.

Stack not found
*/
type StackMigrateNotFound struct {
}

// IsSuccess returns true when this stack migrate not found response has a 2xx status code
func (o *StackMigrateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack migrate not found response has a 3xx status code
func (o *StackMigrateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack migrate not found response has a 4xx status code
func (o *StackMigrateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack migrate not found response has a 5xx status code
func (o *StackMigrateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stack migrate not found response a status code equal to that given
func (o *StackMigrateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stack migrate not found response
func (o *StackMigrateNotFound) Code() int {
	return 404
}

func (o *StackMigrateNotFound) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateNotFound", 404)
}

func (o *StackMigrateNotFound) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateNotFound", 404)
}

func (o *StackMigrateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackMigrateConflict creates a StackMigrateConflict with default headers values
func NewStackMigrateConflict() *StackMigrateConflict {
	return &StackMigrateConflict{}
}

/*
StackMigrateConflict describes a response with status code 409, with default header values.

A stack with the same name is already running on the target environment(endpoint)
*/
type StackMigrateConflict struct {
}

// IsSuccess returns true when this stack migrate conflict response has a 2xx status code
func (o *StackMigrateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack migrate conflict response has a 3xx status code
func (o *StackMigrateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack migrate conflict response has a 4xx status code
func (o *StackMigrateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack migrate conflict response has a 5xx status code
func (o *StackMigrateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this stack migrate conflict response a status code equal to that given
func (o *StackMigrateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the stack migrate conflict response
func (o *StackMigrateConflict) Code() int {
	return 409
}

func (o *StackMigrateConflict) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateConflict", 409)
}

func (o *StackMigrateConflict) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateConflict", 409)
}

func (o *StackMigrateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackMigrateInternalServerError creates a StackMigrateInternalServerError with default headers values
func NewStackMigrateInternalServerError() *StackMigrateInternalServerError {
	return &StackMigrateInternalServerError{}
}

/*
StackMigrateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackMigrateInternalServerError struct {
}

// IsSuccess returns true when this stack migrate internal server error response has a 2xx status code
func (o *StackMigrateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack migrate internal server error response has a 3xx status code
func (o *StackMigrateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack migrate internal server error response has a 4xx status code
func (o *StackMigrateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack migrate internal server error response has a 5xx status code
func (o *StackMigrateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack migrate internal server error response a status code equal to that given
func (o *StackMigrateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stack migrate internal server error response
func (o *StackMigrateInternalServerError) Code() int {
	return 500
}

func (o *StackMigrateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateInternalServerError", 500)
}

func (o *StackMigrateInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks/{id}/migrate][%d] stackMigrateInternalServerError", 500)
}

func (o *StackMigrateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
