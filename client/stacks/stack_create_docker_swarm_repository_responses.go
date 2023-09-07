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

// StackCreateDockerSwarmRepositoryReader is a Reader for the StackCreateDockerSwarmRepository structure.
type StackCreateDockerSwarmRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackCreateDockerSwarmRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStackCreateDockerSwarmRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackCreateDockerSwarmRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackCreateDockerSwarmRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackCreateDockerSwarmRepositoryOK creates a StackCreateDockerSwarmRepositoryOK with default headers values
func NewStackCreateDockerSwarmRepositoryOK() *StackCreateDockerSwarmRepositoryOK {
	return &StackCreateDockerSwarmRepositoryOK{}
}

/*
StackCreateDockerSwarmRepositoryOK describes a response with status code 200, with default header values.

OK
*/
type StackCreateDockerSwarmRepositoryOK struct {
	Payload *models.PortainereeStack
}

// IsSuccess returns true when this stack create docker swarm repository o k response has a 2xx status code
func (o *StackCreateDockerSwarmRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stack create docker swarm repository o k response has a 3xx status code
func (o *StackCreateDockerSwarmRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create docker swarm repository o k response has a 4xx status code
func (o *StackCreateDockerSwarmRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create docker swarm repository o k response has a 5xx status code
func (o *StackCreateDockerSwarmRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create docker swarm repository o k response a status code equal to that given
func (o *StackCreateDockerSwarmRepositoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *StackCreateDockerSwarmRepositoryOK) Error() string {
	return fmt.Sprintf("[POST /stacks/create/swarm/repository][%d] stackCreateDockerSwarmRepositoryOK  %+v", 200, o.Payload)
}

func (o *StackCreateDockerSwarmRepositoryOK) String() string {
	return fmt.Sprintf("[POST /stacks/create/swarm/repository][%d] stackCreateDockerSwarmRepositoryOK  %+v", 200, o.Payload)
}

func (o *StackCreateDockerSwarmRepositoryOK) GetPayload() *models.PortainereeStack {
	return o.Payload
}

func (o *StackCreateDockerSwarmRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStackCreateDockerSwarmRepositoryBadRequest creates a StackCreateDockerSwarmRepositoryBadRequest with default headers values
func NewStackCreateDockerSwarmRepositoryBadRequest() *StackCreateDockerSwarmRepositoryBadRequest {
	return &StackCreateDockerSwarmRepositoryBadRequest{}
}

/*
StackCreateDockerSwarmRepositoryBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackCreateDockerSwarmRepositoryBadRequest struct {
}

// IsSuccess returns true when this stack create docker swarm repository bad request response has a 2xx status code
func (o *StackCreateDockerSwarmRepositoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create docker swarm repository bad request response has a 3xx status code
func (o *StackCreateDockerSwarmRepositoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create docker swarm repository bad request response has a 4xx status code
func (o *StackCreateDockerSwarmRepositoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stack create docker swarm repository bad request response has a 5xx status code
func (o *StackCreateDockerSwarmRepositoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stack create docker swarm repository bad request response a status code equal to that given
func (o *StackCreateDockerSwarmRepositoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StackCreateDockerSwarmRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /stacks/create/swarm/repository][%d] stackCreateDockerSwarmRepositoryBadRequest ", 400)
}

func (o *StackCreateDockerSwarmRepositoryBadRequest) String() string {
	return fmt.Sprintf("[POST /stacks/create/swarm/repository][%d] stackCreateDockerSwarmRepositoryBadRequest ", 400)
}

func (o *StackCreateDockerSwarmRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackCreateDockerSwarmRepositoryInternalServerError creates a StackCreateDockerSwarmRepositoryInternalServerError with default headers values
func NewStackCreateDockerSwarmRepositoryInternalServerError() *StackCreateDockerSwarmRepositoryInternalServerError {
	return &StackCreateDockerSwarmRepositoryInternalServerError{}
}

/*
StackCreateDockerSwarmRepositoryInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackCreateDockerSwarmRepositoryInternalServerError struct {
}

// IsSuccess returns true when this stack create docker swarm repository internal server error response has a 2xx status code
func (o *StackCreateDockerSwarmRepositoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stack create docker swarm repository internal server error response has a 3xx status code
func (o *StackCreateDockerSwarmRepositoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stack create docker swarm repository internal server error response has a 4xx status code
func (o *StackCreateDockerSwarmRepositoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stack create docker swarm repository internal server error response has a 5xx status code
func (o *StackCreateDockerSwarmRepositoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stack create docker swarm repository internal server error response a status code equal to that given
func (o *StackCreateDockerSwarmRepositoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StackCreateDockerSwarmRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stacks/create/swarm/repository][%d] stackCreateDockerSwarmRepositoryInternalServerError ", 500)
}

func (o *StackCreateDockerSwarmRepositoryInternalServerError) String() string {
	return fmt.Sprintf("[POST /stacks/create/swarm/repository][%d] stackCreateDockerSwarmRepositoryInternalServerError ", 500)
}

func (o *StackCreateDockerSwarmRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
