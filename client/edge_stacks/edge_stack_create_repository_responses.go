// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

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

// EdgeStackCreateRepositoryReader is a Reader for the EdgeStackCreateRepository structure.
type EdgeStackCreateRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackCreateRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackCreateRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackCreateRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeStackCreateRepositoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackCreateRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackCreateRepositoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /edge_stacks/create/repository] EdgeStackCreateRepository", response, response.Code())
	}
}

// NewEdgeStackCreateRepositoryOK creates a EdgeStackCreateRepositoryOK with default headers values
func NewEdgeStackCreateRepositoryOK() *EdgeStackCreateRepositoryOK {
	return &EdgeStackCreateRepositoryOK{}
}

/*
EdgeStackCreateRepositoryOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackCreateRepositoryOK struct {
	Payload *models.PortainereeEdgeStack
}

// IsSuccess returns true when this edge stack create repository o k response has a 2xx status code
func (o *EdgeStackCreateRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack create repository o k response has a 3xx status code
func (o *EdgeStackCreateRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create repository o k response has a 4xx status code
func (o *EdgeStackCreateRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create repository o k response has a 5xx status code
func (o *EdgeStackCreateRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack create repository o k response a status code equal to that given
func (o *EdgeStackCreateRepositoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge stack create repository o k response
func (o *EdgeStackCreateRepositoryOK) Code() int {
	return 200
}

func (o *EdgeStackCreateRepositoryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryOK %s", 200, payload)
}

func (o *EdgeStackCreateRepositoryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryOK %s", 200, payload)
}

func (o *EdgeStackCreateRepositoryOK) GetPayload() *models.PortainereeEdgeStack {
	return o.Payload
}

func (o *EdgeStackCreateRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeEdgeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackCreateRepositoryBadRequest creates a EdgeStackCreateRepositoryBadRequest with default headers values
func NewEdgeStackCreateRepositoryBadRequest() *EdgeStackCreateRepositoryBadRequest {
	return &EdgeStackCreateRepositoryBadRequest{}
}

/*
EdgeStackCreateRepositoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type EdgeStackCreateRepositoryBadRequest struct {
}

// IsSuccess returns true when this edge stack create repository bad request response has a 2xx status code
func (o *EdgeStackCreateRepositoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create repository bad request response has a 3xx status code
func (o *EdgeStackCreateRepositoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create repository bad request response has a 4xx status code
func (o *EdgeStackCreateRepositoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack create repository bad request response has a 5xx status code
func (o *EdgeStackCreateRepositoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack create repository bad request response a status code equal to that given
func (o *EdgeStackCreateRepositoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge stack create repository bad request response
func (o *EdgeStackCreateRepositoryBadRequest) Code() int {
	return 400
}

func (o *EdgeStackCreateRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryBadRequest", 400)
}

func (o *EdgeStackCreateRepositoryBadRequest) String() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryBadRequest", 400)
}

func (o *EdgeStackCreateRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackCreateRepositoryConflict creates a EdgeStackCreateRepositoryConflict with default headers values
func NewEdgeStackCreateRepositoryConflict() *EdgeStackCreateRepositoryConflict {
	return &EdgeStackCreateRepositoryConflict{}
}

/*
EdgeStackCreateRepositoryConflict describes a response with status code 409, with default header values.

Webhook ID already exists
*/
type EdgeStackCreateRepositoryConflict struct {
}

// IsSuccess returns true when this edge stack create repository conflict response has a 2xx status code
func (o *EdgeStackCreateRepositoryConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create repository conflict response has a 3xx status code
func (o *EdgeStackCreateRepositoryConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create repository conflict response has a 4xx status code
func (o *EdgeStackCreateRepositoryConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack create repository conflict response has a 5xx status code
func (o *EdgeStackCreateRepositoryConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack create repository conflict response a status code equal to that given
func (o *EdgeStackCreateRepositoryConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge stack create repository conflict response
func (o *EdgeStackCreateRepositoryConflict) Code() int {
	return 409
}

func (o *EdgeStackCreateRepositoryConflict) Error() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryConflict", 409)
}

func (o *EdgeStackCreateRepositoryConflict) String() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryConflict", 409)
}

func (o *EdgeStackCreateRepositoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackCreateRepositoryInternalServerError creates a EdgeStackCreateRepositoryInternalServerError with default headers values
func NewEdgeStackCreateRepositoryInternalServerError() *EdgeStackCreateRepositoryInternalServerError {
	return &EdgeStackCreateRepositoryInternalServerError{}
}

/*
EdgeStackCreateRepositoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type EdgeStackCreateRepositoryInternalServerError struct {
}

// IsSuccess returns true when this edge stack create repository internal server error response has a 2xx status code
func (o *EdgeStackCreateRepositoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create repository internal server error response has a 3xx status code
func (o *EdgeStackCreateRepositoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create repository internal server error response has a 4xx status code
func (o *EdgeStackCreateRepositoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create repository internal server error response has a 5xx status code
func (o *EdgeStackCreateRepositoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack create repository internal server error response a status code equal to that given
func (o *EdgeStackCreateRepositoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge stack create repository internal server error response
func (o *EdgeStackCreateRepositoryInternalServerError) Code() int {
	return 500
}

func (o *EdgeStackCreateRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryInternalServerError", 500)
}

func (o *EdgeStackCreateRepositoryInternalServerError) String() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryInternalServerError", 500)
}

func (o *EdgeStackCreateRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackCreateRepositoryServiceUnavailable creates a EdgeStackCreateRepositoryServiceUnavailable with default headers values
func NewEdgeStackCreateRepositoryServiceUnavailable() *EdgeStackCreateRepositoryServiceUnavailable {
	return &EdgeStackCreateRepositoryServiceUnavailable{}
}

/*
EdgeStackCreateRepositoryServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackCreateRepositoryServiceUnavailable struct {
}

// IsSuccess returns true when this edge stack create repository service unavailable response has a 2xx status code
func (o *EdgeStackCreateRepositoryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create repository service unavailable response has a 3xx status code
func (o *EdgeStackCreateRepositoryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create repository service unavailable response has a 4xx status code
func (o *EdgeStackCreateRepositoryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create repository service unavailable response has a 5xx status code
func (o *EdgeStackCreateRepositoryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack create repository service unavailable response a status code equal to that given
func (o *EdgeStackCreateRepositoryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge stack create repository service unavailable response
func (o *EdgeStackCreateRepositoryServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeStackCreateRepositoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryServiceUnavailable", 503)
}

func (o *EdgeStackCreateRepositoryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /edge_stacks/create/repository][%d] edgeStackCreateRepositoryServiceUnavailable", 503)
}

func (o *EdgeStackCreateRepositoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
