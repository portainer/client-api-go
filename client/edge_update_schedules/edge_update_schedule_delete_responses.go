// Code generated by go-swagger; DO NOT EDIT.

package edge_update_schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EdgeUpdateScheduleDeleteReader is a Reader for the EdgeUpdateScheduleDelete structure.
type EdgeUpdateScheduleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeUpdateScheduleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEdgeUpdateScheduleDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEdgeUpdateScheduleDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /edge_update_schedules/{id}] EdgeUpdateScheduleDelete", response, response.Code())
	}
}

// NewEdgeUpdateScheduleDeleteNoContent creates a EdgeUpdateScheduleDeleteNoContent with default headers values
func NewEdgeUpdateScheduleDeleteNoContent() *EdgeUpdateScheduleDeleteNoContent {
	return &EdgeUpdateScheduleDeleteNoContent{}
}

/*
EdgeUpdateScheduleDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type EdgeUpdateScheduleDeleteNoContent struct {
}

// IsSuccess returns true when this edge update schedule delete no content response has a 2xx status code
func (o *EdgeUpdateScheduleDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge update schedule delete no content response has a 3xx status code
func (o *EdgeUpdateScheduleDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule delete no content response has a 4xx status code
func (o *EdgeUpdateScheduleDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update schedule delete no content response has a 5xx status code
func (o *EdgeUpdateScheduleDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update schedule delete no content response a status code equal to that given
func (o *EdgeUpdateScheduleDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the edge update schedule delete no content response
func (o *EdgeUpdateScheduleDeleteNoContent) Code() int {
	return 204
}

func (o *EdgeUpdateScheduleDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /edge_update_schedules/{id}][%d] edgeUpdateScheduleDeleteNoContent ", 204)
}

func (o *EdgeUpdateScheduleDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /edge_update_schedules/{id}][%d] edgeUpdateScheduleDeleteNoContent ", 204)
}

func (o *EdgeUpdateScheduleDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeUpdateScheduleDeleteInternalServerError creates a EdgeUpdateScheduleDeleteInternalServerError with default headers values
func NewEdgeUpdateScheduleDeleteInternalServerError() *EdgeUpdateScheduleDeleteInternalServerError {
	return &EdgeUpdateScheduleDeleteInternalServerError{}
}

/*
EdgeUpdateScheduleDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EdgeUpdateScheduleDeleteInternalServerError struct {
}

// IsSuccess returns true when this edge update schedule delete internal server error response has a 2xx status code
func (o *EdgeUpdateScheduleDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update schedule delete internal server error response has a 3xx status code
func (o *EdgeUpdateScheduleDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update schedule delete internal server error response has a 4xx status code
func (o *EdgeUpdateScheduleDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update schedule delete internal server error response has a 5xx status code
func (o *EdgeUpdateScheduleDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge update schedule delete internal server error response a status code equal to that given
func (o *EdgeUpdateScheduleDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge update schedule delete internal server error response
func (o *EdgeUpdateScheduleDeleteInternalServerError) Code() int {
	return 500
}

func (o *EdgeUpdateScheduleDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /edge_update_schedules/{id}][%d] edgeUpdateScheduleDeleteInternalServerError ", 500)
}

func (o *EdgeUpdateScheduleDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /edge_update_schedules/{id}][%d] edgeUpdateScheduleDeleteInternalServerError ", 500)
}

func (o *EdgeUpdateScheduleDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
