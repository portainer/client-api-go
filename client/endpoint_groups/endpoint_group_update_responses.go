// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/models"
)

// EndpointGroupUpdateReader is a Reader for the EndpointGroupUpdate structure.
type EndpointGroupUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointGroupUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointGroupUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointGroupUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointGroupUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointGroupUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointGroupUpdateOK creates a EndpointGroupUpdateOK with default headers values
func NewEndpointGroupUpdateOK() *EndpointGroupUpdateOK {
	return &EndpointGroupUpdateOK{}
}

/*
EndpointGroupUpdateOK describes a response with status code 200, with default header values.

Success
*/
type EndpointGroupUpdateOK struct {
	Payload *models.PortainereeEndpointGroup
}

// IsSuccess returns true when this endpoint group update o k response has a 2xx status code
func (o *EndpointGroupUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint group update o k response has a 3xx status code
func (o *EndpointGroupUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group update o k response has a 4xx status code
func (o *EndpointGroupUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group update o k response has a 5xx status code
func (o *EndpointGroupUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group update o k response a status code equal to that given
func (o *EndpointGroupUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointGroupUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointGroupUpdateOK) String() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateOK  %+v", 200, o.Payload)
}

func (o *EndpointGroupUpdateOK) GetPayload() *models.PortainereeEndpointGroup {
	return o.Payload
}

func (o *EndpointGroupUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeEndpointGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointGroupUpdateBadRequest creates a EndpointGroupUpdateBadRequest with default headers values
func NewEndpointGroupUpdateBadRequest() *EndpointGroupUpdateBadRequest {
	return &EndpointGroupUpdateBadRequest{}
}

/*
EndpointGroupUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointGroupUpdateBadRequest struct {
}

// IsSuccess returns true when this endpoint group update bad request response has a 2xx status code
func (o *EndpointGroupUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group update bad request response has a 3xx status code
func (o *EndpointGroupUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group update bad request response has a 4xx status code
func (o *EndpointGroupUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint group update bad request response has a 5xx status code
func (o *EndpointGroupUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group update bad request response a status code equal to that given
func (o *EndpointGroupUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointGroupUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateBadRequest ", 400)
}

func (o *EndpointGroupUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateBadRequest ", 400)
}

func (o *EndpointGroupUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupUpdateNotFound creates a EndpointGroupUpdateNotFound with default headers values
func NewEndpointGroupUpdateNotFound() *EndpointGroupUpdateNotFound {
	return &EndpointGroupUpdateNotFound{}
}

/*
EndpointGroupUpdateNotFound describes a response with status code 404, with default header values.

EndpointGroup not found
*/
type EndpointGroupUpdateNotFound struct {
}

// IsSuccess returns true when this endpoint group update not found response has a 2xx status code
func (o *EndpointGroupUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group update not found response has a 3xx status code
func (o *EndpointGroupUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group update not found response has a 4xx status code
func (o *EndpointGroupUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint group update not found response has a 5xx status code
func (o *EndpointGroupUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group update not found response a status code equal to that given
func (o *EndpointGroupUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EndpointGroupUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateNotFound ", 404)
}

func (o *EndpointGroupUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateNotFound ", 404)
}

func (o *EndpointGroupUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupUpdateInternalServerError creates a EndpointGroupUpdateInternalServerError with default headers values
func NewEndpointGroupUpdateInternalServerError() *EndpointGroupUpdateInternalServerError {
	return &EndpointGroupUpdateInternalServerError{}
}

/*
EndpointGroupUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointGroupUpdateInternalServerError struct {
}

// IsSuccess returns true when this endpoint group update internal server error response has a 2xx status code
func (o *EndpointGroupUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group update internal server error response has a 3xx status code
func (o *EndpointGroupUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group update internal server error response has a 4xx status code
func (o *EndpointGroupUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group update internal server error response has a 5xx status code
func (o *EndpointGroupUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint group update internal server error response a status code equal to that given
func (o *EndpointGroupUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointGroupUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateInternalServerError ", 500)
}

func (o *EndpointGroupUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /endpoint_groups/{id}][%d] endpointGroupUpdateInternalServerError ", 500)
}

func (o *EndpointGroupUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
