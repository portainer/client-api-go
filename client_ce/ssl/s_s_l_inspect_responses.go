// Code generated by go-swagger; DO NOT EDIT.

package ssl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// SSLInspectReader is a Reader for the SSLInspect structure.
type SSLInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSLInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSLInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSSLInspectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSSLInspectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSSLInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSSLInspectOK creates a SSLInspectOK with default headers values
func NewSSLInspectOK() *SSLInspectOK {
	return &SSLInspectOK{}
}

/*
SSLInspectOK describes a response with status code 200, with default header values.

Success
*/
type SSLInspectOK struct {
	Payload *models_ce.PortainerSSLSettings
}

// IsSuccess returns true when this s s l inspect o k response has a 2xx status code
func (o *SSLInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s s l inspect o k response has a 3xx status code
func (o *SSLInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s l inspect o k response has a 4xx status code
func (o *SSLInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s s l inspect o k response has a 5xx status code
func (o *SSLInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s s l inspect o k response a status code equal to that given
func (o *SSLInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *SSLInspectOK) Error() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectOK  %+v", 200, o.Payload)
}

func (o *SSLInspectOK) String() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectOK  %+v", 200, o.Payload)
}

func (o *SSLInspectOK) GetPayload() *models_ce.PortainerSSLSettings {
	return o.Payload
}

func (o *SSLInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.PortainerSSLSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSLInspectBadRequest creates a SSLInspectBadRequest with default headers values
func NewSSLInspectBadRequest() *SSLInspectBadRequest {
	return &SSLInspectBadRequest{}
}

/*
SSLInspectBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type SSLInspectBadRequest struct {
}

// IsSuccess returns true when this s s l inspect bad request response has a 2xx status code
func (o *SSLInspectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s s l inspect bad request response has a 3xx status code
func (o *SSLInspectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s l inspect bad request response has a 4xx status code
func (o *SSLInspectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s s l inspect bad request response has a 5xx status code
func (o *SSLInspectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s s l inspect bad request response a status code equal to that given
func (o *SSLInspectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SSLInspectBadRequest) Error() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectBadRequest ", 400)
}

func (o *SSLInspectBadRequest) String() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectBadRequest ", 400)
}

func (o *SSLInspectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSSLInspectForbidden creates a SSLInspectForbidden with default headers values
func NewSSLInspectForbidden() *SSLInspectForbidden {
	return &SSLInspectForbidden{}
}

/*
SSLInspectForbidden describes a response with status code 403, with default header values.

Permission denied to access settings
*/
type SSLInspectForbidden struct {
}

// IsSuccess returns true when this s s l inspect forbidden response has a 2xx status code
func (o *SSLInspectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s s l inspect forbidden response has a 3xx status code
func (o *SSLInspectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s l inspect forbidden response has a 4xx status code
func (o *SSLInspectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s s l inspect forbidden response has a 5xx status code
func (o *SSLInspectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s s l inspect forbidden response a status code equal to that given
func (o *SSLInspectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SSLInspectForbidden) Error() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectForbidden ", 403)
}

func (o *SSLInspectForbidden) String() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectForbidden ", 403)
}

func (o *SSLInspectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSSLInspectInternalServerError creates a SSLInspectInternalServerError with default headers values
func NewSSLInspectInternalServerError() *SSLInspectInternalServerError {
	return &SSLInspectInternalServerError{}
}

/*
SSLInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SSLInspectInternalServerError struct {
}

// IsSuccess returns true when this s s l inspect internal server error response has a 2xx status code
func (o *SSLInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s s l inspect internal server error response has a 3xx status code
func (o *SSLInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s l inspect internal server error response has a 4xx status code
func (o *SSLInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this s s l inspect internal server error response has a 5xx status code
func (o *SSLInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this s s l inspect internal server error response a status code equal to that given
func (o *SSLInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SSLInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectInternalServerError ", 500)
}

func (o *SSLInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /ssl][%d] sSLInspectInternalServerError ", 500)
}

func (o *SSLInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
