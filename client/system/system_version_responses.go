// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// SystemVersionReader is a Reader for the SystemVersion structure.
type SystemVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemVersionOK creates a SystemVersionOK with default headers values
func NewSystemVersionOK() *SystemVersionOK {
	return &SystemVersionOK{}
}

/*
SystemVersionOK describes a response with status code 200, with default header values.

Success
*/
type SystemVersionOK struct {
	Payload *models.SystemVersionResponse
}

// IsSuccess returns true when this system version o k response has a 2xx status code
func (o *SystemVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system version o k response has a 3xx status code
func (o *SystemVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system version o k response has a 4xx status code
func (o *SystemVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system version o k response has a 5xx status code
func (o *SystemVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system version o k response a status code equal to that given
func (o *SystemVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemVersionOK) Error() string {
	return fmt.Sprintf("[GET /system/version][%d] systemVersionOK  %+v", 200, o.Payload)
}

func (o *SystemVersionOK) String() string {
	return fmt.Sprintf("[GET /system/version][%d] systemVersionOK  %+v", 200, o.Payload)
}

func (o *SystemVersionOK) GetPayload() *models.SystemVersionResponse {
	return o.Payload
}

func (o *SystemVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
