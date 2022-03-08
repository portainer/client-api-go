// Code generated by go-swagger; DO NOT EDIT.

package intel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FdoRegisterDeviceReader is a Reader for the FdoRegisterDevice structure.
type FdoRegisterDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FdoRegisterDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFdoRegisterDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFdoRegisterDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFdoRegisterDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFdoRegisterDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFdoRegisterDeviceOK creates a FdoRegisterDeviceOK with default headers values
func NewFdoRegisterDeviceOK() *FdoRegisterDeviceOK {
	return &FdoRegisterDeviceOK{}
}

/* FdoRegisterDeviceOK describes a response with status code 200, with default header values.

Success
*/
type FdoRegisterDeviceOK struct {
}

func (o *FdoRegisterDeviceOK) Error() string {
	return fmt.Sprintf("[POST /fdo/register][%d] fdoRegisterDeviceOK ", 200)
}

func (o *FdoRegisterDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoRegisterDeviceBadRequest creates a FdoRegisterDeviceBadRequest with default headers values
func NewFdoRegisterDeviceBadRequest() *FdoRegisterDeviceBadRequest {
	return &FdoRegisterDeviceBadRequest{}
}

/* FdoRegisterDeviceBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type FdoRegisterDeviceBadRequest struct {
}

func (o *FdoRegisterDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /fdo/register][%d] fdoRegisterDeviceBadRequest ", 400)
}

func (o *FdoRegisterDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoRegisterDeviceForbidden creates a FdoRegisterDeviceForbidden with default headers values
func NewFdoRegisterDeviceForbidden() *FdoRegisterDeviceForbidden {
	return &FdoRegisterDeviceForbidden{}
}

/* FdoRegisterDeviceForbidden describes a response with status code 403, with default header values.

Permission denied to access settings
*/
type FdoRegisterDeviceForbidden struct {
}

func (o *FdoRegisterDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /fdo/register][%d] fdoRegisterDeviceForbidden ", 403)
}

func (o *FdoRegisterDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoRegisterDeviceInternalServerError creates a FdoRegisterDeviceInternalServerError with default headers values
func NewFdoRegisterDeviceInternalServerError() *FdoRegisterDeviceInternalServerError {
	return &FdoRegisterDeviceInternalServerError{}
}

/* FdoRegisterDeviceInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type FdoRegisterDeviceInternalServerError struct {
}

func (o *FdoRegisterDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fdo/register][%d] fdoRegisterDeviceInternalServerError ", 500)
}

func (o *FdoRegisterDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
