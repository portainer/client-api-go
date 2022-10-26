// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// EndpointDockerhubStatusReader is a Reader for the EndpointDockerhubStatus structure.
type EndpointDockerhubStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointDockerhubStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointDockerhubStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointDockerhubStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEndpointDockerhubStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointDockerhubStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointDockerhubStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointDockerhubStatusOK creates a EndpointDockerhubStatusOK with default headers values
func NewEndpointDockerhubStatusOK() *EndpointDockerhubStatusOK {
	return &EndpointDockerhubStatusOK{}
}

/*
EndpointDockerhubStatusOK describes a response with status code 200, with default header values.

Success
*/
type EndpointDockerhubStatusOK struct {
	Payload *models.EndpointsDockerhubStatusResponse
}

// IsSuccess returns true when this endpoint dockerhub status o k response has a 2xx status code
func (o *EndpointDockerhubStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint dockerhub status o k response has a 3xx status code
func (o *EndpointDockerhubStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint dockerhub status o k response has a 4xx status code
func (o *EndpointDockerhubStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint dockerhub status o k response has a 5xx status code
func (o *EndpointDockerhubStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint dockerhub status o k response a status code equal to that given
func (o *EndpointDockerhubStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointDockerhubStatusOK) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusOK  %+v", 200, o.Payload)
}

func (o *EndpointDockerhubStatusOK) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusOK  %+v", 200, o.Payload)
}

func (o *EndpointDockerhubStatusOK) GetPayload() *models.EndpointsDockerhubStatusResponse {
	return o.Payload
}

func (o *EndpointDockerhubStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointsDockerhubStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointDockerhubStatusBadRequest creates a EndpointDockerhubStatusBadRequest with default headers values
func NewEndpointDockerhubStatusBadRequest() *EndpointDockerhubStatusBadRequest {
	return &EndpointDockerhubStatusBadRequest{}
}

/*
EndpointDockerhubStatusBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointDockerhubStatusBadRequest struct {
}

// IsSuccess returns true when this endpoint dockerhub status bad request response has a 2xx status code
func (o *EndpointDockerhubStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint dockerhub status bad request response has a 3xx status code
func (o *EndpointDockerhubStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint dockerhub status bad request response has a 4xx status code
func (o *EndpointDockerhubStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint dockerhub status bad request response has a 5xx status code
func (o *EndpointDockerhubStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint dockerhub status bad request response a status code equal to that given
func (o *EndpointDockerhubStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointDockerhubStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusBadRequest ", 400)
}

func (o *EndpointDockerhubStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusBadRequest ", 400)
}

func (o *EndpointDockerhubStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointDockerhubStatusForbidden creates a EndpointDockerhubStatusForbidden with default headers values
func NewEndpointDockerhubStatusForbidden() *EndpointDockerhubStatusForbidden {
	return &EndpointDockerhubStatusForbidden{}
}

/*
EndpointDockerhubStatusForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type EndpointDockerhubStatusForbidden struct {
}

// IsSuccess returns true when this endpoint dockerhub status forbidden response has a 2xx status code
func (o *EndpointDockerhubStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint dockerhub status forbidden response has a 3xx status code
func (o *EndpointDockerhubStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint dockerhub status forbidden response has a 4xx status code
func (o *EndpointDockerhubStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint dockerhub status forbidden response has a 5xx status code
func (o *EndpointDockerhubStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint dockerhub status forbidden response a status code equal to that given
func (o *EndpointDockerhubStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EndpointDockerhubStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusForbidden ", 403)
}

func (o *EndpointDockerhubStatusForbidden) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusForbidden ", 403)
}

func (o *EndpointDockerhubStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointDockerhubStatusNotFound creates a EndpointDockerhubStatusNotFound with default headers values
func NewEndpointDockerhubStatusNotFound() *EndpointDockerhubStatusNotFound {
	return &EndpointDockerhubStatusNotFound{}
}

/*
EndpointDockerhubStatusNotFound describes a response with status code 404, with default header values.

registry or endpoint not found
*/
type EndpointDockerhubStatusNotFound struct {
}

// IsSuccess returns true when this endpoint dockerhub status not found response has a 2xx status code
func (o *EndpointDockerhubStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint dockerhub status not found response has a 3xx status code
func (o *EndpointDockerhubStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint dockerhub status not found response has a 4xx status code
func (o *EndpointDockerhubStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint dockerhub status not found response has a 5xx status code
func (o *EndpointDockerhubStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint dockerhub status not found response a status code equal to that given
func (o *EndpointDockerhubStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EndpointDockerhubStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusNotFound ", 404)
}

func (o *EndpointDockerhubStatusNotFound) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusNotFound ", 404)
}

func (o *EndpointDockerhubStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointDockerhubStatusInternalServerError creates a EndpointDockerhubStatusInternalServerError with default headers values
func NewEndpointDockerhubStatusInternalServerError() *EndpointDockerhubStatusInternalServerError {
	return &EndpointDockerhubStatusInternalServerError{}
}

/*
EndpointDockerhubStatusInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointDockerhubStatusInternalServerError struct {
}

// IsSuccess returns true when this endpoint dockerhub status internal server error response has a 2xx status code
func (o *EndpointDockerhubStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint dockerhub status internal server error response has a 3xx status code
func (o *EndpointDockerhubStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint dockerhub status internal server error response has a 4xx status code
func (o *EndpointDockerhubStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint dockerhub status internal server error response has a 5xx status code
func (o *EndpointDockerhubStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint dockerhub status internal server error response a status code equal to that given
func (o *EndpointDockerhubStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointDockerhubStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusInternalServerError ", 500)
}

func (o *EndpointDockerhubStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/dockerhub/{registryId}][%d] endpointDockerhubStatusInternalServerError ", 500)
}

func (o *EndpointDockerhubStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
