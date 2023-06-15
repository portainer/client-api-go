// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/ce/models"
)

// EndpointInspectReader is a Reader for the EndpointInspect structure.
type EndpointInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointInspectOK creates a EndpointInspectOK with default headers values
func NewEndpointInspectOK() *EndpointInspectOK {
	return &EndpointInspectOK{}
}

/*
EndpointInspectOK describes a response with status code 200, with default header values.

Success
*/
type EndpointInspectOK struct {
	Payload *models.PortainerEndpoint
}

// IsSuccess returns true when this endpoint inspect o k response has a 2xx status code
func (o *EndpointInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint inspect o k response has a 3xx status code
func (o *EndpointInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint inspect o k response has a 4xx status code
func (o *EndpointInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint inspect o k response has a 5xx status code
func (o *EndpointInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint inspect o k response a status code equal to that given
func (o *EndpointInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointInspectOK) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectOK  %+v", 200, o.Payload)
}

func (o *EndpointInspectOK) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectOK  %+v", 200, o.Payload)
}

func (o *EndpointInspectOK) GetPayload() *models.PortainerEndpoint {
	return o.Payload
}

func (o *EndpointInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointInspectBadRequest creates a EndpointInspectBadRequest with default headers values
func NewEndpointInspectBadRequest() *EndpointInspectBadRequest {
	return &EndpointInspectBadRequest{}
}

/*
EndpointInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointInspectBadRequest struct {
}

// IsSuccess returns true when this endpoint inspect bad request response has a 2xx status code
func (o *EndpointInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint inspect bad request response has a 3xx status code
func (o *EndpointInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint inspect bad request response has a 4xx status code
func (o *EndpointInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint inspect bad request response has a 5xx status code
func (o *EndpointInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint inspect bad request response a status code equal to that given
func (o *EndpointInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectBadRequest ", 400)
}

func (o *EndpointInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectBadRequest ", 400)
}

func (o *EndpointInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointInspectNotFound creates a EndpointInspectNotFound with default headers values
func NewEndpointInspectNotFound() *EndpointInspectNotFound {
	return &EndpointInspectNotFound{}
}

/*
EndpointInspectNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) not found
*/
type EndpointInspectNotFound struct {
}

// IsSuccess returns true when this endpoint inspect not found response has a 2xx status code
func (o *EndpointInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint inspect not found response has a 3xx status code
func (o *EndpointInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint inspect not found response has a 4xx status code
func (o *EndpointInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint inspect not found response has a 5xx status code
func (o *EndpointInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint inspect not found response a status code equal to that given
func (o *EndpointInspectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EndpointInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectNotFound ", 404)
}

func (o *EndpointInspectNotFound) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectNotFound ", 404)
}

func (o *EndpointInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointInspectInternalServerError creates a EndpointInspectInternalServerError with default headers values
func NewEndpointInspectInternalServerError() *EndpointInspectInternalServerError {
	return &EndpointInspectInternalServerError{}
}

/*
EndpointInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointInspectInternalServerError struct {
}

// IsSuccess returns true when this endpoint inspect internal server error response has a 2xx status code
func (o *EndpointInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint inspect internal server error response has a 3xx status code
func (o *EndpointInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint inspect internal server error response has a 4xx status code
func (o *EndpointInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint inspect internal server error response has a 5xx status code
func (o *EndpointInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint inspect internal server error response a status code equal to that given
func (o *EndpointInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectInternalServerError ", 500)
}

func (o *EndpointInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}][%d] endpointInspectInternalServerError ", 500)
}

func (o *EndpointInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
