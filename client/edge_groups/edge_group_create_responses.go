// Code generated by go-swagger; DO NOT EDIT.

package edge_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// EdgeGroupCreateReader is a Reader for the EdgeGroupCreate structure.
type EdgeGroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeGroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeGroupCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeGroupCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeGroupCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeGroupCreateOK creates a EdgeGroupCreateOK with default headers values
func NewEdgeGroupCreateOK() *EdgeGroupCreateOK {
	return &EdgeGroupCreateOK{}
}

/*
EdgeGroupCreateOK describes a response with status code 200, with default header values.

OK
*/
type EdgeGroupCreateOK struct {
	Payload *models.PortainereeEdgeGroup
}

// IsSuccess returns true when this edge group create o k response has a 2xx status code
func (o *EdgeGroupCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge group create o k response has a 3xx status code
func (o *EdgeGroupCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group create o k response has a 4xx status code
func (o *EdgeGroupCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group create o k response has a 5xx status code
func (o *EdgeGroupCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge group create o k response a status code equal to that given
func (o *EdgeGroupCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeGroupCreateOK) Error() string {
	return fmt.Sprintf("[POST /edge_groups][%d] edgeGroupCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeGroupCreateOK) String() string {
	return fmt.Sprintf("[POST /edge_groups][%d] edgeGroupCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeGroupCreateOK) GetPayload() *models.PortainereeEdgeGroup {
	return o.Payload
}

func (o *EdgeGroupCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainereeEdgeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeGroupCreateInternalServerError creates a EdgeGroupCreateInternalServerError with default headers values
func NewEdgeGroupCreateInternalServerError() *EdgeGroupCreateInternalServerError {
	return &EdgeGroupCreateInternalServerError{}
}

/*
EdgeGroupCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeGroupCreateInternalServerError struct {
}

// IsSuccess returns true when this edge group create internal server error response has a 2xx status code
func (o *EdgeGroupCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge group create internal server error response has a 3xx status code
func (o *EdgeGroupCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group create internal server error response has a 4xx status code
func (o *EdgeGroupCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group create internal server error response has a 5xx status code
func (o *EdgeGroupCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge group create internal server error response a status code equal to that given
func (o *EdgeGroupCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeGroupCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /edge_groups][%d] edgeGroupCreateInternalServerError ", 500)
}

func (o *EdgeGroupCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /edge_groups][%d] edgeGroupCreateInternalServerError ", 500)
}

func (o *EdgeGroupCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeGroupCreateServiceUnavailable creates a EdgeGroupCreateServiceUnavailable with default headers values
func NewEdgeGroupCreateServiceUnavailable() *EdgeGroupCreateServiceUnavailable {
	return &EdgeGroupCreateServiceUnavailable{}
}

/*
EdgeGroupCreateServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeGroupCreateServiceUnavailable struct {
}

// IsSuccess returns true when this edge group create service unavailable response has a 2xx status code
func (o *EdgeGroupCreateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge group create service unavailable response has a 3xx status code
func (o *EdgeGroupCreateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group create service unavailable response has a 4xx status code
func (o *EdgeGroupCreateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group create service unavailable response has a 5xx status code
func (o *EdgeGroupCreateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge group create service unavailable response a status code equal to that given
func (o *EdgeGroupCreateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeGroupCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /edge_groups][%d] edgeGroupCreateServiceUnavailable ", 503)
}

func (o *EdgeGroupCreateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /edge_groups][%d] edgeGroupCreateServiceUnavailable ", 503)
}

func (o *EdgeGroupCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
