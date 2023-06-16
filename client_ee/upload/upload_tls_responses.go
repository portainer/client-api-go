// Code generated by go-swagger; DO NOT EDIT.

package upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UploadTLSReader is a Reader for the UploadTLS structure.
type UploadTLSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadTLSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUploadTLSNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadTLSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadTLSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadTLSNoContent creates a UploadTLSNoContent with default headers values
func NewUploadTLSNoContent() *UploadTLSNoContent {
	return &UploadTLSNoContent{}
}

/*
UploadTLSNoContent describes a response with status code 204, with default header values.

Success
*/
type UploadTLSNoContent struct {
}

// IsSuccess returns true when this upload Tls no content response has a 2xx status code
func (o *UploadTLSNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload Tls no content response has a 3xx status code
func (o *UploadTLSNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload Tls no content response has a 4xx status code
func (o *UploadTLSNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload Tls no content response has a 5xx status code
func (o *UploadTLSNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this upload Tls no content response a status code equal to that given
func (o *UploadTLSNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UploadTLSNoContent) Error() string {
	return fmt.Sprintf("[POST /upload/tls/{certificate}][%d] uploadTlsNoContent ", 204)
}

func (o *UploadTLSNoContent) String() string {
	return fmt.Sprintf("[POST /upload/tls/{certificate}][%d] uploadTlsNoContent ", 204)
}

func (o *UploadTLSNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadTLSBadRequest creates a UploadTLSBadRequest with default headers values
func NewUploadTLSBadRequest() *UploadTLSBadRequest {
	return &UploadTLSBadRequest{}
}

/*
UploadTLSBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UploadTLSBadRequest struct {
}

// IsSuccess returns true when this upload Tls bad request response has a 2xx status code
func (o *UploadTLSBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload Tls bad request response has a 3xx status code
func (o *UploadTLSBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload Tls bad request response has a 4xx status code
func (o *UploadTLSBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload Tls bad request response has a 5xx status code
func (o *UploadTLSBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload Tls bad request response a status code equal to that given
func (o *UploadTLSBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UploadTLSBadRequest) Error() string {
	return fmt.Sprintf("[POST /upload/tls/{certificate}][%d] uploadTlsBadRequest ", 400)
}

func (o *UploadTLSBadRequest) String() string {
	return fmt.Sprintf("[POST /upload/tls/{certificate}][%d] uploadTlsBadRequest ", 400)
}

func (o *UploadTLSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadTLSInternalServerError creates a UploadTLSInternalServerError with default headers values
func NewUploadTLSInternalServerError() *UploadTLSInternalServerError {
	return &UploadTLSInternalServerError{}
}

/*
UploadTLSInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UploadTLSInternalServerError struct {
}

// IsSuccess returns true when this upload Tls internal server error response has a 2xx status code
func (o *UploadTLSInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload Tls internal server error response has a 3xx status code
func (o *UploadTLSInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload Tls internal server error response has a 4xx status code
func (o *UploadTLSInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload Tls internal server error response has a 5xx status code
func (o *UploadTLSInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload Tls internal server error response a status code equal to that given
func (o *UploadTLSInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UploadTLSInternalServerError) Error() string {
	return fmt.Sprintf("[POST /upload/tls/{certificate}][%d] uploadTlsInternalServerError ", 500)
}

func (o *UploadTLSInternalServerError) String() string {
	return fmt.Sprintf("[POST /upload/tls/{certificate}][%d] uploadTlsInternalServerError ", 500)
}

func (o *UploadTLSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
