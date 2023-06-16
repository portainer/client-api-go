// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// EdgeStackCreateReader is a Reader for the EdgeStackCreate structure.
type EdgeStackCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeStackCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeStackCreateOK creates a EdgeStackCreateOK with default headers values
func NewEdgeStackCreateOK() *EdgeStackCreateOK {
	return &EdgeStackCreateOK{}
}

/*
EdgeStackCreateOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackCreateOK struct {
	Payload *models_ce.PortainerEdgeStack
}

// IsSuccess returns true when this edge stack create o k response has a 2xx status code
func (o *EdgeStackCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack create o k response has a 3xx status code
func (o *EdgeStackCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create o k response has a 4xx status code
func (o *EdgeStackCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create o k response has a 5xx status code
func (o *EdgeStackCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack create o k response a status code equal to that given
func (o *EdgeStackCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeStackCreateOK) Error() string {
	return fmt.Sprintf("[POST /edge_stacks][%d] edgeStackCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeStackCreateOK) String() string {
	return fmt.Sprintf("[POST /edge_stacks][%d] edgeStackCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeStackCreateOK) GetPayload() *models_ce.PortainerEdgeStack {
	return o.Payload
}

func (o *EdgeStackCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.PortainerEdgeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackCreateInternalServerError creates a EdgeStackCreateInternalServerError with default headers values
func NewEdgeStackCreateInternalServerError() *EdgeStackCreateInternalServerError {
	return &EdgeStackCreateInternalServerError{}
}

/*
EdgeStackCreateInternalServerError describes a response with status code 500, with default header values.

EdgeStackCreateInternalServerError edge stack create internal server error
*/
type EdgeStackCreateInternalServerError struct {
}

// IsSuccess returns true when this edge stack create internal server error response has a 2xx status code
func (o *EdgeStackCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create internal server error response has a 3xx status code
func (o *EdgeStackCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create internal server error response has a 4xx status code
func (o *EdgeStackCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create internal server error response has a 5xx status code
func (o *EdgeStackCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack create internal server error response a status code equal to that given
func (o *EdgeStackCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeStackCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /edge_stacks][%d] edgeStackCreateInternalServerError ", 500)
}

func (o *EdgeStackCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /edge_stacks][%d] edgeStackCreateInternalServerError ", 500)
}

func (o *EdgeStackCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackCreateServiceUnavailable creates a EdgeStackCreateServiceUnavailable with default headers values
func NewEdgeStackCreateServiceUnavailable() *EdgeStackCreateServiceUnavailable {
	return &EdgeStackCreateServiceUnavailable{}
}

/*
EdgeStackCreateServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackCreateServiceUnavailable struct {
}

// IsSuccess returns true when this edge stack create service unavailable response has a 2xx status code
func (o *EdgeStackCreateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack create service unavailable response has a 3xx status code
func (o *EdgeStackCreateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack create service unavailable response has a 4xx status code
func (o *EdgeStackCreateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack create service unavailable response has a 5xx status code
func (o *EdgeStackCreateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack create service unavailable response a status code equal to that given
func (o *EdgeStackCreateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeStackCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /edge_stacks][%d] edgeStackCreateServiceUnavailable ", 503)
}

func (o *EdgeStackCreateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /edge_stacks][%d] edgeStackCreateServiceUnavailable ", 503)
}

func (o *EdgeStackCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
