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

// StackCreateKubernetesGitReader is a Reader for the StackCreateKubernetesGit structure.
type StackCreateKubernetesGitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackCreateKubernetesGitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackCreateKubernetesGitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackCreateKubernetesGitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewStackCreateKubernetesGitConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackCreateKubernetesGitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /stacks/create/kubernetes/repository] StackCreateKubernetesGit", response, response.Code())
	}
}

// NewStackCreateKubernetesGitOK creates a StackCreateKubernetesGitOK with default headers values
func NewStackCreateKubernetesGitOK() *StackCreateKubernetesGitOK {
	return &StackCreateKubernetesGitOK{}
}

/*
StackCreateKubernetesGitOK describes a response with status code 200, with default header values.

OK
*/
type StackCreateKubernetesGitOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack create kubernetes git o k response has a 2xx status code
func (o *StackCreateKubernetesGitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack create kubernetes git o k response has a 3xx status code
func (o *StackCreateKubernetesGitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create kubernetes git o k response has a 4xx status code
func (o *StackCreateKubernetesGitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create kubernetes git o k response has a 5xx status code
func (o *StackCreateKubernetesGitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create kubernetes git o k response a status code equal to that given
func (o *StackCreateKubernetesGitOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stack create kubernetes git o k response
func (o *StackCreateKubernetesGitOK) Code() int {
	return 200
}

func (o *StackCreateKubernetesGitOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitOK %s", 200, payload)
}

func (o *StackCreateKubernetesGitOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitOK %s", 200, payload)
}

func (o *StackCreateKubernetesGitOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackCreateKubernetesGitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCreateKubernetesGitBadRequest creates a StackCreateKubernetesGitBadRequest with default headers values
func NewStackCreateKubernetesGitBadRequest() *StackCreateKubernetesGitBadRequest {
	return &StackCreateKubernetesGitBadRequest{}
}

/*
StackCreateKubernetesGitBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackCreateKubernetesGitBadRequest struct {
}

// IsSuccess returns true when this stack create kubernetes git bad request response has a 2xx status code
func (o *StackCreateKubernetesGitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create kubernetes git bad request response has a 3xx status code
func (o *StackCreateKubernetesGitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create kubernetes git bad request response has a 4xx status code
func (o *StackCreateKubernetesGitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack create kubernetes git bad request response has a 5xx status code
func (o *StackCreateKubernetesGitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create kubernetes git bad request response a status code equal to that given
func (o *StackCreateKubernetesGitBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stack create kubernetes git bad request response
func (o *StackCreateKubernetesGitBadRequest) Code() int {
	return 400
}

func (o *StackCreateKubernetesGitBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitBadRequest", 400)
}

func (o *StackCreateKubernetesGitBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitBadRequest", 400)
}

func (o *StackCreateKubernetesGitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackCreateKubernetesGitConflict creates a StackCreateKubernetesGitConflict with default headers values
func NewStackCreateKubernetesGitConflict() *StackCreateKubernetesGitConflict {
	return &StackCreateKubernetesGitConflict{}
}

/*
StackCreateKubernetesGitConflict describes a response with status code 409, with default header values.

Webhook ID already exists
*/
type StackCreateKubernetesGitConflict struct {
}

// IsSuccess returns true when this stack create kubernetes git conflict response has a 2xx status code
func (o *StackCreateKubernetesGitConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create kubernetes git conflict response has a 3xx status code
func (o *StackCreateKubernetesGitConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create kubernetes git conflict response has a 4xx status code
func (o *StackCreateKubernetesGitConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack create kubernetes git conflict response has a 5xx status code
func (o *StackCreateKubernetesGitConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create kubernetes git conflict response a status code equal to that given
func (o *StackCreateKubernetesGitConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the stack create kubernetes git conflict response
func (o *StackCreateKubernetesGitConflict) Code() int {
	return 409
}

func (o *StackCreateKubernetesGitConflict) Error() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitConflict", 409)
}

func (o *StackCreateKubernetesGitConflict) String() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitConflict", 409)
}

func (o *StackCreateKubernetesGitConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackCreateKubernetesGitInternalServerError creates a StackCreateKubernetesGitInternalServerError with default headers values
func NewStackCreateKubernetesGitInternalServerError() *StackCreateKubernetesGitInternalServerError {
	return &StackCreateKubernetesGitInternalServerError{}
}

/*
StackCreateKubernetesGitInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackCreateKubernetesGitInternalServerError struct {
}

// IsSuccess returns true when this stack create kubernetes git internal server error response has a 2xx status code
func (o *StackCreateKubernetesGitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create kubernetes git internal server error response has a 3xx status code
func (o *StackCreateKubernetesGitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create kubernetes git internal server error response has a 4xx status code
func (o *StackCreateKubernetesGitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create kubernetes git internal server error response has a 5xx status code
func (o *StackCreateKubernetesGitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack create kubernetes git internal server error response a status code equal to that given
func (o *StackCreateKubernetesGitInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stack create kubernetes git internal server error response
func (o *StackCreateKubernetesGitInternalServerError) Code() int {
	return 500
}

func (o *StackCreateKubernetesGitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitInternalServerError", 500)
}

func (o *StackCreateKubernetesGitInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/repository][%d] stackCreateKubernetesGitInternalServerError", 500)
}

func (o *StackCreateKubernetesGitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
