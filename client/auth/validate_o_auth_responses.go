// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// ValidateOAuthReader is a Reader for the ValidateOAuth structure.
type ValidateOAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateOAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateOAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateOAuthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewValidateOAuthUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateOAuthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/oauth/validate] ValidateOAuth", response, response.Code())
	}
}

// NewValidateOAuthOK creates a ValidateOAuthOK with default headers values
func NewValidateOAuthOK() *ValidateOAuthOK {
	return &ValidateOAuthOK{}
}

/*
ValidateOAuthOK describes a response with status code 200, with default header values.

Success
*/
type ValidateOAuthOK struct {
	Payload *models.AuthAuthenticateResponse
}

// IsSuccess returns true when this validate o auth o k response has a 2xx status code
func (o *ValidateOAuthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate o auth o k response has a 3xx status code
func (o *ValidateOAuthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate o auth o k response has a 4xx status code
func (o *ValidateOAuthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate o auth o k response has a 5xx status code
func (o *ValidateOAuthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate o auth o k response a status code equal to that given
func (o *ValidateOAuthOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the validate o auth o k response
func (o *ValidateOAuthOK) Code() int {
	return 200
}

func (o *ValidateOAuthOK) Error() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthOK  %+v", 200, o.Payload)
}

func (o *ValidateOAuthOK) String() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthOK  %+v", 200, o.Payload)
}

func (o *ValidateOAuthOK) GetPayload() *models.AuthAuthenticateResponse {
	return o.Payload
}

func (o *ValidateOAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthAuthenticateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateOAuthBadRequest creates a ValidateOAuthBadRequest with default headers values
func NewValidateOAuthBadRequest() *ValidateOAuthBadRequest {
	return &ValidateOAuthBadRequest{}
}

/*
ValidateOAuthBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type ValidateOAuthBadRequest struct {
}

// IsSuccess returns true when this validate o auth bad request response has a 2xx status code
func (o *ValidateOAuthBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate o auth bad request response has a 3xx status code
func (o *ValidateOAuthBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate o auth bad request response has a 4xx status code
func (o *ValidateOAuthBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate o auth bad request response has a 5xx status code
func (o *ValidateOAuthBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this validate o auth bad request response a status code equal to that given
func (o *ValidateOAuthBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the validate o auth bad request response
func (o *ValidateOAuthBadRequest) Code() int {
	return 400
}

func (o *ValidateOAuthBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthBadRequest ", 400)
}

func (o *ValidateOAuthBadRequest) String() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthBadRequest ", 400)
}

func (o *ValidateOAuthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateOAuthUnprocessableEntity creates a ValidateOAuthUnprocessableEntity with default headers values
func NewValidateOAuthUnprocessableEntity() *ValidateOAuthUnprocessableEntity {
	return &ValidateOAuthUnprocessableEntity{}
}

/*
ValidateOAuthUnprocessableEntity describes a response with status code 422, with default header values.

Invalid Credentials
*/
type ValidateOAuthUnprocessableEntity struct {
}

// IsSuccess returns true when this validate o auth unprocessable entity response has a 2xx status code
func (o *ValidateOAuthUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate o auth unprocessable entity response has a 3xx status code
func (o *ValidateOAuthUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate o auth unprocessable entity response has a 4xx status code
func (o *ValidateOAuthUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate o auth unprocessable entity response has a 5xx status code
func (o *ValidateOAuthUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this validate o auth unprocessable entity response a status code equal to that given
func (o *ValidateOAuthUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the validate o auth unprocessable entity response
func (o *ValidateOAuthUnprocessableEntity) Code() int {
	return 422
}

func (o *ValidateOAuthUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthUnprocessableEntity ", 422)
}

func (o *ValidateOAuthUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthUnprocessableEntity ", 422)
}

func (o *ValidateOAuthUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateOAuthInternalServerError creates a ValidateOAuthInternalServerError with default headers values
func NewValidateOAuthInternalServerError() *ValidateOAuthInternalServerError {
	return &ValidateOAuthInternalServerError{}
}

/*
ValidateOAuthInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ValidateOAuthInternalServerError struct {
}

// IsSuccess returns true when this validate o auth internal server error response has a 2xx status code
func (o *ValidateOAuthInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate o auth internal server error response has a 3xx status code
func (o *ValidateOAuthInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate o auth internal server error response has a 4xx status code
func (o *ValidateOAuthInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate o auth internal server error response has a 5xx status code
func (o *ValidateOAuthInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate o auth internal server error response a status code equal to that given
func (o *ValidateOAuthInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the validate o auth internal server error response
func (o *ValidateOAuthInternalServerError) Code() int {
	return 500
}

func (o *ValidateOAuthInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthInternalServerError ", 500)
}

func (o *ValidateOAuthInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/oauth/validate][%d] validateOAuthInternalServerError ", 500)
}

func (o *ValidateOAuthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
