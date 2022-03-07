// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointGroupListOK creates a EndpointGroupListOK with default headers values
func NewEndpointGroupListOK() *EndpointGroupListOK {
	return &EndpointGroupListOK{}
}

/* EndpointGroupListOK describes a response with status code 200, with default header values.

Endpoint group
*/
type EndpointGroupListOK struct {
	Payload []*models.PortainerEndpointGroup
}

func (o *EndpointGroupListOK) Error() string {
	return fmt.Sprintf("[GET /endpoint_groups][%d] endpointGroupListOK  %+v", 200, o.Payload)
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

/* EndpointGroupListInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointGroupListInternalServerError struct {
}

func (o *EndpointGroupListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /endpoint_groups][%d] endpointGroupListInternalServerError ", 500)
}

func (o *EndpointGroupListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
