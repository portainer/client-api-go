// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

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
		return nil, runtime.NewAPIError("[GET /edge_jobs/{id}] EdgeJobInspect", response, response.Code())
	}
}

// NewEdgeJobInspectOK creates a EdgeJobInspectOK with default headers values
func NewEdgeJobInspectOK() *EdgeJobInspectOK {
	return &EdgeJobInspectOK{}
}

/*
EdgeJobInspectOK describes a response with status code 200, with default header values.

OK
*/
type EdgeJobInspectOK struct {
	Payload *models.PortainerEdgeJob
}

// IsSuccess returns true when this edge job inspect o k response has a 2xx status code
func (o *EdgeJobInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge job inspect o k response has a 3xx status code
func (o *EdgeJobInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job inspect o k response has a 4xx status code
func (o *EdgeJobInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job inspect o k response has a 5xx status code
func (o *EdgeJobInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job inspect o k response a status code equal to that given
func (o *EdgeJobInspectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge job inspect o k response
func (o *EdgeJobInspectOK) Code() int {
	return 200
}

func (o *EdgeJobInspectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectOK %s", 200, payload)
}

func (o *EdgeJobInspectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectOK %s", 200, payload)
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

/*
EdgeJobInspectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EdgeJobInspectBadRequest struct {
}

// IsSuccess returns true when this edge job inspect bad request response has a 2xx status code
func (o *EdgeJobInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job inspect bad request response has a 3xx status code
func (o *EdgeJobInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job inspect bad request response has a 4xx status code
func (o *EdgeJobInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge job inspect bad request response has a 5xx status code
func (o *EdgeJobInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job inspect bad request response a status code equal to that given
func (o *EdgeJobInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge job inspect bad request response
func (o *EdgeJobInspectBadRequest) Code() int {
	return 400
}

func (o *EdgeJobInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectBadRequest", 400)
}

func (o *EdgeJobInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectBadRequest", 400)
}

func (o *EdgeJobInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobInspectInternalServerError creates a EdgeJobInspectInternalServerError with default headers values
func NewEdgeJobInspectInternalServerError() *EdgeJobInspectInternalServerError {
	return &EdgeJobInspectInternalServerError{}
}

/*
EdgeJobInspectInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeJobInspectInternalServerError struct {
}

// IsSuccess returns true when this edge job inspect internal server error response has a 2xx status code
func (o *EdgeJobInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job inspect internal server error response has a 3xx status code
func (o *EdgeJobInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job inspect internal server error response has a 4xx status code
func (o *EdgeJobInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job inspect internal server error response has a 5xx status code
func (o *EdgeJobInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job inspect internal server error response a status code equal to that given
func (o *EdgeJobInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge job inspect internal server error response
func (o *EdgeJobInspectInternalServerError) Code() int {
	return 500
}

func (o *EdgeJobInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectInternalServerError", 500)
}

func (o *EdgeJobInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectInternalServerError", 500)
}

func (o *EdgeJobInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobInspectServiceUnavailable creates a EdgeJobInspectServiceUnavailable with default headers values
func NewEdgeJobInspectServiceUnavailable() *EdgeJobInspectServiceUnavailable {
	return &EdgeJobInspectServiceUnavailable{}
}

/*
EdgeJobInspectServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobInspectServiceUnavailable struct {
}

// IsSuccess returns true when this edge job inspect service unavailable response has a 2xx status code
func (o *EdgeJobInspectServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job inspect service unavailable response has a 3xx status code
func (o *EdgeJobInspectServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job inspect service unavailable response has a 4xx status code
func (o *EdgeJobInspectServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job inspect service unavailable response has a 5xx status code
func (o *EdgeJobInspectServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job inspect service unavailable response a status code equal to that given
func (o *EdgeJobInspectServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge job inspect service unavailable response
func (o *EdgeJobInspectServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeJobInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectServiceUnavailable", 503)
}

func (o *EdgeJobInspectServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}][%d] edgeJobInspectServiceUnavailable", 503)
}

func (o *EdgeJobInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
