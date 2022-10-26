// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/models"
)

// EndpointUpdateReader is a Reader for the EndpointUpdate structure.
type EndpointUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointUpdateOK creates a EndpointUpdateOK with default headers values
func NewEndpointUpdateOK() *EndpointUpdateOK {
	return &EndpointUpdateOK{}
}

/*
EndpointUpdateOK describes a response with status code 200, with default header values.

Success
*/
type EndpointUpdateOK struct {
	Payload *models.PortainereeEndpoint
}

// IsSuccess returns true when this endpoint update o k response has a 2xx status code
func (o *EndpointUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint update o k response has a 3xx status code
func (o *EndpointUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update o k response has a 4xx status code
func (o *EndpointUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint update o k response has a 5xx status code
func (o *EndpointUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint update o k response a status code equal to that given
func (o *EndpointUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointUpdateOK) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointUpdateOK) GetPayload() *models.PortainereeEndpoint {
	return o.Payload
}

func (o *EndpointUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointUpdateBadRequest creates a EndpointUpdateBadRequest with default headers values
func NewEndpointUpdateBadRequest() *EndpointUpdateBadRequest {
	return &EndpointUpdateBadRequest{}
}

/*
EndpointUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointUpdateBadRequest struct {
}

// IsSuccess returns true when this endpoint update bad request response has a 2xx status code
func (o *EndpointUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint update bad request response has a 3xx status code
func (o *EndpointUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update bad request response has a 4xx status code
func (o *EndpointUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint update bad request response has a 5xx status code
func (o *EndpointUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint update bad request response a status code equal to that given
func (o *EndpointUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateBadRequest ", 400)
}

func (o *EndpointUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateBadRequest ", 400)
}

func (o *EndpointUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointUpdateNotFound creates a EndpointUpdateNotFound with default headers values
func NewEndpointUpdateNotFound() *EndpointUpdateNotFound {
	return &EndpointUpdateNotFound{}
}

/*
EndpointUpdateNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) not found
*/
type EndpointUpdateNotFound struct {
}

// IsSuccess returns true when this endpoint update not found response has a 2xx status code
func (o *EndpointUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint update not found response has a 3xx status code
func (o *EndpointUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update not found response has a 4xx status code
func (o *EndpointUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint update not found response has a 5xx status code
func (o *EndpointUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint update not found response a status code equal to that given
func (o *EndpointUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EndpointUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateNotFound ", 404)
}

func (o *EndpointUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateNotFound ", 404)
}

func (o *EndpointUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointUpdateInternalServerError creates a EndpointUpdateInternalServerError with default headers values
func NewEndpointUpdateInternalServerError() *EndpointUpdateInternalServerError {
	return &EndpointUpdateInternalServerError{}
}

/*
EndpointUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointUpdateInternalServerError struct {
}

// IsSuccess returns true when this endpoint update internal server error response has a 2xx status code
func (o *EndpointUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint update internal server error response has a 3xx status code
func (o *EndpointUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update internal server error response has a 4xx status code
func (o *EndpointUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint update internal server error response has a 5xx status code
func (o *EndpointUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint update internal server error response a status code equal to that given
func (o *EndpointUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateInternalServerError ", 500)
}

func (o *EndpointUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}][%d] endpointUpdateInternalServerError ", 500)
}

func (o *EndpointUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
