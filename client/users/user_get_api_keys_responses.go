// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/portainer/client-api/models"
)

// UserGetAPIKeysReader is a Reader for the UserGetAPIKeys structure.
type UserGetAPIKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetAPIKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetAPIKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGetAPIKeysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGetAPIKeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGetAPIKeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGetAPIKeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGetAPIKeysOK creates a UserGetAPIKeysOK with default headers values
func NewUserGetAPIKeysOK() *UserGetAPIKeysOK {
	return &UserGetAPIKeysOK{}
}

/* UserGetAPIKeysOK describes a response with status code 200, with default header values.

Success
*/
type UserGetAPIKeysOK struct {
	Payload []*models.PortainerAPIKey
}

func (o *UserGetAPIKeysOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/tokens][%d] userGetApiKeysOK  %+v", 200, o.Payload)
}
func (o *UserGetAPIKeysOK) GetPayload() []*models.PortainerAPIKey {
	return o.Payload
}

func (o *UserGetAPIKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetAPIKeysBadRequest creates a UserGetAPIKeysBadRequest with default headers values
func NewUserGetAPIKeysBadRequest() *UserGetAPIKeysBadRequest {
	return &UserGetAPIKeysBadRequest{}
}

/* UserGetAPIKeysBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type UserGetAPIKeysBadRequest struct {
}

func (o *UserGetAPIKeysBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}/tokens][%d] userGetApiKeysBadRequest ", 400)
}

func (o *UserGetAPIKeysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetAPIKeysForbidden creates a UserGetAPIKeysForbidden with default headers values
func NewUserGetAPIKeysForbidden() *UserGetAPIKeysForbidden {
	return &UserGetAPIKeysForbidden{}
}

/* UserGetAPIKeysForbidden describes a response with status code 403, with default header values.

Permission denied
*/
type UserGetAPIKeysForbidden struct {
}

func (o *UserGetAPIKeysForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}/tokens][%d] userGetApiKeysForbidden ", 403)
}

func (o *UserGetAPIKeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetAPIKeysNotFound creates a UserGetAPIKeysNotFound with default headers values
func NewUserGetAPIKeysNotFound() *UserGetAPIKeysNotFound {
	return &UserGetAPIKeysNotFound{}
}

/* UserGetAPIKeysNotFound describes a response with status code 404, with default header values.

User not found
*/
type UserGetAPIKeysNotFound struct {
}

func (o *UserGetAPIKeysNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}/tokens][%d] userGetApiKeysNotFound ", 404)
}

func (o *UserGetAPIKeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGetAPIKeysInternalServerError creates a UserGetAPIKeysInternalServerError with default headers values
func NewUserGetAPIKeysInternalServerError() *UserGetAPIKeysInternalServerError {
	return &UserGetAPIKeysInternalServerError{}
}

/* UserGetAPIKeysInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UserGetAPIKeysInternalServerError struct {
}

func (o *UserGetAPIKeysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{id}/tokens][%d] userGetApiKeysInternalServerError ", 500)
}

func (o *UserGetAPIKeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
