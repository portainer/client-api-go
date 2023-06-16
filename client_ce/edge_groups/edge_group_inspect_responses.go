// Code generated by go-swagger; DO NOT EDIT.

package edge_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// EdgeGroupInspectReader is a Reader for the EdgeGroupInspect structure.
type EdgeGroupInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeGroupInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeGroupInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeGroupInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeGroupInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeGroupInspectOK creates a EdgeGroupInspectOK with default headers values
func NewEdgeGroupInspectOK() *EdgeGroupInspectOK {
	return &EdgeGroupInspectOK{}
}

/*
EdgeGroupInspectOK describes a response with status code 200, with default header values.

OK
*/
type EdgeGroupInspectOK struct {
	Payload *models_ce.PortainerEdgeGroup
}

// IsSuccess returns true when this edge group inspect o k response has a 2xx status code
func (o *EdgeGroupInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge group inspect o k response has a 3xx status code
func (o *EdgeGroupInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group inspect o k response has a 4xx status code
func (o *EdgeGroupInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group inspect o k response has a 5xx status code
func (o *EdgeGroupInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge group inspect o k response a status code equal to that given
func (o *EdgeGroupInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeGroupInspectOK) Error() string {
	return fmt.Sprintf("[GET /edge_groups/{id}][%d] edgeGroupInspectOK  %+v", 200, o.Payload)
}

func (o *EdgeGroupInspectOK) String() string {
	return fmt.Sprintf("[GET /edge_groups/{id}][%d] edgeGroupInspectOK  %+v", 200, o.Payload)
}

func (o *EdgeGroupInspectOK) GetPayload() *models_ce.PortainerEdgeGroup {
	return o.Payload
}

func (o *EdgeGroupInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.PortainerEdgeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeGroupInspectInternalServerError creates a EdgeGroupInspectInternalServerError with default headers values
func NewEdgeGroupInspectInternalServerError() *EdgeGroupInspectInternalServerError {
	return &EdgeGroupInspectInternalServerError{}
}

/*
EdgeGroupInspectInternalServerError describes a response with status code 500, with default header values.

EdgeGroupInspectInternalServerError edge group inspect internal server error
*/
type EdgeGroupInspectInternalServerError struct {
}

// IsSuccess returns true when this edge group inspect internal server error response has a 2xx status code
func (o *EdgeGroupInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge group inspect internal server error response has a 3xx status code
func (o *EdgeGroupInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group inspect internal server error response has a 4xx status code
func (o *EdgeGroupInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group inspect internal server error response has a 5xx status code
func (o *EdgeGroupInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge group inspect internal server error response a status code equal to that given
func (o *EdgeGroupInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeGroupInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_groups/{id}][%d] edgeGroupInspectInternalServerError ", 500)
}

func (o *EdgeGroupInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_groups/{id}][%d] edgeGroupInspectInternalServerError ", 500)
}

func (o *EdgeGroupInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeGroupInspectServiceUnavailable creates a EdgeGroupInspectServiceUnavailable with default headers values
func NewEdgeGroupInspectServiceUnavailable() *EdgeGroupInspectServiceUnavailable {
	return &EdgeGroupInspectServiceUnavailable{}
}

/*
EdgeGroupInspectServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeGroupInspectServiceUnavailable struct {
}

// IsSuccess returns true when this edge group inspect service unavailable response has a 2xx status code
func (o *EdgeGroupInspectServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge group inspect service unavailable response has a 3xx status code
func (o *EdgeGroupInspectServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge group inspect service unavailable response has a 4xx status code
func (o *EdgeGroupInspectServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge group inspect service unavailable response has a 5xx status code
func (o *EdgeGroupInspectServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge group inspect service unavailable response a status code equal to that given
func (o *EdgeGroupInspectServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeGroupInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_groups/{id}][%d] edgeGroupInspectServiceUnavailable ", 503)
}

func (o *EdgeGroupInspectServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /edge_groups/{id}][%d] edgeGroupInspectServiceUnavailable ", 503)
}

func (o *EdgeGroupInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
