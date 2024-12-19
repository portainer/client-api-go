// Code generated by go-swagger; DO NOT EDIT.

package kaas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// KaasVersionReader is a Reader for the KaasVersion structure.
type KaasVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KaasVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKaasVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKaasVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKaasVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKaasVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewKaasVersionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud/endpoints/{environmentId}/version] kaasVersion", response, response.Code())
	}
}

// NewKaasVersionOK creates a KaasVersionOK with default headers values
func NewKaasVersionOK() *KaasVersionOK {
	return &KaasVersionOK{}
}

/*
KaasVersionOK describes a response with status code 200, with default header values.

Success
*/
type KaasVersionOK struct {
}

// IsSuccess returns true when this kaas version o k response has a 2xx status code
func (o *KaasVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kaas version o k response has a 3xx status code
func (o *KaasVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kaas version o k response has a 4xx status code
func (o *KaasVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kaas version o k response has a 5xx status code
func (o *KaasVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kaas version o k response a status code equal to that given
func (o *KaasVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kaas version o k response
func (o *KaasVersionOK) Code() int {
	return 200
}

func (o *KaasVersionOK) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionOK", 200)
}

func (o *KaasVersionOK) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionOK", 200)
}

func (o *KaasVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKaasVersionBadRequest creates a KaasVersionBadRequest with default headers values
func NewKaasVersionBadRequest() *KaasVersionBadRequest {
	return &KaasVersionBadRequest{}
}

/*
KaasVersionBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type KaasVersionBadRequest struct {
}

// IsSuccess returns true when this kaas version bad request response has a 2xx status code
func (o *KaasVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kaas version bad request response has a 3xx status code
func (o *KaasVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kaas version bad request response has a 4xx status code
func (o *KaasVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kaas version bad request response has a 5xx status code
func (o *KaasVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kaas version bad request response a status code equal to that given
func (o *KaasVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kaas version bad request response
func (o *KaasVersionBadRequest) Code() int {
	return 400
}

func (o *KaasVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionBadRequest", 400)
}

func (o *KaasVersionBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionBadRequest", 400)
}

func (o *KaasVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKaasVersionForbidden creates a KaasVersionForbidden with default headers values
func NewKaasVersionForbidden() *KaasVersionForbidden {
	return &KaasVersionForbidden{}
}

/*
KaasVersionForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type KaasVersionForbidden struct {
}

// IsSuccess returns true when this kaas version forbidden response has a 2xx status code
func (o *KaasVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kaas version forbidden response has a 3xx status code
func (o *KaasVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kaas version forbidden response has a 4xx status code
func (o *KaasVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kaas version forbidden response has a 5xx status code
func (o *KaasVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kaas version forbidden response a status code equal to that given
func (o *KaasVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kaas version forbidden response
func (o *KaasVersionForbidden) Code() int {
	return 403
}

func (o *KaasVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionForbidden", 403)
}

func (o *KaasVersionForbidden) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionForbidden", 403)
}

func (o *KaasVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKaasVersionInternalServerError creates a KaasVersionInternalServerError with default headers values
func NewKaasVersionInternalServerError() *KaasVersionInternalServerError {
	return &KaasVersionInternalServerError{}
}

/*
KaasVersionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type KaasVersionInternalServerError struct {
}

// IsSuccess returns true when this kaas version internal server error response has a 2xx status code
func (o *KaasVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kaas version internal server error response has a 3xx status code
func (o *KaasVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kaas version internal server error response has a 4xx status code
func (o *KaasVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kaas version internal server error response has a 5xx status code
func (o *KaasVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kaas version internal server error response a status code equal to that given
func (o *KaasVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kaas version internal server error response
func (o *KaasVersionInternalServerError) Code() int {
	return 500
}

func (o *KaasVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionInternalServerError", 500)
}

func (o *KaasVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionInternalServerError", 500)
}

func (o *KaasVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKaasVersionServiceUnavailable creates a KaasVersionServiceUnavailable with default headers values
func NewKaasVersionServiceUnavailable() *KaasVersionServiceUnavailable {
	return &KaasVersionServiceUnavailable{}
}

/*
KaasVersionServiceUnavailable describes a response with status code 503, with default header values.

Missing configuration
*/
type KaasVersionServiceUnavailable struct {
}

// IsSuccess returns true when this kaas version service unavailable response has a 2xx status code
func (o *KaasVersionServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kaas version service unavailable response has a 3xx status code
func (o *KaasVersionServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kaas version service unavailable response has a 4xx status code
func (o *KaasVersionServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this kaas version service unavailable response has a 5xx status code
func (o *KaasVersionServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this kaas version service unavailable response a status code equal to that given
func (o *KaasVersionServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the kaas version service unavailable response
func (o *KaasVersionServiceUnavailable) Code() int {
	return 503
}

func (o *KaasVersionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionServiceUnavailable", 503)
}

func (o *KaasVersionServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /cloud/endpoints/{environmentId}/version][%d] kaasVersionServiceUnavailable", 503)
}

func (o *KaasVersionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
