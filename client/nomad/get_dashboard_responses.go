// Code generated by go-swagger; DO NOT EDIT.

package nomad

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDashboardReader is a Reader for the GetDashboard structure.
type GetDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDashboardOK creates a GetDashboardOK with default headers values
func NewGetDashboardOK() *GetDashboardOK {
	return &GetDashboardOK{}
}

/*
GetDashboardOK describes a response with status code 200, with default header values.

Success
*/
type GetDashboardOK struct {
}

// IsSuccess returns true when this get dashboard o k response has a 2xx status code
func (o *GetDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboard o k response has a 3xx status code
func (o *GetDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard o k response has a 4xx status code
func (o *GetDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard o k response has a 5xx status code
func (o *GetDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard o k response a status code equal to that given
func (o *GetDashboardOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDashboardOK) Error() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{endpointID}/dashboard][%d] getDashboardOK ", 200)
}

func (o *GetDashboardOK) String() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{endpointID}/dashboard][%d] getDashboardOK ", 200)
}

func (o *GetDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDashboardNotFound creates a GetDashboardNotFound with default headers values
func NewGetDashboardNotFound() *GetDashboardNotFound {
	return &GetDashboardNotFound{}
}

/*
GetDashboardNotFound describes a response with status code 404, with default header values.

Endpoint not found
*/
type GetDashboardNotFound struct {
}

// IsSuccess returns true when this get dashboard not found response has a 2xx status code
func (o *GetDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard not found response has a 3xx status code
func (o *GetDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard not found response has a 4xx status code
func (o *GetDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard not found response has a 5xx status code
func (o *GetDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard not found response a status code equal to that given
func (o *GetDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{endpointID}/dashboard][%d] getDashboardNotFound ", 404)
}

func (o *GetDashboardNotFound) String() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{endpointID}/dashboard][%d] getDashboardNotFound ", 404)
}

func (o *GetDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDashboardInternalServerError creates a GetDashboardInternalServerError with default headers values
func NewGetDashboardInternalServerError() *GetDashboardInternalServerError {
	return &GetDashboardInternalServerError{}
}

/*
GetDashboardInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetDashboardInternalServerError struct {
}

// IsSuccess returns true when this get dashboard internal server error response has a 2xx status code
func (o *GetDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard internal server error response has a 3xx status code
func (o *GetDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard internal server error response has a 4xx status code
func (o *GetDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard internal server error response has a 5xx status code
func (o *GetDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dashboard internal server error response a status code equal to that given
func (o *GetDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{endpointID}/dashboard][%d] getDashboardInternalServerError ", 500)
}

func (o *GetDashboardInternalServerError) String() string {
	return fmt.Sprintf("[GET /nomad/endpoints/{endpointID}/dashboard][%d] getDashboardInternalServerError ", 500)
}

func (o *GetDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
