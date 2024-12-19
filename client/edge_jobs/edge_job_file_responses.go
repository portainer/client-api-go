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

	"github.com/portainer/client-api-go/v2/models"
)

// EdgeJobFileReader is a Reader for the EdgeJobFile structure.
type EdgeJobFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeJobFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeJobFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeJobFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeJobFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeJobFileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /edge_jobs/{id}/file] EdgeJobFile", response, response.Code())
	}
}

// NewEdgeJobFileOK creates a EdgeJobFileOK with default headers values
func NewEdgeJobFileOK() *EdgeJobFileOK {
	return &EdgeJobFileOK{}
}

/*
EdgeJobFileOK describes a response with status code 200, with default header values.

OK
*/
type EdgeJobFileOK struct {
	Payload *models.EdgejobsEdgeJobFileResponse
}

// IsSuccess returns true when this edge job file o k response has a 2xx status code
func (o *EdgeJobFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge job file o k response has a 3xx status code
func (o *EdgeJobFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job file o k response has a 4xx status code
func (o *EdgeJobFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job file o k response has a 5xx status code
func (o *EdgeJobFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job file o k response a status code equal to that given
func (o *EdgeJobFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge job file o k response
func (o *EdgeJobFileOK) Code() int {
	return 200
}

func (o *EdgeJobFileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileOK %s", 200, payload)
}

func (o *EdgeJobFileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileOK %s", 200, payload)
}

func (o *EdgeJobFileOK) GetPayload() *models.EdgejobsEdgeJobFileResponse {
	return o.Payload
}

func (o *EdgeJobFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgejobsEdgeJobFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeJobFileBadRequest creates a EdgeJobFileBadRequest with default headers values
func NewEdgeJobFileBadRequest() *EdgeJobFileBadRequest {
	return &EdgeJobFileBadRequest{}
}

/*
EdgeJobFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EdgeJobFileBadRequest struct {
}

// IsSuccess returns true when this edge job file bad request response has a 2xx status code
func (o *EdgeJobFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job file bad request response has a 3xx status code
func (o *EdgeJobFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job file bad request response has a 4xx status code
func (o *EdgeJobFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge job file bad request response has a 5xx status code
func (o *EdgeJobFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job file bad request response a status code equal to that given
func (o *EdgeJobFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge job file bad request response
func (o *EdgeJobFileBadRequest) Code() int {
	return 400
}

func (o *EdgeJobFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileBadRequest", 400)
}

func (o *EdgeJobFileBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileBadRequest", 400)
}

func (o *EdgeJobFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobFileInternalServerError creates a EdgeJobFileInternalServerError with default headers values
func NewEdgeJobFileInternalServerError() *EdgeJobFileInternalServerError {
	return &EdgeJobFileInternalServerError{}
}

/*
EdgeJobFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeJobFileInternalServerError struct {
}

// IsSuccess returns true when this edge job file internal server error response has a 2xx status code
func (o *EdgeJobFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job file internal server error response has a 3xx status code
func (o *EdgeJobFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job file internal server error response has a 4xx status code
func (o *EdgeJobFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job file internal server error response has a 5xx status code
func (o *EdgeJobFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job file internal server error response a status code equal to that given
func (o *EdgeJobFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge job file internal server error response
func (o *EdgeJobFileInternalServerError) Code() int {
	return 500
}

func (o *EdgeJobFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileInternalServerError", 500)
}

func (o *EdgeJobFileInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileInternalServerError", 500)
}

func (o *EdgeJobFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobFileServiceUnavailable creates a EdgeJobFileServiceUnavailable with default headers values
func NewEdgeJobFileServiceUnavailable() *EdgeJobFileServiceUnavailable {
	return &EdgeJobFileServiceUnavailable{}
}

/*
EdgeJobFileServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobFileServiceUnavailable struct {
}

// IsSuccess returns true when this edge job file service unavailable response has a 2xx status code
func (o *EdgeJobFileServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job file service unavailable response has a 3xx status code
func (o *EdgeJobFileServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job file service unavailable response has a 4xx status code
func (o *EdgeJobFileServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job file service unavailable response has a 5xx status code
func (o *EdgeJobFileServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job file service unavailable response a status code equal to that given
func (o *EdgeJobFileServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the edge job file service unavailable response
func (o *EdgeJobFileServiceUnavailable) Code() int {
	return 503
}

func (o *EdgeJobFileServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileServiceUnavailable", 503)
}

func (o *EdgeJobFileServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /edge_jobs/{id}/file][%d] edgeJobFileServiceUnavailable", 503)
}

func (o *EdgeJobFileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
