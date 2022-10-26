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

// EndpointSettingsUpdateReader is a Reader for the EndpointSettingsUpdate structure.
type EndpointSettingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointSettingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointSettingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointSettingsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointSettingsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointSettingsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointSettingsUpdateOK creates a EndpointSettingsUpdateOK with default headers values
func NewEndpointSettingsUpdateOK() *EndpointSettingsUpdateOK {
	return &EndpointSettingsUpdateOK{}
}

/*
EndpointSettingsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type EndpointSettingsUpdateOK struct {
	Payload *models.PortainereeEndpoint
}

// IsSuccess returns true when this endpoint settings update o k response has a 2xx status code
func (o *EndpointSettingsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint settings update o k response has a 3xx status code
func (o *EndpointSettingsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint settings update o k response has a 4xx status code
func (o *EndpointSettingsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint settings update o k response has a 5xx status code
func (o *EndpointSettingsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint settings update o k response a status code equal to that given
func (o *EndpointSettingsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointSettingsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointSettingsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointSettingsUpdateOK) GetPayload() *models.PortainereeEndpoint {
	return o.Payload
}

func (o *EndpointSettingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointSettingsUpdateBadRequest creates a EndpointSettingsUpdateBadRequest with default headers values
func NewEndpointSettingsUpdateBadRequest() *EndpointSettingsUpdateBadRequest {
	return &EndpointSettingsUpdateBadRequest{}
}

/*
EndpointSettingsUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointSettingsUpdateBadRequest struct {
}

// IsSuccess returns true when this endpoint settings update bad request response has a 2xx status code
func (o *EndpointSettingsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint settings update bad request response has a 3xx status code
func (o *EndpointSettingsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint settings update bad request response has a 4xx status code
func (o *EndpointSettingsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint settings update bad request response has a 5xx status code
func (o *EndpointSettingsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint settings update bad request response a status code equal to that given
func (o *EndpointSettingsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointSettingsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateBadRequest ", 400)
}

func (o *EndpointSettingsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateBadRequest ", 400)
}

func (o *EndpointSettingsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointSettingsUpdateNotFound creates a EndpointSettingsUpdateNotFound with default headers values
func NewEndpointSettingsUpdateNotFound() *EndpointSettingsUpdateNotFound {
	return &EndpointSettingsUpdateNotFound{}
}

/*
EndpointSettingsUpdateNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) not found
*/
type EndpointSettingsUpdateNotFound struct {
}

// IsSuccess returns true when this endpoint settings update not found response has a 2xx status code
func (o *EndpointSettingsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint settings update not found response has a 3xx status code
func (o *EndpointSettingsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint settings update not found response has a 4xx status code
func (o *EndpointSettingsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint settings update not found response has a 5xx status code
func (o *EndpointSettingsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint settings update not found response a status code equal to that given
func (o *EndpointSettingsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EndpointSettingsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateNotFound ", 404)
}

func (o *EndpointSettingsUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateNotFound ", 404)
}

func (o *EndpointSettingsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointSettingsUpdateInternalServerError creates a EndpointSettingsUpdateInternalServerError with default headers values
func NewEndpointSettingsUpdateInternalServerError() *EndpointSettingsUpdateInternalServerError {
	return &EndpointSettingsUpdateInternalServerError{}
}

/*
EndpointSettingsUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointSettingsUpdateInternalServerError struct {
}

// IsSuccess returns true when this endpoint settings update internal server error response has a 2xx status code
func (o *EndpointSettingsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint settings update internal server error response has a 3xx status code
func (o *EndpointSettingsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint settings update internal server error response has a 4xx status code
func (o *EndpointSettingsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint settings update internal server error response has a 5xx status code
func (o *EndpointSettingsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint settings update internal server error response a status code equal to that given
func (o *EndpointSettingsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointSettingsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateInternalServerError ", 500)
}

func (o *EndpointSettingsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/settings][%d] endpointSettingsUpdateInternalServerError ", 500)
}

func (o *EndpointSettingsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
