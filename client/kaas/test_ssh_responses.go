// Code generated by go-swagger; DO NOT EDIT.

package kaas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// TestSSHReader is a Reader for the TestSSH structure.
type TestSSHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestSSHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestSSHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestSSHBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTestSSHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewTestSSHServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTestSSHOK creates a TestSSHOK with default headers values
func NewTestSSHOK() *TestSSHOK {
	return &TestSSHOK{}
}

/*
TestSSHOK describes a response with status code 200, with default header values.

Success
*/
type TestSSHOK struct {
	Payload []*models.SshtestSSHTestResult
}

// IsSuccess returns true when this test Ssh o k response has a 2xx status code
func (o *TestSSHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this test Ssh o k response has a 3xx status code
func (o *TestSSHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test Ssh o k response has a 4xx status code
func (o *TestSSHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this test Ssh o k response has a 5xx status code
func (o *TestSSHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this test Ssh o k response a status code equal to that given
func (o *TestSSHOK) IsCode(code int) bool {
	return code == 200
}

func (o *TestSSHOK) Error() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshOK  %+v", 200, o.Payload)
}

func (o *TestSSHOK) String() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshOK  %+v", 200, o.Payload)
}

func (o *TestSSHOK) GetPayload() []*models.SshtestSSHTestResult {
	return o.Payload
}

func (o *TestSSHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestSSHBadRequest creates a TestSSHBadRequest with default headers values
func NewTestSSHBadRequest() *TestSSHBadRequest {
	return &TestSSHBadRequest{}
}

/*
TestSSHBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type TestSSHBadRequest struct {
}

// IsSuccess returns true when this test Ssh bad request response has a 2xx status code
func (o *TestSSHBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test Ssh bad request response has a 3xx status code
func (o *TestSSHBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test Ssh bad request response has a 4xx status code
func (o *TestSSHBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this test Ssh bad request response has a 5xx status code
func (o *TestSSHBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this test Ssh bad request response a status code equal to that given
func (o *TestSSHBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TestSSHBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshBadRequest ", 400)
}

func (o *TestSSHBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshBadRequest ", 400)
}

func (o *TestSSHBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestSSHInternalServerError creates a TestSSHInternalServerError with default headers values
func NewTestSSHInternalServerError() *TestSSHInternalServerError {
	return &TestSSHInternalServerError{}
}

/*
TestSSHInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type TestSSHInternalServerError struct {
}

// IsSuccess returns true when this test Ssh internal server error response has a 2xx status code
func (o *TestSSHInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test Ssh internal server error response has a 3xx status code
func (o *TestSSHInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test Ssh internal server error response has a 4xx status code
func (o *TestSSHInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this test Ssh internal server error response has a 5xx status code
func (o *TestSSHInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this test Ssh internal server error response a status code equal to that given
func (o *TestSSHInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TestSSHInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshInternalServerError ", 500)
}

func (o *TestSSHInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshInternalServerError ", 500)
}

func (o *TestSSHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestSSHServiceUnavailable creates a TestSSHServiceUnavailable with default headers values
func NewTestSSHServiceUnavailable() *TestSSHServiceUnavailable {
	return &TestSSHServiceUnavailable{}
}

/*
TestSSHServiceUnavailable describes a response with status code 503, with default header values.

Missing configuration
*/
type TestSSHServiceUnavailable struct {
}

// IsSuccess returns true when this test Ssh service unavailable response has a 2xx status code
func (o *TestSSHServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test Ssh service unavailable response has a 3xx status code
func (o *TestSSHServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test Ssh service unavailable response has a 4xx status code
func (o *TestSSHServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this test Ssh service unavailable response has a 5xx status code
func (o *TestSSHServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this test Ssh service unavailable response a status code equal to that given
func (o *TestSSHServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *TestSSHServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshServiceUnavailable ", 503)
}

func (o *TestSSHServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /cloud/testssh][%d] testSshServiceUnavailable ", 503)
}

func (o *TestSSHServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
