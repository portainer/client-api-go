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

// StackCreateKubernetesFileReader is a Reader for the StackCreateKubernetesFile structure.
type StackCreateKubernetesFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackCreateKubernetesFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackCreateKubernetesFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackCreateKubernetesFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackCreateKubernetesFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /stacks/create/kubernetes/string] StackCreateKubernetesFile", response, response.Code())
	}
}

// NewStackCreateKubernetesFileOK creates a StackCreateKubernetesFileOK with default headers values
func NewStackCreateKubernetesFileOK() *StackCreateKubernetesFileOK {
	return &StackCreateKubernetesFileOK{}
}

/*
StackCreateKubernetesFileOK describes a response with status code 200, with default header values.

OK
*/
type StackCreateKubernetesFileOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack create kubernetes file o k response has a 2xx status code
func (o *StackCreateKubernetesFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack create kubernetes file o k response has a 3xx status code
func (o *StackCreateKubernetesFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create kubernetes file o k response has a 4xx status code
func (o *StackCreateKubernetesFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create kubernetes file o k response has a 5xx status code
func (o *StackCreateKubernetesFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create kubernetes file o k response a status code equal to that given
func (o *StackCreateKubernetesFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stack create kubernetes file o k response
func (o *StackCreateKubernetesFileOK) Code() int {
	return 200
}

func (o *StackCreateKubernetesFileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stacks/create/kubernetes/string][%d] stackCreateKubernetesFileOK %s", 200, payload)
}

func (o *StackCreateKubernetesFileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stacks/create/kubernetes/string][%d] stackCreateKubernetesFileOK %s", 200, payload)
}

func (o *StackCreateKubernetesFileOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackCreateKubernetesFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCreateKubernetesFileBadRequest creates a StackCreateKubernetesFileBadRequest with default headers values
func NewStackCreateKubernetesFileBadRequest() *StackCreateKubernetesFileBadRequest {
	return &StackCreateKubernetesFileBadRequest{}
}

/*
StackCreateKubernetesFileBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackCreateKubernetesFileBadRequest struct {
}

// IsSuccess returns true when this stack create kubernetes file bad request response has a 2xx status code
func (o *StackCreateKubernetesFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create kubernetes file bad request response has a 3xx status code
func (o *StackCreateKubernetesFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create kubernetes file bad request response has a 4xx status code
func (o *StackCreateKubernetesFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack create kubernetes file bad request response has a 5xx status code
func (o *StackCreateKubernetesFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create kubernetes file bad request response a status code equal to that given
func (o *StackCreateKubernetesFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stack create kubernetes file bad request response
func (o *StackCreateKubernetesFileBadRequest) Code() int {
	return 400
}

func (o *StackCreateKubernetesFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/string][%d] stackCreateKubernetesFileBadRequest", 400)
}

func (o *StackCreateKubernetesFileBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/string][%d] stackCreateKubernetesFileBadRequest", 400)
}

func (o *StackCreateKubernetesFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackCreateKubernetesFileInternalServerError creates a StackCreateKubernetesFileInternalServerError with default headers values
func NewStackCreateKubernetesFileInternalServerError() *StackCreateKubernetesFileInternalServerError {
	return &StackCreateKubernetesFileInternalServerError{}
}

/*
StackCreateKubernetesFileInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackCreateKubernetesFileInternalServerError struct {
}

// IsSuccess returns true when this stack create kubernetes file internal server error response has a 2xx status code
func (o *StackCreateKubernetesFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create kubernetes file internal server error response has a 3xx status code
func (o *StackCreateKubernetesFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create kubernetes file internal server error response has a 4xx status code
func (o *StackCreateKubernetesFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create kubernetes file internal server error response has a 5xx status code
func (o *StackCreateKubernetesFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack create kubernetes file internal server error response a status code equal to that given
func (o *StackCreateKubernetesFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stack create kubernetes file internal server error response
func (o *StackCreateKubernetesFileInternalServerError) Code() int {
	return 500
}

func (o *StackCreateKubernetesFileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/string][%d] stackCreateKubernetesFileInternalServerError", 500)
}

func (o *StackCreateKubernetesFileInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks/create/kubernetes/string][%d] stackCreateKubernetesFileInternalServerError", 500)
}

func (o *StackCreateKubernetesFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
