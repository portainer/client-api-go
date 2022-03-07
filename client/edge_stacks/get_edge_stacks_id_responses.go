// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
)

// GetEdgeStacksIDReader is a Reader for the GetEdgeStacksID structure.
type GetEdgeStacksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeStacksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeStacksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEdgeStacksIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeStacksIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetEdgeStacksIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeStacksIDOK creates a GetEdgeStacksIDOK with default headers values
func NewGetEdgeStacksIDOK() *GetEdgeStacksIDOK {
	return &GetEdgeStacksIDOK{}
}

/* GetEdgeStacksIDOK describes a response with status code 200, with default header values.

OK
*/
type GetEdgeStacksIDOK struct {
	Payload *models.PortainerEdgeStack
}

func (o *GetEdgeStacksIDOK) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}][%d] getEdgeStacksIdOK  %+v", 200, o.Payload)
}
func (o *GetEdgeStacksIDOK) GetPayload() *models.PortainerEdgeStack {
	return o.Payload
}

func (o *GetEdgeStacksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEdgeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeStacksIDBadRequest creates a GetEdgeStacksIDBadRequest with default headers values
func NewGetEdgeStacksIDBadRequest() *GetEdgeStacksIDBadRequest {
	return &GetEdgeStacksIDBadRequest{}
}

/* GetEdgeStacksIDBadRequest describes a response with status code 400, with default header values.

GetEdgeStacksIDBadRequest get edge stacks Id bad request
*/
type GetEdgeStacksIDBadRequest struct {
}

func (o *GetEdgeStacksIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}][%d] getEdgeStacksIdBadRequest ", 400)
}

func (o *GetEdgeStacksIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEdgeStacksIDInternalServerError creates a GetEdgeStacksIDInternalServerError with default headers values
func NewGetEdgeStacksIDInternalServerError() *GetEdgeStacksIDInternalServerError {
	return &GetEdgeStacksIDInternalServerError{}
}

/* GetEdgeStacksIDInternalServerError describes a response with status code 500, with default header values.

GetEdgeStacksIDInternalServerError get edge stacks Id internal server error
*/
type GetEdgeStacksIDInternalServerError struct {
}

func (o *GetEdgeStacksIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}][%d] getEdgeStacksIdInternalServerError ", 500)
}

func (o *GetEdgeStacksIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEdgeStacksIDServiceUnavailable creates a GetEdgeStacksIDServiceUnavailable with default headers values
func NewGetEdgeStacksIDServiceUnavailable() *GetEdgeStacksIDServiceUnavailable {
	return &GetEdgeStacksIDServiceUnavailable{}
}

/* GetEdgeStacksIDServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type GetEdgeStacksIDServiceUnavailable struct {
}

func (o *GetEdgeStacksIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}][%d] getEdgeStacksIdServiceUnavailable ", 503)
}

func (o *GetEdgeStacksIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
