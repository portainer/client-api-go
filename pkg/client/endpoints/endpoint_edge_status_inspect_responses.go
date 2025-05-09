// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// EndpointEdgeStatusInspectReader is a Reader for the EndpointEdgeStatusInspect structure.
type EndpointEdgeStatusInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointEdgeStatusInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointEdgeStatusInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointEdgeStatusInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEndpointEdgeStatusInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointEdgeStatusInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointEdgeStatusInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /endpoints/{id}/edge/status] EndpointEdgeStatusInspect", response, response.Code())
	}
}

// NewEndpointEdgeStatusInspectOK creates a EndpointEdgeStatusInspectOK with default headers values
func NewEndpointEdgeStatusInspectOK() *EndpointEdgeStatusInspectOK {
	return &EndpointEdgeStatusInspectOK{}
}

/*
EndpointEdgeStatusInspectOK describes a response with status code 200, with default header values.

Success
*/
type EndpointEdgeStatusInspectOK struct {
	Payload *models.EndpointedgeEndpointEdgeStatusInspectResponse
}

// IsSuccess returns true when this endpoint edge status inspect o k response has a 2xx status code
func (o *EndpointEdgeStatusInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint edge status inspect o k response has a 3xx status code
func (o *EndpointEdgeStatusInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint edge status inspect o k response has a 4xx status code
func (o *EndpointEdgeStatusInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint edge status inspect o k response has a 5xx status code
func (o *EndpointEdgeStatusInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint edge status inspect o k response a status code equal to that given
func (o *EndpointEdgeStatusInspectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the endpoint edge status inspect o k response
func (o *EndpointEdgeStatusInspectOK) Code() int {
	return 200
}

func (o *EndpointEdgeStatusInspectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectOK %s", 200, payload)
}

func (o *EndpointEdgeStatusInspectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectOK %s", 200, payload)
}

func (o *EndpointEdgeStatusInspectOK) GetPayload() *models.EndpointedgeEndpointEdgeStatusInspectResponse {
	return o.Payload
}

func (o *EndpointEdgeStatusInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointedgeEndpointEdgeStatusInspectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointEdgeStatusInspectBadRequest creates a EndpointEdgeStatusInspectBadRequest with default headers values
func NewEndpointEdgeStatusInspectBadRequest() *EndpointEdgeStatusInspectBadRequest {
	return &EndpointEdgeStatusInspectBadRequest{}
}

/*
EndpointEdgeStatusInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointEdgeStatusInspectBadRequest struct {
}

// IsSuccess returns true when this endpoint edge status inspect bad request response has a 2xx status code
func (o *EndpointEdgeStatusInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint edge status inspect bad request response has a 3xx status code
func (o *EndpointEdgeStatusInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint edge status inspect bad request response has a 4xx status code
func (o *EndpointEdgeStatusInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint edge status inspect bad request response has a 5xx status code
func (o *EndpointEdgeStatusInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint edge status inspect bad request response a status code equal to that given
func (o *EndpointEdgeStatusInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the endpoint edge status inspect bad request response
func (o *EndpointEdgeStatusInspectBadRequest) Code() int {
	return 400
}

func (o *EndpointEdgeStatusInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectBadRequest", 400)
}

func (o *EndpointEdgeStatusInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectBadRequest", 400)
}

func (o *EndpointEdgeStatusInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointEdgeStatusInspectForbidden creates a EndpointEdgeStatusInspectForbidden with default headers values
func NewEndpointEdgeStatusInspectForbidden() *EndpointEdgeStatusInspectForbidden {
	return &EndpointEdgeStatusInspectForbidden{}
}

/*
EndpointEdgeStatusInspectForbidden describes a response with status code 403, with default header values.

Permission denied to access environment(endpoint)
*/
type EndpointEdgeStatusInspectForbidden struct {
}

// IsSuccess returns true when this endpoint edge status inspect forbidden response has a 2xx status code
func (o *EndpointEdgeStatusInspectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint edge status inspect forbidden response has a 3xx status code
func (o *EndpointEdgeStatusInspectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint edge status inspect forbidden response has a 4xx status code
func (o *EndpointEdgeStatusInspectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint edge status inspect forbidden response has a 5xx status code
func (o *EndpointEdgeStatusInspectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint edge status inspect forbidden response a status code equal to that given
func (o *EndpointEdgeStatusInspectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the endpoint edge status inspect forbidden response
func (o *EndpointEdgeStatusInspectForbidden) Code() int {
	return 403
}

func (o *EndpointEdgeStatusInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectForbidden", 403)
}

func (o *EndpointEdgeStatusInspectForbidden) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectForbidden", 403)
}

func (o *EndpointEdgeStatusInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointEdgeStatusInspectNotFound creates a EndpointEdgeStatusInspectNotFound with default headers values
func NewEndpointEdgeStatusInspectNotFound() *EndpointEdgeStatusInspectNotFound {
	return &EndpointEdgeStatusInspectNotFound{}
}

/*
EndpointEdgeStatusInspectNotFound describes a response with status code 404, with default header values.

Environment(Endpoint) not found
*/
type EndpointEdgeStatusInspectNotFound struct {
}

// IsSuccess returns true when this endpoint edge status inspect not found response has a 2xx status code
func (o *EndpointEdgeStatusInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint edge status inspect not found response has a 3xx status code
func (o *EndpointEdgeStatusInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint edge status inspect not found response has a 4xx status code
func (o *EndpointEdgeStatusInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint edge status inspect not found response has a 5xx status code
func (o *EndpointEdgeStatusInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint edge status inspect not found response a status code equal to that given
func (o *EndpointEdgeStatusInspectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the endpoint edge status inspect not found response
func (o *EndpointEdgeStatusInspectNotFound) Code() int {
	return 404
}

func (o *EndpointEdgeStatusInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectNotFound", 404)
}

func (o *EndpointEdgeStatusInspectNotFound) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectNotFound", 404)
}

func (o *EndpointEdgeStatusInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointEdgeStatusInspectInternalServerError creates a EndpointEdgeStatusInspectInternalServerError with default headers values
func NewEndpointEdgeStatusInspectInternalServerError() *EndpointEdgeStatusInspectInternalServerError {
	return &EndpointEdgeStatusInspectInternalServerError{}
}

/*
EndpointEdgeStatusInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointEdgeStatusInspectInternalServerError struct {
}

// IsSuccess returns true when this endpoint edge status inspect internal server error response has a 2xx status code
func (o *EndpointEdgeStatusInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint edge status inspect internal server error response has a 3xx status code
func (o *EndpointEdgeStatusInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint edge status inspect internal server error response has a 4xx status code
func (o *EndpointEdgeStatusInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint edge status inspect internal server error response has a 5xx status code
func (o *EndpointEdgeStatusInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint edge status inspect internal server error response a status code equal to that given
func (o *EndpointEdgeStatusInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the endpoint edge status inspect internal server error response
func (o *EndpointEdgeStatusInspectInternalServerError) Code() int {
	return 500
}

func (o *EndpointEdgeStatusInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectInternalServerError", 500)
}

func (o *EndpointEdgeStatusInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoints/{id}/edge/status][%d] endpointEdgeStatusInspectInternalServerError", 500)
}

func (o *EndpointEdgeStatusInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
