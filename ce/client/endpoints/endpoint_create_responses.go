// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/ce/v2/models"
)

// EndpointCreateReader is a Reader for the EndpointCreate structure.
type EndpointCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointCreateOK creates a EndpointCreateOK with default headers values
func NewEndpointCreateOK() *EndpointCreateOK {
	return &EndpointCreateOK{}
}

/*
EndpointCreateOK describes a response with status code 200, with default header values.

Success
*/
type EndpointCreateOK struct {
	Payload *models.PortainerEndpoint
}

// IsSuccess returns true when this endpoint create o k response has a 2xx status code
func (o *EndpointCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint create o k response has a 3xx status code
func (o *EndpointCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint create o k response has a 4xx status code
func (o *EndpointCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint create o k response has a 5xx status code
func (o *EndpointCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint create o k response a status code equal to that given
func (o *EndpointCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointCreateOK) Error() string {
	return fmt.Sprintf("[POST /endpoints][%d] endpointCreateOK  %+v", 200, o.Payload)
}

func (o *EndpointCreateOK) String() string {
	return fmt.Sprintf("[POST /endpoints][%d] endpointCreateOK  %+v", 200, o.Payload)
}

func (o *EndpointCreateOK) GetPayload() *models.PortainerEndpoint {
	return o.Payload
}

func (o *EndpointCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointCreateBadRequest creates a EndpointCreateBadRequest with default headers values
func NewEndpointCreateBadRequest() *EndpointCreateBadRequest {
	return &EndpointCreateBadRequest{}
}

/*
EndpointCreateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointCreateBadRequest struct {
}

// IsSuccess returns true when this endpoint create bad request response has a 2xx status code
func (o *EndpointCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint create bad request response has a 3xx status code
func (o *EndpointCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint create bad request response has a 4xx status code
func (o *EndpointCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint create bad request response has a 5xx status code
func (o *EndpointCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint create bad request response a status code equal to that given
func (o *EndpointCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /endpoints][%d] endpointCreateBadRequest ", 400)
}

func (o *EndpointCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /endpoints][%d] endpointCreateBadRequest ", 400)
}

func (o *EndpointCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointCreateInternalServerError creates a EndpointCreateInternalServerError with default headers values
func NewEndpointCreateInternalServerError() *EndpointCreateInternalServerError {
	return &EndpointCreateInternalServerError{}
}

/*
EndpointCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointCreateInternalServerError struct {
}

// IsSuccess returns true when this endpoint create internal server error response has a 2xx status code
func (o *EndpointCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint create internal server error response has a 3xx status code
func (o *EndpointCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint create internal server error response has a 4xx status code
func (o *EndpointCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint create internal server error response has a 5xx status code
func (o *EndpointCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint create internal server error response a status code equal to that given
func (o *EndpointCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /endpoints][%d] endpointCreateInternalServerError ", 500)
}

func (o *EndpointCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /endpoints][%d] endpointCreateInternalServerError ", 500)
}

func (o *EndpointCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
