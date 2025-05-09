// Code generated by go-swagger; DO NOT EDIT.

package endpoint_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EndpointGroupDeleteEndpointReader is a Reader for the EndpointGroupDeleteEndpoint structure.
type EndpointGroupDeleteEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointGroupDeleteEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEndpointGroupDeleteEndpointNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointGroupDeleteEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointGroupDeleteEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointGroupDeleteEndpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}] EndpointGroupDeleteEndpoint", response, response.Code())
	}
}

// NewEndpointGroupDeleteEndpointNoContent creates a EndpointGroupDeleteEndpointNoContent with default headers values
func NewEndpointGroupDeleteEndpointNoContent() *EndpointGroupDeleteEndpointNoContent {
	return &EndpointGroupDeleteEndpointNoContent{}
}

/*
EndpointGroupDeleteEndpointNoContent describes a response with status code 204, with default header values.

Success
*/
type EndpointGroupDeleteEndpointNoContent struct {
}

// IsSuccess returns true when this endpoint group delete endpoint no content response has a 2xx status code
func (o *EndpointGroupDeleteEndpointNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint group delete endpoint no content response has a 3xx status code
func (o *EndpointGroupDeleteEndpointNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete endpoint no content response has a 4xx status code
func (o *EndpointGroupDeleteEndpointNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group delete endpoint no content response has a 5xx status code
func (o *EndpointGroupDeleteEndpointNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group delete endpoint no content response a status code equal to that given
func (o *EndpointGroupDeleteEndpointNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the endpoint group delete endpoint no content response
func (o *EndpointGroupDeleteEndpointNoContent) Code() int {
	return 204
}

func (o *EndpointGroupDeleteEndpointNoContent) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointNoContent", 204)
}

func (o *EndpointGroupDeleteEndpointNoContent) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointNoContent", 204)
}

func (o *EndpointGroupDeleteEndpointNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupDeleteEndpointBadRequest creates a EndpointGroupDeleteEndpointBadRequest with default headers values
func NewEndpointGroupDeleteEndpointBadRequest() *EndpointGroupDeleteEndpointBadRequest {
	return &EndpointGroupDeleteEndpointBadRequest{}
}

/*
EndpointGroupDeleteEndpointBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointGroupDeleteEndpointBadRequest struct {
}

// IsSuccess returns true when this endpoint group delete endpoint bad request response has a 2xx status code
func (o *EndpointGroupDeleteEndpointBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group delete endpoint bad request response has a 3xx status code
func (o *EndpointGroupDeleteEndpointBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete endpoint bad request response has a 4xx status code
func (o *EndpointGroupDeleteEndpointBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint group delete endpoint bad request response has a 5xx status code
func (o *EndpointGroupDeleteEndpointBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group delete endpoint bad request response a status code equal to that given
func (o *EndpointGroupDeleteEndpointBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the endpoint group delete endpoint bad request response
func (o *EndpointGroupDeleteEndpointBadRequest) Code() int {
	return 400
}

func (o *EndpointGroupDeleteEndpointBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointBadRequest", 400)
}

func (o *EndpointGroupDeleteEndpointBadRequest) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointBadRequest", 400)
}

func (o *EndpointGroupDeleteEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupDeleteEndpointNotFound creates a EndpointGroupDeleteEndpointNotFound with default headers values
func NewEndpointGroupDeleteEndpointNotFound() *EndpointGroupDeleteEndpointNotFound {
	return &EndpointGroupDeleteEndpointNotFound{}
}

/*
EndpointGroupDeleteEndpointNotFound describes a response with status code 404, with default header values.

EndpointGroup not found
*/
type EndpointGroupDeleteEndpointNotFound struct {
}

// IsSuccess returns true when this endpoint group delete endpoint not found response has a 2xx status code
func (o *EndpointGroupDeleteEndpointNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group delete endpoint not found response has a 3xx status code
func (o *EndpointGroupDeleteEndpointNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete endpoint not found response has a 4xx status code
func (o *EndpointGroupDeleteEndpointNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint group delete endpoint not found response has a 5xx status code
func (o *EndpointGroupDeleteEndpointNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint group delete endpoint not found response a status code equal to that given
func (o *EndpointGroupDeleteEndpointNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the endpoint group delete endpoint not found response
func (o *EndpointGroupDeleteEndpointNotFound) Code() int {
	return 404
}

func (o *EndpointGroupDeleteEndpointNotFound) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointNotFound", 404)
}

func (o *EndpointGroupDeleteEndpointNotFound) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointNotFound", 404)
}

func (o *EndpointGroupDeleteEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointGroupDeleteEndpointInternalServerError creates a EndpointGroupDeleteEndpointInternalServerError with default headers values
func NewEndpointGroupDeleteEndpointInternalServerError() *EndpointGroupDeleteEndpointInternalServerError {
	return &EndpointGroupDeleteEndpointInternalServerError{}
}

/*
EndpointGroupDeleteEndpointInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointGroupDeleteEndpointInternalServerError struct {
}

// IsSuccess returns true when this endpoint group delete endpoint internal server error response has a 2xx status code
func (o *EndpointGroupDeleteEndpointInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint group delete endpoint internal server error response has a 3xx status code
func (o *EndpointGroupDeleteEndpointInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint group delete endpoint internal server error response has a 4xx status code
func (o *EndpointGroupDeleteEndpointInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint group delete endpoint internal server error response has a 5xx status code
func (o *EndpointGroupDeleteEndpointInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint group delete endpoint internal server error response a status code equal to that given
func (o *EndpointGroupDeleteEndpointInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the endpoint group delete endpoint internal server error response
func (o *EndpointGroupDeleteEndpointInternalServerError) Code() int {
	return 500
}

func (o *EndpointGroupDeleteEndpointInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointInternalServerError", 500)
}

func (o *EndpointGroupDeleteEndpointInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /endpoint_groups/{id}/endpoints/{endpointId}][%d] endpointGroupDeleteEndpointInternalServerError", 500)
}

func (o *EndpointGroupDeleteEndpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
