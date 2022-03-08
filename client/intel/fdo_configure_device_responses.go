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

/* FdoConfigureDeviceOK describes a response with status code 200, with default header values.

Success
*/
type FdoConfigureDeviceOK struct {
}

func (o *FdoConfigureDeviceOK) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceOK ", 200)
}

func (o *FdoConfigureDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoConfigureDeviceBadRequest creates a FdoConfigureDeviceBadRequest with default headers values
func NewFdoConfigureDeviceBadRequest() *FdoConfigureDeviceBadRequest {
	return &FdoConfigureDeviceBadRequest{}
}

/* FdoConfigureDeviceBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type FdoConfigureDeviceBadRequest struct {
}

func (o *FdoConfigureDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceBadRequest ", 400)
}

func (o *FdoConfigureDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoConfigureDeviceForbidden creates a FdoConfigureDeviceForbidden with default headers values
func NewFdoConfigureDeviceForbidden() *FdoConfigureDeviceForbidden {
	return &FdoConfigureDeviceForbidden{}
}

/* FdoConfigureDeviceForbidden describes a response with status code 403, with default header values.

Permission denied to access settings
*/
type FdoConfigureDeviceForbidden struct {
}

func (o *FdoConfigureDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceForbidden ", 403)
}

func (o *FdoConfigureDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFdoConfigureDeviceInternalServerError creates a FdoConfigureDeviceInternalServerError with default headers values
func NewFdoConfigureDeviceInternalServerError() *FdoConfigureDeviceInternalServerError {
	return &FdoConfigureDeviceInternalServerError{}
}

/* FdoConfigureDeviceInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type FdoConfigureDeviceInternalServerError struct {
}

func (o *FdoConfigureDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fdo/configure/{guid}][%d] fdoConfigureDeviceInternalServerError ", 500)
}

func (o *FdoConfigureDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
