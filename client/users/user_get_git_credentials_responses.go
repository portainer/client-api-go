// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models"
)

// UserGetGitCredentialsReader is a Reader for the UserGetGitCredentials structure.
type UserGetGitCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetGitCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetGitCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGetGitCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGetGitCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGetGitCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGetGitCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGetGitCredentialsOK creates a UserGetGitCredentialsOK with default headers values
func NewUserGetGitCredentialsOK() *UserGetGitCredentialsOK {
	return &UserGetGitCredentialsOK{}
}

/*
UserGetGitCredentialsOK describes a response with status code 200, with default header values.

Success
*/
type UserGetGitCredentialsOK struct {
	Payload []*models.PortainereeGitCredential
}

// IsSuccess returns true when this user get git credentials o k response has a 2xx status code
func (o *UserGetGitCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user get git credentials o k response has a 3xx status code
func (o *UserGetGitCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get git credentials o k response has a 4xx status code
func (o *UserGetGitCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user get git credentials o k response has a 5xx status code
func (o *UserGetGitCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user get git credentials o k response a status code equal to that given
func (o *UserGetGitCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserGetGitCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsOK  %+v", 200, o.Payload)
}

func (o *UserGetGitCredentialsOK) String() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsOK  %+v", 200, o.Payload)
}

func (o *UserGetGitCredentialsOK) GetPayload() []*models.PortainereeGitCredential {
	return o.Payload
}

func (o *UserGetGitCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetGitCredentialsBadRequest creates a UserGetGitCredentialsBadRequest with default headers values
func NewUserGetGitCredentialsBadRequest() *UserGetGitCredentialsBadRequest {
	return &UserGetGitCredentialsBadRequest{}
}

/*
UserGetGitCredentialsBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UserGetGitCredentialsBadRequest struct {
}

// IsSuccess returns true when this user get git credentials bad request response has a 2xx status code
func (o *UserGetGitCredentialsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user get git credentials bad request response has a 3xx status code
func (o *UserGetGitCredentialsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get git credentials bad request response has a 4xx status code
func (o *UserGetGitCredentialsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user get git credentials bad request response has a 5xx status code
func (o *UserGetGitCredentialsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user get git credentials bad request response a status code equal to that given
func (o *UserGetGitCredentialsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserGetGitCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsBadRequest ", 400)
}

func (o *UserGetGitCredentialsBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsBadRequest ", 400)
}

func (o *UserGetGitCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetGitCredentialsForbidden creates a UserGetGitCredentialsForbidden with default headers values
func NewUserGetGitCredentialsForbidden() *UserGetGitCredentialsForbidden {
	return &UserGetGitCredentialsForbidden{}
}

/*
UserGetGitCredentialsForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type UserGetGitCredentialsForbidden struct {
}

// IsSuccess returns true when this user get git credentials forbidden response has a 2xx status code
func (o *UserGetGitCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user get git credentials forbidden response has a 3xx status code
func (o *UserGetGitCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get git credentials forbidden response has a 4xx status code
func (o *UserGetGitCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user get git credentials forbidden response has a 5xx status code
func (o *UserGetGitCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user get git credentials forbidden response a status code equal to that given
func (o *UserGetGitCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserGetGitCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsForbidden ", 403)
}

func (o *UserGetGitCredentialsForbidden) String() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsForbidden ", 403)
}

func (o *UserGetGitCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetGitCredentialsNotFound creates a UserGetGitCredentialsNotFound with default headers values
func NewUserGetGitCredentialsNotFound() *UserGetGitCredentialsNotFound {
	return &UserGetGitCredentialsNotFound{}
}

/*
UserGetGitCredentialsNotFound describes a response with status code 404, with default header values.

User not found
*/
type UserGetGitCredentialsNotFound struct {
}

// IsSuccess returns true when this user get git credentials not found response has a 2xx status code
func (o *UserGetGitCredentialsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user get git credentials not found response has a 3xx status code
func (o *UserGetGitCredentialsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get git credentials not found response has a 4xx status code
func (o *UserGetGitCredentialsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user get git credentials not found response has a 5xx status code
func (o *UserGetGitCredentialsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user get git credentials not found response a status code equal to that given
func (o *UserGetGitCredentialsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserGetGitCredentialsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsNotFound ", 404)
}

func (o *UserGetGitCredentialsNotFound) String() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsNotFound ", 404)
}

func (o *UserGetGitCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetGitCredentialsInternalServerError creates a UserGetGitCredentialsInternalServerError with default headers values
func NewUserGetGitCredentialsInternalServerError() *UserGetGitCredentialsInternalServerError {
	return &UserGetGitCredentialsInternalServerError{}
}

/*
UserGetGitCredentialsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UserGetGitCredentialsInternalServerError struct {
}

// IsSuccess returns true when this user get git credentials internal server error response has a 2xx status code
func (o *UserGetGitCredentialsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user get git credentials internal server error response has a 3xx status code
func (o *UserGetGitCredentialsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get git credentials internal server error response has a 4xx status code
func (o *UserGetGitCredentialsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user get git credentials internal server error response has a 5xx status code
func (o *UserGetGitCredentialsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user get git credentials internal server error response a status code equal to that given
func (o *UserGetGitCredentialsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserGetGitCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsInternalServerError ", 500)
}

func (o *UserGetGitCredentialsInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{id}/gitcredentials][%d] userGetGitCredentialsInternalServerError ", 500)
}

func (o *UserGetGitCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
