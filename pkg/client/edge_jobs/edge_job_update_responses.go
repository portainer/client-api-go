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

// EdgeJobUpdateReader is a Reader for the EdgeJobUpdate structure.
type EdgeJobUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeJobUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeJobUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeJobUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeJobUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeJobUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /edge_jobs/{id}] EdgeJobUpdate", response, response.Code())
	}
}

// NewEdgeJobUpdateOK creates a EdgeJobUpdateOK with default headers values
func NewEdgeJobUpdateOK() *EdgeJobUpdateOK {
	return &EdgeJobUpdateOK{}
}

/*
EdgeJobUpdateOK describes a response with status code 200, with default header values.

OK
*/
type EdgeJobUpdateOK struct {
	Payload *models.PortainerEdgeJob
}

// IsSuccess returns true when this edge job update o k response has a 2xx status code
func (o *EdgeJobUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge job update o k response has a 3xx status code
func (o *EdgeJobUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job update o k response has a 4xx status code
func (o *EdgeJobUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job update o k response has a 5xx status code
func (o *EdgeJobUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job update o k response a status code equal to that given
func (o *EdgeJobUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge job update o k response
func (o *EdgeJobUpdateOK) Code() int {
	return 200
}

func (o *EdgeJobUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateOK %s", 200, payload)
}

func (o *EdgeJobUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateOK %s", 200, payload)
}

func (o *EdgeJobUpdateOK) GetPayload() *models.PortainerEdgeJob {
	return o.Payload
}

func (o *EdgeJobUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortainerEdgeJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeJobUpdateBadRequest creates a EdgeJobUpdateBadRequest with default headers values
func NewEdgeJobUpdateBadRequest() *EdgeJobUpdateBadRequest {
	return &EdgeJobUpdateBadRequest{}
}

/*
EdgeJobUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EdgeJobUpdateBadRequest struct {
}

// IsSuccess returns true when this edge job update bad request response has a 2xx status code
func (o *EdgeJobUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job update bad request response has a 3xx status code
func (o *EdgeJobUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job update bad request response has a 4xx status code
func (o *EdgeJobUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge job update bad request response has a 5xx status code
func (o *EdgeJobUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job update bad request response a status code equal to that given
func (o *EdgeJobUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge job update bad request response
func (o *EdgeJobUpdateBadRequest) Code() int {
	return 400
}

func (o *EdgeJobUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateBadRequest", 400)
}

func (o *EdgeJobUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateBadRequest", 400)
}

func (o *EdgeJobUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobUpdateInternalServerError creates a EdgeJobUpdateInternalServerError with default headers values
func NewEdgeJobUpdateInternalServerError() *EdgeJobUpdateInternalServerError {
	return &EdgeJobUpdateInternalServerError{}
}

/*
EdgeJobUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeJobUpdateInternalServerError struct {
}

// IsSuccess returns true when this edge job update internal server error response has a 2xx status code
func (o *EdgeJobUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job update internal server error response has a 3xx status code
func (o *EdgeJobUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job update internal server error response has a 4xx status code
func (o *EdgeJobUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job update internal server error response has a 5xx status code
func (o *EdgeJobUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job update internal server error response a status code equal to that given
func (o *EdgeJobUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge job update internal server error response
func (o *EdgeJobUpdateInternalServerError) Code() int {
	return 500
}

func (o *EdgeJobUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateInternalServerError", 500)
}

func (o *EdgeJobUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateInternalServerError", 500)
}

func (o *EdgeJobUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobUpdateServiceUnavailable creates a EdgeJobUpdateServiceUnavailable with default headers values
func NewEdgeJobUpdateServiceUnavailable() *EdgeJobUpdateServiceUnavailable {
	return &EdgeJobUpdateServiceUnavailable{}
}

/*
EdgeJobUpdateServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobUpdateServiceUnavailable struct {
}

// IsSuccess returns true when this edge job update service unavailable response has a 2xx status code
func (o *EdgeJobUpdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job update service unavailable response has a 3xx status code
func (o *EdgeJobUpdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job update service unavailable response has a 4xx status code
func (o *EdgeJobUpdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job update service unavailable response has a 5xx status code
func (o *EdgeJobUpdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job update service unavailable response a status code equal to that given
func (o *EdgeJobUpdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge job update service unavailable response
func (o *EdgeJobUpdateServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeJobUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateServiceUnavailable", 503)
}

func (o *EdgeJobUpdateServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /edge_jobs/{id}][%d] edgeJobUpdateServiceUnavailable", 503)
}

func (o *EdgeJobUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
