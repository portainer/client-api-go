// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserUpdateGitCredentialReader is a Reader for the UserUpdateGitCredential structure.
type UserUpdateGitCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateGitCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserUpdateGitCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserUpdateGitCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserUpdateGitCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserUpdateGitCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserUpdateGitCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /users/{id}/gitcredentials/{credentialID}] UserUpdateGitCredential", response, response.Code())
	}
}

// NewUserUpdateGitCredentialNoContent creates a UserUpdateGitCredentialNoContent with default headers values
func NewUserUpdateGitCredentialNoContent() *UserUpdateGitCredentialNoContent {
	return &UserUpdateGitCredentialNoContent{}
}

/*
UserUpdateGitCredentialNoContent describes a response with status code 204, with default header values.

Success
*/
type UserUpdateGitCredentialNoContent struct {
}

// IsSuccess returns true when this user update git credential no content response has a 2xx status code
func (o *UserUpdateGitCredentialNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user update git credential no content response has a 3xx status code
func (o *UserUpdateGitCredentialNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update git credential no content response has a 4xx status code
func (o *UserUpdateGitCredentialNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this user update git credential no content response has a 5xx status code
func (o *UserUpdateGitCredentialNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this user update git credential no content response a status code equal to that given
func (o *UserUpdateGitCredentialNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the user update git credential no content response
func (o *UserUpdateGitCredentialNoContent) Code() int {
	return 204
}

func (o *UserUpdateGitCredentialNoContent) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialNoContent", 204)
}

func (o *UserUpdateGitCredentialNoContent) String() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialNoContent", 204)
}

func (o *UserUpdateGitCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateGitCredentialBadRequest creates a UserUpdateGitCredentialBadRequest with default headers values
func NewUserUpdateGitCredentialBadRequest() *UserUpdateGitCredentialBadRequest {
	return &UserUpdateGitCredentialBadRequest{}
}

/*
UserUpdateGitCredentialBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UserUpdateGitCredentialBadRequest struct {
}

// IsSuccess returns true when this user update git credential bad request response has a 2xx status code
func (o *UserUpdateGitCredentialBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update git credential bad request response has a 3xx status code
func (o *UserUpdateGitCredentialBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update git credential bad request response has a 4xx status code
func (o *UserUpdateGitCredentialBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user update git credential bad request response has a 5xx status code
func (o *UserUpdateGitCredentialBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user update git credential bad request response a status code equal to that given
func (o *UserUpdateGitCredentialBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user update git credential bad request response
func (o *UserUpdateGitCredentialBadRequest) Code() int {
	return 400
}

func (o *UserUpdateGitCredentialBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialBadRequest", 400)
}

func (o *UserUpdateGitCredentialBadRequest) String() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialBadRequest", 400)
}

func (o *UserUpdateGitCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateGitCredentialForbidden creates a UserUpdateGitCredentialForbidden with default headers values
func NewUserUpdateGitCredentialForbidden() *UserUpdateGitCredentialForbidden {
	return &UserUpdateGitCredentialForbidden{}
}

/*
UserUpdateGitCredentialForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type UserUpdateGitCredentialForbidden struct {
}

// IsSuccess returns true when this user update git credential forbidden response has a 2xx status code
func (o *UserUpdateGitCredentialForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update git credential forbidden response has a 3xx status code
func (o *UserUpdateGitCredentialForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update git credential forbidden response has a 4xx status code
func (o *UserUpdateGitCredentialForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user update git credential forbidden response has a 5xx status code
func (o *UserUpdateGitCredentialForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user update git credential forbidden response a status code equal to that given
func (o *UserUpdateGitCredentialForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user update git credential forbidden response
func (o *UserUpdateGitCredentialForbidden) Code() int {
	return 403
}

func (o *UserUpdateGitCredentialForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialForbidden", 403)
}

func (o *UserUpdateGitCredentialForbidden) String() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialForbidden", 403)
}

func (o *UserUpdateGitCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateGitCredentialNotFound creates a UserUpdateGitCredentialNotFound with default headers values
func NewUserUpdateGitCredentialNotFound() *UserUpdateGitCredentialNotFound {
	return &UserUpdateGitCredentialNotFound{}
}

/*
UserUpdateGitCredentialNotFound describes a response with status code 404, with default header values.

Not found
*/
type UserUpdateGitCredentialNotFound struct {
}

// IsSuccess returns true when this user update git credential not found response has a 2xx status code
func (o *UserUpdateGitCredentialNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update git credential not found response has a 3xx status code
func (o *UserUpdateGitCredentialNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update git credential not found response has a 4xx status code
func (o *UserUpdateGitCredentialNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user update git credential not found response has a 5xx status code
func (o *UserUpdateGitCredentialNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user update git credential not found response a status code equal to that given
func (o *UserUpdateGitCredentialNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user update git credential not found response
func (o *UserUpdateGitCredentialNotFound) Code() int {
	return 404
}

func (o *UserUpdateGitCredentialNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialNotFound", 404)
}

func (o *UserUpdateGitCredentialNotFound) String() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialNotFound", 404)
}

func (o *UserUpdateGitCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserUpdateGitCredentialInternalServerError creates a UserUpdateGitCredentialInternalServerError with default headers values
func NewUserUpdateGitCredentialInternalServerError() *UserUpdateGitCredentialInternalServerError {
	return &UserUpdateGitCredentialInternalServerError{}
}

/*
UserUpdateGitCredentialInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UserUpdateGitCredentialInternalServerError struct {
}

// IsSuccess returns true when this user update git credential internal server error response has a 2xx status code
func (o *UserUpdateGitCredentialInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user update git credential internal server error response has a 3xx status code
func (o *UserUpdateGitCredentialInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update git credential internal server error response has a 4xx status code
func (o *UserUpdateGitCredentialInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user update git credential internal server error response has a 5xx status code
func (o *UserUpdateGitCredentialInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user update git credential internal server error response a status code equal to that given
func (o *UserUpdateGitCredentialInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user update git credential internal server error response
func (o *UserUpdateGitCredentialInternalServerError) Code() int {
	return 500
}

func (o *UserUpdateGitCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialInternalServerError", 500)
}

func (o *UserUpdateGitCredentialInternalServerError) String() string {
	return fmt.Sprintf("[PUT /users/{id}/gitcredentials/{credentialID}][%d] userUpdateGitCredentialInternalServerError", 500)
}

func (o *UserUpdateGitCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
