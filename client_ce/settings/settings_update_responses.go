// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// SettingsUpdateReader is a Reader for the SettingsUpdate structure.
type SettingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSettingsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSettingsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSettingsUpdateOK creates a SettingsUpdateOK with default headers values
func NewSettingsUpdateOK() *SettingsUpdateOK {
	return &SettingsUpdateOK{}
}

/*
SettingsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type SettingsUpdateOK struct {
	Payload *models_ce.PortainerSettings
}

// IsSuccess returns true when this settings update o k response has a 2xx status code
func (o *SettingsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this settings update o k response has a 3xx status code
func (o *SettingsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings update o k response has a 4xx status code
func (o *SettingsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this settings update o k response has a 5xx status code
func (o *SettingsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this settings update o k response a status code equal to that given
func (o *SettingsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SettingsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /settings][%d] settingsUpdateOK  %+v", 200, o.Payload)
}

func (o *SettingsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /settings][%d] settingsUpdateOK  %+v", 200, o.Payload)
}

func (o *SettingsUpdateOK) GetPayload() *models_ce.PortainerSettings {
	return o.Payload
}

func (o *SettingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.PortainerSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSettingsUpdateBadRequest creates a SettingsUpdateBadRequest with default headers values
func NewSettingsUpdateBadRequest() *SettingsUpdateBadRequest {
	return &SettingsUpdateBadRequest{}
}

/*
SettingsUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type SettingsUpdateBadRequest struct {
}

// IsSuccess returns true when this settings update bad request response has a 2xx status code
func (o *SettingsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this settings update bad request response has a 3xx status code
func (o *SettingsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings update bad request response has a 4xx status code
func (o *SettingsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this settings update bad request response has a 5xx status code
func (o *SettingsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this settings update bad request response a status code equal to that given
func (o *SettingsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SettingsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /settings][%d] settingsUpdateBadRequest ", 400)
}

func (o *SettingsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /settings][%d] settingsUpdateBadRequest ", 400)
}

func (o *SettingsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSettingsUpdateInternalServerError creates a SettingsUpdateInternalServerError with default headers values
func NewSettingsUpdateInternalServerError() *SettingsUpdateInternalServerError {
	return &SettingsUpdateInternalServerError{}
}

/*
SettingsUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SettingsUpdateInternalServerError struct {
}

// IsSuccess returns true when this settings update internal server error response has a 2xx status code
func (o *SettingsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this settings update internal server error response has a 3xx status code
func (o *SettingsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings update internal server error response has a 4xx status code
func (o *SettingsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this settings update internal server error response has a 5xx status code
func (o *SettingsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this settings update internal server error response a status code equal to that given
func (o *SettingsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SettingsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /settings][%d] settingsUpdateInternalServerError ", 500)
}

func (o *SettingsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /settings][%d] settingsUpdateInternalServerError ", 500)
}

func (o *SettingsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
