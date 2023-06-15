// Code generated by go-swagger; DO NOT EDIT.

package edge_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/ce/models"
)

// EdgeStackFileReader is a Reader for the EdgeStackFile structure.
type EdgeStackFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeStackFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeStackFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeStackFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeStackFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeStackFileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeStackFileOK creates a EdgeStackFileOK with default headers values
func NewEdgeStackFileOK() *EdgeStackFileOK {
	return &EdgeStackFileOK{}
}

/*
EdgeStackFileOK describes a response with status code 200, with default header values.

OK
*/
type EdgeStackFileOK struct {
	Payload *models.EdgestacksStackFileResponse
}

// IsSuccess returns true when this edge stack file o k response has a 2xx status code
func (o *EdgeStackFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge stack file o k response has a 3xx status code
func (o *EdgeStackFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack file o k response has a 4xx status code
func (o *EdgeStackFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack file o k response has a 5xx status code
func (o *EdgeStackFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack file o k response a status code equal to that given
func (o *EdgeStackFileOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeStackFileOK) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileOK  %+v", 200, o.Payload)
}

func (o *EdgeStackFileOK) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileOK  %+v", 200, o.Payload)
}

func (o *EdgeStackFileOK) GetPayload() *models.EdgestacksStackFileResponse {
	return o.Payload
}

func (o *EdgeStackFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgestacksStackFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeStackFileBadRequest creates a EdgeStackFileBadRequest with default headers values
func NewEdgeStackFileBadRequest() *EdgeStackFileBadRequest {
	return &EdgeStackFileBadRequest{}
}

/*
EdgeStackFileBadRequest describes a response with status code 400, with default header values.

EdgeStackFileBadRequest edge stack file bad request
*/
type EdgeStackFileBadRequest struct {
}

// IsSuccess returns true when this edge stack file bad request response has a 2xx status code
func (o *EdgeStackFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack file bad request response has a 3xx status code
func (o *EdgeStackFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack file bad request response has a 4xx status code
func (o *EdgeStackFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge stack file bad request response has a 5xx status code
func (o *EdgeStackFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge stack file bad request response a status code equal to that given
func (o *EdgeStackFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeStackFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileBadRequest ", 400)
}

func (o *EdgeStackFileBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileBadRequest ", 400)
}

func (o *EdgeStackFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackFileInternalServerError creates a EdgeStackFileInternalServerError with default headers values
func NewEdgeStackFileInternalServerError() *EdgeStackFileInternalServerError {
	return &EdgeStackFileInternalServerError{}
}

/*
EdgeStackFileInternalServerError describes a response with status code 500, with default header values.

EdgeStackFileInternalServerError edge stack file internal server error
*/
type EdgeStackFileInternalServerError struct {
}

// IsSuccess returns true when this edge stack file internal server error response has a 2xx status code
func (o *EdgeStackFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack file internal server error response has a 3xx status code
func (o *EdgeStackFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack file internal server error response has a 4xx status code
func (o *EdgeStackFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack file internal server error response has a 5xx status code
func (o *EdgeStackFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack file internal server error response a status code equal to that given
func (o *EdgeStackFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeStackFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileInternalServerError ", 500)
}

func (o *EdgeStackFileInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileInternalServerError ", 500)
}

func (o *EdgeStackFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeStackFileServiceUnavailable creates a EdgeStackFileServiceUnavailable with default headers values
func NewEdgeStackFileServiceUnavailable() *EdgeStackFileServiceUnavailable {
	return &EdgeStackFileServiceUnavailable{}
}

/*
EdgeStackFileServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeStackFileServiceUnavailable struct {
}

// IsSuccess returns true when this edge stack file service unavailable response has a 2xx status code
func (o *EdgeStackFileServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge stack file service unavailable response has a 3xx status code
func (o *EdgeStackFileServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge stack file service unavailable response has a 4xx status code
func (o *EdgeStackFileServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge stack file service unavailable response has a 5xx status code
func (o *EdgeStackFileServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge stack file service unavailable response a status code equal to that given
func (o *EdgeStackFileServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeStackFileServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileServiceUnavailable ", 503)
}

func (o *EdgeStackFileServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /edge_stacks/{id}/file][%d] edgeStackFileServiceUnavailable ", 503)
}

func (o *EdgeStackFileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
