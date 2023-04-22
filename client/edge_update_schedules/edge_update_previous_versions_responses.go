// Code generated by go-swagger; DO NOT EDIT.

package edge_update_schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EdgeUpdatePreviousVersionsReader is a Reader for the EdgeUpdatePreviousVersions structure.
type EdgeUpdatePreviousVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeUpdatePreviousVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeUpdatePreviousVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeUpdatePreviousVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeUpdatePreviousVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEdgeUpdatePreviousVersionsOK creates a EdgeUpdatePreviousVersionsOK with default headers values
func NewEdgeUpdatePreviousVersionsOK() *EdgeUpdatePreviousVersionsOK {
	return &EdgeUpdatePreviousVersionsOK{}
}

/*
EdgeUpdatePreviousVersionsOK describes a response with status code 200, with default header values.

OK
*/
type EdgeUpdatePreviousVersionsOK struct {
	Payload []string
}

// IsSuccess returns true when this edge update previous versions o k response has a 2xx status code
func (o *EdgeUpdatePreviousVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge update previous versions o k response has a 3xx status code
func (o *EdgeUpdatePreviousVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update previous versions o k response has a 4xx status code
func (o *EdgeUpdatePreviousVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update previous versions o k response has a 5xx status code
func (o *EdgeUpdatePreviousVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update previous versions o k response a status code equal to that given
func (o *EdgeUpdatePreviousVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeUpdatePreviousVersionsOK) Error() string {
	return fmt.Sprintf("[GET /edge_update_schedules/previous_versions][%d] edgeUpdatePreviousVersionsOK  %+v", 200, o.Payload)
}

func (o *EdgeUpdatePreviousVersionsOK) String() string {
	return fmt.Sprintf("[GET /edge_update_schedules/previous_versions][%d] edgeUpdatePreviousVersionsOK  %+v", 200, o.Payload)
}

func (o *EdgeUpdatePreviousVersionsOK) GetPayload() []string {
	return o.Payload
}

func (o *EdgeUpdatePreviousVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeUpdatePreviousVersionsBadRequest creates a EdgeUpdatePreviousVersionsBadRequest with default headers values
func NewEdgeUpdatePreviousVersionsBadRequest() *EdgeUpdatePreviousVersionsBadRequest {
	return &EdgeUpdatePreviousVersionsBadRequest{}
}

/*
EdgeUpdatePreviousVersionsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type EdgeUpdatePreviousVersionsBadRequest struct {
}

// IsSuccess returns true when this edge update previous versions bad request response has a 2xx status code
func (o *EdgeUpdatePreviousVersionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update previous versions bad request response has a 3xx status code
func (o *EdgeUpdatePreviousVersionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update previous versions bad request response has a 4xx status code
func (o *EdgeUpdatePreviousVersionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge update previous versions bad request response has a 5xx status code
func (o *EdgeUpdatePreviousVersionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge update previous versions bad request response a status code equal to that given
func (o *EdgeUpdatePreviousVersionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeUpdatePreviousVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge_update_schedules/previous_versions][%d] edgeUpdatePreviousVersionsBadRequest ", 400)
}

func (o *EdgeUpdatePreviousVersionsBadRequest) String() string {
	return fmt.Sprintf("[GET /edge_update_schedules/previous_versions][%d] edgeUpdatePreviousVersionsBadRequest ", 400)
}

func (o *EdgeUpdatePreviousVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEdgeUpdatePreviousVersionsInternalServerError creates a EdgeUpdatePreviousVersionsInternalServerError with default headers values
func NewEdgeUpdatePreviousVersionsInternalServerError() *EdgeUpdatePreviousVersionsInternalServerError {
	return &EdgeUpdatePreviousVersionsInternalServerError{}
}

/*
EdgeUpdatePreviousVersionsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type EdgeUpdatePreviousVersionsInternalServerError struct {
}

// IsSuccess returns true when this edge update previous versions internal server error response has a 2xx status code
func (o *EdgeUpdatePreviousVersionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge update previous versions internal server error response has a 3xx status code
func (o *EdgeUpdatePreviousVersionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge update previous versions internal server error response has a 4xx status code
func (o *EdgeUpdatePreviousVersionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge update previous versions internal server error response has a 5xx status code
func (o *EdgeUpdatePreviousVersionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge update previous versions internal server error response a status code equal to that given
func (o *EdgeUpdatePreviousVersionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeUpdatePreviousVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /edge_update_schedules/previous_versions][%d] edgeUpdatePreviousVersionsInternalServerError ", 500)
}

func (o *EdgeUpdatePreviousVersionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /edge_update_schedules/previous_versions][%d] edgeUpdatePreviousVersionsInternalServerError ", 500)
}

func (o *EdgeUpdatePreviousVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
