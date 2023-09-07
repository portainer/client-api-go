// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SettingsExperimentalUpdateReader is a Reader for the SettingsExperimentalUpdate structure.
type SettingsExperimentalUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsExperimentalUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSettingsExperimentalUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSettingsExperimentalUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSettingsExperimentalUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSettingsExperimentalUpdateNoContent creates a SettingsExperimentalUpdateNoContent with default headers values
func NewSettingsExperimentalUpdateNoContent() *SettingsExperimentalUpdateNoContent {
	return &SettingsExperimentalUpdateNoContent{}
}

/*
SettingsExperimentalUpdateNoContent describes a response with status code 204, with default header values.

Success
*/
type SettingsExperimentalUpdateNoContent struct {
}

// IsSuccess returns true when this settings experimental update no content response has a 2xx status code
func (o *SettingsExperimentalUpdateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this settings experimental update no content response has a 3xx status code
func (o *SettingsExperimentalUpdateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings experimental update no content response has a 4xx status code
func (o *SettingsExperimentalUpdateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this settings experimental update no content response has a 5xx status code
func (o *SettingsExperimentalUpdateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this settings experimental update no content response a status code equal to that given
func (o *SettingsExperimentalUpdateNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *SettingsExperimentalUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /settings/experimental][%d] settingsExperimentalUpdateNoContent ", 204)
}

func (o *SettingsExperimentalUpdateNoContent) String() string {
	return fmt.Sprintf("[PUT /settings/experimental][%d] settingsExperimentalUpdateNoContent ", 204)
}

func (o *SettingsExperimentalUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSettingsExperimentalUpdateBadRequest creates a SettingsExperimentalUpdateBadRequest with default headers values
func NewSettingsExperimentalUpdateBadRequest() *SettingsExperimentalUpdateBadRequest {
	return &SettingsExperimentalUpdateBadRequest{}
}

/*
SettingsExperimentalUpdateBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type SettingsExperimentalUpdateBadRequest struct {
}

// IsSuccess returns true when this settings experimental update bad request response has a 2xx status code
func (o *SettingsExperimentalUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this settings experimental update bad request response has a 3xx status code
func (o *SettingsExperimentalUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings experimental update bad request response has a 4xx status code
func (o *SettingsExperimentalUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this settings experimental update bad request response has a 5xx status code
func (o *SettingsExperimentalUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this settings experimental update bad request response a status code equal to that given
func (o *SettingsExperimentalUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SettingsExperimentalUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /settings/experimental][%d] settingsExperimentalUpdateBadRequest ", 400)
}

func (o *SettingsExperimentalUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /settings/experimental][%d] settingsExperimentalUpdateBadRequest ", 400)
}

func (o *SettingsExperimentalUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSettingsExperimentalUpdateInternalServerError creates a SettingsExperimentalUpdateInternalServerError with default headers values
func NewSettingsExperimentalUpdateInternalServerError() *SettingsExperimentalUpdateInternalServerError {
	return &SettingsExperimentalUpdateInternalServerError{}
}

/*
SettingsExperimentalUpdateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SettingsExperimentalUpdateInternalServerError struct {
}

// IsSuccess returns true when this settings experimental update internal server error response has a 2xx status code
func (o *SettingsExperimentalUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this settings experimental update internal server error response has a 3xx status code
func (o *SettingsExperimentalUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings experimental update internal server error response has a 4xx status code
func (o *SettingsExperimentalUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this settings experimental update internal server error response has a 5xx status code
func (o *SettingsExperimentalUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this settings experimental update internal server error response a status code equal to that given
func (o *SettingsExperimentalUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SettingsExperimentalUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /settings/experimental][%d] settingsExperimentalUpdateInternalServerError ", 500)
}

func (o *SettingsExperimentalUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /settings/experimental][%d] settingsExperimentalUpdateInternalServerError ", 500)
}

func (o *SettingsExperimentalUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
