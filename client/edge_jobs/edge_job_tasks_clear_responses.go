// Code generated by go-swagger; DO NOT EDIT.

package edge_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EdgeJobTasksClearReader is a Reader for the EdgeJobTasksClear structure.
type EdgeJobTasksClearReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeJobTasksClearReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEdgeJobTasksClearNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeJobTasksClearBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeJobTasksClearInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewEdgeJobTasksClearServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeJobTasksClearNoContent creates a EdgeJobTasksClearNoContent with default headers values
func NewEdgeJobTasksClearNoContent() *EdgeJobTasksClearNoContent {
	return &EdgeJobTasksClearNoContent{}
}

/*
EdgeJobTasksClearNoContent describes a response with status code 204, with default header values.

EdgeJobTasksClearNoContent edge job tasks clear no content
*/
type EdgeJobTasksClearNoContent struct {
}

// IsSuccess returns true when this edge job tasks clear no content response has a 2xx status code
func (o *EdgeJobTasksClearNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge job tasks clear no content response has a 3xx status code
func (o *EdgeJobTasksClearNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job tasks clear no content response has a 4xx status code
func (o *EdgeJobTasksClearNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job tasks clear no content response has a 5xx status code
func (o *EdgeJobTasksClearNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job tasks clear no content response a status code equal to that given
func (o *EdgeJobTasksClearNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *EdgeJobTasksClearNoContent) Error() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearNoContent ", 204)
}

func (o *EdgeJobTasksClearNoContent) String() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearNoContent ", 204)
}

func (o *EdgeJobTasksClearNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobTasksClearBadRequest creates a EdgeJobTasksClearBadRequest with default headers values
func NewEdgeJobTasksClearBadRequest() *EdgeJobTasksClearBadRequest {
	return &EdgeJobTasksClearBadRequest{}
}

/*
EdgeJobTasksClearBadRequest describes a response with status code 400, with default header values.

EdgeJobTasksClearBadRequest edge job tasks clear bad request
*/
type EdgeJobTasksClearBadRequest struct {
}

// IsSuccess returns true when this edge job tasks clear bad request response has a 2xx status code
func (o *EdgeJobTasksClearBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job tasks clear bad request response has a 3xx status code
func (o *EdgeJobTasksClearBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job tasks clear bad request response has a 4xx status code
func (o *EdgeJobTasksClearBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge job tasks clear bad request response has a 5xx status code
func (o *EdgeJobTasksClearBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge job tasks clear bad request response a status code equal to that given
func (o *EdgeJobTasksClearBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeJobTasksClearBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearBadRequest ", 400)
}

func (o *EdgeJobTasksClearBadRequest) String() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearBadRequest ", 400)
}

func (o *EdgeJobTasksClearBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobTasksClearInternalServerError creates a EdgeJobTasksClearInternalServerError with default headers values
func NewEdgeJobTasksClearInternalServerError() *EdgeJobTasksClearInternalServerError {
	return &EdgeJobTasksClearInternalServerError{}
}

/*
EdgeJobTasksClearInternalServerError describes a response with status code 500, with default header values.

EdgeJobTasksClearInternalServerError edge job tasks clear internal server error
*/
type EdgeJobTasksClearInternalServerError struct {
}

// IsSuccess returns true when this edge job tasks clear internal server error response has a 2xx status code
func (o *EdgeJobTasksClearInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job tasks clear internal server error response has a 3xx status code
func (o *EdgeJobTasksClearInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job tasks clear internal server error response has a 4xx status code
func (o *EdgeJobTasksClearInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job tasks clear internal server error response has a 5xx status code
func (o *EdgeJobTasksClearInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job tasks clear internal server error response a status code equal to that given
func (o *EdgeJobTasksClearInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeJobTasksClearInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearInternalServerError ", 500)
}

func (o *EdgeJobTasksClearInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearInternalServerError ", 500)
}

func (o *EdgeJobTasksClearInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeJobTasksClearServiceUnavailable creates a EdgeJobTasksClearServiceUnavailable with default headers values
func NewEdgeJobTasksClearServiceUnavailable() *EdgeJobTasksClearServiceUnavailable {
	return &EdgeJobTasksClearServiceUnavailable{}
}

/*
EdgeJobTasksClearServiceUnavailable describes a response with status code 503, with default header values.

Edge compute features are disabled
*/
type EdgeJobTasksClearServiceUnavailable struct {
}

// IsSuccess returns true when this edge job tasks clear service unavailable response has a 2xx status code
func (o *EdgeJobTasksClearServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge job tasks clear service unavailable response has a 3xx status code
func (o *EdgeJobTasksClearServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge job tasks clear service unavailable response has a 4xx status code
func (o *EdgeJobTasksClearServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge job tasks clear service unavailable response has a 5xx status code
func (o *EdgeJobTasksClearServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this edge job tasks clear service unavailable response a status code equal to that given
func (o *EdgeJobTasksClearServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *EdgeJobTasksClearServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearServiceUnavailable ", 503)
}

func (o *EdgeJobTasksClearServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /edge_jobs/{id}/tasks/{taskID}/logs][%d] edgeJobTasksClearServiceUnavailable ", 503)
}

func (o *EdgeJobTasksClearServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
