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

// EdgeStackListReader is a Reader for the EdgeStackList structure.
type EdgeStackListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeStackListOK creates a EdgeStackListOK with default headers values
func NewEdgeStackListOK() *EdgeStackListOK {
	return &EdgeStackListOK{}
}

/* EdgeStackListOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackListOK struct {
	Payload []*models.PortainerEdgeStack
}

func (o *EdgeStackListOK) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListOK  %+v", 200, o.Payload)
}
func (o *EdgeStackListOK) GetPayload() []*models.PortainerEdgeStack {
	return o.Payload
}

func (o *EdgeStackListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackListBadRequest creates a EdgeStackListBadRequest with default headers values
func NewEdgeStackListBadRequest() *EdgeStackListBadRequest {
	return &EdgeStackListBadRequest{}
}

/* EdgeStackListBadRequest describes a response with status code 400, with default header values.

EdgeStackListBadRequest edge stack list bad request
*/
type EdgeStackListBadRequest struct {
}

func (o *EdgeStackListBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListBadRequest ", 400)
}

func (o *EdgeStackListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackListInternalServerError creates a EdgeStackListInternalServerError with default headers values
func NewEdgeStackListInternalServerError() *EdgeStackListInternalServerError {
	return &EdgeStackListInternalServerError{}
}

/* EdgeStackListInternalServerError describes a response with status code 500, with default header values.

EdgeStackListInternalServerError edge stack list internal server error
*/
type EdgeStackListInternalServerError struct {
}

func (o *EdgeStackListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListInternalServerError ", 500)
}

func (o *EdgeStackListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackListServiceUnavailable creates a EdgeStackListServiceUnavailable with default headers values
func NewEdgeStackListServiceUnavailable() *EdgeStackListServiceUnavailable {
	return &EdgeStackListServiceUnavailable{}
}

/* EdgeStackListServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackListServiceUnavailable struct {
}

func (o *EdgeStackListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_stacks][%d] edgeStackListServiceUnavailable ", 503)
}

func (o *EdgeStackListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
