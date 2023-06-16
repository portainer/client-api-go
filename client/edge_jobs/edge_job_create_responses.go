// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// EdgeJobCreateReader is a Reader for the EdgeJobCreate structure.
type EdgeJobCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeJobCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeJobCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeJobCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeJobCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeJobCreateOK creates a EdgeJobCreateOK with default headers values
func NewEdgeJobCreateOK() *EdgeJobCreateOK {
	return &EdgeJobCreateOK{}
}

/*
EdgeJobCreateOK describes a response with status code 200, with default header values.

OK
*/
type EdgeJobCreateOK struct {
	Payload *models.PortainerEdgeGroup
}

// IsSuccess returns true when this edge job create o k response has a 2xx status code
func (o *EdgeJobCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge job create o k response has a 3xx status code
func (o *EdgeJobCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job create o k response has a 4xx status code
func (o *EdgeJobCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job create o k response has a 5xx status code
func (o *EdgeJobCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job create o k response a status code equal to that given
func (o *EdgeJobCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeJobCreateOK) Error() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] edgeJobCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeJobCreateOK) String() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] edgeJobCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeJobCreateOK) GetPayload() *models.PortainerEdgeGroup {
	return o.Payload
}

func (o *EdgeJobCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEdgeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeJobCreateInternalServerError creates a EdgeJobCreateInternalServerError with default headers values
func NewEdgeJobCreateInternalServerError() *EdgeJobCreateInternalServerError {
	return &EdgeJobCreateInternalServerError{}
}

/*
EdgeJobCreateInternalServerError describes a response with status code 500, with default header values.

EdgeJobCreateInternalServerError edge job create internal server error
*/
type EdgeJobCreateInternalServerError struct {
}

// IsSuccess returns true when this edge job create internal server error response has a 2xx status code
func (o *EdgeJobCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job create internal server error response has a 3xx status code
func (o *EdgeJobCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job create internal server error response has a 4xx status code
func (o *EdgeJobCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job create internal server error response has a 5xx status code
func (o *EdgeJobCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job create internal server error response a status code equal to that given
func (o *EdgeJobCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeJobCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] edgeJobCreateInternalServerError ", 500)
}

func (o *EdgeJobCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] edgeJobCreateInternalServerError ", 500)
}

func (o *EdgeJobCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobCreateServiceUnavailable creates a EdgeJobCreateServiceUnavailable with default headers values
func NewEdgeJobCreateServiceUnavailable() *EdgeJobCreateServiceUnavailable {
	return &EdgeJobCreateServiceUnavailable{}
}

/*
EdgeJobCreateServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobCreateServiceUnavailable struct {
}

// IsSuccess returns true when this edge job create service unavailable response has a 2xx status code
func (o *EdgeJobCreateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job create service unavailable response has a 3xx status code
func (o *EdgeJobCreateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job create service unavailable response has a 4xx status code
func (o *EdgeJobCreateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job create service unavailable response has a 5xx status code
func (o *EdgeJobCreateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job create service unavailable response a status code equal to that given
func (o *EdgeJobCreateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeJobCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] edgeJobCreateServiceUnavailable ", 503)
}

func (o *EdgeJobCreateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /edge_jobs][%d] edgeJobCreateServiceUnavailable ", 503)
}

func (o *EdgeJobCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
