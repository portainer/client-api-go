// Code generated by go-swagger; DO NOT EDIT.

package kaas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddNodesReader is a Reader for the AddNodes structure.
type AddNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddNodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddNodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewAddNodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddNodesOK creates a AddNodesOK with default headers values
func NewAddNodesOK() *AddNodesOK {
	return &AddNodesOK{}
}

/*
AddNodesOK describes a response with status code 200, with default header values.

Success
*/
type AddNodesOK struct {
}

// IsSuccess returns true when this add nodes o k response has a 2xx status code
func (o *AddNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add nodes o k response has a 3xx status code
func (o *AddNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add nodes o k response has a 4xx status code
func (o *AddNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add nodes o k response has a 5xx status code
func (o *AddNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add nodes o k response a status code equal to that given
func (o *AddNodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddNodesOK) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesOK ", 200)
}

func (o *AddNodesOK) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesOK ", 200)
}

func (o *AddNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodesBadRequest creates a AddNodesBadRequest with default headers values
func NewAddNodesBadRequest() *AddNodesBadRequest {
	return &AddNodesBadRequest{}
}

/*
AddNodesBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type AddNodesBadRequest struct {
}

// IsSuccess returns true when this add nodes bad request response has a 2xx status code
func (o *AddNodesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add nodes bad request response has a 3xx status code
func (o *AddNodesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add nodes bad request response has a 4xx status code
func (o *AddNodesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add nodes bad request response has a 5xx status code
func (o *AddNodesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add nodes bad request response a status code equal to that given
func (o *AddNodesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddNodesBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesBadRequest ", 400)
}

func (o *AddNodesBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesBadRequest ", 400)
}

func (o *AddNodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodesForbidden creates a AddNodesForbidden with default headers values
func NewAddNodesForbidden() *AddNodesForbidden {
	return &AddNodesForbidden{}
}

/*
AddNodesForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type AddNodesForbidden struct {
}

// IsSuccess returns true when this add nodes forbidden response has a 2xx status code
func (o *AddNodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add nodes forbidden response has a 3xx status code
func (o *AddNodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add nodes forbidden response has a 4xx status code
func (o *AddNodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add nodes forbidden response has a 5xx status code
func (o *AddNodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add nodes forbidden response a status code equal to that given
func (o *AddNodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddNodesForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesForbidden ", 403)
}

func (o *AddNodesForbidden) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesForbidden ", 403)
}

func (o *AddNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodesInternalServerError creates a AddNodesInternalServerError with default headers values
func NewAddNodesInternalServerError() *AddNodesInternalServerError {
	return &AddNodesInternalServerError{}
}

/*
AddNodesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type AddNodesInternalServerError struct {
}

// IsSuccess returns true when this add nodes internal server error response has a 2xx status code
func (o *AddNodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add nodes internal server error response has a 3xx status code
func (o *AddNodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add nodes internal server error response has a 4xx status code
func (o *AddNodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add nodes internal server error response has a 5xx status code
func (o *AddNodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add nodes internal server error response a status code equal to that given
func (o *AddNodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddNodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesInternalServerError ", 500)
}

func (o *AddNodesInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesInternalServerError ", 500)
}

func (o *AddNodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodesServiceUnavailable creates a AddNodesServiceUnavailable with default headers values
func NewAddNodesServiceUnavailable() *AddNodesServiceUnavailable {
	return &AddNodesServiceUnavailable{}
}

/*
AddNodesServiceUnavailable describes a response with status code 503, with default header values.

Missing configuration
*/
type AddNodesServiceUnavailable struct {
}

// IsSuccess returns true when this add nodes service unavailable response has a 2xx status code
func (o *AddNodesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add nodes service unavailable response has a 3xx status code
func (o *AddNodesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add nodes service unavailable response has a 4xx status code
func (o *AddNodesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this add nodes service unavailable response has a 5xx status code
func (o *AddNodesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this add nodes service unavailable response a status code equal to that given
func (o *AddNodesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *AddNodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesServiceUnavailable ", 503)
}

func (o *AddNodesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentid}/nodes/add][%d] addNodesServiceUnavailable ", 503)
}

func (o *AddNodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
