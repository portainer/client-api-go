// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ee/v2/models"
)

// EndpointListReader is a Reader for the EndpointList structure.
type EndpointListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEndpointListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointListOK creates a EndpointListOK with default headers values
func NewEndpointListOK() *EndpointListOK {
	return &EndpointListOK{}
}

/*
EndpointListOK describes a response with status code 200, with default header values.

Endpoints
*/
type EndpointListOK struct {
	Payload []*models.PortainereeEndpoint
}

// IsSuccess returns true when this endpoint list o k response has a 2xx status code
func (o *EndpointListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint list o k response has a 3xx status code
func (o *EndpointListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint list o k response has a 4xx status code
func (o *EndpointListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint list o k response has a 5xx status code
func (o *EndpointListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint list o k response a status code equal to that given
func (o *EndpointListOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointListOK) Error() string {
	return fmt.Sprintf("[GET /endpoints][%d] endpointListOK  %+v", 200, o.Payload)
}

func (o *EndpointListOK) String() string {
	return fmt.Sprintf("[GET /endpoints][%d] endpointListOK  %+v", 200, o.Payload)
}

func (o *EndpointListOK) GetPayload() []*models.PortainereeEndpoint {
	return o.Payload
}

func (o *EndpointListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointListInternalServerError creates a EndpointListInternalServerError with default headers values
func NewEndpointListInternalServerError() *EndpointListInternalServerError {
	return &EndpointListInternalServerError{}
}

/*
EndpointListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointListInternalServerError struct {
}

// IsSuccess returns true when this endpoint list internal server error response has a 2xx status code
func (o *EndpointListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint list internal server error response has a 3xx status code
func (o *EndpointListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint list internal server error response has a 4xx status code
func (o *EndpointListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint list internal server error response has a 5xx status code
func (o *EndpointListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint list internal server error response a status code equal to that given
func (o *EndpointListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoints][%d] endpointListInternalServerError ", 500)
}

func (o *EndpointListInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoints][%d] endpointListInternalServerError ", 500)
}

func (o *EndpointListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
