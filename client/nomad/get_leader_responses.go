// Code generated by go-swagger; DO NOT EDIT.

package nomad

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetLeaderReader is a Reader for the GetLeader structure.
type GetLeaderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLeaderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLeaderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLeaderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLeaderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /nomad/endpoints/{environmentId}/leader] getLeader", response, response.Code())
	}
}

// NewGetLeaderOK creates a GetLeaderOK with default headers values
func NewGetLeaderOK() *GetLeaderOK {
	return &GetLeaderOK{}
}

/*
GetLeaderOK describes a response with status code 200, with default header values.

Success
*/
type GetLeaderOK struct {
}

// IsSuccess returns true when this get leader o k response has a 2xx status code
func (o *GetLeaderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get leader o k response has a 3xx status code
func (o *GetLeaderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get leader o k response has a 4xx status code
func (o *GetLeaderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get leader o k response has a 5xx status code
func (o *GetLeaderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get leader o k response a status code equal to that given
func (o *GetLeaderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get leader o k response
func (o *GetLeaderOK) Code() int {
	return 200
}

func (o *GetLeaderOK) Error() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{environmentId}/leader][%d] getLeaderOK ", 200)
}

func (o *GetLeaderOK) String() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{environmentId}/leader][%d] getLeaderOK ", 200)
}

func (o *GetLeaderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLeaderNotFound creates a GetLeaderNotFound with default headers values
func NewGetLeaderNotFound() *GetLeaderNotFound {
	return &GetLeaderNotFound{}
}

/*
GetLeaderNotFound describes a response with status code 404, with default header values.

Endpoint not found
*/
type GetLeaderNotFound struct {
}

// IsSuccess returns true when this get leader not found response has a 2xx status code
func (o *GetLeaderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get leader not found response has a 3xx status code
func (o *GetLeaderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get leader not found response has a 4xx status code
func (o *GetLeaderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get leader not found response has a 5xx status code
func (o *GetLeaderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get leader not found response a status code equal to that given
func (o *GetLeaderNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get leader not found response
func (o *GetLeaderNotFound) Code() int {
	return 404
}

func (o *GetLeaderNotFound) Error() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{environmentId}/leader][%d] getLeaderNotFound ", 404)
}

func (o *GetLeaderNotFound) String() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{environmentId}/leader][%d] getLeaderNotFound ", 404)
}

func (o *GetLeaderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLeaderInternalServerError creates a GetLeaderInternalServerError with default headers values
func NewGetLeaderInternalServerError() *GetLeaderInternalServerError {
	return &GetLeaderInternalServerError{}
}

/*
GetLeaderInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetLeaderInternalServerError struct {
}

// IsSuccess returns true when this get leader internal server error response has a 2xx status code
func (o *GetLeaderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get leader internal server error response has a 3xx status code
func (o *GetLeaderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get leader internal server error response has a 4xx status code
func (o *GetLeaderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get leader internal server error response has a 5xx status code
func (o *GetLeaderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get leader internal server error response a status code equal to that given
func (o *GetLeaderInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get leader internal server error response
func (o *GetLeaderInternalServerError) Code() int {
	return 500
}

func (o *GetLeaderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{environmentId}/leader][%d] getLeaderInternalServerError ", 500)
}

func (o *GetLeaderInternalServerError) String() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{environmentId}/leader][%d] getLeaderInternalServerError ", 500)
}

func (o *GetLeaderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
