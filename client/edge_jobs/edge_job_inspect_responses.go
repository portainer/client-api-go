// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
)

// EdgeJobInspectReader is a Reader for the EdgeJobInspect structure.
type EdgeJobInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeJobInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeJobInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeJobInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeJobInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeJobInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeJobInspectOK creates a EdgeJobInspectOK with default headers values
func NewEdgeJobInspectOK() *EdgeJobInspectOK {
	return &EdgeJobInspectOK{}
}

/* EdgeJobInspectOK describes a response with status code 200, with default header values.

OK
*/
type EdgeJobInspectOK struct {
	Payload *models.PortainerEdgeJob
}

func (o *EdgeJobInspectOK) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectOK  %+v", 200, o.Payload)
}
func (o *EdgeJobInspectOK) GetPayload() *models.PortainerEdgeJob {
	return o.Payload
}

func (o *EdgeJobInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEdgeJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeJobInspectBadRequest creates a EdgeJobInspectBadRequest with default headers values
func NewEdgeJobInspectBadRequest() *EdgeJobInspectBadRequest {
	return &EdgeJobInspectBadRequest{}
}

/* EdgeJobInspectBadRequest describes a response with status code 400, with default header values.

EdgeJobInspectBadRequest edge job inspect bad request
*/
type EdgeJobInspectBadRequest struct {
}

func (o *EdgeJobInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectBadRequest ", 400)
}

func (o *EdgeJobInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobInspectInternalServerError creates a EdgeJobInspectInternalServerError with default headers values
func NewEdgeJobInspectInternalServerError() *EdgeJobInspectInternalServerError {
	return &EdgeJobInspectInternalServerError{}
}

/* EdgeJobInspectInternalServerError describes a response with status code 500, with default header values.

EdgeJobInspectInternalServerError edge job inspect internal server error
*/
type EdgeJobInspectInternalServerError struct {
}

func (o *EdgeJobInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectInternalServerError ", 500)
}

func (o *EdgeJobInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobInspectServiceUnavailable creates a EdgeJobInspectServiceUnavailable with default headers values
func NewEdgeJobInspectServiceUnavailable() *EdgeJobInspectServiceUnavailable {
	return &EdgeJobInspectServiceUnavailable{}
}

/* EdgeJobInspectServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobInspectServiceUnavailable struct {
}

func (o *EdgeJobInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectServiceUnavailable ", 503)
}

func (o *EdgeJobInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
