// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StackDeleteReader is a Reader for the StackDelete structure.
type StackDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StackDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStackDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStackDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStackDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStackDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStackDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStackDeleteNoContent creates a StackDeleteNoContent with default headers values
func NewStackDeleteNoContent() *StackDeleteNoContent {
	return &StackDeleteNoContent{}
}

/* StackDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type StackDeleteNoContent struct {
}

func (o *StackDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /stacks/{id}][%d] stackDeleteNoContent ", 204)
}

func (o *StackDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackDeleteBadRequest creates a StackDeleteBadRequest with default headers values
func NewStackDeleteBadRequest() *StackDeleteBadRequest {
	return &StackDeleteBadRequest{}
}

/* StackDeleteBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type StackDeleteBadRequest struct {
}

func (o *StackDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /stacks/{id}][%d] stackDeleteBadRequest ", 400)
}

func (o *StackDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackDeleteForbidden creates a StackDeleteForbidden with default headers values
func NewStackDeleteForbidden() *StackDeleteForbidden {
	return &StackDeleteForbidden{}
}

/* StackDeleteForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type StackDeleteForbidden struct {
}

func (o *StackDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /stacks/{id}][%d] stackDeleteForbidden ", 403)
}

func (o *StackDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackDeleteNotFound creates a StackDeleteNotFound with default headers values
func NewStackDeleteNotFound() *StackDeleteNotFound {
	return &StackDeleteNotFound{}
}

/* StackDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type StackDeleteNotFound struct {
}

func (o *StackDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /stacks/{id}][%d] stackDeleteNotFound ", 404)
}

func (o *StackDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStackDeleteInternalServerError creates a StackDeleteInternalServerError with default headers values
func NewStackDeleteInternalServerError() *StackDeleteInternalServerError {
	return &StackDeleteInternalServerError{}
}

/* StackDeleteInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StackDeleteInternalServerError struct {
}

func (o *StackDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /stacks/{id}][%d] stackDeleteInternalServerError ", 500)
}

func (o *StackDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
