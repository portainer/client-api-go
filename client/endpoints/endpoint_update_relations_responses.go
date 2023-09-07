// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EndpointUpdateRelationsReader is a Reader for the EndpointUpdateRelations structure.
type EndpointUpdateRelationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointUpdateRelationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEndpointUpdateRelationsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEndpointUpdateRelationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEndpointUpdateRelationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointUpdateRelationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointUpdateRelationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEndpointUpdateRelationsNoContent creates a EndpointUpdateRelationsNoContent with default headers values
func NewEndpointUpdateRelationsNoContent() *EndpointUpdateRelationsNoContent {
	return &EndpointUpdateRelationsNoContent{}
}

/*
EndpointUpdateRelationsNoContent describes a response with status code 204, with default header values.

Success
*/
type EndpointUpdateRelationsNoContent struct {
}

// IsSuccess returns true when this endpoint update relations no content response has a 2xx status code
func (o *EndpointUpdateRelationsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoint update relations no content response has a 3xx status code
func (o *EndpointUpdateRelationsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update relations no content response has a 4xx status code
func (o *EndpointUpdateRelationsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint update relations no content response has a 5xx status code
func (o *EndpointUpdateRelationsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint update relations no content response a status code equal to that given
func (o *EndpointUpdateRelationsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *EndpointUpdateRelationsNoContent) Error() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsNoContent ", 204)
}

func (o *EndpointUpdateRelationsNoContent) String() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsNoContent ", 204)
}

func (o *EndpointUpdateRelationsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointUpdateRelationsBadRequest creates a EndpointUpdateRelationsBadRequest with default headers values
func NewEndpointUpdateRelationsBadRequest() *EndpointUpdateRelationsBadRequest {
	return &EndpointUpdateRelationsBadRequest{}
}

/*
EndpointUpdateRelationsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EndpointUpdateRelationsBadRequest struct {
}

// IsSuccess returns true when this endpoint update relations bad request response has a 2xx status code
func (o *EndpointUpdateRelationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint update relations bad request response has a 3xx status code
func (o *EndpointUpdateRelationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update relations bad request response has a 4xx status code
func (o *EndpointUpdateRelationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint update relations bad request response has a 5xx status code
func (o *EndpointUpdateRelationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint update relations bad request response a status code equal to that given
func (o *EndpointUpdateRelationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EndpointUpdateRelationsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsBadRequest ", 400)
}

func (o *EndpointUpdateRelationsBadRequest) String() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsBadRequest ", 400)
}

func (o *EndpointUpdateRelationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointUpdateRelationsUnauthorized creates a EndpointUpdateRelationsUnauthorized with default headers values
func NewEndpointUpdateRelationsUnauthorized() *EndpointUpdateRelationsUnauthorized {
	return &EndpointUpdateRelationsUnauthorized{}
}

/*
EndpointUpdateRelationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type EndpointUpdateRelationsUnauthorized struct {
}

// IsSuccess returns true when this endpoint update relations unauthorized response has a 2xx status code
func (o *EndpointUpdateRelationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint update relations unauthorized response has a 3xx status code
func (o *EndpointUpdateRelationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update relations unauthorized response has a 4xx status code
func (o *EndpointUpdateRelationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint update relations unauthorized response has a 5xx status code
func (o *EndpointUpdateRelationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint update relations unauthorized response a status code equal to that given
func (o *EndpointUpdateRelationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EndpointUpdateRelationsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsUnauthorized ", 401)
}

func (o *EndpointUpdateRelationsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsUnauthorized ", 401)
}

func (o *EndpointUpdateRelationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointUpdateRelationsNotFound creates a EndpointUpdateRelationsNotFound with default headers values
func NewEndpointUpdateRelationsNotFound() *EndpointUpdateRelationsNotFound {
	return &EndpointUpdateRelationsNotFound{}
}

/*
EndpointUpdateRelationsNotFound describes a response with status code 404, with default header values.

Not found
*/
type EndpointUpdateRelationsNotFound struct {
}

// IsSuccess returns true when this endpoint update relations not found response has a 2xx status code
func (o *EndpointUpdateRelationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint update relations not found response has a 3xx status code
func (o *EndpointUpdateRelationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update relations not found response has a 4xx status code
func (o *EndpointUpdateRelationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoint update relations not found response has a 5xx status code
func (o *EndpointUpdateRelationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoint update relations not found response a status code equal to that given
func (o *EndpointUpdateRelationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EndpointUpdateRelationsNotFound) Error() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsNotFound ", 404)
}

func (o *EndpointUpdateRelationsNotFound) String() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsNotFound ", 404)
}

func (o *EndpointUpdateRelationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEndpointUpdateRelationsInternalServerError creates a EndpointUpdateRelationsInternalServerError with default headers values
func NewEndpointUpdateRelationsInternalServerError() *EndpointUpdateRelationsInternalServerError {
	return &EndpointUpdateRelationsInternalServerError{}
}

/*
EndpointUpdateRelationsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EndpointUpdateRelationsInternalServerError struct {
}

// IsSuccess returns true when this endpoint update relations internal server error response has a 2xx status code
func (o *EndpointUpdateRelationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoint update relations internal server error response has a 3xx status code
func (o *EndpointUpdateRelationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoint update relations internal server error response has a 4xx status code
func (o *EndpointUpdateRelationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoint update relations internal server error response has a 5xx status code
func (o *EndpointUpdateRelationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoint update relations internal server error response a status code equal to that given
func (o *EndpointUpdateRelationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EndpointUpdateRelationsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsInternalServerError ", 500)
}

func (o *EndpointUpdateRelationsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /endpoints/relations][%d] endpointUpdateRelationsInternalServerError ", 500)
}

func (o *EndpointUpdateRelationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
