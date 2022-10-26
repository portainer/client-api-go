// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EndpointGroupDeleteReader is a Reader for the EndpointGroupDelete structure.
type EndpointGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEndpointGroupDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointGroupDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointGroupDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointGroupDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointGroupDeleteNoContent creates a EndpointGroupDeleteNoContent with default headers values
func NewEndpointGroupDeleteNoContent() *EndpointGroupDeleteNoContent {
	return &EndpointGroupDeleteNoContent{}
}

/*
EndpointGroupDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type EndpointGroupDeleteNoContent struct {
}

// IsSuccess returns true when this endpoint group delete no content response has a 2xx status code
func (o *EndpointGroupDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint group delete no content response has a 3xx status code
func (o *EndpointGroupDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete no content response has a 4xx status code
func (o *EndpointGroupDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group delete no content response has a 5xx status code
func (o *EndpointGroupDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group delete no content response a status code equal to that given
func (o *EndpointGroupDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *EndpointGroupDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteNoContent ", 204)
}

func (o *EndpointGroupDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteNoContent ", 204)
}

func (o *EndpointGroupDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupDeleteBadRequest creates a EndpointGroupDeleteBadRequest with default headers values
func NewEndpointGroupDeleteBadRequest() *EndpointGroupDeleteBadRequest {
	return &EndpointGroupDeleteBadRequest{}
}

/*
EndpointGroupDeleteBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointGroupDeleteBadRequest struct {
}

// IsSuccess returns true when this endpoint group delete bad request response has a 2xx status code
func (o *EndpointGroupDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group delete bad request response has a 3xx status code
func (o *EndpointGroupDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete bad request response has a 4xx status code
func (o *EndpointGroupDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint group delete bad request response has a 5xx status code
func (o *EndpointGroupDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group delete bad request response a status code equal to that given
func (o *EndpointGroupDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointGroupDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteBadRequest ", 400)
}

func (o *EndpointGroupDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteBadRequest ", 400)
}

func (o *EndpointGroupDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupDeleteNotFound creates a EndpointGroupDeleteNotFound with default headers values
func NewEndpointGroupDeleteNotFound() *EndpointGroupDeleteNotFound {
	return &EndpointGroupDeleteNotFound{}
}

/*
EndpointGroupDeleteNotFound describes a response with status code 404, with default header values.

EndpointGroup not found
*/
type EndpointGroupDeleteNotFound struct {
}

// IsSuccess returns true when this endpoint group delete not found response has a 2xx status code
func (o *EndpointGroupDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group delete not found response has a 3xx status code
func (o *EndpointGroupDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete not found response has a 4xx status code
func (o *EndpointGroupDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint group delete not found response has a 5xx status code
func (o *EndpointGroupDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group delete not found response a status code equal to that given
func (o *EndpointGroupDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EndpointGroupDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteNotFound ", 404)
}

func (o *EndpointGroupDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteNotFound ", 404)
}

func (o *EndpointGroupDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupDeleteInternalServerError creates a EndpointGroupDeleteInternalServerError with default headers values
func NewEndpointGroupDeleteInternalServerError() *EndpointGroupDeleteInternalServerError {
	return &EndpointGroupDeleteInternalServerError{}
}

/*
EndpointGroupDeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointGroupDeleteInternalServerError struct {
}

// IsSuccess returns true when this endpoint group delete internal server error response has a 2xx status code
func (o *EndpointGroupDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group delete internal server error response has a 3xx status code
func (o *EndpointGroupDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete internal server error response has a 4xx status code
func (o *EndpointGroupDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group delete internal server error response has a 5xx status code
func (o *EndpointGroupDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint group delete internal server error response a status code equal to that given
func (o *EndpointGroupDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointGroupDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteInternalServerError ", 500)
}

func (o *EndpointGroupDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}][%d] endpointGroupDeleteInternalServerError ", 500)
}

func (o *EndpointGroupDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
