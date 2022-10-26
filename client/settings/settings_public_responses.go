// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// SettingsPublicReader is a Reader for the SettingsPublic structure.
type SettingsPublicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsPublicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingsPublicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSettingsPublicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSettingsPublicOK creates a SettingsPublicOK with default headers values
func NewSettingsPublicOK() *SettingsPublicOK {
	return &SettingsPublicOK{}
}

/*
SettingsPublicOK describes a response with status code 200, with default header values.

Success
*/
type SettingsPublicOK struct {
	Payload *models.SettingsPublicSettingsResponse
}

// IsSuccess returns true when this settings public o k response has a 2xx status code
func (o *SettingsPublicOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this settings public o k response has a 3xx status code
func (o *SettingsPublicOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings public o k response has a 4xx status code
func (o *SettingsPublicOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this settings public o k response has a 5xx status code
func (o *SettingsPublicOK) IsServerError() bool {
	return false
}

// IsCode returns true when this settings public o k response a status code equal to that given
func (o *SettingsPublicOK) IsCode(code int) bool {
	return code == 200
}

func (o *SettingsPublicOK) Error() string {
	return fmt.Sprintf("[GET /settings/public][%d] settingsPublicOK  %+v", 200, o.Payload)
}

func (o *SettingsPublicOK) String() string {
	return fmt.Sprintf("[GET /settings/public][%d] settingsPublicOK  %+v", 200, o.Payload)
}

func (o *SettingsPublicOK) GetPayload() *models.SettingsPublicSettingsResponse {
	return o.Payload
}

func (o *SettingsPublicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsPublicSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSettingsPublicInternalServerError creates a SettingsPublicInternalServerError with default headers values
func NewSettingsPublicInternalServerError() *SettingsPublicInternalServerError {
	return &SettingsPublicInternalServerError{}
}

/*
SettingsPublicInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SettingsPublicInternalServerError struct {
}

// IsSuccess returns true when this settings public internal server error response has a 2xx status code
func (o *SettingsPublicInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this settings public internal server error response has a 3xx status code
func (o *SettingsPublicInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings public internal server error response has a 4xx status code
func (o *SettingsPublicInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this settings public internal server error response has a 5xx status code
func (o *SettingsPublicInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this settings public internal server error response a status code equal to that given
func (o *SettingsPublicInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SettingsPublicInternalServerError) Error() string {
	return fmt.Sprintf("[GET /settings/public][%d] settingsPublicInternalServerError ", 500)
}

func (o *SettingsPublicInternalServerError) String() string {
	return fmt.Sprintf("[GET /settings/public][%d] settingsPublicInternalServerError ", 500)
}

func (o *SettingsPublicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
