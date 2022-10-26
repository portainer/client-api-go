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

// AuthenticateUserReader is a Reader for the AuthenticateUser structure.
type AuthenticateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthenticateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthenticateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAuthenticateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthenticateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthenticateUserOK creates a AuthenticateUserOK with default headers values
func NewAuthenticateUserOK() *AuthenticateUserOK {
	return &AuthenticateUserOK{}
}

/*
AuthenticateUserOK describes a response with status code 200, with default header values.

Success
*/
type AuthenticateUserOK struct {
	Payload *models.AuthAuthenticateResponse
}

// IsSuccess returns true when this authenticate user o k response has a 2xx status code
func (o *AuthenticateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authenticate user o k response has a 3xx status code
func (o *AuthenticateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authenticate user o k response has a 4xx status code
func (o *AuthenticateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authenticate user o k response has a 5xx status code
func (o *AuthenticateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authenticate user o k response a status code equal to that given
func (o *AuthenticateUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthenticateUserOK) Error() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserOK  %+v", 200, o.Payload)
}

func (o *AuthenticateUserOK) String() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserOK  %+v", 200, o.Payload)
}

func (o *AuthenticateUserOK) GetPayload() *models.AuthAuthenticateResponse {
	return o.Payload
}

func (o *AuthenticateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthAuthenticateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthenticateUserBadRequest creates a AuthenticateUserBadRequest with default headers values
func NewAuthenticateUserBadRequest() *AuthenticateUserBadRequest {
	return &AuthenticateUserBadRequest{}
}

/*
AuthenticateUserBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type AuthenticateUserBadRequest struct {
}

// IsSuccess returns true when this authenticate user bad request response has a 2xx status code
func (o *AuthenticateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authenticate user bad request response has a 3xx status code
func (o *AuthenticateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authenticate user bad request response has a 4xx status code
func (o *AuthenticateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this authenticate user bad request response has a 5xx status code
func (o *AuthenticateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this authenticate user bad request response a status code equal to that given
func (o *AuthenticateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AuthenticateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserBadRequest ", 400)
}

func (o *AuthenticateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserBadRequest ", 400)
}

func (o *AuthenticateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthenticateUserUnprocessableEntity creates a AuthenticateUserUnprocessableEntity with default headers values
func NewAuthenticateUserUnprocessableEntity() *AuthenticateUserUnprocessableEntity {
	return &AuthenticateUserUnprocessableEntity{}
}

/*
AuthenticateUserUnprocessableEntity describes a response with status code 422, with default header values.

Invalid Credentials
*/
type AuthenticateUserUnprocessableEntity struct {
}

// IsSuccess returns true when this authenticate user unprocessable entity response has a 2xx status code
func (o *AuthenticateUserUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authenticate user unprocessable entity response has a 3xx status code
func (o *AuthenticateUserUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authenticate user unprocessable entity response has a 4xx status code
func (o *AuthenticateUserUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this authenticate user unprocessable entity response has a 5xx status code
func (o *AuthenticateUserUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this authenticate user unprocessable entity response a status code equal to that given
func (o *AuthenticateUserUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *AuthenticateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserUnprocessableEntity ", 422)
}

func (o *AuthenticateUserUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserUnprocessableEntity ", 422)
}

func (o *AuthenticateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthenticateUserInternalServerError creates a AuthenticateUserInternalServerError with default headers values
func NewAuthenticateUserInternalServerError() *AuthenticateUserInternalServerError {
	return &AuthenticateUserInternalServerError{}
}

/*
AuthenticateUserInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type AuthenticateUserInternalServerError struct {
}

// IsSuccess returns true when this authenticate user internal server error response has a 2xx status code
func (o *AuthenticateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authenticate user internal server error response has a 3xx status code
func (o *AuthenticateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authenticate user internal server error response has a 4xx status code
func (o *AuthenticateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this authenticate user internal server error response has a 5xx status code
func (o *AuthenticateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this authenticate user internal server error response a status code equal to that given
func (o *AuthenticateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AuthenticateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserInternalServerError ", 500)
}

func (o *AuthenticateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth][%d] authenticateUserInternalServerError ", 500)
}

func (o *AuthenticateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
