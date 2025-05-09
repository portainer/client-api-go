// Code generated by go-swagger; DO NOT EDIT.

package kaas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Microk8sGetAddonsReader is a Reader for the Microk8sGetAddons structure.
type Microk8sGetAddonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Microk8sGetAddonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMicrok8sGetAddonsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMicrok8sGetAddonsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMicrok8sGetAddonsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMicrok8sGetAddonsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewMicrok8sGetAddonsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud/endpoints/{environmentId}/addons] microk8sGetAddons", response, response.Code())
	}
}

// NewMicrok8sGetAddonsOK creates a Microk8sGetAddonsOK with default headers values
func NewMicrok8sGetAddonsOK() *Microk8sGetAddonsOK {
	return &Microk8sGetAddonsOK{}
}

/*
Microk8sGetAddonsOK describes a response with status code 200, with default header values.

Success
*/
type Microk8sGetAddonsOK struct {
}

// IsSuccess returns true when this microk8s get addons o k response has a 2xx status code
func (o *Microk8sGetAddonsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this microk8s get addons o k response has a 3xx status code
func (o *Microk8sGetAddonsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get addons o k response has a 4xx status code
func (o *Microk8sGetAddonsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this microk8s get addons o k response has a 5xx status code
func (o *Microk8sGetAddonsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this microk8s get addons o k response a status code equal to that given
func (o *Microk8sGetAddonsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the microk8s get addons o k response
func (o *Microk8sGetAddonsOK) Code() int {
	return 200
}

func (o *Microk8sGetAddonsOK) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsOK", 200)
}

func (o *Microk8sGetAddonsOK) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsOK", 200)
}

func (o *Microk8sGetAddonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetAddonsBadRequest creates a Microk8sGetAddonsBadRequest with default headers values
func NewMicrok8sGetAddonsBadRequest() *Microk8sGetAddonsBadRequest {
	return &Microk8sGetAddonsBadRequest{}
}

/*
Microk8sGetAddonsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type Microk8sGetAddonsBadRequest struct {
}

// IsSuccess returns true when this microk8s get addons bad request response has a 2xx status code
func (o *Microk8sGetAddonsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get addons bad request response has a 3xx status code
func (o *Microk8sGetAddonsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get addons bad request response has a 4xx status code
func (o *Microk8sGetAddonsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this microk8s get addons bad request response has a 5xx status code
func (o *Microk8sGetAddonsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this microk8s get addons bad request response a status code equal to that given
func (o *Microk8sGetAddonsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the microk8s get addons bad request response
func (o *Microk8sGetAddonsBadRequest) Code() int {
	return 400
}

func (o *Microk8sGetAddonsBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsBadRequest", 400)
}

func (o *Microk8sGetAddonsBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsBadRequest", 400)
}

func (o *Microk8sGetAddonsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetAddonsForbidden creates a Microk8sGetAddonsForbidden with default headers values
func NewMicrok8sGetAddonsForbidden() *Microk8sGetAddonsForbidden {
	return &Microk8sGetAddonsForbidden{}
}

/*
Microk8sGetAddonsForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type Microk8sGetAddonsForbidden struct {
}

// IsSuccess returns true when this microk8s get addons forbidden response has a 2xx status code
func (o *Microk8sGetAddonsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get addons forbidden response has a 3xx status code
func (o *Microk8sGetAddonsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get addons forbidden response has a 4xx status code
func (o *Microk8sGetAddonsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this microk8s get addons forbidden response has a 5xx status code
func (o *Microk8sGetAddonsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this microk8s get addons forbidden response a status code equal to that given
func (o *Microk8sGetAddonsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the microk8s get addons forbidden response
func (o *Microk8sGetAddonsForbidden) Code() int {
	return 403
}

func (o *Microk8sGetAddonsForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsForbidden", 403)
}

func (o *Microk8sGetAddonsForbidden) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsForbidden", 403)
}

func (o *Microk8sGetAddonsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetAddonsInternalServerError creates a Microk8sGetAddonsInternalServerError with default headers values
func NewMicrok8sGetAddonsInternalServerError() *Microk8sGetAddonsInternalServerError {
	return &Microk8sGetAddonsInternalServerError{}
}

/*
Microk8sGetAddonsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type Microk8sGetAddonsInternalServerError struct {
}

// IsSuccess returns true when this microk8s get addons internal server error response has a 2xx status code
func (o *Microk8sGetAddonsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get addons internal server error response has a 3xx status code
func (o *Microk8sGetAddonsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get addons internal server error response has a 4xx status code
func (o *Microk8sGetAddonsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this microk8s get addons internal server error response has a 5xx status code
func (o *Microk8sGetAddonsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this microk8s get addons internal server error response a status code equal to that given
func (o *Microk8sGetAddonsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the microk8s get addons internal server error response
func (o *Microk8sGetAddonsInternalServerError) Code() int {
	return 500
}

func (o *Microk8sGetAddonsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsInternalServerError", 500)
}

func (o *Microk8sGetAddonsInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsInternalServerError", 500)
}

func (o *Microk8sGetAddonsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMicrok8sGetAddonsServiceUnavailable creates a Microk8sGetAddonsServiceUnavailable with default headers values
func NewMicrok8sGetAddonsServiceUnavailable() *Microk8sGetAddonsServiceUnavailable {
	return &Microk8sGetAddonsServiceUnavailable{}
}

/*
Microk8sGetAddonsServiceUnavailable describes a response with status code 503, with default header values.

Missing configuration
*/
type Microk8sGetAddonsServiceUnavailable struct {
}

// IsSuccess returns true when this microk8s get addons service unavailable response has a 2xx status code
func (o *Microk8sGetAddonsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this microk8s get addons service unavailable response has a 3xx status code
func (o *Microk8sGetAddonsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this microk8s get addons service unavailable response has a 4xx status code
func (o *Microk8sGetAddonsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this microk8s get addons service unavailable response has a 5xx status code
func (o *Microk8sGetAddonsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this microk8s get addons service unavailable response a status code equal to that given
func (o *Microk8sGetAddonsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the microk8s get addons service unavailable response
func (o *Microk8sGetAddonsServiceUnavailable) Code() int {
	return 503
}

func (o *Microk8sGetAddonsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsServiceUnavailable", 503)
}

func (o *Microk8sGetAddonsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/addons][%d] microk8sGetAddonsServiceUnavailable", 503)
}

func (o *Microk8sGetAddonsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
