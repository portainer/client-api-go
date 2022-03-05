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

// EdgeJobListReader is a Reader for the EdgeJobList structure.
type EdgeJobListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeJobListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeJobListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeJobListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeJobListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeJobListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeJobListOK creates a EdgeJobListOK with default headers values
func NewEdgeJobListOK() *EdgeJobListOK {
	return &EdgeJobListOK{}
}

/* EdgeJobListOK describes a response with status code 200, with default header values.

OK
*/
type EdgeJobListOK struct {
	Payload []*models.PortainerEdgeJob
}

func (o *EdgeJobListOK) Error() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListOK  %+v", 200, o.Payload)
}
func (o *EdgeJobListOK) GetPayload() []*models.PortainerEdgeJob {
	return o.Payload
}

func (o *EdgeJobListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeJobListBadRequest creates a EdgeJobListBadRequest with default headers values
func NewEdgeJobListBadRequest() *EdgeJobListBadRequest {
	return &EdgeJobListBadRequest{}
}

/* EdgeJobListBadRequest describes a response with status code 400, with default header values.

EdgeJobListBadRequest edge job list bad request
*/
type EdgeJobListBadRequest struct {
}

func (o *EdgeJobListBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListBadRequest ", 400)
}

func (o *EdgeJobListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobListInternalServerError creates a EdgeJobListInternalServerError with default headers values
func NewEdgeJobListInternalServerError() *EdgeJobListInternalServerError {
	return &EdgeJobListInternalServerError{}
}

/* EdgeJobListInternalServerError describes a response with status code 500, with default header values.

EdgeJobListInternalServerError edge job list internal server error
*/
type EdgeJobListInternalServerError struct {
}

func (o *EdgeJobListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListInternalServerError ", 500)
}

func (o *EdgeJobListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobListServiceUnavailable creates a EdgeJobListServiceUnavailable with default headers values
func NewEdgeJobListServiceUnavailable() *EdgeJobListServiceUnavailable {
	return &EdgeJobListServiceUnavailable{}
}

/* EdgeJobListServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobListServiceUnavailable struct {
}

func (o *EdgeJobListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_jobs][%d] edgeJobListServiceUnavailable ", 503)
}

func (o *EdgeJobListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
