// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NamespacesAccessUpdateReader is a Reader for the NamespacesAccessUpdate structure.
type NamespacesAccessUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespacesAccessUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNamespacesAccessUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNamespacesAccessUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNamespacesAccessUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNamespacesAccessUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNamespacesAccessUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /endpoints/{id}/pools/{rpn}/access] namespacesAccessUpdate", response, response.Code())
	}
}

// NewNamespacesAccessUpdateNoContent creates a NamespacesAccessUpdateNoContent with default headers values
func NewNamespacesAccessUpdateNoContent() *NamespacesAccessUpdateNoContent {
	return &NamespacesAccessUpdateNoContent{}
}

/*
NamespacesAccessUpdateNoContent describes a response with status code 204, with default header values.

Success
*/
type NamespacesAccessUpdateNoContent struct {
}

// IsSuccess returns true when this namespaces access update no content response has a 2xx status code
func (o *NamespacesAccessUpdateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this namespaces access update no content response has a 3xx status code
func (o *NamespacesAccessUpdateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this namespaces access update no content response has a 4xx status code
func (o *NamespacesAccessUpdateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this namespaces access update no content response has a 5xx status code
func (o *NamespacesAccessUpdateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this namespaces access update no content response a status code equal to that given
func (o *NamespacesAccessUpdateNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the namespaces access update no content response
func (o *NamespacesAccessUpdateNoContent) Code() int {
	return 204
}

func (o *NamespacesAccessUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateNoContent", 204)
}

func (o *NamespacesAccessUpdateNoContent) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateNoContent", 204)
}

func (o *NamespacesAccessUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNamespacesAccessUpdateBadRequest creates a NamespacesAccessUpdateBadRequest with default headers values
func NewNamespacesAccessUpdateBadRequest() *NamespacesAccessUpdateBadRequest {
	return &NamespacesAccessUpdateBadRequest{}
}

/*
NamespacesAccessUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type NamespacesAccessUpdateBadRequest struct {
}

// IsSuccess returns true when this namespaces access update bad request response has a 2xx status code
func (o *NamespacesAccessUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this namespaces access update bad request response has a 3xx status code
func (o *NamespacesAccessUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this namespaces access update bad request response has a 4xx status code
func (o *NamespacesAccessUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this namespaces access update bad request response has a 5xx status code
func (o *NamespacesAccessUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this namespaces access update bad request response a status code equal to that given
func (o *NamespacesAccessUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the namespaces access update bad request response
func (o *NamespacesAccessUpdateBadRequest) Code() int {
	return 400
}

func (o *NamespacesAccessUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateBadRequest", 400)
}

func (o *NamespacesAccessUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateBadRequest", 400)
}

func (o *NamespacesAccessUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNamespacesAccessUpdateForbidden creates a NamespacesAccessUpdateForbidden with default headers values
func NewNamespacesAccessUpdateForbidden() *NamespacesAccessUpdateForbidden {
	return &NamespacesAccessUpdateForbidden{}
}

/*
NamespacesAccessUpdateForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type NamespacesAccessUpdateForbidden struct {
}

// IsSuccess returns true when this namespaces access update forbidden response has a 2xx status code
func (o *NamespacesAccessUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this namespaces access update forbidden response has a 3xx status code
func (o *NamespacesAccessUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this namespaces access update forbidden response has a 4xx status code
func (o *NamespacesAccessUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this namespaces access update forbidden response has a 5xx status code
func (o *NamespacesAccessUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this namespaces access update forbidden response a status code equal to that given
func (o *NamespacesAccessUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the namespaces access update forbidden response
func (o *NamespacesAccessUpdateForbidden) Code() int {
	return 403
}

func (o *NamespacesAccessUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateForbidden", 403)
}

func (o *NamespacesAccessUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateForbidden", 403)
}

func (o *NamespacesAccessUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNamespacesAccessUpdateNotFound creates a NamespacesAccessUpdateNotFound with default headers values
func NewNamespacesAccessUpdateNotFound() *NamespacesAccessUpdateNotFound {
	return &NamespacesAccessUpdateNotFound{}
}

/*
NamespacesAccessUpdateNotFound describes a response with status code 404, with default header values.

Endpoint not found
*/
type NamespacesAccessUpdateNotFound struct {
}

// IsSuccess returns true when this namespaces access update not found response has a 2xx status code
func (o *NamespacesAccessUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this namespaces access update not found response has a 3xx status code
func (o *NamespacesAccessUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this namespaces access update not found response has a 4xx status code
func (o *NamespacesAccessUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this namespaces access update not found response has a 5xx status code
func (o *NamespacesAccessUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this namespaces access update not found response a status code equal to that given
func (o *NamespacesAccessUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the namespaces access update not found response
func (o *NamespacesAccessUpdateNotFound) Code() int {
	return 404
}

func (o *NamespacesAccessUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateNotFound", 404)
}

func (o *NamespacesAccessUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateNotFound", 404)
}

func (o *NamespacesAccessUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNamespacesAccessUpdateInternalServerError creates a NamespacesAccessUpdateInternalServerError with default headers values
func NewNamespacesAccessUpdateInternalServerError() *NamespacesAccessUpdateInternalServerError {
	return &NamespacesAccessUpdateInternalServerError{}
}

/*
NamespacesAccessUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type NamespacesAccessUpdateInternalServerError struct {
}

// IsSuccess returns true when this namespaces access update internal server error response has a 2xx status code
func (o *NamespacesAccessUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this namespaces access update internal server error response has a 3xx status code
func (o *NamespacesAccessUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this namespaces access update internal server error response has a 4xx status code
func (o *NamespacesAccessUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this namespaces access update internal server error response has a 5xx status code
func (o *NamespacesAccessUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this namespaces access update internal server error response a status code equal to that given
func (o *NamespacesAccessUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the namespaces access update internal server error response
func (o *NamespacesAccessUpdateInternalServerError) Code() int {
	return 500
}

func (o *NamespacesAccessUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateInternalServerError", 500)
}

func (o *NamespacesAccessUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /endpoints/{id}/pools/{rpn}/access][%d] namespacesAccessUpdateInternalServerError", 500)
}

func (o *NamespacesAccessUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
