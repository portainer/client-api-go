// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api-go/v2/models_ce"
)

// UserGenerateAPIKeyReader is a Reader for the UserGenerateAPIKey structure.
type UserGenerateAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGenerateAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUserGenerateAPIKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGenerateAPIKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserGenerateAPIKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGenerateAPIKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGenerateAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGenerateAPIKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGenerateAPIKeyCreated creates a UserGenerateAPIKeyCreated with default headers values
func NewUserGenerateAPIKeyCreated() *UserGenerateAPIKeyCreated {
	return &UserGenerateAPIKeyCreated{}
}

/*
UserGenerateAPIKeyCreated describes a response with status code 201, with default header values.

Created
*/
type UserGenerateAPIKeyCreated struct {
	Payload *models_ce.UsersAccessTokenResponse
}

// IsSuccess returns true when this user generate Api key created response has a 2xx status code
func (o *UserGenerateAPIKeyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user generate Api key created response has a 3xx status code
func (o *UserGenerateAPIKeyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user generate Api key created response has a 4xx status code
func (o *UserGenerateAPIKeyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this user generate Api key created response has a 5xx status code
func (o *UserGenerateAPIKeyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this user generate Api key created response a status code equal to that given
func (o *UserGenerateAPIKeyCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UserGenerateAPIKeyCreated) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyCreated  %+v", 201, o.Payload)
}

func (o *UserGenerateAPIKeyCreated) String() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyCreated  %+v", 201, o.Payload)
}

func (o *UserGenerateAPIKeyCreated) GetPayload() *models_ce.UsersAccessTokenResponse {
	return o.Payload
}

func (o *UserGenerateAPIKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_ce.UsersAccessTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGenerateAPIKeyBadRequest creates a UserGenerateAPIKeyBadRequest with default headers values
func NewUserGenerateAPIKeyBadRequest() *UserGenerateAPIKeyBadRequest {
	return &UserGenerateAPIKeyBadRequest{}
}

/*
UserGenerateAPIKeyBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UserGenerateAPIKeyBadRequest struct {
}

// IsSuccess returns true when this user generate Api key bad request response has a 2xx status code
func (o *UserGenerateAPIKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user generate Api key bad request response has a 3xx status code
func (o *UserGenerateAPIKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user generate Api key bad request response has a 4xx status code
func (o *UserGenerateAPIKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user generate Api key bad request response has a 5xx status code
func (o *UserGenerateAPIKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user generate Api key bad request response a status code equal to that given
func (o *UserGenerateAPIKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserGenerateAPIKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyBadRequest ", 400)
}

func (o *UserGenerateAPIKeyBadRequest) String() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyBadRequest ", 400)
}

func (o *UserGenerateAPIKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGenerateAPIKeyUnauthorized creates a UserGenerateAPIKeyUnauthorized with default headers values
func NewUserGenerateAPIKeyUnauthorized() *UserGenerateAPIKeyUnauthorized {
	return &UserGenerateAPIKeyUnauthorized{}
}

/*
UserGenerateAPIKeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGenerateAPIKeyUnauthorized struct {
}

// IsSuccess returns true when this user generate Api key unauthorized response has a 2xx status code
func (o *UserGenerateAPIKeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user generate Api key unauthorized response has a 3xx status code
func (o *UserGenerateAPIKeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user generate Api key unauthorized response has a 4xx status code
func (o *UserGenerateAPIKeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user generate Api key unauthorized response has a 5xx status code
func (o *UserGenerateAPIKeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user generate Api key unauthorized response a status code equal to that given
func (o *UserGenerateAPIKeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserGenerateAPIKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyUnauthorized ", 401)
}

func (o *UserGenerateAPIKeyUnauthorized) String() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyUnauthorized ", 401)
}

func (o *UserGenerateAPIKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGenerateAPIKeyForbidden creates a UserGenerateAPIKeyForbidden with default headers values
func NewUserGenerateAPIKeyForbidden() *UserGenerateAPIKeyForbidden {
	return &UserGenerateAPIKeyForbidden{}
}

/*
UserGenerateAPIKeyForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type UserGenerateAPIKeyForbidden struct {
}

// IsSuccess returns true when this user generate Api key forbidden response has a 2xx status code
func (o *UserGenerateAPIKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user generate Api key forbidden response has a 3xx status code
func (o *UserGenerateAPIKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user generate Api key forbidden response has a 4xx status code
func (o *UserGenerateAPIKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user generate Api key forbidden response has a 5xx status code
func (o *UserGenerateAPIKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user generate Api key forbidden response a status code equal to that given
func (o *UserGenerateAPIKeyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserGenerateAPIKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyForbidden ", 403)
}

func (o *UserGenerateAPIKeyForbidden) String() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyForbidden ", 403)
}

func (o *UserGenerateAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGenerateAPIKeyNotFound creates a UserGenerateAPIKeyNotFound with default headers values
func NewUserGenerateAPIKeyNotFound() *UserGenerateAPIKeyNotFound {
	return &UserGenerateAPIKeyNotFound{}
}

/*
UserGenerateAPIKeyNotFound describes a response with status code 404, with default header values.

User not found
*/
type UserGenerateAPIKeyNotFound struct {
}

// IsSuccess returns true when this user generate Api key not found response has a 2xx status code
func (o *UserGenerateAPIKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user generate Api key not found response has a 3xx status code
func (o *UserGenerateAPIKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user generate Api key not found response has a 4xx status code
func (o *UserGenerateAPIKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user generate Api key not found response has a 5xx status code
func (o *UserGenerateAPIKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user generate Api key not found response a status code equal to that given
func (o *UserGenerateAPIKeyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserGenerateAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyNotFound ", 404)
}

func (o *UserGenerateAPIKeyNotFound) String() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyNotFound ", 404)
}

func (o *UserGenerateAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGenerateAPIKeyInternalServerError creates a UserGenerateAPIKeyInternalServerError with default headers values
func NewUserGenerateAPIKeyInternalServerError() *UserGenerateAPIKeyInternalServerError {
	return &UserGenerateAPIKeyInternalServerError{}
}

/*
UserGenerateAPIKeyInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UserGenerateAPIKeyInternalServerError struct {
}

// IsSuccess returns true when this user generate Api key internal server error response has a 2xx status code
func (o *UserGenerateAPIKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user generate Api key internal server error response has a 3xx status code
func (o *UserGenerateAPIKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user generate Api key internal server error response has a 4xx status code
func (o *UserGenerateAPIKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user generate Api key internal server error response has a 5xx status code
func (o *UserGenerateAPIKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user generate Api key internal server error response a status code equal to that given
func (o *UserGenerateAPIKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserGenerateAPIKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyInternalServerError ", 500)
}

func (o *UserGenerateAPIKeyInternalServerError) String() string {
	return fmt.Sprintf("[POST /users/{id}/tokens][%d] userGenerateApiKeyInternalServerError ", 500)
}

func (o *UserGenerateAPIKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
