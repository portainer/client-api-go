// Code generated by go-swagger; DO NOT EDIT.

package kaas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveNodesReader is a Reader for the RemoveNodes structure.
type RemoveNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveNodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveNodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewRemoveNodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloud/endpoints/{environmentId}/nodes/remove] removeNodes", response, response.Code())
	}
}

// NewRemoveNodesOK creates a RemoveNodesOK with default headers values
func NewRemoveNodesOK() *RemoveNodesOK {
	return &RemoveNodesOK{}
}

/*
RemoveNodesOK describes a response with status code 200, with default header values.

Success
*/
type RemoveNodesOK struct {
}

// IsSuccess returns true when this remove nodes o k response has a 2xx status code
func (o *RemoveNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove nodes o k response has a 3xx status code
func (o *RemoveNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove nodes o k response has a 4xx status code
func (o *RemoveNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove nodes o k response has a 5xx status code
func (o *RemoveNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove nodes o k response a status code equal to that given
func (o *RemoveNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove nodes o k response
func (o *RemoveNodesOK) Code() int {
	return 200
}

func (o *RemoveNodesOK) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesOK", 200)
}

func (o *RemoveNodesOK) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesOK", 200)
}

func (o *RemoveNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveNodesBadRequest creates a RemoveNodesBadRequest with default headers values
func NewRemoveNodesBadRequest() *RemoveNodesBadRequest {
	return &RemoveNodesBadRequest{}
}

/*
RemoveNodesBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type RemoveNodesBadRequest struct {
}

// IsSuccess returns true when this remove nodes bad request response has a 2xx status code
func (o *RemoveNodesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove nodes bad request response has a 3xx status code
func (o *RemoveNodesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove nodes bad request response has a 4xx status code
func (o *RemoveNodesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove nodes bad request response has a 5xx status code
func (o *RemoveNodesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove nodes bad request response a status code equal to that given
func (o *RemoveNodesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove nodes bad request response
func (o *RemoveNodesBadRequest) Code() int {
	return 400
}

func (o *RemoveNodesBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesBadRequest", 400)
}

func (o *RemoveNodesBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesBadRequest", 400)
}

func (o *RemoveNodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveNodesForbidden creates a RemoveNodesForbidden with default headers values
func NewRemoveNodesForbidden() *RemoveNodesForbidden {
	return &RemoveNodesForbidden{}
}

/*
RemoveNodesForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type RemoveNodesForbidden struct {
}

// IsSuccess returns true when this remove nodes forbidden response has a 2xx status code
func (o *RemoveNodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove nodes forbidden response has a 3xx status code
func (o *RemoveNodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove nodes forbidden response has a 4xx status code
func (o *RemoveNodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove nodes forbidden response has a 5xx status code
func (o *RemoveNodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove nodes forbidden response a status code equal to that given
func (o *RemoveNodesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove nodes forbidden response
func (o *RemoveNodesForbidden) Code() int {
	return 403
}

func (o *RemoveNodesForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesForbidden", 403)
}

func (o *RemoveNodesForbidden) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesForbidden", 403)
}

func (o *RemoveNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveNodesInternalServerError creates a RemoveNodesInternalServerError with default headers values
func NewRemoveNodesInternalServerError() *RemoveNodesInternalServerError {
	return &RemoveNodesInternalServerError{}
}

/*
RemoveNodesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RemoveNodesInternalServerError struct {
}

// IsSuccess returns true when this remove nodes internal server error response has a 2xx status code
func (o *RemoveNodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove nodes internal server error response has a 3xx status code
func (o *RemoveNodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove nodes internal server error response has a 4xx status code
func (o *RemoveNodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove nodes internal server error response has a 5xx status code
func (o *RemoveNodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove nodes internal server error response a status code equal to that given
func (o *RemoveNodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove nodes internal server error response
func (o *RemoveNodesInternalServerError) Code() int {
	return 500
}

func (o *RemoveNodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesInternalServerError", 500)
}

func (o *RemoveNodesInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesInternalServerError", 500)
}

func (o *RemoveNodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveNodesServiceUnavailable creates a RemoveNodesServiceUnavailable with default headers values
func NewRemoveNodesServiceUnavailable() *RemoveNodesServiceUnavailable {
	return &RemoveNodesServiceUnavailable{}
}

/*
RemoveNodesServiceUnavailable describes a response with status code 503, with default header values.

Missing configuration
*/
type RemoveNodesServiceUnavailable struct {
}

// IsSuccess returns true when this remove nodes service unavailable response has a 2xx status code
func (o *RemoveNodesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove nodes service unavailable response has a 3xx status code
func (o *RemoveNodesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove nodes service unavailable response has a 4xx status code
func (o *RemoveNodesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove nodes service unavailable response has a 5xx status code
func (o *RemoveNodesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this remove nodes service unavailable response a status code equal to that given
func (o *RemoveNodesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the remove nodes service unavailable response
func (o *RemoveNodesServiceUnavailable) Code() int {
	return 503
}

func (o *RemoveNodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesServiceUnavailable", 503)
}

func (o *RemoveNodesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /cloud/endpoints/{environmentId}/nodes/remove][%d] removeNodesServiceUnavailable", 503)
}

func (o *RemoveNodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
