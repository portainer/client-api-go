// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ee"
)

// EndpointCreateGlobalKeyReader is a Reader for the EndpointCreateGlobalKey structure.
type EndpointCreateGlobalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointCreateGlobalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointCreateGlobalKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointCreateGlobalKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointCreateGlobalKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointCreateGlobalKeyOK creates a EndpointCreateGlobalKeyOK with default headers values
func NewEndpointCreateGlobalKeyOK() *EndpointCreateGlobalKeyOK {
	return &EndpointCreateGlobalKeyOK{}
}

/*
EndpointCreateGlobalKeyOK describes a response with status code 200, with default header values.

Success
*/
type EndpointCreateGlobalKeyOK struct {
	Payload *models_ee.EndpointsEndpointCreateGlobalKeyResponse
}

// IsSuccess returns true when this endpoint create global key o k response has a 2xx status code
func (o *EndpointCreateGlobalKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint create global key o k response has a 3xx status code
func (o *EndpointCreateGlobalKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint create global key o k response has a 4xx status code
func (o *EndpointCreateGlobalKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint create global key o k response has a 5xx status code
func (o *EndpointCreateGlobalKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint create global key o k response a status code equal to that given
func (o *EndpointCreateGlobalKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *EndpointCreateGlobalKeyOK) Error() string {
	return fmt.Sprintf("[POST /endpoints/global-key][%d] endpointCreateGlobalKeyOK  %+v", 200, o.Payload)
}

func (o *EndpointCreateGlobalKeyOK) String() string {
	return fmt.Sprintf("[POST /endpoints/global-key][%d] endpointCreateGlobalKeyOK  %+v", 200, o.Payload)
}

func (o *EndpointCreateGlobalKeyOK) GetPayload() *models_ee.EndpointsEndpointCreateGlobalKeyResponse {
	return o.Payload
}

func (o *EndpointCreateGlobalKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ee.EndpointsEndpointCreateGlobalKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointCreateGlobalKeyBadRequest creates a EndpointCreateGlobalKeyBadRequest with default headers values
func NewEndpointCreateGlobalKeyBadRequest() *EndpointCreateGlobalKeyBadRequest {
	return &EndpointCreateGlobalKeyBadRequest{}
}

/*
EndpointCreateGlobalKeyBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointCreateGlobalKeyBadRequest struct {
}

// IsSuccess returns true when this endpoint create global key bad request response has a 2xx status code
func (o *EndpointCreateGlobalKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint create global key bad request response has a 3xx status code
func (o *EndpointCreateGlobalKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint create global key bad request response has a 4xx status code
func (o *EndpointCreateGlobalKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint create global key bad request response has a 5xx status code
func (o *EndpointCreateGlobalKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint create global key bad request response a status code equal to that given
func (o *EndpointCreateGlobalKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointCreateGlobalKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /endpoints/global-key][%d] endpointCreateGlobalKeyBadRequest ", 400)
}

func (o *EndpointCreateGlobalKeyBadRequest) String() string {
	return fmt.Sprintf("[POST /endpoints/global-key][%d] endpointCreateGlobalKeyBadRequest ", 400)
}

func (o *EndpointCreateGlobalKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointCreateGlobalKeyInternalServerError creates a EndpointCreateGlobalKeyInternalServerError with default headers values
func NewEndpointCreateGlobalKeyInternalServerError() *EndpointCreateGlobalKeyInternalServerError {
	return &EndpointCreateGlobalKeyInternalServerError{}
}

/*
EndpointCreateGlobalKeyInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointCreateGlobalKeyInternalServerError struct {
}

// IsSuccess returns true when this endpoint create global key internal server error response has a 2xx status code
func (o *EndpointCreateGlobalKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint create global key internal server error response has a 3xx status code
func (o *EndpointCreateGlobalKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint create global key internal server error response has a 4xx status code
func (o *EndpointCreateGlobalKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint create global key internal server error response has a 5xx status code
func (o *EndpointCreateGlobalKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint create global key internal server error response a status code equal to that given
func (o *EndpointCreateGlobalKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointCreateGlobalKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /endpoints/global-key][%d] endpointCreateGlobalKeyInternalServerError ", 500)
}

func (o *EndpointCreateGlobalKeyInternalServerError) String() string {
	return fmt.Sprintf("[POST /endpoints/global-key][%d] endpointCreateGlobalKeyInternalServerError ", 500)
}

func (o *EndpointCreateGlobalKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
