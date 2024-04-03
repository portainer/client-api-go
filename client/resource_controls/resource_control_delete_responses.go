// Code generated by go-swagger; DO NOT EDIT.

package resource_controls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ResourceControlDeleteReader is a Reader for the ResourceControlDelete structure.
type ResourceControlDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceControlDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewResourceControlDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResourceControlDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResourceControlDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResourceControlDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /resource_controls/{id}] ResourceControlDelete", response, response.Code())
	}
}

// NewResourceControlDeleteNoContent creates a ResourceControlDeleteNoContent with default headers values
func NewResourceControlDeleteNoContent() *ResourceControlDeleteNoContent {
	return &ResourceControlDeleteNoContent{}
}

/*
ResourceControlDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type ResourceControlDeleteNoContent struct {
}

// IsSuccess returns true when this resource control delete no content response has a 2xx status code
func (o *ResourceControlDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource control delete no content response has a 3xx status code
func (o *ResourceControlDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control delete no content response has a 4xx status code
func (o *ResourceControlDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource control delete no content response has a 5xx status code
func (o *ResourceControlDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this resource control delete no content response a status code equal to that given
func (o *ResourceControlDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the resource control delete no content response
func (o *ResourceControlDeleteNoContent) Code() int {
	return 204
}

func (o *ResourceControlDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteNoContent ", 204)
}

func (o *ResourceControlDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteNoContent ", 204)
}

func (o *ResourceControlDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourceControlDeleteBadRequest creates a ResourceControlDeleteBadRequest with default headers values
func NewResourceControlDeleteBadRequest() *ResourceControlDeleteBadRequest {
	return &ResourceControlDeleteBadRequest{}
}

/*
ResourceControlDeleteBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type ResourceControlDeleteBadRequest struct {
}

// IsSuccess returns true when this resource control delete bad request response has a 2xx status code
func (o *ResourceControlDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource control delete bad request response has a 3xx status code
func (o *ResourceControlDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control delete bad request response has a 4xx status code
func (o *ResourceControlDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource control delete bad request response has a 5xx status code
func (o *ResourceControlDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this resource control delete bad request response a status code equal to that given
func (o *ResourceControlDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the resource control delete bad request response
func (o *ResourceControlDeleteBadRequest) Code() int {
	return 400
}

func (o *ResourceControlDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteBadRequest ", 400)
}

func (o *ResourceControlDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteBadRequest ", 400)
}

func (o *ResourceControlDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourceControlDeleteNotFound creates a ResourceControlDeleteNotFound with default headers values
func NewResourceControlDeleteNotFound() *ResourceControlDeleteNotFound {
	return &ResourceControlDeleteNotFound{}
}

/*
ResourceControlDeleteNotFound describes a response with status code 404, with default header values.

Resource control not found
*/
type ResourceControlDeleteNotFound struct {
}

// IsSuccess returns true when this resource control delete not found response has a 2xx status code
func (o *ResourceControlDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource control delete not found response has a 3xx status code
func (o *ResourceControlDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control delete not found response has a 4xx status code
func (o *ResourceControlDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this resource control delete not found response has a 5xx status code
func (o *ResourceControlDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this resource control delete not found response a status code equal to that given
func (o *ResourceControlDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the resource control delete not found response
func (o *ResourceControlDeleteNotFound) Code() int {
	return 404
}

func (o *ResourceControlDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteNotFound ", 404)
}

func (o *ResourceControlDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteNotFound ", 404)
}

func (o *ResourceControlDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourceControlDeleteInternalServerError creates a ResourceControlDeleteInternalServerError with default headers values
func NewResourceControlDeleteInternalServerError() *ResourceControlDeleteInternalServerError {
	return &ResourceControlDeleteInternalServerError{}
}

/*
ResourceControlDeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ResourceControlDeleteInternalServerError struct {
}

// IsSuccess returns true when this resource control delete internal server error response has a 2xx status code
func (o *ResourceControlDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resource control delete internal server error response has a 3xx status code
func (o *ResourceControlDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource control delete internal server error response has a 4xx status code
func (o *ResourceControlDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource control delete internal server error response has a 5xx status code
func (o *ResourceControlDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resource control delete internal server error response a status code equal to that given
func (o *ResourceControlDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resource control delete internal server error response
func (o *ResourceControlDeleteInternalServerError) Code() int {
	return 500
}

func (o *ResourceControlDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteInternalServerError ", 500)
}

func (o *ResourceControlDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /resource_controls/{id}][%d] resourceControlDeleteInternalServerError ", 500)
}

func (o *ResourceControlDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
