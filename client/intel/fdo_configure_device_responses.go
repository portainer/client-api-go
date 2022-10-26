// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FdoConfigureDeviceReader is a Reader for the FdoConfigureDevice structure.
type FdoConfigureDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FdoConfigureDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFdoConfigureDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFdoConfigureDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFdoConfigureDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFdoConfigureDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFdoConfigureDeviceOK creates a FdoConfigureDeviceOK with default headers values
func NewFdoConfigureDeviceOK() *FdoConfigureDeviceOK {
	return &FdoConfigureDeviceOK{}
}

/*
FdoConfigureDeviceOK describes a response with status code 200, with default header values.

Success
*/
type FdoConfigureDeviceOK struct {
}

// IsSuccess returns true when this fdo configure device o k response has a 2xx status code
func (o *FdoConfigureDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fdo configure device o k response has a 3xx status code
func (o *FdoConfigureDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo configure device o k response has a 4xx status code
func (o *FdoConfigureDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fdo configure device o k response has a 5xx status code
func (o *FdoConfigureDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fdo configure device o k response a status code equal to that given
func (o *FdoConfigureDeviceOK) IsCode(code int) bool {
	return code == 200
}

func (o *FdoConfigureDeviceOK) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceOK ", 200)
}

func (o *FdoConfigureDeviceOK) String() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceOK ", 200)
}

func (o *FdoConfigureDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoConfigureDeviceBadRequest creates a FdoConfigureDeviceBadRequest with default headers values
func NewFdoConfigureDeviceBadRequest() *FdoConfigureDeviceBadRequest {
	return &FdoConfigureDeviceBadRequest{}
}

/*
FdoConfigureDeviceBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type FdoConfigureDeviceBadRequest struct {
}

// IsSuccess returns true when this fdo configure device bad request response has a 2xx status code
func (o *FdoConfigureDeviceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fdo configure device bad request response has a 3xx status code
func (o *FdoConfigureDeviceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo configure device bad request response has a 4xx status code
func (o *FdoConfigureDeviceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this fdo configure device bad request response has a 5xx status code
func (o *FdoConfigureDeviceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this fdo configure device bad request response a status code equal to that given
func (o *FdoConfigureDeviceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *FdoConfigureDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceBadRequest ", 400)
}

func (o *FdoConfigureDeviceBadRequest) String() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceBadRequest ", 400)
}

func (o *FdoConfigureDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoConfigureDeviceForbidden creates a FdoConfigureDeviceForbidden with default headers values
func NewFdoConfigureDeviceForbidden() *FdoConfigureDeviceForbidden {
	return &FdoConfigureDeviceForbidden{}
}

/*
FdoConfigureDeviceForbidden describes a response with status code 403, with default header values.

Permission denied to access settings
*/
type FdoConfigureDeviceForbidden struct {
}

// IsSuccess returns true when this fdo configure device forbidden response has a 2xx status code
func (o *FdoConfigureDeviceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fdo configure device forbidden response has a 3xx status code
func (o *FdoConfigureDeviceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo configure device forbidden response has a 4xx status code
func (o *FdoConfigureDeviceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this fdo configure device forbidden response has a 5xx status code
func (o *FdoConfigureDeviceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this fdo configure device forbidden response a status code equal to that given
func (o *FdoConfigureDeviceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *FdoConfigureDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceForbidden ", 403)
}

func (o *FdoConfigureDeviceForbidden) String() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceForbidden ", 403)
}

func (o *FdoConfigureDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoConfigureDeviceInternalServerError creates a FdoConfigureDeviceInternalServerError with default headers values
func NewFdoConfigureDeviceInternalServerError() *FdoConfigureDeviceInternalServerError {
	return &FdoConfigureDeviceInternalServerError{}
}

/*
FdoConfigureDeviceInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type FdoConfigureDeviceInternalServerError struct {
}

// IsSuccess returns true when this fdo configure device internal server error response has a 2xx status code
func (o *FdoConfigureDeviceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fdo configure device internal server error response has a 3xx status code
func (o *FdoConfigureDeviceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fdo configure device internal server error response has a 4xx status code
func (o *FdoConfigureDeviceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this fdo configure device internal server error response has a 5xx status code
func (o *FdoConfigureDeviceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this fdo configure device internal server error response a status code equal to that given
func (o *FdoConfigureDeviceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *FdoConfigureDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceInternalServerError ", 500)
}

func (o *FdoConfigureDeviceInternalServerError) String() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceInternalServerError ", 500)
}

func (o *FdoConfigureDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
