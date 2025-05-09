// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

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

// EndpointGroupListReader is a Reader for the EndpointGroupList structure.
type EndpointGroupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointGroupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointGroupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEndpointGroupListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /endpoint_groups] EndpointGroupList", response, response.Code())
	}
}

// NewEndpointGroupListOK creates a EndpointGroupListOK with default headers values
func NewEndpointGroupListOK() *EndpointGroupListOK {
	return &EndpointGroupListOK{}
}

/*
EndpointGroupListOK describes a response with status code 200, with default header values.

Environment(Endpoint) group
*/
type EndpointGroupListOK struct {
	Payload []*models.PortainerEndpointGroup
}

// IsSuccess returns true when this endpoint group list o k response has a 2xx status code
func (o *EndpointGroupListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint group list o k response has a 3xx status code
func (o *EndpointGroupListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group list o k response has a 4xx status code
func (o *EndpointGroupListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group list o k response has a 5xx status code
func (o *EndpointGroupListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group list o k response a status code equal to that given
func (o *EndpointGroupListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the endpoint group list o k response
func (o *EndpointGroupListOK) Code() int {
	return 200
}

func (o *EndpointGroupListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoint_groups][%d] endpointGroupListOK %s", 200, payload)
}

func (o *EndpointGroupListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoint_groups][%d] endpointGroupListOK %s", 200, payload)
}

func (o *EndpointGroupListOK) GetPayload() []*models.PortainerEndpointGroup {
	return o.Payload
}

func (o *EndpointGroupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointGroupListInternalServerError creates a EndpointGroupListInternalServerError with default headers values
func NewEndpointGroupListInternalServerError() *EndpointGroupListInternalServerError {
	return &EndpointGroupListInternalServerError{}
}

/*
EndpointGroupListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointGroupListInternalServerError struct {
}

// IsSuccess returns true when this endpoint group list internal server error response has a 2xx status code
func (o *EndpointGroupListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group list internal server error response has a 3xx status code
func (o *EndpointGroupListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group list internal server error response has a 4xx status code
func (o *EndpointGroupListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group list internal server error response has a 5xx status code
func (o *EndpointGroupListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint group list internal server error response a status code equal to that given
func (o *EndpointGroupListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the endpoint group list internal server error response
func (o *EndpointGroupListInternalServerError) Code() int {
	return 500
}

func (o *EndpointGroupListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoint_groups][%d] endpointGroupListInternalServerError", 500)
}

func (o *EndpointGroupListInternalServerError) String() string {
	return fmt.Sprintf("[GET /endpoint_groups][%d] endpointGroupListInternalServerError", 500)
}

func (o *EndpointGroupListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
